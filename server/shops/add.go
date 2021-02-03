package shops

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"serverInGo/database"
	"serverInGo/types"
)

func saveShop(s types.Shop) error {
	QueryString := fmt.Sprintf("insert into shopdata(name,status,accountid,type,description) values($1,$2,$3,$4,$5)")
	_, err := database.Db.Query(QueryString, s.Name, s.Status, s.AccountID, s.Type, s.Description)
	if err != nil {
		// panic(err)
		return err
	}
	return nil
}
func getAccountID(token string) (int, error) {
	fmt.Printf("-----token-----", token)
	var accountID int
	QueryString := fmt.Sprintf("select index from userdata where token=$1")
	err := database.Db.QueryRow(QueryString, token).Scan(&accountID)
	if err != nil {
		// return 0, err
		panic(err)
	}
	fmt.Printf("\n -----AccountID-----", accountID)
	return accountID, nil
}
func AddShop(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// fmt.Println("----ACCOUNTID---", r.Header.Get("Token"))
	if r.Header["Token"] != nil {

		token := r.Header.Get("Token")
		accountID, err := getAccountID(token)
		if err != nil {
			panic(err)
		}

		fmt.Println("----ACCOUNTID---", accountID)
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		var body types.Shop
		err = json.Unmarshal(b, &body)
		if err != nil {
			fmt.Println("errrrrrr", err)
		}

		var shop types.Shop

		shop = types.Shop{
			Name:        body.Name,
			Status:      body.Status,
			AccountID:   accountID,
			Type:        body.Type,
			Description: body.Description,
		}
		err = saveShop(shop)
		if err != nil {
			fmt.Errorf("Error on saving data", err.Error())
			panic(err)
		}
		fmt.Fprintf(w, "successfully shopdata saved .. ! \n %v", shop)

	}
	fmt.Fprintf(w, "invalid token ID .. ! \n ")
	// var shop types.Shop

	// shop = types.Shop{
	// 	Name:        "test",
	// 	Status:      "active",
	// 	AccountID:   accountID,
	// 	Type:        "general store",
	// 	Description: "medium scale store",
	// }

	// QueryString := fmt.Sprintf("insert into shopdata(itemgroupid,name,rating,status,accountid,type,description) values($1,$2,$3,$4,$5,$6,$7)")
	// // createTable := "create table shopdata(id INT,itemgroupid VARCHAR, name VARCHAR,rating VARCHAR,status VARCHAR,accountid VARCHAR)"
	// _, err := database.Db.Query(QueryString, shop.ItemGroupID, shop.Name, shop.Rating, shop.Status, shop.AccountID)
	// if err != nil {
	// 	panic(err)
	// }
}
