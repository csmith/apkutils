package main

import (
	"github.com/csmith/apkutils/v2"
	"github.com/csmith/apkutils/v2/keys"
	"log"
	"net/http"
)

func main() {
	// Obtain a copy of the index
	const exampleIndex = "https://dl-cdn.alpinelinux.org/alpine/latest-stable/main/x86_64/APKINDEX.tar.gz"
	res, err := http.Get(exampleIndex)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// Use apkutils to read the index. The second argument specifies where to look for the public keys used to
	// verify the index signature. apkutil bundles these keys for convenience, but you can implement your own
	// KeyProvider if you have an alternative source (such as /usr/share/apk/keys/ on alpine systems).
	packages, err := apkutils.ReadApkIndex(res.Body, keys.X86_64)
	if err != nil {
		log.Fatal(err)
	}

	// The result is a map of package names and provided labels to details about the package
	log.Printf("Found %d packages", len(packages))
	log.Printf("Details of the alpine-keys package: %#v", packages["alpine-keys"])
	log.Printf("A package that provides the go binary: %#v", packages["cmd:go"])
}
