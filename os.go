package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

// https://golang.hotexamples.com/examples/os/-/OpenFile/golang-openfile-function-examples.html

func main() {
	ready("wow-some.html")
}

func ready(filename string) {
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		bytes, err := ioutil.ReadFile(filename)
		bytesClean := CleanHtml(bytes)
		if err != nil {
			log.Fatal(err)
		}
		f, _ := os.OpenFile(filename, os.O_TRUNC|os.O_RDWR, 0644)
		if _, err := f.Write(bytesClean); err != nil {
			f.Close()
			log.Fatal(err)
		}
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}
}

func getNeedFile(filename string) {
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		fmt.Println("Good!")
		f, _ := os.OpenFile(filename, os.O_TRUNC|os.O_RDWR, 0644)
		if _, err := f.Write([]byte("appended some data\n")); err != nil {
			f.Close()
			log.Fatal(err)
		}
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}
}

func CleanHtml(content []byte) []byte {
	cleanComments := regexp.MustCompile(`<!--[^>]*-->`)
	cleanLineSeparator := regexp.MustCompile(`\n`)
	cleanSpaces1 := regexp.MustCompile(`\s{2,}`)
	cleanSpaces2 := regexp.MustCompile(`>\s{1,}<`)
	noComment := cleanComments.ReplaceAll(content, []byte(""))
	noSeparators := cleanLineSeparator.ReplaceAll(noComment, []byte(""))
	noSpaces1 := cleanSpaces1.ReplaceAll(noSeparators, []byte(" "))
	return cleanSpaces2.ReplaceAll(noSpaces1, []byte("><"))
}

func CleanJsonBrackets(html string) string {
	htmlClear := strings.ReplaceAll(html, " } <", "}<")
	htmlClear = strings.ReplaceAll(htmlClear, " { ", "{")
	return htmlClear
}

func readHtmlFile(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
