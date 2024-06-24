/*
 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU General Public License as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.

 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU General Public License for more details.

 You should have received a copy of the GNU General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.

 Written by Frederic PONT.
 (c) Frederic Pont 2024
*/

package fileutil

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

type Args struct {
	Length         int
	Interactive    bool
	Dir            string
	EnablePrinting bool
	Exclude        []string
}

func ParseDir(args Args) {
	dir := args.Dir
	excludedDirs := args.Exclude // Directories to exclude
	pathCounter := 0
	// Create a file for writing
	outfile, err := os.Create("results/output.tsv")
	if err != nil {
		log.Println(err)
	}
	defer outfile.Close()

	// Create a CSV writer with a tab delimiter for TSV format
	writer := csv.NewWriter(outfile)
	writer.Comma = '\t' // Set the delimiter to tab

	err = filepath.WalkDir(dir, func(path string, info os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && contains(excludedDirs, info.Name()) {
			return filepath.SkipDir // Skip directory if it's in the excluded list
		}

		// Process the file or directory here
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		pathLength := len(path)

		if pathLength >= args.Length {
			pathCounter++
			writeLine(writer, []string{path, strconv.Itoa(pathLength)})
			if args.EnablePrinting {
				fmt.Println(pathCounter, "- ", path, pathLength)
			}

		}

		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
	}
	// Flush writes any buffered data to the underlying io.Writer
	writer.Flush()

	// Check if there have been any errors during Write or Flush
	if err := writer.Error(); err != nil {
		fmt.Println(err) // Handle errors after flushing
	}
	fmt.Println("\n\n──────────────────────────────────────────────────")
	fmt.Println(pathCounter, " path are above the limit of ", args.Length, " char")
}

func writeLine(writer *csv.Writer, data []string) {
	// Write the []string as a row to the file
	err := writer.Write(data)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func contains(arr []string, str string) bool {
	for _, s := range arr {
		if s == str {
			return true
		}
	}
	return false
}
