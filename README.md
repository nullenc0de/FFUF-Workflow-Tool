# FFUF Workflow Tool

A streamlined workflow integration tool that combines the power of ffuf and ffufPostprocessing for efficient web fuzzing.

## Prerequisites

Before installing FFUF Workflow Tool, ensure you have:

```bash
# 1. Go 1.21 or later
go version  # Should show 1.21 or higher

# 2. Install ffuf
go install github.com/ffuf/ffuf@latest

# 3. Install ffufPostprocessing
git clone https://github.com/dsecuredcom/ffufPostprocessing.git
cd ffufPostprocessing
go build -o ffufPostprocessing main.go
sudo mv ffufPostprocessing /usr/local/bin/
```

## Installation

Install FFUF Workflow Tool using one of these methods:

### 1. Using go install (Recommended)
```bash
go install github.com/nullenc0de/FFUF-Workflow-Tool/cmd/ffuf-workflow@latest
```

### 2. Building from source
```bash
git clone https://github.com/nullenc0de/FFUF-Workflow-Tool.git
cd FFUF-Workflow-Tool
go build -o ffuf-workflow ./cmd/ffuf-workflow
sudo mv ffuf-workflow /usr/local/bin/
```

### Verify Installation
```bash
# Check if installation was successful
which ffuf-workflow
ffuf-workflow -h
```

## Usage

### Basic Usage
```bash
# Using a wordlist
ffuf-workflow -wordlist /path/to/wordlist.txt -output results.txt

# Using pipe input
cat urls.txt | ffuf-workflow -output results.txt
```

### Command Line Options
```bash
ffuf-workflow -h
  -wordlist string    Path to wordlist file
  -output string      Path to output file
  [additional options will be listed here]
```

## Example Workflows

### Single Target Scan
```bash
ffuf-workflow -wordlist wordlist.txt -output target1_results.txt
```

### Multiple Targets
```bash
echo "example.com" > urls.txt
echo "example2.com" >> urls.txt
cat urls.txt | ffuf-workflow -output multi_results.txt
```

## Integration with Other Tools

### Using with Custom Scripts
```bash
#!/bin/bash
# Example integration script
ffuf-workflow -wordlist custom.txt -output scan.txt | other-tool
```

## Troubleshooting

### Common Issues

1. Dependencies not found:
```bash
# Verify dependencies
which ffuf
which ffufPostprocessing
```

2. Permission issues:
```bash
# Fix permissions
chmod +x $(which ffuf-workflow)
```

3. Go path issues:
```bash
# Add to ~/.bashrc
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Compatibility

- Go version: 1.21 or later
- Operating Systems: Linux, macOS, Windows
- Required tools: ffuf, ffufPostprocessing

## Support

- Issues: Report on [GitHub Issues](https://github.com/nullenc0de/FFUF-Workflow-Tool/issues)
- Questions: Open a discussion on GitHub
- Contributions: Pull requests welcome

## License

[License Type] - see LICENSE file for details

## Acknowledgments

- ffuf team for the excellent fuzzing tool
- dsecuredcom for ffufPostprocessing
- Contributors and community members
