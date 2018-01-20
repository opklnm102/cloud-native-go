package api

import (
	"encoding/json"
	"net/http"
)

// go는 struct의 var는 Upper case로 시작
type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

// func (struct) funcName(paras) return type
// (struct) -> struct에 포함된 func다, struct의 member func라고 생각
func (book Book) ToJSON() []byte { // book은 this라고 생각하자
	ToJSON, err := json.Marshal(book)
	if err != nil {
		panic(err)
	}
	return ToJSON
}

func FromJSON(data []byte) Book {
	book := Book{} // create empty struct
	err := json.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}
	return book
}

var Books = []Book{
	Book{Title: "The Hitchhiker's Guide", Author: "Douglas Adams", ISBN: "0345391802"},
	Book{Title: "Cloud Native Go", Author: "M.-Leander Reimer", ISBN: "0000000000"},
}

func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	books, err := json.Marshal(Books)
	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(books)
}
