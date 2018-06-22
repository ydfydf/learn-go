package poster

type Retriever struct {
	Contents string
}

//注意这里是指针接收者，所以r.Contents才能被改变
func (r *Retriever) Post(url string,
	form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}

func (r *Retriever) Connet() {

}

func (r *Retriever) Get(url string) string {
	return r.Contents
}

