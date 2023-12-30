package main

import (
	"html/template"
	"os"
)

type userMeta struct {
	Visits   int
	Articles []string
}

type User struct {
	Name     string
	Age      int
	Bio      string
	Meta     userMeta
	IsActive bool
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
			Visits:   1500,
			Articles: []string{"The first article", "The second article", "The third article", "Fourth article"},
		},
		IsActive: true,
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}

}
