package main

import (
	"fmt"
	"pogoda/mypkg"
)

func main() {
	fmt.Println("we", mypkg.ApiKey)
	title, lat, lon := mypkg.GetPlaceDetail("nartkala")

	url := mypkg.GetApiUrl(lat, lon)
	json := mypkg.SendRequest(url)
	mypkg.SaveJsonFile(json, "nartkala-cool.json")

	fmt.Println("url: ", url)
	fmt.Println("nartkala: title: ", title, ", latitude: ", lat, ", longitude: ", lon)
	fmt.Println("json: ", json)
}
