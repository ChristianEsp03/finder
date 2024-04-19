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

func searchFile(elem string, items []fs.DirEntry, path string) {
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
				
				searchFile(elem, newItems, fullpath)			
		}
	}
}

// func searchDir() {}

func main() {
	drives := getdrives()
	elem, elem_type := os.Args[1], os.Args[2]
	fmt.Println("Ricerca...")

	switch elem_type {
		case "-f":
			for _, drive := range drives {
				items, err := os.ReadDir(drive)
				funcErr(err)
				searchFile(elem, items, drive)
			}
		case "-d":
			
		default:
			fmt.Println("The element type isn't correct") // Modificare la stringa
	}
}

