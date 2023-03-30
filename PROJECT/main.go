package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var Data = map[string]interface{}{
	"title":   "lorem ipsum",
	"isLogin": false,
}

func main() {

	router := mux.NewRouter()

	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/project", project).Methods("GET")
	router.HandleFunc("/mainblog/{id}", mainblog).Methods("GET")
	router.HandleFunc("/new-blog", newblog).Methods("POST")

	fmt.Println("server running on port 5000")
	http.ListenAndServe("localhost:5000", router)

}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html ; Charset=utf-8")
	w.WriteHeader(http.StatusOK)
	templ, err := template.ParseFiles("html/index.html")
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("message" + err.Error()))
		return
	}
	templ.Execute(w, Data)
}

func project(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html ; Charset=utf-8")
	w.WriteHeader(http.StatusOK)

	templ, err := template.ParseFiles("html/blog.html")

	if err != nil {

		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("message  " + err.Error()))
		return
	}

	templ.Execute(w, Data)
}

func mainblog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html ; Charset=utf-8")
	w.WriteHeader(http.StatusOK)

	templ, err := template.ParseFiles("html/mainblog.html")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {

		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("message  " + err.Error()))
		return
	}

	var resp = map[string]interface{}{
		"data": Data,
		"id":   id,
	}

	templ.Execute(w, resp)
}


func newblog(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	 fmt.Println( "perojctname : " + r.PostForm.Get("projectname : ")) 
	 fmt.Println("deskripsi" + r.PostForm.Get("description")) 


	http.Redirect(w ,r ,"/" , http.StatusMovedPermanently)
} 