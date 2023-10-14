package main

import (
	"fmt"
	"net/http"
	"strconv"

	//"net/url"
	"github.com/sbani/go-humanizer/numbers"
)

// Map represents a map with a number, an author and a finished status
type Map struct {
	Number   int    `json:"number"`
	Author   string `json:"author"`
	Finished bool   `json:"finished"`
}

// MapInfo represents a map with a number, an author, a time left, a time until, a finished status, a server and a difficulty
type MapInfo struct {
	Number     int
	Author     string
	TimeLeft   int
	TimeUntil  int
	Finished   bool
	Server     int
	Difficulty string
}

// Server represents a server with a number, a difficulty, a slice of maps, a time limit and a time left
type Server struct {
	ServerNumber     int    `json:"serverNumber"`
	ServerDifficulty string `json:"serverDifficulty"`
	Maps             []Map  `json:"maps"`
	TimeLimit        int    `json:"timeLimit"`
	TimeLeft         int    `json:"timeLeft"`
}

// Data represents the whole data with a slice of servers and a competition time left
type Data struct {
	Servers      []Server `json:"servers"`
	ComptimeLeft int      `json:"comptimeLeft"`
}

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		argString := r.URL.Query().Get("num")
		if argString == "" {
			fmt.Fprintf(w, "Error: No number provided")
			return
		}
		arg, err := strconv.Atoi(argString)
		if err != nil {
			fmt.Fprintf(w, "Error: %s", err)
			return
		}
		ordinal := getOrdinal(arg)
		fmt.Fprintf(w, ordinal)
	}
	fmt.Println("Listening on http://localhost:8026/ordinal")
	http.HandleFunc("/ordinal", handler)
	http.ListenAndServe(":8026", nil)
}

func getOrdinal(num int) string {
	return numbers.Ordinalize(num)
}
