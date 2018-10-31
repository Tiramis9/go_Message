package server

import (
	"encoding/json"
	"fmt"
	"regexp"
)

//处理client 消息，通过正则匹配，返回IP
func MustIP(Msg string) string {
	myexp := regexp.MustCompile(`\d+\.\d+.\d+\.\d+`)
	result := myexp.FindAllStringSubmatch(Msg, 1)
	var str string
	for _, value := range result {
		fmt.Println("start run MustIP")
		if value != nil {
			for _, str = range value {
				fmt.Println(str)
			}
		}
	}
	return str
}

//根据国家匹配本地语言
func MatchMapTOJson(Names map[string]string, national string) []byte {
	result, err := json.MarshalIndent(Names[national], "", "")
	//result, err := json.MarshalIndent(Names, "", "")
	//result, err := json.Marshal(Names)
	if err != nil {
		fmt.Println("err = ", err)
		return nil
	}
	//	fmt.Println(string(result))
	return result
}

//返回全部语言
func MapTOJson(Names map[string]string) []byte {
	//result, err := json.MarshalIndent(Names["zh-CN"], "", "")
	result, err := json.MarshalIndent(Names, "", "")
	//result, err := json.Marshal(Names)
	if err != nil {
		fmt.Println("err = ", err)
		return nil
	}
	//	fmt.Println(string(result))
	return result
}
func FloatTOJson(location float64) []byte {

	//result, err := json.Marshal(Msg)
	result, err := json.MarshalIndent(location, "", " ")
	if err != nil {
		fmt.Println("err = ", err)
		return nil
	}
	//	fmt.Println("FloatTOJson", string(result))
	return result
}
func StringTOJson(TimeZone string) []byte {

	//result, err := json.Marshal(Msg)
	result, err := json.MarshalIndent(TimeZone, "", " ")
	if err != nil {
		fmt.Println("err = ", err)
		return nil
	}
	//	fmt.Println("FloatTOJson", string(result))
	return result
}
