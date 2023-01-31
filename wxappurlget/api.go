package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

func main() {
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		time.Sleep(1)
		r, _ := ioutil.ReadFile(path)
		exp, _ := regexp.Compile(`http[s]{0,1}://(([a-zA-Z0-9\._-]+\.[a-zA-Z]{2,6})|([0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}))(:[0-9]{1,4})*(/[a-zA-Z0-9\&%_\./-~-]*)?`)
		//exp, _ := regexp.Compile(`"[.][.]/[.][.]/(\w+\W+){4}\w+"`) ../../
		//exp, _ := regexp.Compile("[.]/(\\w+\\W+){3}\\w+")  ././
		//exp, _ := regexp.Compile("/(\\w+\\W+){2}\\w+") /x/x/
		match := exp.FindAllString(string(r), -1)
		for z, _ := range match {
			a := match[z]
			fmt.Println(a)
		}
		return nil
	})
	if err != nil {
		return
	}
}
