package main

import (
	"bytes"
	"crypto/rand"
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/artdarek/go-unzip"

	_ "github.com/lib/pq"
)

const (
	port = 5432
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("DATABASE_HOST"), port, os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_NAME"))
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// Adding crypto mod which has CVE-2020-7919 in go v1.12.5
	c := 10
	b := make([]byte, c)
	_, err1 := rand.Read(b)
	if err1 != nil {
		fmt.Println("error:", err)
		return
	}
	// The slice should now contain random bytes instead of only zeroes.
	fmt.Println(bytes.Equal(b, make([]byte, c)))

	// Another CVE
	uz := unzip.New("file.zip", "directory/")
	err2 := uz.Extract()
	if err2 != nil {
		fmt.Println(err)
	}

	//
	fmt.Println("Successfully connected to:", os.Getenv("DATABASE_HOST"))
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Successfully connected to: %s Postgres host", os.Getenv("DATABASE_HOST"))
}
