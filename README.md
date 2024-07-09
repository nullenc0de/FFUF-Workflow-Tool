# FFUF Workflow Tool

## Overview
This tool automates the process of running [FFUF](https://github.com/ffuf/ffuf) (Fuzz Faster U Fool) and post-processing its results to extract valid URLs. It supports both direct file input and piped input for wordlists.

## Features
- **FFUF Execution**: Runs FFUF with specified parameters, including a wordlist of URLs.
- **Post-processing**: Uses a custom post-processing tool (`ffufPostprocessing`) to handle FFUF output for further analysis.
- **URL Extraction**: Extracts valid URLs from FFUF's JSON output.
- **Output Handling**: Optionally saves extracted URLs to a specified output file.

## Usage
```bash
# Direct file input
./ffuf-workflow -wordlist wordlist.txt -output valid_urls.txt

# Piped input
cat wordlist.txt | ./ffuf-workflow -output valid_urls.txt

## Parameters
- **wordlist**: Path to the wordlist file containing URLs.
- **output**: (Optional) Path to save the extracted valid URLs.

## Dependencies
- **FFUF**: Fuzzing tool used for URL fuzzing and testing.
- **ffufPostprocessing**: Custom tool for post-processing FFUF results.

## Notes
- Ensure FFUF and `ffufPostprocessing` are installed and accessible in your PATH.
- The tool handles both direct file input and piped input for convenience.

## Example Workflow
1. **Run FFUF**: Executes FFUF with the specified wordlist and saves results.
2. **Post-process**: Uses `ffufPostprocessing` to handle FFUF results for extraction.
3. **Extract URLs**: Retrieves valid URLs from the processed output.
4. **Output**: Optionally saves extracted URLs to a specified output file for further use.

## License
This tool is licensed under the MIT License. See LICENSE for more details.
