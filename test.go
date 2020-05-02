package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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
		writeLog(err)
	}

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var root Root
	xml.Unmarshal(byteValue, &root)

	path := root.Config.Path
	fmt.Println("path: ", path)

	_, er := os.Stat(path)
	if er != nil {
		writeLog(er)
	}

	readAllDir(path)

	targetFile := filepath.Join(path, "test_result.txt")

	formatted := getNowDate()

	data := []byte("This content created at: " + formatted)
	e := ioutil.WriteFile(targetFile, data, 0644)
	if e != nil {
		writeLog(e)
	}
}

func getNowDate() string {
	time := time.Now()
	formatted := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		time.Year(), time.Month(), time.Day(),
		time.Hour(), time.Minute(), time.Second())

	return formatted
}

func writeLog(err error)  {
	e := err.Error()
	filename := os.Args
	fn := filepath.Base(filename[0])
	fnm := strings.Replace(fn, ".exe", "",4) + ".log"
	f, _ := os.OpenFile(fnm, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	output := fmt.Sprintf("%s: %s \n", getNowDate(), e)
	f.WriteString(output)
}

func readAllDir(path string) {
	c, err := ioutil.ReadDir(path)
	if err != nil {
		writeLog(err)
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
