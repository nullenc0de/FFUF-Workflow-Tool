# FFUF Workflow Tool

This tool integrates ffuf and ffufPostprocessing to streamline your workflow.

## Requirements

- Go 1.21 or later
- `ffuf` and `ffufPostprocessing` tools installed and available in your PATH

## Installation

To install the latest version of the FFUF Workflow Tool, use the following command:

```
go install github.com/nullenc0de/FFUF-Workflow-Tool/cmd/ffuf-workflow@latest
```

Make sure your Go bin directory is in your PATH. You can verify the installation by running:

```
which ffuf-workflow
ffuf-workflow -h
```

## Usage

```
ffuf-workflow -wordlist path/to/wordlist.txt -output path/to/output.txt
```

Or pipe input:

```
cat urls.txt | ffuf-workflow -output path/to/output.txt
```

## Building from Source

If you prefer to build from source or want to contribute to the project, you can clone the repository and build manually:

```
git clone https://github.com/nullenc0de/FFUF-Workflow-Tool.git
cd FFUF-Workflow-Tool
go build -o ffuf-workflow ./cmd/ffuf-workflow
```

You can then move the built binary to a directory in your PATH or run it directly from the build location.

## Compatibility

This tool is designed to be compatible with Go 1.18 and later versions. If you encounter any issues with your Go version, please report them in the GitHub issues section.

## Reporting Issues

If you encounter any problems or have suggestions for improvements, please open an issue on the [GitHub repository](https://github.com/nullenc0de/FFUF-Workflow-Tool/issues).
