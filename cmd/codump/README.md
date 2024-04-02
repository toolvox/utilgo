# Codump

`codump` is a command-line utility designed to traverse a directory tree and compile the contents of files that match specified criteria into a single output file or stream. It leverages flexible include and exclude filtering capabilities, allowing users to specify precisely which files should be aggregated based on glob patterns. This utility is versatile for various scripting, logging, and data aggregation purposes.

## Features

- Traverse directories starting from a specified root or the current working directory.
- Include or exclude files based on glob patterns for precise control over which files are aggregated.
- Supports output to both files and standard output/error, allowing for flexible integration into workflows and scripts.

## Installation

To install `codump`, you'll need to build it from the source:

```sh
# Clone the repository
git clone https://github.com/toolvox/utilgo.git
# Navigate to the directory
cd ./toolvox/utilgo
# Build and install codump
go install ./cmd/codump
```

Make sure you have Go installed on your system to build the tool.

## Usage

Here are some examples demonstrating how to use `codump`.

### Basic Usage

To compile all files from the current directory to standard output:

```sh
$ codump
```

### Specifying Root Directory

To start traversal from a specific directory:

```sh
$ codump -root "./project/target"
```

### Including and Excluding Files

To only include `.go` and `.md` files, excluding any files in a `.git` directory:

```sh
$ codump -include "*.go,*.md" -exclude ".git/*"
```

### Output to a File

To direct the aggregated content to a file named `aggregate.txt`:

```sh
$ codump -out "aggregate.txt"
```

## Example Output

When aggregating `.go` and `.md` files, the output might look like this (assuming output to stdout):

```
--- ./project/target/main.go
package main

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}
---
--- ./project/target/README.md
# Project Title

This is a sample project.
---
```

Each file's content is prefixed and suffixed with `---` and its path, clearly delineating where each file's content begins and ends.

## Reporting Issues

If you encounter any problems or have feature suggestions, please open an issue on the GitHub repository.

## License

`codump` is released under the MIT License. For more details, see the [LICENSE](LICENSE) file in the repository.