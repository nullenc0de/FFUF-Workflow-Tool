package main

import (
    "bufio"
    "encoding/json"
    "flag"
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
)

type FfufResult struct {
    Results []struct {
        URL string `json:"url"`
    } `json:"results"`
}

func runFfuf(wordlist, outputDir string) error {
    cmd := exec.Command("ffuf",
        "-w", wordlist,
        "-u", "FUZZ",
        "-o", filepath.Join(outputDir, "results.json"),
        "-od", filepath.Join(outputDir, "bodies"),
        "-of", "json")
    return cmd.Run()
}

func runFfufPostprocessing(outputDir string) error {
    cmd := exec.Command("ffufPostprocessing",
        "-result-file", filepath.Join(outputDir, "results.json"),
        "-bodies-folder", filepath.Join(outputDir, "bodies"),
        "-delete-bodies",
        "-overwrite-result-file")
    return cmd.Run()
}

func extractUrls(resultsFile string) ([]string, error) {
    data, err := os.ReadFile(resultsFile)
    if err != nil {
        return nil, err
    }
    var result FfufResult
    if err := json.Unmarshal(data, &result); err != nil {
        return nil, err
    }
    var urls []string
    for _, r := range result.Results {
        urls = append(urls, r.URL)
    }
    return urls, nil
}

func main() {
    var wordlist string
    var output string
    flag.StringVar(&wordlist, "wordlist", "", "Path to the wordlist file containing URLs")
    flag.StringVar(&output, "output", "", "Path to save the extracted URLs (optional)")
    flag.Parse()

    // Check if input is being piped
    stat, _ := os.Stdin.Stat()
    if (stat.Mode() & os.ModeCharDevice) == 0 {
        // Input is being piped
        scanner := bufio.NewScanner(os.Stdin)
        tempFile, err := os.CreateTemp("", "piped-wordlist")
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error creating temp file: %v\n", err)
            os.Exit(1)
        }
        defer os.Remove(tempFile.Name())
        for scanner.Scan() {
            tempFile.WriteString(scanner.Text() + "\n")
        }
        tempFile.Close()
        wordlist = tempFile.Name()
    } else if wordlist == "" {
        fmt.Println("Please provide a wordlist file or pipe input")
        flag.Usage()
        os.Exit(1)
    }

    tempDir, err := os.MkdirTemp("", "ffuf-workflow")
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error creating temp directory: %v\n", err)
        os.Exit(1)
    }
    defer os.RemoveAll(tempDir)

    fmt.Println("Running ffuf...")
    if err := runFfuf(wordlist, tempDir); err != nil {
        fmt.Fprintf(os.Stderr, "Error running ffuf: %v\n", err)
        os.Exit(1)
    }

    fmt.Println("Running ffuf post-processing...")
    if err := runFfufPostprocessing(tempDir); err != nil {
        fmt.Fprintf(os.Stderr, "Error running ffuf post-processing: %v\n", err)
        os.Exit(1)
    }

    fmt.Println("Extracting valid URLs...")
    urls, err := extractUrls(filepath.Join(tempDir, "results.json"))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error extracting URLs: %v\n", err)
        os.Exit(1)
    }

    if output != "" {
        file, err := os.Create(output)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
            os.Exit(1)
        }
        defer file.Close()
        for _, url := range urls {
            fmt.Fprintln(file, url)
        }
        fmt.Printf("Valid URLs saved to: %s\n", output)
    } else {
        fmt.Println("Valid URLs:")
        for _, url := range urls {
            fmt.Println(url)
        }
    }
}
