package apkutils

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"
	"io"
	"io/fs"
	"strings"
)

// A KeyProvider supplies public keys for use in package verification
type KeyProvider interface {
	Key(name string) (*rsa.PublicKey, error)
}

type fileSystemKeyProvider struct {
	fs fs.FS
}

// NewFileSystemKeyProvider creates a new KeyProvider that will load PEM
// encoded public keys from the root of the given filesystem.
//
// No validation is performed on key names; the filesystem should be
// appropriately rooted to ensure only key material is accessible.
func NewFileSystemKeyProvider(fs fs.FS) KeyProvider {
	return &fileSystemKeyProvider{
		fs: fs,
	}
}

func (f *fileSystemKeyProvider) Key(name string) (*rsa.PublicKey, error) {
	b, err := fs.ReadFile(f.fs, name)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(b)
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return pub.(*rsa.PublicKey), nil
}

// Verify checks the embedded signature within an alpine APK or APKINDEX file.
//
// These files are concatenations of two or more gzip streams, the first of
// which contains a signature of the (compressed) second stream.
//
// A valid signature is indicated by returning a nil error.
func Verify(reader io.Reader, keyProvider KeyProvider) error {
	b, err := io.ReadAll(reader)
	byteBuffer := bytes.NewBuffer(b)

	gz, err := gzip.NewReader(byteBuffer)
	if err != nil {
		return err
	}

	gz.Multistream(false)
	sigBytes, err := io.ReadAll(gz)
	if err != nil {
		return err
	}

	key, sig, err := readSignature(sigBytes)
	if err != nil {
		return err
	}

	publicKey, err := keyProvider.Key(key)
	if err != nil {
		return err
	}

	// We want to read the second stream to figure out the end position,
	// but we actually want the raw (compressed) data to verify the signature.
	start := len(b) - len(byteBuffer.Bytes())
	err = gz.Reset(byteBuffer)
	if err != nil {
		return err
	}

	gz.Multistream(false)
	_, err = io.ReadAll(gz)
	if err != nil {
		return err
	}
	end := len(b) - len(byteBuffer.Bytes())

	hash, err := calculateHash(b[start:end])
	if err != nil {
		return err
	}

	return rsa.VerifyPKCS1v15(publicKey, crypto.SHA1, hash, sig)
}

// calculateHash returns the SHA-1 hash of the data.
func calculateHash(data []byte) ([]byte, error) {
	h := sha1.New()
	_, err := io.Copy(h, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

// readSignature reads the details of the APK signature from the given tar segment.
func readSignature(segment []byte) (key string, sig []byte, err error) {
	tarReader := tar.NewReader(bytes.NewReader(segment))
	h, err := tarReader.Next()
	if err != nil {
		return "", nil, err
	}

	key = strings.TrimPrefix(h.Name, ".SIGN.RSA.")
	sig, err = io.ReadAll(tarReader)
	if err != nil {
		return "", nil, err
	}

	return
}
