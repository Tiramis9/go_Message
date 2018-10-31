package server

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/oschwald/geoip2-golang"
	"public.info/config"
)

//CLIENT
type client struct {
	C    chan string
	Name string
}

//用户登录
var login = make(chan client)
var message = make(chan string)
var logout = make(chan client)
var onlineClientMap map[client]bool
var onlineClientNameMap map[string]client

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
func dealQuerycmd(conn net.Conn, connMsg string) {
	log := config.Log.WithField("package", "server")
	log.Info("start running dealQuerycmd ")
	fmt.Println(connMsg)
	IP := MustIP(connMsg)
	if IP == "" {
		fmt.Println("IP is nil", IP)
		fmt.Fprintln(conn, "输入的IP有误，请重新输入!")
		return
	}
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

//给当前客户端写信息，消息回显给当前客户端
func ReadFromMyclient(conn net.Conn, myCli client) {
	for msg := range myCli.C {
		conn.Write([]byte(msg + "\n"))
	}
}

//处理查询任务指令
func dealWhoCmd(conn net.Conn) {
	for cli := range onlineClientMap {
		fmt.Fprintln(conn, "name=", cli.Name)
	}
	fmt.Fprintln(conn, "目前在线任务数:", len(onlineClientMap))
}

//处理已连接的函数，接收用户发送过来的数据
func HandleConnect(conn net.Conn) {
	log := config.Log.WithField("package", "server")
	log.Info("start running HandleConnect")
	defer conn.Close()
	cliTimeOutTicker := time.NewTicker(time.Second * 60)
	//当前用户是否发送消息channel
	revcData := make(chan interface{})
	//对端是否关闭，用户不是超时关闭而是主动关闭
	removeClose := make(chan interface{})
	myName := conn.RemoteAddr().String()
	fmt.Println(myName)
	myCli := client{make(chan string), myName}
	go ReadFromMyclient(conn, myCli)
	login <- myCli
	message <- myName + "已经登入"
	myCli.C <- myName + " Connection successful"
	go func() {
		input_scanf := bufio.NewScanner(conn)
		for input_scanf.Scan() {
			connMsg := input_scanf.Text()
			switch {
			case connMsg == "ip":
				fmt.Println("Query local IP")
				dealQuerycmd(conn, myName)
				break
			case connMsg == "who":
				fmt.Println("Query on line IP")
				dealWhoCmd(conn)
			default:
				fmt.Println("default dealMsg")
				dealQuerycmd(conn, connMsg)
				break
			}
			//用户正常输入了数据，更新定时器
			revcData <- struct{}{}
		}
		//用户正常退出
		removeClose <- struct{}{}
	}()
	for {
		select {
		case <-cliTimeOutTicker.C:
			logout <- myCli
			fmt.Fprintln(conn, "由于你长时间没有活跃，已断开连接")
			fmt.Println(myName, "长时间没有活跃，已断开连接")
			message <- myName + "已经超时，被退出"
			return
		case <-revcData:
			//用户正常数据 刷新定时器
			cliTimeOutTicker.Stop()
			cliTimeOutTicker = time.NewTicker(time.Second * 60)
		case <-removeClose:
			//用户主动退出
			logout <- myCli
			message <- myName + "主动退出"
			return
		}
	}

}

//管理当前任务情况，
func ManagerInfo() {
	//要把客户端的信息存储到map中
	onlineClientMap = make(map[client]bool)
	onlineClientNameMap = make(map[string]client)
	for {
		select {
		case cli := <-login:
			onlineClientMap[cli] = true
			onlineClientNameMap[cli.Name] = cli
		case cli := <-logout:
			delete(onlineClientMap, cli)
			delete(onlineClientNameMap, cli.Name)
		case msg := <-message:
			for cli, isOnline := range onlineClientMap {
				if isOnline == true {
					cli.C <- msg
				}
			}
		}
	}
}
