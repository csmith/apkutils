package main

import (
	"github.com/csmith/apkutils"
	"github.com/csmith/apkutils/keys"
	"log"
	"net/http"
)

func main() {
	p := packages()

	deps, err := apkutils.FlattenDependencies(p, "bash", "git")
	if err != nil {
		log.Fatal(err)
	}

	for i := range deps {
		log.Printf("Dependency: %s", deps[i].Name)
	}
}

// packages retrieves and parses a copy of the index. See the verify-and-parse-index example for further comment.
func packages() map[string]*apkutils.PackageInfo {
	const exampleIndex = "https://dl-cdn.alpinelinux.org/alpine/latest-stable/main/x86_64/APKINDEX.tar.gz"
	r, err := http.Get(exampleIndex)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	res, err := apkutils.ReadApkIndex(r.Body, apkutils.NewFileSystemKeyProvider(keys.X86_64))
	if err != nil {
		log.Fatal(err)
	}
	return res
}
