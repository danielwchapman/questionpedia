package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/danielwchapman/questionpedia-backend/pkg/question"
)

func main() {
	http.HandleFunc("/question", questionHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func questionHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	switch req.Method {
	case "GET":
		fmt.Fprintf(w, "GET\n\n")
		query := req.URL.Query()
		id := query.Get("id")
		q := getQuestion(id)
		fmt.Fprintf(w, "%v", q)
	case "POST":
		fmt.Fprintf(w, "POST")
	case "PUT":
		fmt.Fprintf(w, "PUT")
	case "DELETE":
		fmt.Fprintf(w, "DELETE")
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func getQuestion(id string) *question.Question {
	if id == "0" {
		return &question.Question{
			id: id,
			text: "quetsion 0",
		}
	} else {
		return &question.Question{
			id: id,
			text: "doesn't exist",
		}
	}
}