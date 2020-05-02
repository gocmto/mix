package mypkg

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const ApiKey = "some-api-key"
const ApiUrl = "https://api.weather.yandex.ru/v1/informers?lat=#LAT&lon=#LON&lang=ru_RU"
const ApiUrlShort = "https://api.weather.yandex.ru/v1/informers"

// Describe some location
type location struct {
	title string
	lat   string
	lon   string
}

// Save API Yandex Weather response to a json-file
func SaveJsonFile(content string, fileName string) {
	separator := string(filepath.Separator)
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	handleError(err)
	filePath := dir + separator + fileName

	f, err := os.Create(filePath)
	defer f.Close()

	w := bufio.NewWriter(f)
	_, er := w.WriteString(content)
	handleError(er)
	w.Flush()
	fmt.Println("filePath: ", filePath)
}

// Get place/location details
func GetPlaceDetail(key string) (string, string, string) {
	places := map[string]location{
		"nartkala":    location{title: "Нарткала", lat: "43.5575", lon: "43.85083"},
		"makhachkala": location{title: "Махачкала", lat: "42.97638", lon: "47.50236"},
	}
	place := places[key]
	return place.title, place.lat, place.lon
}

// Get Yandex Weather API URL
func GetApiUrl(lat string, lon string) string {
	url := strings.ReplaceAll(ApiUrl, "#LAT", lat)
	return strings.ReplaceAll(url, "#LON", lon)
}

// Get Yandex Weather API URL from float lat/lon values
func GetApiUrlFloat(lat float64, lon float64) string {
	stringLat, stringLon := GetLatLot(lat, lon)
	url := strings.ReplaceAll(ApiUrl, "#LAT", stringLat)
	return strings.ReplaceAll(url, "#LON", stringLon)
}

// Get latitude and longitude as string
func GetLatLot(lat float64, lon float64) (string, string) {
	strLat := strconv.FormatFloat(lat, 'f', 6, 64)
	strLon := strconv.FormatFloat(lon, 'f', 6, 64)
	return strLat, strLon
}

// Send request and get json response
func SendRequest(url string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	handleError(err)
	req.Header.Add("X-Yandex-API-Key", ApiKey)
	return getResponse(client, req)
}

// Get string response from client request
func getResponse(client *http.Client, req *http.Request) string {
	resp, err := client.Do(req)
	handleError(err)
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	handleError(err)
	return string(html)
}

// Error handler
func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
