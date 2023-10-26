package main

import (
	"fmt"
	"net/http"

	spinhttp "github.com/fermyon/spin/sdk/go/http"
	"github.com/fermyon/spin/sdk/go/sqlite"
)

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {

		router := spinhttp.NewRouter()
		router.POST("/:bite", func(writer http.ResponseWriter, req *http.Request, params spinhttp.Params) {
			db := sqlite.Open("default")
			defer db.Close()

			_, err := db.Exec("INSERT INTO bites (content) VALUES (?)", params.ByName("bite"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			fmt.Fprintf(writer, "Created and stored a bite: %s!", params.ByName("bite"))
		})

		router.ServeHTTP(w, r)

	})
}

func main() {}
