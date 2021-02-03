package main

import (
	"flag"
	"net/http"
	"serverInGo/database"
	"serverInGo/shops"
	"serverInGo/user"

	// _ "serverInGo/handler"

	"github.com/golang/glog"
	_ "github.com/lib/pq"
)

func main() {
	flag.Parse()
	flag.Lookup("logtostderr").Value.Set("true")
	// Initailze Db connection
	database.InitDb()
	http.HandleFunc("/token", user.GenToken)
	http.HandleFunc("/user/new", user.NewUser)
	http.HandleFunc("/shops", shops.UserShops)
	http.HandleFunc("/shop/add", shops.AddShop)

	glog.Infof("Listening on http://0.0.0.0:3000")

	// go handler.UpdateDatabase()
	glog.Info(http.ListenAndServe(":"+"3000", nil))
}
