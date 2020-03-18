package main

import (
	"bitbucket.org/Hacktive8/FinalProjectGO/src/user"

	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main()  {
	fmt.Println("hello")
	Template := template.Must(template.ParseGlob("files/var/templates/*"))
	user := user.New(Template)

	http.HandleFunc("/", user.HomeUser)
	//http.HandleFunc("/about", user.AboutUser)
	http.HandleFunc("/contact", user.ContactUser)
	//http.HandleFunc("/articles", user.ArticlesUser)
	http.HandleFunc("/articles/add", user.ArticlesUser)
	//http.HandleFunc("/articles/edit", user.ArticlesUser)
	//http.HandleFunc("/articles/remove", user.ArticlesUser)
	http.HandleFunc("/register", user.RegisterUser)
	http.HandleFunc("/login", user.LoginUser)
	//http.HandleFunc("/logout", user.LogoutUser)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("error listern")
	}
}