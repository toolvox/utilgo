# Countula

`countula` is a command-line tool designed for counting lines of code (LOC) across various file types in a directory tree. 
It provides users with granular control to include or exclude specific files based on glob patterns and offers options to ignore lines with certain prefixes, such as comments.

## Features

- Flexible root directory specification.
- Inclusion and exclusion of files through glob patterns.
- Ignoring lines with specified prefixes.
- Output can be directed to a file or stdout.
- Option to group output by directory.
- Merge mode for combined counts instead of per-extension details.

## Installation

Currently, `countula` needs to be built from source:

```sh
# clone the repo
git clone https://github.com/toolvox/utilgo.git
# navigate to new .../utilgo directory
cd ./toolvox/utilgo
# go install
go install ./cmd/countula
```

Ensure you have Go installed on your system.

## Usage

Below are some examples of how to use `countula`.

### Basic Usage

To count all lines of code starting from the current directory and output to stdout:

```sh
$ countula
```

### Specifying Root Directory

To specify a root directory other than the current directory:

```sh
$ countula -root "./my_project"
```

### Including and Excluding Files

To include only `.go`, `.json`, and `.md` files, and exclude any files in a `.git` directory:

```sh
$ countula -include "*.go,*.json,*.md" -exclude ".git/*"
```

### Ignoring Lines

To ignore lines that start with `//` or `#` (common comment patterns in many languages):

```sh
$ countula -ignore-prefix "//,#"
```

### Output to a File

To write the output to a file named `loc_report.txt`:

```sh
$ countula -out "loc_report.txt"
```

### Directory and Merge Modes

To group counts by directory and not split counts by file extension:

```sh
$ countula -dir-mode -merge-mode
```

## Example Output

When running with `-dir-mode` and without `-merge-mode`:

```
Directory `src`:
    `.go`: 1200
    `.js`: 300

Directory `tests`:
    `.go`: 400

  ===  
Total lines: 1900
```

With `-merge-mode` enabled, the output might look like this:

```
Total lines: 1900
```

Or, with `-dir-mode` and `-merge-mode`:

```
Directory `src`:
    lines: 1500

Directory `tests`:
    lines: 400

  ===  
Total lines: 1900
```

Without any modes enabled, showing counts by extension globally:

```
Counts by extension:
    `.go`: 1600
    `.js`: 300

  ===  
Total lines: 1900
```

## Reporting Issues

For any issues or feature requests, please open an issue on the GitHub repository.

## License

`countula` is released under the MIT License. See the [LICENSE](LICENSE) file for more details.