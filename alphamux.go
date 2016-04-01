package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"github.com/gorilla/mux"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

var lines, _ = readLines("stdmsg.txt")

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, Welcome to Golang!")
    })
	r.HandleFunc("/query/{msgNo}", findMsg)
	r.HandleFunc("/queryall", displayAllMsg)
	http.ListenAndServe(":8090", r)
}

func lookUp(stdMsgNo string) []string {
	var results []string
	found := false
	for _, line := range lines {
		if strings.Contains(line, stdMsgNo) {
			results = append(results, line)
			found = true
		}
	}
	if found == false {
		results = append(results, "Not found")
	}
	return results
}

func findMsg(w http.ResponseWriter, r *http.Request) {
	//stdMsgs := lookUp(r.URL.Query().Get("msgNo"))
	stdMsgs := lookUp(mux.Vars(r)["msgNo"])
	for _, stdMsg := range stdMsgs {
		fmt.Fprintln(w, stdMsg)
	}
}

func displayAllMsg(w http.ResponseWriter, r *http.Request) {
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
}
