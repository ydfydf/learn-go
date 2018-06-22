package parser

import (
	"testing"
	"io/ioutil"
)

//由于做测试的时候，可能测试机器没有联网，或者网页访问不了等等外在因素，所以需要将测试的网页数据先保存到文件citylist_test_data.html里

func TestParseCityList(t *testing.T) {
	//contents,err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	contents , err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s\n",contents)
	result := ParseCityList(contents)

	const resultSize = 470
	expectedUrls := []string {
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}

	for i, url := range expectedUrls{
		if result.Requests[i].Url != url{
			t.Errorf("Expected url #%d: %s; but was %s",i,url,result.Requests[i].Url)
		}
	}

	if len(result.Requests) != resultSize{
		t.Errorf("result should have %d requests;but had %d",resultSize,len(result.Requests))
	}



}