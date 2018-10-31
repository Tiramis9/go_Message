package config

import "time"

var (
	Path     = string("./log/")
	File     = string("std.log")
	MaxAge   = time.Hour * 24
	Interval = time.Second * 2000
)
var (
	ADDRPORT = "192.168.197.1:9999"
	//ADDRPORT1 = "172.16.101.8:9999"
)

var (
	MmdbFile = "./GeoIP2_data/GeoLite2-City_20181023/GeoLite2-City.mmdb"
)
var (
	CN          = "zh-CN"
	US          = "en"
	Japanese    = "ja"
	Korea       = "ru"
	Filipino    = "fr"
	Brazil      = "pt-BR"
	Spain       = "es"
	Deutschland = "de"
)
