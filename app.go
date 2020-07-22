package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/gorilla/mux"
)

type EnvironmentVariable struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Parameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Command struct {
	Command              string                `json:"command"`
	Session              bool                  `json:"session"`
	Parameters           []Parameter           `json:"parameter"`
	EnvironmentVariables []EnvironmentVariable `json:"environment_variables"`
	Break                bool                  `json:"break"`
}

type Commands struct {
	Commands []Command `json:"commands"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response Commands
	_ = json.NewDecoder(r.Body).Decode(&response)
	json.NewEncoder(w).Encode(response)
	format(response)
}

func format(commands Commands) {
	for _, command := range commands.Commands {
		env := map[string]string{}
		for _, envVar := range command.EnvironmentVariables {
			env[envVar.Name] = envVar.Value
		}
		rgx := regexp.MustCompile(`\{([^}]+)\}`)
		out := rgx.FindAllStringSubmatch(command.Command, -1)

		for _, i := range out {
			command.Command = strings.Replace(command.Command, i[0], env[i[1]], -1)
		}
		fmt.Println(command.Command)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/commands", handler).Methods("POST")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":3001", r))

}
