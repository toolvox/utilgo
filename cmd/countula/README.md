# countula

`countula` is a command-line tool designed to count lines of code across a directory tree, offering control over the inclusion and exclusion of files through glob patterns.

## Installation

Ensure you have a Go environment set up, then clone the repository and build the tool:

```sh
git clone https://github.com/toolvox/utilgo
cd countula
go build
```

## Usage

### Flags

- `-root`: Root directory for file traversal. Defaults to the current directory.
- `-include`: Comma-separated list of glob patterns for files to include.
- `-exclude`: Comma-separated list of glob patterns for files to exclude.
- `-out`: Output file path for the report. Defaults to stdout if not specified.
- `-ignore-prefix`: Comma-separated list of line prefixes to ignore.
- `-dir-mode`: Enables grouping of counts by directory.
- `-merge-mode`: Merges counts across file types without splitting by extension.

### Example Command

Count lines in `.go` and `.js` files, excluding the `vendor/` directory and `.test.js` files, ignoring lines starting with `//` or `#`, and group counts by directory:

```sh
countula -root "./source" -out "count_report.txt" -include ".go,.js" -exclude "vendor/,.test.js" -ignore-prefix "//,#" -dir-mode
```

## License

`codump` is made available under the [MIT License](LICENSE). For more details, see the LICENSE file in the repository.

## Support

For support, questions, or more information about using `codump`, please visit the project's issues page on GitHub.
