package main

import(
    "net/http"
	"github.com/gorilla/mux"
	"html/template"
	"github.com/go-redis/redis"

)
var client *redis.Client
var templates *template.Template

func main(){
	client=redis.NewClient(&redis.Options{
		Addr:"localhost:6379",


	})
	templates=template.Must(template.ParseGlob("templates/*.html"))
	r:=mux.NewRouter()
	r.HandleFunc("/",indexhandler).Methods("GET")
	http.Handle("/",r)
	http.ListenAndServe(":8080",nil)
}

func indexhandler(w http.ResponseWriter,r *http.Request){
	comments,err:=client.LRange("comments",0,10).Result()
	templates.ExecuteTemplate(w,"index.html",nil)
}

