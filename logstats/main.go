// Problem:
// given a directory, read all .log files in that directory.
// The log files are Cloudwatch JSON events.
// Step 1: return all the http methods and the total count (from all log files)
// of each method (e.g. GET)

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func locateFilesBySuffix(dir string, suffix string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", dir, err)
			return err
		}
		if !info.IsDir() && info.Name()[len(info.Name())-len(suffix):] == suffix {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func main() {

}
