package parser

import (
	"regexp"
	"strconv"
	"learngo/crawler/model"
	"learngo/crawler/engine"
	"learngo/crawler_distributed/config"
)

//const ageRe = `<td><span class="label">年龄：</span>([\d]+)岁</td>`//[\d]+，[\d]代表一个数字，+代表前面表达式执行1此或多次
//const marriageRe = `<td><span class="label">婚况：</span>([^<]+)</td>`

//由于需要很多正则表达式进行compile，会很花时间，所以在定义的时候就将正则表达式compile了，预先编译
var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+)CM</td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var weightRe = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([^<])kg</span></td>`)
var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var xinzuoRe = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(`<td><span class="label">职业：</span><span field="">([^<])+</span></td>`)
var hokouRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var houseRe = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
var gurssRe = regexp.MustCompile(`<a class="exp-user-name"[^>]*href="(http://album.zhenai.com/u/[\d]+)">([^<]+)</a>`)
var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

func extractString(contents []byte,re *regexp.Regexp) string{
	match := re.FindSubmatch(contents) //只获取第一个匹配的数据
	if len(match) >= 2 { //因为match是[][]byte，[0]表示整个字符串，[1]表示提取出来的值
		return string(match[1])
	} else {
			return ""
		}
}


func parseProfile(contents []byte,url string,name string) engine.ParserResult{
	profile := model.Profile{}
	profile.Name = name

	age , err := strconv.Atoi(extractString(contents,ageRe))
	if err == nil {
		profile.Age = age
	}
	height , err := strconv.Atoi(extractString(contents,heightRe))
	if err == nil {
		profile.Height = height
	}
	weight , err := strconv.Atoi(extractString(contents,weightRe))
	if err == nil {
		profile.Weight = weight
	}
	profile.Gender = extractString(contents,genderRe)
	profile.Income = extractString(contents,incomeRe)
	profile.Marriage = extractString(contents,marriageRe)
	profile.Education = extractString(contents,educationRe)
	profile.Occupation = extractString(contents,occupationRe)
	profile.Hokou = extractString(contents,hokouRe)
	profile.Xinzuo = extractString(contents,xinzuoRe)
	profile.House = extractString(contents,houseRe)
	profile.Car = extractString(contents,carRe)

	result := engine.ParserResult{
		Items: []engine.Item{
			{//注意这里赋值的打括号
				Url: url,
				Type:"zhenai",
				Id : extractString([]byte(url),idUrlRe),
				Payload: profile,
			},
		},
	} //注意这里的slice赋值方式，不需要一个一个append

	matches := gurssRe.FindAllSubmatch(contents,-1)
	for _,m := range matches {
		url := string(m[1])//这里会开辟栈，会将此时的m[1]的值拷贝给url，如果不拷贝直接使用Url:string(m[1])，由于m的值在不断的累加，所以不能得到预想的值
		name := string(m[2])
		result.Requests = append(result.Requests,engine.Request{
			Url:url,
			Parser: NewProfileParser(name),
		})
	}
	return result
}

type ProfileParser struct {
	userName string
}

func (p *ProfileParser) Parse(contents []byte, url string) engine.ParserResult {
	return parseProfile(contents,url,p.userName)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return config.ParseProfile,p.userName
}

func NewProfileParser(name string) *ProfileParser {
	return &ProfileParser{
		userName:name,
	}
}
