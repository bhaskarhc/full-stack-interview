package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"serverInGo/database"
)

type mobile struct {
	Number string `json:"mobile"`
}

//CreateAccount will create a user data
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// body :=
	if r.Method == "GET" {
		userMobile := r.Form.Get("user")
		pipelineQuery := fmt.Sprintf("SELECT * FROM accounts WHERE mobile=" + userMobile)
		row, err := database.Db.Query(pipelineQuery)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("row", row)

	}

	if r.Method == "POST" {
		b, err := ioutil.ReadAll(r.Body)
		// body := string(b)
		if err != nil {
			panic(err)
		}
		var msg mobile
		err = json.Unmarshal(b, &msg)
		if err != nil {
			fmt.Println("errrrrrr", err)
		}
		fmt.Println("msg", msg.Number)

		insertData(msg.Number)

	}

	if r.Method == "PUT" {
		b, err := ioutil.ReadAll(r.Body)
		// body := string(b)
		if err != nil {
			panic(err)
		}
		var msg mobile
		err = json.Unmarshal(b, &msg)
		if err != nil {
			fmt.Println("errrrrrr", err)
		}
		updateUserDetails(msg.Number)

	}
}

func insertData(b string) {
	// PrimaryId := string(b)

	fmt.Println(string(b))
	pipelineQuery := fmt.Sprintf("insert into accounts(mobile,shoptype) values (%s , 'init')", b)
	row, err := database.Db.Query(pipelineQuery)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(row)

}

func updateUserDetails(n string) {

}
