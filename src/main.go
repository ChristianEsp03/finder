package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"
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
	start := time.Now()
	drives := getdrives()
	elem, elem_type := os.Args[1], os.Args[2]
	
	fmt.Println("Ricerca...")

	switch elem_type {
		case "-f":
			for _, drive := range drives {
				items, err := os.ReadDir(drive)
				funcErr(err)
				findFile(elem, items, drive)
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
			fmt.Println("The element type isn't correct") // Modificare la stringa
	}

	finish := time.Since(start)

	fmt.Println(finish)
}

