package main

import (
	"bufio"
	"fmt"
	//"log"
	"net/http"
	"os"
	"strings"
)

var lines, _ = readLines("stdmsg.txt")

func main() {
    http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, Welcome to Golang!")
    })
	http.HandleFunc("/query", findMsg)
	http.HandleFunc("/queryall", displayAllMsg)
	http.ListenAndServe(":8080", nil)
}

func readLines(path string) ([]string, error) {  //function multiple returns
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
	stdMsgs := lookUp(r.URL.Query().Get("msgNo"))
	for _, stdMsg := range stdMsgs {
		fmt.Fprintln(w, stdMsg)
	}
}

func displayAllMsg(w http.ResponseWriter, r *http.Request) {
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
}
