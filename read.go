package apkutils

import (
	"archive/tar"
	"bufio"
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"strings"
)

// ReadTarball takes an Alpine APK, verifies the embedded signature using the
// public keys provided by the keyProvider, verifies the hash of the tarball,
// and returns a reader over the contents of the APK tarball.
func ReadTarball(reader io.Reader, keyProvider KeyProvider) (*tar.Reader, error) {
	controlTar, tarballGz, err := read(reader, keyProvider)
	if err != nil {
		return nil, err
	}

	pkginfo, err := extractPkgInfo(controlTar)
	if err != nil {
		return nil, fmt.Errorf("failed to read PKGINFO: %v", err)
	}

	h := sha256.New()
	h.Write(tarballGz)
	hash := fmt.Sprintf("%x", h.Sum(nil))

	if hash != pkginfo["datahash"] {
		return nil, fmt.Errorf("payload inconsistent with datahash")
	}

	gz, err := gzip.NewReader(bytes.NewBuffer(tarballGz))
	if err != nil {
		return nil, err
	}

	return tar.NewReader(gz), nil
}

// read takes an Alpine APK or APKINDEX file, verifies the embedded signature,
// and returns the content segment and any remaining bytes (which may be
// additional segments not covered by the signature).
func read(reader io.Reader, keyProvider KeyProvider) (content, remaining []byte, err error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, nil, err
	}

	if err = Verify(bytes.NewBuffer(b), keyProvider); err != nil {
		return nil, nil, fmt.Errorf("verification failed: %v", err)
	}

	buffer := bytes.NewBuffer(b)
	gz, err := gzip.NewReader(buffer)
	if err != nil {
		return nil, nil, err
	}

	// Skip the signature stream
	gz.Multistream(false)
	if _, err = io.ReadAll(gz); err != nil {
		return nil, nil, err
	}

	// Reset the reader for the main stream
	if err = gz.Reset(buffer); err != nil {
		return nil, nil, err
	}
	gz.Multistream(false)

	content, err = io.ReadAll(gz)
	if err != nil {
		return nil, nil, err
	}

	return content, buffer.Bytes(), nil
}

// readFile reads the given tar and returns the contents of the file at
// the given path.
func readFile(tarBytes []byte, path string) ([]byte, error) {
	t := tar.NewReader(bytes.NewReader(tarBytes))
	for {
		h, err := t.Next()
		if err == io.EOF {
			return nil, fmt.Errorf("file %s not found", path)
		}

		if h.Name == path {
			return io.ReadAll(t)
		}

		_, err = io.ReadAll(t)
		if err != nil {
			return nil, err
		}
	}
}

// extractPkgInfo reads the given tar archive, extracts the PKGINFO file,
// and parses it into a map.
func extractPkgInfo(tarBytes []byte) (map[string]string, error) {
	fileBytes, err := readFile(tarBytes, ".PKGINFO")
	if err != nil {
		return nil, err
	}

	scan := bufio.NewScanner(bytes.NewReader(fileBytes))
	res := make(map[string]string)
	for scan.Scan() {
		line := scan.Text()

		if strings.HasPrefix(line, "#") || len(line) == 0 {
			continue
		}

		parts := strings.SplitN(scan.Text(), " = ", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid pkginfo line: %s", line)
		}

		res[parts[0]] = parts[1]
	}

	return res, scan.Err()
}
