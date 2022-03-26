package adapter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	modeladapter "demoproject/src/adapter/model"
)

func ApiIntegrationGet() modeladapter.User {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}
	var response modeladapter.User
	json.Unmarshal(responseData, &response)
	return response
}
