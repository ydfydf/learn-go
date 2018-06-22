package queue

//interface{}表示任何类型
type Queue []interface{}

func (q *Queue) Push(v interface{}){
	*q = append(*q,v)
}

//限定参数为int
func (q *Queue) Push2(v int){
	*q = append(*q,v)
}

//参数不限定类型
func (q *Queue) Push3(v interface{}){
	//函数实现时限定类型
	*q = append(*q,v.(int))
}

func (q *Queue) Pop() interface{}{
	head := (*q)[0]
	*q = (*q)[1:]
	return head//注意head是interface{}，所以需要head.(int)
}

func (q *Queue) Pop1() int{
	head := (*q)[0]
	*q = (*q)[1:]
	return head.(int)//注意head是interface{}，所以需要head.(int)
}

//参数不限定类型
func (q *Queue) Pop2() interface{}{
	head := (*q)[0]
	*q = (*q)[1:]
	return head.(int)//函数实现时限定类型,注意head是interface{}，所以需要head.(int)
}


func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
