package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/risset/go-rest-api/data"
	"github.com/risset/go-rest-api/routes"
)

func main() {
	port := flag.Int("p", 3333, "port to open server on")
	flag.Parse()

	dbParams := data.DbParam{
		User:    os.Getenv("DB_USER"),
		Pass:    os.Getenv("DB_PASS"),
		Name:    os.Getenv("DB_NAME"),
		Address: os.Getenv("DB_ADDRESS"),
	}

	store, err := data.NewDataStore(dbParams)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	r := routes.NewRouter(store)

	http.ListenAndServe(fmt.Sprintf(":%d", *port), r)
}
