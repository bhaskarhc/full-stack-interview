package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"serverInGo/database"
	"serverInGo/types"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// createUser();
// login();

/* login user by his mobile number
- enter mobile number, get OTP from his/her number
- successful login create a jwt token
- using JWT token they can access user credentials .
*/
// Create a struct that models the structure of a user, both in the request body, and in the DB
type Credentials struct {
	Mobile int `json:"mobile"`
	OTP    int `json:"otp"`
}

type token struct {
	JWT        string `json:"jwt"`
	ExpireTime string `json:"expireTime"`
}
type mobile struct {
	Number int `json:"mobile"`
}

func NewUser(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	b, err := ioutil.ReadAll(r.Body)
	// body := string(b)
	if err != nil {
		panic(err)
	}

	var userNumber mobile
	err = json.Unmarshal(b, &userNumber)
	if err != nil {
		panic(err)
	}
	var user types.User
	err = addUser(userNumber.Number)
	if err != nil {
		panic(err)
	}
	user.Mobile = userNumber.Number
	user.OTP = createOTP(user)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Successfully created user \n %v", user)
}
func addUser(number int) error {
	pipelineQuery := fmt.Sprintf("insert into userdata(mobile) values($1)")
	_, err := database.Db.Query(pipelineQuery, number)
	if err != nil {
		return err
	}
	return nil
}

// func updateUser(user types.User) error {
// 	pipelineQuery := fmt.Sprintf("UPDATE userdata SET ")
// 	_, err := database.Db.Query(pipelineQuery, user.Mobile, user.Token)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}
func createOTP(u types.User) int {

	randomOTP := random(0000, 9999)
	pipelineQuery := fmt.Sprintf("UPDATE userdata SET otp=$1 WHERE mobile=$2")
	_, err := database.Db.Query(pipelineQuery, randomOTP, u.Mobile)
	if err != nil {
		panic(err)
	}
	var genOTPfromDB int
	queryForOTP := fmt.Sprintf("select otp from userdata where mobile=$1")
	_ = database.Db.QueryRow(queryForOTP, u.Mobile).Scan(&genOTPfromDB)

	fmt.Printf("---OTP_DB_", genOTPfromDB)
	return randomOTP

}

func ValidateUser(c Credentials) (string, error) {
	var cred Credentials
	cred = Credentials{
		OTP:    c.OTP,
		Mobile: c.Mobile,
	}
	// pipelineQuery := fmt.Sprintf("select otp from userdata where mobile=$1")
	var otp int
	var token string
	err := database.Db.QueryRow("select otp from userdata where mobile=$1", cred.Mobile).Scan(&otp)
	if err != nil {
		panic(err)
	}
	// fmt.Println("DBOTP : %d \n APIActual : %d", otp, cred.OTP)

	if otp == cred.OTP {
		token, err = CreateToken(cred.Mobile)
		if err != nil {
			panic(err)
		}

		// return token
	}
	fmt.Println("\n \t  Token :  ", token)
	return token, nil

}

var mySigningKey = []byte("dukaanlogintoken")

func CreateToken(mobile int) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = mobile
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString([]byte(fmt.Sprint(mobile)))

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	_, err = database.Db.Query("update userdata set token=$1 where mobile=$2", tokenString, mobile)
	if err != nil {
		panic(err)
	}
	fmt.Println(tokenString)
	return tokenString, nil
}

func GenToken(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// w.WriteHeader(http.StatusOK)
	// fmt.Fprintf(w, "Category: %v\n", vars["category"])

	r.ParseForm()
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
	}
	b, err := ioutil.ReadAll(r.Body)
	// body := string(b)
	if err != nil {
		panic(err)
	}
	var cred Credentials
	err = json.Unmarshal(b, &cred)
	if err != nil {
		panic(err)
	}
	var JWTToken token
	JWTToken.JWT, err = ValidateUser(cred)
	if err != nil {

	}
	fmt.Fprintf(w, " {token : %v}", cred)
	fmt.Fprintf(w, " {token : %v}", JWTToken)

}
