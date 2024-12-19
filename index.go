package apkutils

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

// ReadApkIndex verifies the embedded signature in the file, then extracts
// and parses the APKINDEX contents.
func ReadApkIndex(reader io.Reader, keyProvider KeyProvider) (map[string]*PackageInfo, error) {
	tarBytes, _, err := read(reader, keyProvider)
	if err != nil {
		return nil, err
	}

	indexBytes, err := readFile(tarBytes, "APKINDEX")
	if err != nil {
		return nil, err
	}

	return readApkIndexContent(bytes.NewReader(indexBytes))
}

// readApkIndexContent reads an APKINDEX file, parsing out the contained packages.
func readApkIndexContent(reader io.Reader) (map[string]*PackageInfo, error) {
	res := make(map[string]*PackageInfo)
	scanner := bufio.NewScanner(reader)
	buf := make([]byte, 0, 1024*1024)
	scanner.Buffer(buf, 1024*1024)

	current := &PackageInfo{}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			res[current.Name] = current

			for i := range current.Provides {
				// Don't overwrite real packages with provides info
				if _, ok := res[current.Provides[i]]; !ok {
					res[current.Provides[i]] = current
				}
			}

			current = &PackageInfo{}
		} else if strings.HasPrefix(line, "P:") {
			current.Name = strings.TrimPrefix(line, "P:")
		} else if strings.HasPrefix(line, "D:") {
			d := strings.Fields(strings.TrimPrefix(line, "D:"))
			for i := range d {
				current.Dependencies = append(current.Dependencies, stripVersion(d[i]))
			}
		} else if strings.HasPrefix(line, "p:") {
			p := strings.Fields(strings.TrimPrefix(line, "p:"))
			for i := range p {
				current.Provides = append(current.Provides, stripVersion(p[i]))
			}
		} else if strings.HasPrefix(line, "V:") {
			current.Version = strings.TrimPrefix(line, "V:")
		}
	}

	if scanner.Err() != nil {
		return nil, fmt.Errorf("unable to read index: %v", scanner.Err())
	}

	return res, nil
}

// PackageInfo describes a package available in a repository.
type PackageInfo struct {
	Name         string
	Version      string
	Dependencies []string
	Provides     []string
}

// stripVersion removes version qualifiers from a package name such as `foo>=1.2`.
func stripVersion(name string) string {
	i := strings.IndexAny(name, ">=<~")
	if i > -1 {
		return name[0:i]
	}
	return name
}
