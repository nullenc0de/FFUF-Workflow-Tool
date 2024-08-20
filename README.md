# FFUF Workflow Tool

This tool integrates ffuf and ffufPostprocessing to streamline your workflow.

## Requirements

- Go 1.21 or later
- `ffuf` and `ffufPostprocessing` tools installed and available in your PATH

## Installation

### Option 1: Using `go install`

```
go install github.com/yourusername/ffuf-workflow/cmd/ffuf-workflow@latest
```

Make sure your Go bin directory is in your PATH.

### Option 2: Cloning the repository

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/ffuf-workflow.git
   cd ffuf-workflow
   ```

2. Build and install:
   ```
   make install
   ```

## Usage

```
ffuf-workflow -wordlist path/to/wordlist.txt -output path/to/output.txt
```

Or pipe input:

```
cat urls.txt | ffuf-workflow -output path/to/output.txt
```

## Compatibility

This tool is designed to be compatible with Go 1.21 and later versions. If you encounter any issues with newer Go versions, please report them in the GitHub issues section.
