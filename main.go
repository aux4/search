package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"bufio"
	"io/fs"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: aux4-search <files|content> <args...>")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "files":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "Usage: aux4-search files <pattern> [searchPath] [exclude] [maxDepth]")
			os.Exit(1)
		}
		pattern := os.Args[2]
		searchPath := "."
		exclude := "node_modules,.git"
		maxDepth := -1

		if len(os.Args) > 3 && os.Args[3] != "" {
			searchPath = os.Args[3]
		}
		if len(os.Args) > 4 && os.Args[4] != "" {
			exclude = os.Args[4]
		}
		if len(os.Args) > 5 && os.Args[5] != "" {
			d, err := strconv.Atoi(os.Args[5])
			if err == nil {
				maxDepth = d
			}
		}

		searchFiles(pattern, searchPath, exclude, maxDepth)

	case "content":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "Usage: aux4-search content <pattern> [searchPath] [include] [exclude] [limit] [context]")
			os.Exit(1)
		}
		pattern := os.Args[2]
		searchPath := "."
		include := ""
		exclude := "node_modules,.git"
		limit := 50
		context := 0

		if len(os.Args) > 3 && os.Args[3] != "" {
			searchPath = os.Args[3]
		}
		if len(os.Args) > 4 && os.Args[4] != "" {
			include = os.Args[4]
		}
		if len(os.Args) > 5 && os.Args[5] != "" {
			exclude = os.Args[5]
		}
		if len(os.Args) > 6 && os.Args[6] != "" {
			l, err := strconv.Atoi(os.Args[6])
			if err == nil {
				limit = l
			}
		}
		if len(os.Args) > 7 && os.Args[7] != "" {
			c, err := strconv.Atoi(os.Args[7])
			if err == nil {
				context = c
			}
		}

		searchContent(pattern, searchPath, include, exclude, limit, context)

	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}
}

func searchFiles(pattern, searchPath, exclude string, maxDepth int) {
	excludeDirs := parseCSV(exclude)
	absRoot, err := filepath.Abs(searchPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Check if pattern contains glob characters
	isGlob := strings.ContainsAny(pattern, "*?[{")

	filepath.WalkDir(absRoot, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		rel, _ := filepath.Rel(absRoot, path)
		if rel == "." {
			return nil
		}

		// Check depth
		if maxDepth >= 0 {
			depth := strings.Count(rel, string(filepath.Separator))
			if depth > maxDepth {
				if d.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}
		}

		// Skip excluded directories
		if d.IsDir() {
			base := d.Name()
			for _, ex := range excludeDirs {
				if base == ex {
					return filepath.SkipDir
				}
			}
			return nil
		}

		// Match against pattern
		if isGlob {
			matched, _ := filepath.Match(pattern, d.Name())
			if !matched {
				// Try matching against full relative path for ** patterns
				matched = matchDoubleGlob(pattern, rel)
			}
			if matched {
				fmt.Println(rel)
			}
		} else {
			// Plain text: substring match on filename
			if strings.Contains(strings.ToLower(d.Name()), strings.ToLower(pattern)) {
				fmt.Println(rel)
			}
		}

		return nil
	})
}

func matchDoubleGlob(pattern, path string) bool {
	// Handle ** patterns by trying each path segment combination
	if !strings.Contains(pattern, "**") {
		matched, _ := filepath.Match(pattern, path)
		return matched
	}

	parts := strings.SplitN(pattern, "**", 2)
	prefix := parts[0]
	suffix := parts[1]

	if prefix != "" {
		prefix = strings.TrimSuffix(prefix, "/")
		if !strings.HasPrefix(path, prefix) {
			return false
		}
		path = strings.TrimPrefix(path, prefix)
		path = strings.TrimPrefix(path, "/")
	}

	if suffix == "" {
		return true
	}

	suffix = strings.TrimPrefix(suffix, "/")

	// Try matching suffix against every possible tail of the path
	pathParts := strings.Split(path, "/")
	for i := range pathParts {
		candidate := strings.Join(pathParts[i:], "/")
		matched, _ := filepath.Match(suffix, candidate)
		if matched {
			return true
		}
		// Also try matching just the filename part
		if i == len(pathParts)-1 {
			matched, _ = filepath.Match(suffix, pathParts[i])
			if matched {
				return true
			}
		}
	}

	return false
}

func searchContent(pattern, searchPath, include, exclude string, limit, contextLines int) {
	excludeDirs := parseCSV(exclude)
	absRoot, err := filepath.Abs(searchPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	re, err := regexp.Compile("(?i)" + pattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid regex pattern: %v\n", err)
		os.Exit(1)
	}

	var includeGlob string
	if include != "" {
		includeGlob = include
	}

	matchCount := 0
	prevFile := ""

	filepath.WalkDir(absRoot, func(path string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return nil
		}

		if matchCount >= limit {
			return filepath.SkipAll
		}

		if d.IsDir() {
			base := d.Name()
			for _, ex := range excludeDirs {
				if base == ex {
					return filepath.SkipDir
				}
			}
			return nil
		}

		// Skip binary files based on extension
		if isBinaryExtension(d.Name()) {
			return nil
		}

		rel, _ := filepath.Rel(absRoot, path)

		// Apply include filter
		if includeGlob != "" {
			matched, _ := filepath.Match(includeGlob, d.Name())
			if !matched {
				return nil
			}
		}

		file, err := os.Open(path)
		if err != nil {
			return nil
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Buffer(make([]byte, 1024*1024), 1024*1024) // 1MB line buffer

		var lines []string
		lineNumbers := []int{}
		matchIndices := []int{}

		lineNum := 0
		for scanner.Scan() {
			lineNum++
			line := scanner.Text()
			lines = append(lines, line)

			if re.MatchString(line) {
				matchIndices = append(matchIndices, len(lines)-1)
				lineNumbers = append(lineNumbers, lineNum)
			}
		}

		for i, idx := range matchIndices {
			if matchCount >= limit {
				break
			}

			if contextLines > 0 {
				// Print separator between files or between non-contiguous matches
				if prevFile != rel || (i > 0 && idx-matchIndices[i-1] > contextLines*2+1) {
					if prevFile != "" {
						fmt.Println("--")
					}
				}
				prevFile = rel

				start := idx - contextLines
				if start < 0 {
					start = 0
				}
				end := idx + contextLines
				if end >= len(lines) {
					end = len(lines) - 1
				}

				for j := start; j <= end; j++ {
					actualLine := lineNumbers[i] - (idx - j)
					if j == idx {
						fmt.Printf("%s:%d: %s\n", rel, actualLine, lines[j])
					} else {
						fmt.Printf("%s-%d- %s\n", rel, actualLine, lines[j])
					}
				}
			} else {
				fmt.Printf("%s:%d: %s\n", rel, lineNumbers[i], lines[idx])
			}

			matchCount++
		}

		return nil
	})

	if matchCount >= limit {
		fmt.Fprintf(os.Stderr, "[Results limited to %d matches]\n", limit)
	}
}

func parseCSV(s string) []string {
	if s == "" {
		return nil
	}
	parts := strings.Split(s, ",")
	result := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			result = append(result, p)
		}
	}
	return result
}

func isBinaryExtension(name string) bool {
	ext := strings.ToLower(filepath.Ext(name))
	binaryExts := map[string]bool{
		".exe": true, ".dll": true, ".so": true, ".dylib": true,
		".bin": true, ".o": true, ".a": true,
		".zip": true, ".tar": true, ".gz": true, ".bz2": true, ".xz": true, ".7z": true,
		".png": true, ".jpg": true, ".jpeg": true, ".gif": true, ".bmp": true, ".ico": true, ".webp": true,
		".mp3": true, ".mp4": true, ".avi": true, ".mov": true, ".wav": true, ".flac": true,
		".pdf": true, ".doc": true, ".docx": true, ".xls": true, ".xlsx": true,
		".wasm": true, ".class": true, ".pyc": true,
	}
	return binaryExts[ext]
}
