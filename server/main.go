package main

import (
	"flag"
	"net/http"
	"serverInGo/database"
	"serverInGo/handler"
	"serverInGo/user"

	"github.com/golang/glog"
	_ "github.com/lib/pq"
)

func main() {
	flag.Parse()
	flag.Lookup("logtostderr").Value.Set("true")
	// Initailze Db connection
	database.InitDb()
	// user.GenerateToken("9666334149")
	user.ValidateUser(999)

	http.HandleFunc("/account", handler.CreateAccount)

	glog.Infof("Listening on http://0.0.0.0:3000")

	// go handler.UpdateDatabase()
	glog.Info(http.ListenAndServe(":"+"3000", nil))
}
