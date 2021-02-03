package shops

import (
	"encoding/json"
	"fmt"
	"net/http"
	"serverInGo/database"
	"serverInGo/types"
)

func queryShops(shops *types.Shops, AID int) error {
	QueryString := fmt.Sprintf("SELECT id,name,status,type,description FROM shopdata where accountid=$1")
	shopRows, err := database.Db.Query(QueryString, AID)
	if err != nil {
		return err
	}
	defer shopRows.Close()
	// rows := []types.Shop{}
	for shopRows.Next() {
		// fmt.Printf("------> \n %v", rows)
		row := types.Shop{}
		err := shopRows.Scan(
			&row.ID,
			&row.Name,
			&row.Status,
			&row.Type,
			&row.Description,
		)
		if err != nil {
			return err
		}
		shops.Shops = append(shops.Shops, row)
	}
	err = shopRows.Err()
	if err != nil {
		return err
	}
	return nil

}
func UserShops(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm()
	fmt.Println(r)
	token := r.Header.Get("Token")
	AID, err := getAccountID(token)
	if err != nil {
		panic(err)
	}
	// defer getAcc.Close()
	repos := types.Shops{}
	err = queryShops(&repos, AID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	out, err := json.Marshal(repos)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintf(w, string(out))
	// b, err := json.Marshal(rows)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%v", rows)
	// fmt.Fprintf(w, "%v", rows)

}
