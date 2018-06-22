package model

import "encoding/json"

//这个profile文件不要在zhanai目录下，要放在公用的目录下，因为此爬虫项目可能不止爬珍爱网，也有可能爬其他网站，这样的话大家就都能使用

type Profile struct {
	Name        string
	Gender      string
	Age         int
	Height      int
	Weight      int
	Income      string
	Marriage    string
	Education   string
	Occupation  string
	Hokou       string
	Xinzuo      string
	House       string
	Car         string
}

func FromJsonObj(o interface{}) (Profile, error){
	var profile Profile
	s ,err := json.Marshal(o)
	if err != nil {
		return profile,err
	}
	err = json.Unmarshal(s,&profile)
	return profile,err

}