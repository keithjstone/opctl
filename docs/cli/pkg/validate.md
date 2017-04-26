## `pkg validate PKG_REF` (since v0.1.15)

Validates a package according to:

- manifest existence
- manifest validity (per
  [schema](https://opspec.io/spec/0.1.4/package-manifest.schema.json))

## Arguments

### `PKG_REF`
Package reference (either `relative/path`, `/absolute/path`, or `host/path/repo#tag` (since v0.1.19))