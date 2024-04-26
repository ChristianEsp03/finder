package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// get all drive in the pc
func getdrives() (r []string) {
	for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		f, err := os.Open(string(drive) + ":\\")

		if err == nil {
			r = append(r, string(drive)+":\\")
		}

		defer f.Close()
	}
	return
}

func usage() {
	fmt.Print(
		`usage: find [<name>] [<type>] [<optional>]

in the 'type' field insert:
  -f	to find a file
  -d	to find a directory

in the 'optional' field you can put the name of single drive

`)
}

func findFile(elem string, items []fs.DirEntry, path string) {
	for _, obj := range items {
		fullpath := filepath.Join(path, obj.Name())

		if !obj.IsDir() {
			if strings.Contains(obj.Name(), elem) {
				fmt.Println(fullpath)
			}
		} else {
			newItems, err := os.ReadDir(fullpath)

			if err != nil {
				continue
			}

			findFile(elem, newItems, fullpath)
		}
	}
}

func findDir(elem string, items []fs.DirEntry, path string) {
	for _, obj := range items {
		fullpath := filepath.Join(path, obj.Name())

		if obj.IsDir() {
			if strings.Contains(obj.Name(), elem) {
				fmt.Println(fullpath)
			}

			newItems, err := os.ReadDir(fullpath)

			if err != nil {
				continue
			}

			findDir(elem, newItems, fullpath)
		}
	}
}

func main() {
	// start := time.Now()
	drives := getdrives()
	argLen := len(os.Args)

	if argLen == 1 {
		usage()
		os.Exit(1)
	} else if argLen == 2 {
		usage()
		os.Exit(1)
	} else {
		elem, elem_type := os.Args[1], os.Args[2]
		fmt.Println("Searching...")

		switch elem_type {
		case "-f":
			if argLen < 4 {
				for _, drive := range drives {
					items, err := os.ReadDir(drive)

					if err != nil {
						fmt.Println("Error searching file")
					}

					findFile(elem, items, drive)
				}
			} else {
				for _, drive := range drives {
					if strings.EqualFold(os.Args[3], drive) {
						items, err := os.ReadDir(drive)

						if err != nil {
							fmt.Println("Error searching file")
						}

						findFile(elem, items, drive)
					}
				}
			}

		case "-d":
			if argLen < 4 {
				for _, drive := range drives {
					items, err := os.ReadDir(drive)

					if err != nil {
						fmt.Println("Error searching directory")
					}

					findDir(elem, items, drive)
				}
			} else {
				for _, drive := range drives {
					if strings.EqualFold(os.Args[3], drive) {
						items, err := os.ReadDir(drive)

						if err != nil {
							fmt.Println("Error searching directory")
						}

						findDir(elem, items, drive)
					}
				}
			}

		default:
			usage()
		}
	}

	// finish := time.Since(start)
	// fmt.Println(finish)
}
