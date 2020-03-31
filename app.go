package main

import (
	user2 "bitbucket.org/Hacktive8/FinalProjectGO/src/user"
	"fmt"
	"github.com/gorilla/sessions"
	"html/template"
	"log"
	"net/http"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func main()  {
	fmt.Println("hello")
	Template := template.Must(template.ParseGlob("files/var/templates/*"))
	user := user2.New(Template)

	http.HandleFunc("/", user.HomeUser)
	http.HandleFunc("/homelogged", user.HomeUserLogged)
	http.HandleFunc("/about", user.AboutUser)
	http.HandleFunc("/aboutlogged", user.AboutUserLogged)
	http.HandleFunc("/contact", user.ContactUser)
	http.HandleFunc("/contactlogged", user.ContactUserLogged)
	http.HandleFunc("/articles", user.ArticlesUser)
	http.HandleFunc("/articles/add", user.AddArticlesUser)
	http.HandleFunc("/articles/edit", user.EditArticlesUser)
	http.HandleFunc("/articles/remove", user.RemoveArticlesUser)
	http.HandleFunc("/articles/publish", user.PublishArticlesUser)
	http.HandleFunc("/articles/unpublish", user.UnpublishArticlesUser)
	http.HandleFunc("/register", user.RegisterUser)
	http.HandleFunc("/login", user.LoginUser)
	http.HandleFunc("/logout", user.LogoutUser)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))




	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("error listern")
	}
}
