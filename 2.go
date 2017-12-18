package main

import (
	"log"
	"net/http"
	"time"
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
)

//server=10.25.4.35;database=isosdb_30;user id=iaas;password=sdt108
type sqlHandler struct{}

func  (sh *sqlHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {

}
type Struct struct{
	Greeting 	string
	Punct 		string
	Who 		string
}
type timeHandler struct {
	format string
}

func (th *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	tm := time.Now().Format(th.format)
	w.Write([]byte("The time is " + tm))
}
//func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request){}
//func (ts Struct) ServeHTTP(w http.ResponseWriter, r *http.Request){}
func main() {

	mux := http.NewServeMux()

	th1123 := &timeHandler{format: time.RFC1123}
	mux.Handle("/time/rfc1123", th1123)

	th3339 := &timeHandler{format: time.RFC3339}
	mux.Handle("/time/rfc3339", th3339)

	sql := &sqlHandler{}
	mux.Handle("/logicGroupCount", sql)

	log.Println("Listening")
	http.ListenAndServe("localhost:10086", mux)
}