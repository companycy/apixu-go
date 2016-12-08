package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	ApiKey = "a85d495390c24c5cbaf15016160812" // api key from apixu
)

// refer to apixu-nodejs and below link for API info:
// https://www.apixu.com/doc/request.aspx
const (
	ApixuHost       = "http://api.apixu.com/"
	ApixuCurVersion = "v1"
	ApixuPath       = ApixuCurVersion + "/"

	CurPath     = ApixuHost + ApixuPath + "/current.json"
	ForcastPath = ApixuHost + ApixuPath + "/forecast.json"
	SearchPath  = ApixuHost + ApixuPath + "/search.json"
	HistoryPath = ApixuHost + ApixuPath + "/history.json"
)

// https://www.apixu.com/doc/location.aspx
type TLocation struct {
	Lat, Lon                     float32
	Name, Region, Country, Tz_id string
	Localtime_epoch              int
	Localtime                    string
}

// https://www.apixu.com/doc/current.aspx
type TCondition struct {
	Text, Icon string
	Code       int
}

type TCurWeather struct {
	Last_updated                             string
	Last_updated_epoch                       int
	Temp_c, Temp_f, Feelslike_c, Feelslike_f float32

	Condition                                      TCondition
	Wind_mph, Wind_kph                             float32
	Wind_degree                                    int
	Wind_dir                                       string
	Pressure_mb, Pressure_in, Precip_mm, Precip_in float32
	Humidity, Cloud                                int
	Is_day                                         int
}

type CurWeatherInfo struct {
	// common location part
	Location TLocation

	// current weather info
	Current TCurWeather
}

// TODO: similar to current weather info
type ForcastWeatherInfo struct {
}

func main() {
	cur := CurWeatherInfo{}
	// http://api.apixu.com/v1/current.json?key=a85d495390c24c5cbaf15016160812&q=Paris
	city := "Paris"
	url := fmt.Sprintf("%s?key=%s&q=%s", CurPath, ApiKey, city)
	err := getJson(url, &cur)
	if err != nil {
		fmt.Println("Failed to get current weather %s", err)
	}

	fmt.Println(cur)
}

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to get url %s", url)
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}
