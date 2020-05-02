package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const ApiKey = "some-cool-yandex-weather-api-key"
const ApiUrl = "https://api.weather.yandex.ru/v1/informers?lat=#LAT&lon=#LON&lang=ru_RU"

func main() {
	//title, lat, lon := getPlaceDetail("nartkala")

	//url := getApiUrl(lat, lon)
	json := getJson()
	unix := strconv.FormatInt(time.Now().Unix(), 10)
	saveJsonFile(json, "nartkala-original-"+unix+".json")
	bjson := []byte(json)
	bjson, _ = prettyPrint(bjson)
	savePrettyJsonFile(bjson, "nartkala-pretty-"+unix+".json")
}

type location struct {
	title string
	lat   string
	lon   string
}

func prettyPrint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}

func savePrettyJsonFile(content []byte, fileName string) {
	separator := string(filepath.Separator)
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	dir = "/home/styleroom/work/pogoda"
	hError(err)
	filePath := dir + separator + fileName

	f, e := os.Create(filePath)
	hError(e)
	defer f.Close()

	w := bufio.NewWriter(f)
	_, er := w.Write(content)
	hError(er)
	w.Flush()
	fmt.Println("filePath: ", filePath)
}

func saveJsonFile(content string, fileName string) {
	separator := string(filepath.Separator)
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	dir = "/home/styleroom/work/pogoda"
	hError(err)
	filePath := dir + separator + fileName

	f, e := os.Create(filePath)
	hError(e)
	defer f.Close()

	w := bufio.NewWriter(f)
	_, er := w.WriteString(content)
	hError(er)
	w.Flush()
	fmt.Println("filePath: ", filePath)
}

func getApiUrl(lat string, lon string) string {
	url := strings.ReplaceAll(ApiUrl, "#LAT", lat)
	return strings.ReplaceAll(url, "#LON", lon)
}

func getPlaceDetail(key string) (string, string, string) {
	places := map[string]location{
		"nartkala":    location{title: "Нарткала", lat: "43.5575", lon: "43.85083"},
		"makhachkala": location{title: "Махачкала", lat: "42.97638", lon: "47.50236"},
	}
	place := places[key]
	return place.title, place.lat, place.lon
}

func hError(err error) {
	if err != nil {
		panic(err)
	}
}

func getJson() string {
	return `{"now":1587718912,"now_dt":"2020-04-24T09:01:52.234Z","info":{"lat":43.5575,"lon":43.85083,"url":"https://yandex.ru/pogoda/?lat=43.5575&lon=43.85083"},"fact":{"temp":15,"feels_like":11,"icon":"bkn_d","condition":"cloudy","wind_speed":4,"wind_gust":7.3,"wind_dir":"e","pressure_mm":737,"pressure_pa":983,"humidity":30,"daytime":"d","polar":false,"season":"spring","obs_time":1587718404},"forecast":{"date":"2020-04-24","date_ts":1587675600,"week":17,"sunrise":"05:07","sunset":"18:58","moon_code":8,"moon_text":"new-moon","parts":[{"part_name":"evening","temp_min":6,"temp_max":15,"temp_avg":11,"feels_like":7,"icon":"skc_n","condition":"clear","daytime":"n","polar":false,"wind_speed":2.5,"wind_gust":7.1,"wind_dir":"ne","pressure_mm":736,"pressure_pa":982,"humidity":48,"prec_mm":0,"prec_period":360,"prec_prob":0},{"part_name":"night","temp_min":2,"temp_max":5,"temp_avg":4,"feels_like":1,"icon":"skc_n","condition":"clear","daytime":"n","polar":false,"wind_speed":1.7,"wind_gust":4.7,"wind_dir":"s","pressure_mm":736,"pressure_pa":982,"humidity":76,"prec_mm":0,"prec_period":360,"prec_prob":0}]}}`
}
