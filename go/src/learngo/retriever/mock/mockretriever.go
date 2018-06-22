package mock

//实现者(定义者)文件
type Retriever struct {
	Contents string
}

//接口方法的实现，实现了Get方法
//实现者不关心使用者的接口名字叫什么
func (r *Retriever) Get(url string) string {
//func (r Retriever) Get(url string) string {
	return r.Contents
}
