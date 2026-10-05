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
	entries, err := ReadApkIndexEntries(reader, keyProvider)
	if err != nil {
		return nil, err
	}

	res := make(map[string]*PackageInfo)
	for _, entry := range entries {
		res[entry.Name] = entry
		for i := range entry.Provides {
			// Don't overwrite real packages with provides info
			if _, ok := res[entry.Provides[i]]; !ok {
				res[entry.Provides[i]] = entry
			}
		}
	}

	return res, nil
}

// ReadApkIndexEntries verifies the embedded signature in the file, then
// extracts and parses the APKINDEX contents, returning every package entry
// in the order they appear in the file.
func ReadApkIndexEntries(reader io.Reader, keyProvider KeyProvider) ([]*PackageInfo, error) {
	tarBytes, _, err := read(reader, keyProvider)
	if err != nil {
		return nil, err
	}

	indexBytes, err := readFile(tarBytes, "APKINDEX")
	if err != nil {
		return nil, err
	}

	return readApkIndexEntries(bytes.NewReader(indexBytes))
}

// readApkIndexEntries reads an APKINDEX file, parsing out every contained
// package entry in file order.
func readApkIndexEntries(reader io.Reader) ([]*PackageInfo, error) {
	var res []*PackageInfo
	scanner := bufio.NewScanner(reader)
	buf := make([]byte, 0, 1024*1024)
	scanner.Buffer(buf, 1024*1024)

	current := &PackageInfo{}
	started := false
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			if started {
				res = append(res, current)
			}
			current = &PackageInfo{}
			started = false
			continue
		}
		started = true
		if strings.HasPrefix(line, "P:") {
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

	if started {
		res = append(res, current)
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
