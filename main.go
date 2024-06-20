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
package main

import (
	"flag"
	"fmt"
	"pathInspector/src/fileutil"
	"time"
)

func main() {
	fileutil.Title()

	t0 := time.Now()
	args := parseARG()

	if args.Interactive {
		args = interMode()
	}
	fmt.Println("Option :", args)

	// start a new goroutine that runs the spinner function
	// Create a channel called stop
	stop := make(chan struct{})
	go fileutil.Spinner(stop) // enable spinner
	fileutil.ParseDir(args)   // start dir analysis
	close(stop)               // closing the channel stop the goroutine

	fmt.Println("\ndone !")
	fmt.Println("Results are in results/output.tsv !")
	fmt.Println("Elapsed time : ", time.Since(t0))

	// Define the duration for the countdown
	// before closing the window
	// Set the countdown time in seconds
	countdownFrom := 3
	fileutil.Timer(countdownFrom)
}

// parse arg of the command line and return the argument struct
func parseARG() fileutil.Args {
	args := fileutil.Args{}
	flag.IntVar(&args.Length, "l", 255, "maximal path length")
	flag.StringVar(&args.Dir, "d", ".", "path")
	flag.BoolVar(&args.Interactive, "i", true, "Interactive mode. Important syntax is -i=false")
	flag.BoolVar(&args.EnablePrinting, "p", true, "Enable printing path to terminal. Important syntax is -p=false")
	flag.Parse()
	return args
}

func interMode() fileutil.Args {
	fmt.Println("Enter the maximal path size")
	var size int
	_, err := fmt.Scan(&size)
	if err != nil {
		fmt.Println("Error reading size:", err)
		return fileutil.Args{Length: 255, Dir: "."}
	}

	fmt.Println("Enter the path")
	var path string
	_, err = fmt.Scan(&path)
	if err != nil {
		fmt.Println("Error reading path:", err)
		return fileutil.Args{Length: size, Dir: "."}
	}

	fmt.Println("Enable printing path in terminal ? (y/n)")
	var printing string
	enablePrinting := true
	_, err = fmt.Scan(&printing)
	if err != nil {
		fmt.Println("Error reading path:", err)
		return fileutil.Args{Length: size, Dir: ".", EnablePrinting: true}
	}
	if printing != "y" {
		enablePrinting = false
	}

	return fileutil.Args{Length: size, Dir: path, EnablePrinting: enablePrinting}
}
