# apkutils

apkutils contains utilities for operating on Alpine packages and package indices
from Go.

Specifically it can:

- Verify the public key signature in an APKINDEX.tar.gz bundle, then parse out
  the contents of the index. See `examples/verify-and-parse-index`.
- Verify the public key signature in an APK file, and then verify the hash
  of the contents that is stored in the metadata, then return the content of
  the APK. See `examples/verify-and-read-apk`
- Given the package information from an APKINDEX, recursively resolve
  dependencies of packages to generate a bill of materials/list of packages
  to install. See `examples/flatten-dependencies`

Contributions of further functions or enhancements are welcome!
