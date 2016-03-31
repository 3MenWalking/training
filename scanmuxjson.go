package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"github.com/gorilla/mux"
	"encoding/json"
)

type StdMsg struct {
	MsgNo   string   `json:"msgno"`   //Important to capitalize first letter to make it public
	MsgText string   `json:"msgtext"` //use json to map to idomatic lower case element name
}

type StdMsgs []StdMsg

func readLines(path string) (StdMsgs, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines StdMsgs
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(),"\t")
		lines = append(lines, StdMsg{s[0],s[1]})
	}
	return lines, scanner.Err()
}

var lines, _ = readLines("stdmsg.txt")

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, Welcome to Golang!")
    })
	r.HandleFunc("/querybyno/{msgNo}", findMsgByNum)
	r.HandleFunc("/queryall", displayAllMsg)
	http.ListenAndServe(":8099", r)
}

func lookUp(stdMsgNo string) StdMsgs {
	var results StdMsgs
	found := false
	for _, line := range lines {
		if strings.Contains(line.MsgNo, stdMsgNo) {
			results = append(results, line)
			found = true
		}
	}
	if found == false {
		results = append(results, StdMsg{"NotFound", "Query Failed!"})
	}
	return results
}

func findMsgByNum(w http.ResponseWriter, r *http.Request) {
	stdMsgs := lookUp(mux.Vars(r)["msgNo"])
	json.NewEncoder(w).Encode(stdMsgs)
}

func displayAllMsg(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(lines)
}
