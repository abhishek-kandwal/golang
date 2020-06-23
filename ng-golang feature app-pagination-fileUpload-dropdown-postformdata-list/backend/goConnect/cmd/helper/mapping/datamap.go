package mapping

import (
	"goConnect/cmd/models"
	// "strconv"
)

func DataMap(data map[string]interface{}) (models.Weather, error) {
	var err error
	newModel := models.Weather{}
	// newModel.City, _ = strconv.ParseFloat(fmt.Sprintf("%v", data[""].(map[string]interface{})["lat"]), 32)
	// newModel.City, _ = strconv.ParseFloat(fmt.Sprintf("%v", data["coord"].(map[string]interface{})["lat"]), 32)
	// newModel.City, _ = strconv.ParseFloat(fmt.Sprintf("%v", data["coord"].(map[string]interface{})["lat"]), 32)
	// newModel.City, _ = strconv.ParseFloat(fmt.Sprintf("%v", data["coord"].(map[string]interface{})["lat"]), 32)
	// // newModel.Lat, _ = strconv.ParseFloat(fmt.Sprintf("%v", data["coord"].(map[string]interface{})["lat"]), 32)
	// newModel.Lon, _ = strconv.ParseFloat(fmt.Sprintf("%v", data["coord"].(map[string]interface{})["lon"]), 32)
	// newModel.Temp_max, _ = strconv.ParseFloat(fmt.Sprintf("%v", data["main"].(map[string]interface{})["temp_max"]), 32)
	// newModel.Temp_min, _ = strconv.ParseFloat(fmt.Sprintf("%v", data["main"].(map[string]interface{})["temp_min"]), 32)
	// newModel.Speed, _ = strconv.ParseFloat(fmt.Sprintf("%v", data["coord"].(map[string]interface{})["speed"]), 32)
	// newModel.Message, _ = strconv.ParseFloat(fmt.Sprintf("%v", data["wind"].(map[string]interface{})["lon"]), 32)
	// newModel.Temp, _ = strconv.ParseFloat(fmt.Sprintf("%v", data["main"].(map[string]interface{})["temp"]), 32)

	// newModel.Humidity, _ = strconv.ParseInt(fmt.Sprintf("%v", data["main"].(map[string]interface{})["humidity"]), 10, 64)
	// newModel.Id1, _ = strconv.ParseInt(fmt.Sprintf("%v", data["sys"].(map[string]interface{})["id"]), 10, 64)
	// newModel.Pressure, _ = strconv.ParseInt(fmt.Sprintf("%v", data["main"].(map[string]interface{})["pressure"]), 10, 64)
	// newModel.Deg, _ = strconv.ParseInt(fmt.Sprintf("%v", data["wind"].(map[string]interface{})["deg"]), 10, 64)
	// newModel.All, _ = strconv.ParseInt(fmt.Sprintf("%v", data["clouds"].(map[string]interface{})["all"]), 10, 64)
	// //newModel.Id2, _ = strconv.ParseInt(fmt.Sprintf("%v", data["weather"].(map[string]interface{})["id"]), 10, 64)
	// newModel.Sunrise, _ = strconv.ParseInt(fmt.Sprintf("%v", data["sys"].(map[string]interface{})["sunrise"]), 10, 64)
	// newModel.Sunset, _ = strconv.ParseInt(fmt.Sprintf("%v", data["sys"].(map[string]interface{})["sunset"]), 10, 64)
	// newModel.Visibility, _ = strconv.ParseInt(fmt.Sprintf("%v", data["main"].(map[string]interface{})["visibility"]), 10, 64)
	// newModel.Id3, _ = strconv.ParseInt(fmt.Sprintf("%v", data["id"]), 10, 64)
	// newModel.Cod, _ = strconv.ParseInt(fmt.Sprintf("%v", data["cod"]), 10, 64)
	// newModel.Dt, _ = strconv.ParseInt(fmt.Sprintf("%v", data["dt"]), 10, 64)

	// newModel.Country = fmt.Sprintf("%v", data["sys"].(map[string]interface{})["country"])
	// //newModel.Main = fmt.Sprintf("%v", data["weather"].(map[string]interface{})["main"])
	// //newModel.Description = fmt.Sprintf("%v", data["weather"].(map[string]interface{})["description"])
	// newModel.Name = fmt.Sprintf("%v", data["name"])
	// newModel.Base = fmt.Sprintf("%v", data["base"])
	return newModel, err
}
