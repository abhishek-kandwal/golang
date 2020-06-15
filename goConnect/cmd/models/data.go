package models

// type Data struct {

// }

type Weather struct {
	Id1         int64
	Main        string
	Description string

	Temp     float64
	Pressure int64
	Humidity int64
	Temp_min float64
	Temp_max float64

	Speed float64
	Deg   int64

	All int64

	Id2     int64
	Message float64
	Country string
	Sunrise int64
	Sunset  int64

	Lon float64
	Lat float64

	Base       string
	Visibility int64

	Dt int64

	Id3  int64
	Name string
	Cod  int64
}
