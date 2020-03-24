package main

import (
	suivant "github.com/gypsydave5/au-suivant"
	"net/http"
	"time"
)

func main() {
	s := suivant.New([]string{"Dave", "Chris", "Lisa", "Riya"}, 5*time.Second)
	ss := suivant.NewServer(s)

	http.ListenAndServe(":8080", ss)
}
