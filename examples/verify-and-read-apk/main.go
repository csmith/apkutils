package main

import (
	"github.com/csmith/apkutils/v2"
	"github.com/csmith/apkutils/v2/keys"
	"io"
	"log"
	"net/http"
)

func main() {
	// Obtain a copy of a package
	const exampleIndex = "https://dl-cdn.alpinelinux.org/alpine/latest-stable/main/x86_64/alpine-keys-2.5-r0.apk"
	res, err := http.Get(exampleIndex)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// Use apkutils to read the package. The second argument specifies where to look for the public keys used to
	// verify the index signature. apkutil bundles these keys for convenience, but you can implement your own
	// KeyProvider if you have an alternative source (such as /usr/share/apk/keys/ on alpine systems).
	reader, err := apkutils.ReadTarball(res.Body, keys.X86_64)
	if err != nil {
		log.Fatal(err)
	}

	for {
		header, err := reader.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		log.Printf("APK contains file: %s", header.Name)
		_, _ = io.ReadAll(reader)
	}
}
