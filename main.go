package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

// User (more like an object)
type User struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func userHandler(write http.ResponseWriter, read *http.Request) {
	if read.Method == "POST" {
		decoder := json.NewDecoder(read.Body)
		var user User
		err := decoder.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		connStr := "connect_timeout=2 host=192.168.0.2 user=postgres sslmode=disable dbname=postgres password=password"
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			log.Fatal(err)
		}

		var userid int
		rows := db.QueryRow("INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id", user.Name, user.Age).Scan(&userid)
		if rows != nil {
			error := rows.Error()
			fmt.Fprintf(write, error)
		}
		fmt.Fprintf(write, "insert done")
	} else {
		connStr := "connect_timeout=2 host=192.168.0.2 user=postgres sslmode=disable dbname=postgres password=password"
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			log.Fatal(err)
		}

		age := 15
		rows, err := db.Query("SELECT name, age FROM users WHERE age > $1 ", age)
		if err != nil {
			log.Fatal(err)
		}

		var users []User

		for rows.Next() {
			var name string
			var myage int
			if err := rows.Scan(&name, &myage); err != nil {
				log.Fatal(err)
			}

			if err != nil {
				log.Fatal(err)
			}

			users = append(users, User{Name: name, Age: myage})

		}

		write.Header().Set("Content-Type", "application/json")
		write.Header().Set("Access-Control-Allow-Origin", "*")
		write.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		json.NewEncoder(write).Encode(users)

		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
	}

}

func main() {
	http.HandleFunc("/users", userHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
