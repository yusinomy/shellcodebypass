package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

var a = []string{"c", "d", "e", "g", "h", "i"}
var file = []string{".txt", ".xml", ".mdb", ".sql", ".mdf", ".eml", ".pst", "conf", "bak", "pwd", "pass", "login", "user"}
var q string
var k string

func main() {
	for i := 0; i < len(a); i++ {
		q = a[i]
		for pass := 0; pass < len(file); pass++ {
			k = file[pass]
			fmt.Println(k)
			filepath.Walk(q+":\\", func(path string, info fs.FileInfo, err error) error {
				if strings.HasSuffix(path, k) {
					fmt.Println(path)
				}
				return nil
			})
		}
	}
}
