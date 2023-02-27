package main

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type User struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Salary int    `json:"salary"`
}

// var user User
var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root:Malli@252000@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/users", Getval).Methods("GET")
	r.HandleFunc("/users", Insertval).Methods("POST")
	log.Println("Strting server ..........................")
	log.Fatal(http.ListenAndServe(":8080", r))

}

func Insertval(w http.ResponseWriter, r *http.Request) {
	csvfile, err := os.Open("sample.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvfile.Close()
	a := csv.NewReader(csvfile)
	record, _ := a.ReadAll()

	for _, val := range record {
		firs, _ := strconv.Atoi(val[0])
		third, _ := strconv.Atoi(val[2])
		_, err := db.Exec("INSERT INTO emp VALUES (?,?,?)", firs, val[1], third)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func Getval(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	result, err := db.Query("SELECT * FROM emp")
	if err != nil {
		log.Fatal(err)
	}
	defer result.Close()

	var uslice []User
	for result.Next() {
		var user User
		result.Scan(&user.Id, &user.Name, &user.Salary)
		uslice = append(uslice, user)
		log.Println(uslice)
	}
	json.NewEncoder(w).Encode(uslice)

	// log.Println(record)
}
