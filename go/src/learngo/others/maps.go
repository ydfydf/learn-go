package others

import "fmt"

func main() {
	//map创建方式一
	m := map[string]string {
		"name" : "ydf",
		"course" : "golang",
		"site" : "young",
		"quality" : "not bad",
	}
	//map是hashmap，key是无序的，打印出的key和value也不是固定的
	fmt.Println(m)

	//map创建方式二，m2 = empty map，常用方式
	m2 := make(map[string]int)
	fmt.Println(m2)

	//map创建方式三，m3 = nil
	var m3 map[string]int
	fmt.Println(m3)

	//map的遍历
	fmt.Println("Traversing map")
	for k,v := range m {
		fmt.Println(k,v)
	}
	/*for k := range m {
		fmt.Println(k)
	}*/

	//获取map的值
	fmt.Println("Getting Values")
	courseName := m["course"]
	fmt.Println(courseName)
	courseName1 := m["cours"]//当去取map中不存在的key时，也会得到value，value为空
	fmt.Println(courseName1)

	//判断key值再map中存不存在
	/*who,ok := m["name"]
	fmt.Println(who,ok)*/
	if who, ok := m["name"]; ok != true{
		fmt.Println("key does not exsit")
	}else {
		fmt.Println(who,ok)
	}

	//删除元素
	fmt.Println("Delete Values")
	if _, ok := m["course"]; ok != true{
		fmt.Println("key does not exsit")
	}else {
		delete(m,"course")
		fmt.Println(m)
	}

}
