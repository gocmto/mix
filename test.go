package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type Root struct {
	XMLName xml.Name `xml:"root"`
	Config  Config   `xml:"config"`
}

type Config struct {
	XMLName xml.Name `xml:"config"`
	Path    string   `xml:"path"`
}

func main() {
	xmlFile, err := os.Open("test_config.xml")
	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var root Root
	xml.Unmarshal(byteValue, &root)

	path := root.Config.Path
	fmt.Println("path: ", path)

	readAllDir(path)

	targetFile := filepath.Join(path, "test_result.txt")

	time := time.Now()
	formatted := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		time.Year(), time.Month(), time.Day(),
		time.Hour(), time.Minute(), time.Second())

	data := []byte("This content created at: " + formatted)
	e := ioutil.WriteFile(targetFile, data, 0644)
	if e != nil {
		fmt.Println(err)
	}
}

func readAllDir(path string) {
	_, er := os.Stat(path)
	if er != nil {
		fmt.Println(er)
	}

	c, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
	}

	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}
}

/*
<root>
    <config>
        <path>C:\Users\79161\go\src\pogoda\</path>
    </config>
</root>
*/
