package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

func main() {

	pwd, _ := os.Getwd()
	f, err := ioutil.ReadDir(pwd)

	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(f); i++ {
		r, _ := ioutil.ReadFile(f[i].Name())
		exp, _ := regexp.Compile(`[a-zA-z]+://[^\s]*`)
		match := exp.FindAllString(string(r), -1)
		for z, _ := range match {
			a := match[z]
			fmt.Println(a)

		}
	}
}