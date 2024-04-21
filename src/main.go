package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func funcErr(err error) {
	if err != nil {
		panic(err)
	}
}

// get all drive in the pc
func getdrives() (r []string) {
	for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		f, err := os.Open(string(drive) + ":\\")
		
		if err == nil {
			r = append(r, string(drive) + ":\\")
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

	if len(os.Args) == 1 {
		usage()
		os.Exit(1)
	}

	elem, elem_type := os.Args[1], os.Args[2]
	fmt.Println("Ricerca...")

	switch elem_type {
		case "-f":
			if len(os.Args) < 4 {
				for _, drive := range drives {
					items, err := os.ReadDir(drive)
					funcErr(err)
					findFile(elem, items, drive)
				}
			} else {
				for _, drive := range drives {
					if strings.EqualFold(os.Args[3], drive) {
						items, err := os.ReadDir(drive)
						funcErr(err)
						findFile(elem, items, drive)
					}
				}
			}
			
		case "-d":
			if len(os.Args) < 4 {
				for _, drive := range drives {
					items, err := os.ReadDir(drive)
					funcErr(err)
					findDir(elem, items, drive)
				}
			} else {
				for _, drive := range drives {
					if strings.EqualFold(os.Args[3], drive) {
						items, err := os.ReadDir(drive)
						funcErr(err)
						findDir(elem, items, drive)
					}
				}
			}
			
		default:
			usage()
	}
	
	// finish := time.Since(start)
	// fmt.Println(finish)
}

