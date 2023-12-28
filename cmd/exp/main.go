package main

import (
	"fmt"
	"html/template"
	"os"
)

type User struct {
	Name string
	Age  int
	Bio  string
	Meta userMeta
}

type userMeta struct {
	Visits int
}

func main() {
	t, err := template.ParseFiles("hello.gotmpl")
	if err != nil {
		panic(err)
	}

	bio := `<script>alert("hahahah, you have been hacked!");</script>`

	user := User{
		Name: "John Smith",
		Age:  32,
		Bio:  bio,
		Meta: userMeta{
			Visits: 1500,
		},
	}
	// user1 := struct {
	// 	Name string
	// 	Age  int
	// 	Meta struct {
	// 		Visits int
	// 	}
	// }{
	// 	Name: "Ramiro",
	// 	Age:  33,
	// 	Meta: struct{ Visits int }{
	// 		Visits: 66,
	// 	},
	// }

	fmt.Println(user.Meta.Visits)
	// fmt.Println(user1.Meta.Visits)

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}

}
