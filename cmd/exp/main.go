package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Bio  string
	Age  int
}

type UserMeta struct {
	Visits int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "Sunny Motiani",
		Bio:  `<script>alert("Hahah you have been hacked")</script>`,
		Age:  123,
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}

}
#commit