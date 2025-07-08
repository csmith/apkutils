# Changelog

## 2.1.0 - 2025-07-08

### Changes

- Minor dependency updates.

## 2.0.0 - 2024-12-20

### Breaking changes

- The exported vars in the `keys` package are now KeyProviders instead of
  filesystems containing keys, so they can be used directly without wrapping.

### Changes

- Added `keys.All` key provider that contains all known Alpine Linux keys,
  regardless of architecture.

### Fixes

- The file system key provider now actually uses the passed file system, instead
  of ignoring it and reading from disk.
- When trying to use a key provider, the name of the key is no longer
  incorrectly prefixed with "key/".

## 1.0.0 - 2024-12-19

_Initial release._
