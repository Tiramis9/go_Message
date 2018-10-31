package server

import (
	"errors"
	"fmt"
	"net"
	"net/http"

	"github.com/oschwald/geoip2-golang"
	"public.info/config"
)

// 查询IP库返回详细信息
func GeoliteLookup(IP string) (*geoip2.City, error) {
	log := config.Log.WithField("package", "server")
	log.Info("start running GeoliteLookup ")
	log.Info(IP)
	db, err := geoip2.Open(config.MmdbFile)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Open File failed..")
	}
	defer db.Close()
	ip := net.ParseIP(IP)
	record, err := db.City(ip)
	if err != nil {
		log.Println(err)
		return record, errors.New("查询的IP信息不存在")
	}
	if record.City.GeoNameID == 0 {
		return record, errors.New("查询的IP信息为空")
	}
	//log.Println(record)
	log.Info("GeoliteLookup running over")
	return record, nil
}

//获取客户端的ip，给当前客户端返回IP详细信息
func DealQuerycmd(conn http.ResponseWriter, connMsg string) {
	log := config.Log.WithField("package", "server")
	log.Info("start running dealQuerycmd ")
	log.Info(connMsg)
	IP := MustIP(connMsg)
	if IP == "" {
		fmt.Println("IP is nil", IP)
		fmt.Fprintln(conn, "输入的IP有误，请重新输入!")
		return
	}
	ip := StringTOJson(IP)
	fmt.Fprintln(conn, "查询的IP:", string(ip))
	City, err := GeoliteLookup(IP)
	if err != nil {
		log.Println(err)
		fmt.Fprintln(conn, err)
		return
	}
	switch {
	case City.Country.IsoCode == "CN":
		log.Info("dealQuerycmd  Country CN ")
		//result, err := json.MarshalIndent(City.Continent.Names["zh-CN"], "", "")
		Continent := MatchMapTOJson(City.Continent.Names, config.CN)
		Country := MatchMapTOJson(City.Country.Names, config.CN)
		Subdivisions := MatchMapTOJson(City.Subdivisions[0].Names, config.CN)
		city := MatchMapTOJson(City.City.Names, config.CN)
		Latitude := FloatTOJson(City.Location.Latitude)
		Longitude := FloatTOJson(City.Location.Longitude)
		TimeZone := StringTOJson(City.Location.TimeZone)
		fmt.Fprintf(conn, "Continent: %v\n", string(Continent))
		fmt.Fprintf(conn, "Country: %v\n", string(Country))
		fmt.Fprintf(conn, "Subdivisions: %v\n", string(Subdivisions))
		fmt.Fprintf(conn, "City: %v\n", string(city))
		fmt.Fprintf(conn, "Time zone: %v\n", string(TimeZone))
		fmt.Fprintf(conn, "location:%v,%v\n", string(Latitude), string(Longitude))
		break
	case City.Country.IsoCode == "US":
		log.Info("dealQuerycmd  Country US ")
		Continent := MatchMapTOJson(City.Continent.Names, config.US)
		Country := MatchMapTOJson(City.Country.Names, config.US)
		Subdivisions := MatchMapTOJson(City.Subdivisions[0].Names, config.US)
		city := MatchMapTOJson(City.City.Names, config.US)
		Latitude := FloatTOJson(City.Location.Latitude)
		Longitude := FloatTOJson(City.Location.Longitude)
		TimeZone := StringTOJson(City.Location.TimeZone)
		fmt.Fprintf(conn, "Continent: %v\n", string(Continent))
		fmt.Fprintf(conn, "Country: %v\n", string(Country))
		fmt.Fprintf(conn, "Subdivisions: %v\n", string(Subdivisions))
		fmt.Fprintf(conn, "City: %v\n", string(city))
		fmt.Fprintf(conn, "Time zone: %v\n", string(TimeZone))
		fmt.Fprintf(conn, "location:%v,%v\n", string(Latitude), string(Longitude))
		break
	case City.Country.IsoCode == "PH":
		log.Info("dealQuerycmd  Country PH ")
		Continent := MatchMapTOJson(City.Continent.Names, config.Filipino)
		Country := MatchMapTOJson(City.Country.Names, config.Filipino)
		Subdivisions := MatchMapTOJson(City.Subdivisions[0].Names, config.Filipino)
		city := MatchMapTOJson(City.City.Names, config.Filipino)
		Latitude := FloatTOJson(City.Location.Latitude)
		Longitude := FloatTOJson(City.Location.Longitude)
		TimeZone := StringTOJson(City.Location.TimeZone)
		fmt.Fprintf(conn, "Continent: %v\n", string(Continent))
		fmt.Fprintf(conn, "Country: %v\n", string(Country))
		fmt.Fprintf(conn, "Subdivisions: %v\n", string(Subdivisions))
		fmt.Fprintf(conn, "City: %v\n", string(city))
		fmt.Fprintf(conn, "Time zone: %v\n", string(TimeZone))
		fmt.Fprintf(conn, "location:%v,%v\n", string(Latitude), string(Longitude))
		break
	case City.Country.IsoCode == "KR":
		log.Info("dealQuerycmd  Country KR ")
		Continent := MatchMapTOJson(City.Continent.Names, config.Korea)
		Country := MatchMapTOJson(City.Country.Names, config.Korea)
		Subdivisions := MatchMapTOJson(City.Subdivisions[0].Names, config.Korea)
		city := MatchMapTOJson(City.City.Names, config.Korea)
		Latitude := FloatTOJson(City.Location.Latitude)
		Longitude := FloatTOJson(City.Location.Longitude)
		TimeZone := StringTOJson(City.Location.TimeZone)
		fmt.Fprintf(conn, "Continent: %v\n", string(Continent))
		fmt.Fprintf(conn, "Country: %v\n", string(Country))
		fmt.Fprintf(conn, "Subdivisions: %v\n", string(Subdivisions))
		fmt.Fprintf(conn, "City: %v\n", string(city))
		fmt.Fprintf(conn, "Time zone: %v\n", string(TimeZone))
		fmt.Fprintf(conn, "location:%v,%v\n", string(Latitude), string(Longitude))
		break
	case City.Country.IsoCode == "JP":
		log.Info("dealQuerycmd  Country JP ")
		Continent := MatchMapTOJson(City.Continent.Names, config.Japanese)
		Country := MatchMapTOJson(City.Country.Names, config.Japanese)
		Subdivisions := MatchMapTOJson(City.Subdivisions[0].Names, config.Japanese)
		city := MatchMapTOJson(City.City.Names, config.Japanese)
		Latitude := FloatTOJson(City.Location.Latitude)
		Longitude := FloatTOJson(City.Location.Longitude)
		TimeZone := StringTOJson(City.Location.TimeZone)
		fmt.Fprintf(conn, "Continent: %v\n", string(Continent))
		fmt.Fprintf(conn, "Country: %v\n", string(Country))
		fmt.Fprintf(conn, "Subdivisions: %v\n", string(Subdivisions))
		fmt.Fprintf(conn, "City: %v\n", string(city))
		fmt.Fprintf(conn, "Time zone: %v\n", string(TimeZone))
		fmt.Fprintf(conn, "location:%v,%v\n", string(Latitude), string(Longitude))
		break
	case City.Country.IsoCode == "BR":
		log.Info("dealQuerycmd  Country BR ")
		Continent := MatchMapTOJson(City.Continent.Names, config.Brazil)
		Country := MatchMapTOJson(City.Country.Names, config.Brazil)
		Subdivisions := MatchMapTOJson(City.Subdivisions[0].Names, config.Brazil)
		city := MatchMapTOJson(City.City.Names, config.Brazil)
		Latitude := FloatTOJson(City.Location.Latitude)
		Longitude := FloatTOJson(City.Location.Longitude)
		TimeZone := StringTOJson(City.Location.TimeZone)
		fmt.Fprintf(conn, "Continent: %v\n", string(Continent))
		fmt.Fprintf(conn, "Country: %v\n", string(Country))
		fmt.Fprintf(conn, "Subdivisions: %v\n", string(Subdivisions))
		fmt.Fprintf(conn, "City: %v\n", string(city))
		fmt.Fprintf(conn, "Time zone: %v\n", string(TimeZone))
		fmt.Fprintf(conn, "location:%v,%v\n", string(Latitude), string(Longitude))
		break
	case City.Country.IsoCode == "ES":
		log.Info("dealQuerycmd  Country ES ")
		Continent := MatchMapTOJson(City.Continent.Names, config.Spain)
		Country := MatchMapTOJson(City.Country.Names, config.Spain)
		Subdivisions := MatchMapTOJson(City.Subdivisions[0].Names, config.Spain)
		city := MatchMapTOJson(City.City.Names, config.Spain)
		Latitude := FloatTOJson(City.Location.Latitude)
		Longitude := FloatTOJson(City.Location.Longitude)
		TimeZone := StringTOJson(City.Location.TimeZone)
		fmt.Fprintf(conn, "Continent: %v\n", string(Continent))
		fmt.Fprintf(conn, "Country: %v\n", string(Country))
		fmt.Fprintf(conn, "Subdivisions: %v\n", string(Subdivisions))
		fmt.Fprintf(conn, "City: %v\n", string(city))
		fmt.Fprintf(conn, "Time zone: %v\n", string(TimeZone))
		fmt.Fprintf(conn, "location:%v,%v\n", string(Latitude), string(Longitude))
		break
	case City.Country.IsoCode == "DE":
		log.Info("dealQuerycmd  Country DE ")
		Continent := MatchMapTOJson(City.Continent.Names, config.Deutschland)
		Country := MatchMapTOJson(City.Country.Names, config.Deutschland)
		Subdivisions := MatchMapTOJson(City.Subdivisions[0].Names, config.Deutschland)
		city := MatchMapTOJson(City.City.Names, config.Deutschland)
		Latitude := FloatTOJson(City.Location.Latitude)
		Longitude := FloatTOJson(City.Location.Longitude)
		TimeZone := StringTOJson(City.Location.TimeZone)
		fmt.Fprintf(conn, "Continent: %v\n", string(Continent))
		fmt.Fprintf(conn, "Country: %v\n", string(Country))
		fmt.Fprintf(conn, "Subdivisions: %v\n", string(Subdivisions))
		fmt.Fprintf(conn, "City: %v\n", string(city))
		fmt.Fprintf(conn, "Time zone: %v\n", string(TimeZone))
		fmt.Fprintf(conn, "location:%v,%v\n", string(Latitude), string(Longitude))
		break
	default:
		log.Info("dealQuerycmd  Country default ")
		Continent := MapTOJson(City.Continent.Names)
		Country := MapTOJson(City.Country.Names)
		Subdivisions := MapTOJson(City.Subdivisions[0].Names)
		city := MapTOJson(City.City.Names)
		Latitude := FloatTOJson(City.Location.Latitude)
		Longitude := FloatTOJson(City.Location.Longitude)
		TimeZone := StringTOJson(City.Location.TimeZone)
		fmt.Fprintf(conn, "Continent: %v\n", string(Continent))
		fmt.Fprintf(conn, "Country: %v\n", string(Country))
		fmt.Fprintf(conn, "Subdivisions: %v\n", string(Subdivisions))
		fmt.Fprintf(conn, "City: %v\n", string(city))
		fmt.Fprintf(conn, "Time zone: %v\n", string(TimeZone))
		fmt.Fprintf(conn, "location:%v,%v\n", string(Latitude), string(Longitude))
		break
	}
	Continent := MapTOJson(City.Continent.Names)
	Country := MapTOJson(City.Country.Names)
	Subdivisions := MapTOJson(City.Subdivisions[0].Names)
	city := MapTOJson(City.City.Names)
	Latitude := FloatTOJson(City.Location.Latitude)
	Longitude := FloatTOJson(City.Location.Longitude)
	TimeZone := StringTOJson(City.Location.TimeZone)
	fmt.Println("______________________")
	fmt.Println(string(Continent))
	fmt.Println(string(Country))
	fmt.Println(string(Subdivisions))
	fmt.Println(string(city))
	fmt.Println(string(TimeZone))
	fmt.Println(string(Latitude), string(Longitude))
	fmt.Println(string(City.Country.IsoCode))
	fmt.Println("______________________")

}
