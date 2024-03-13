# codump

`codump` is a command-line utility designed to aggregate the contents of files from a directory tree into a single output file or stream.

## Installation

To install `codump`, you will need a Go development environment. Follow these steps:

1. Clone the repository to your local machine.
2. Navigate to the cloned directory.
3. Build the project using `go build`, which will produce the `codump` executable.

```sh
git clone https://github.com/toolvox/utilgo
cd cmd/codump
go install
```

## Usage

`codump` operates with several command-line flags that allow users to specify the root directory, include and exclude patterns, and the output destination.

### Flags

- `-root <directory>`: The root directory from which to start the traversal. Defaults to the current directory.
- `-include <patterns>`: A comma-separated list of glob patterns for files to include.
- `-exclude <patterns>`: A comma-separated list of glob patterns for files to exclude.
- `-out <filepath>`: The output file path for the aggregated content. Defaults to standard output if unspecified.

### Example

Compile `.go` and `.md` files from the `./project/target` directory, excluding anything in the `.git` directory:

```sh
$ codump -root "./project/target" -exclude ".git" -include "*.go,*.md"
```

## License

`codump` is made available under the [MIT License](LICENSE). For more details, see the LICENSE file in the repository.

## Support

For support, questions, or more information about using `codump`, please visit the project's issues page on GitHub.
