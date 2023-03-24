package model

type Book struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

var BookDatas = []Book{}
