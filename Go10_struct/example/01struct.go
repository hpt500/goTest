package example

import (
	. "fmt"
)

type Books struct {
	title   string
	author  string
	subject string
	bookid  int
}

func Struct() {
	Println("一.定义结构体")
	var book Books
	book.title = "struct"
	book.author = "Xuyixian"
	book.subject = "go语言"
	book.bookid = 565420423

	ShowBook(book)
	Println()

}

func ShowBook(book Books) {
	Println("book of title:", book.title)
	Println("book of author:", book.author)
	Println("book of subject:", book.subject)
	Println("book of bookid:", book.bookid)
}
