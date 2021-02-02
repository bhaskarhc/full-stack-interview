package user

import (
	"fmt"
	"math/rand"
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
	mobile int `json:"mobile"`
	OTP    int `json:"otp"`
}

type token struct {
	JWT        string `json:"jwt"`
	ExpireTime string `json:"expireTime"`
}

func NewUser() {
	var number int
	number = 909090
	var user types.User
	user.Mobile = number
	user.OTP = createOTP(user.Mobile)

	err := saveUser(user)
	if err != nil {
		panic(err)
	}
}
func saveUser(user types.User) error {
	pipelineQuery := fmt.Sprintf("insert into userdata(mobile,otp) values($1,$2)")
	_, err := database.Db.Query(pipelineQuery, user.Mobile, user.Token)
	if err != nil {
		return err
	}
	return nil
}

func ValidateUser(num int) {

	var cred Credentials
	cred = Credentials{
		OTP:    1234,
		mobile: 9666334149,
	}
	// pipelineQuery := fmt.Sprintf("select otp from userdata where mobile=$1")
	var counter int
	database.Db.QueryRow("select otp from userdata where mobile=999").Scan(&counter)
	fmt.Println("---------------------****** ", counter)
	// _, err := database.Db.Query(pipelineQuery, cred.mobile)
	// if err != nil {
	// 	panic(err)
	// }
	CreateToken(cred.mobile)

	// createOTP(cred.mobile)
	// validateOTP()
	// random(1000, 2000)
	fmt.Println("----------------------------", cred)

}
func random(min int, max int) int {
	return rand.Intn(max-min) + min
}
func createOTP(otp int) int {

	pipelineQuery := fmt.Sprintf("UPDATE userdata SET token=$1 WHERE mobile=$2 RETURNING token")
	row, err := database.Db.Query(pipelineQuery, otp, 999)
	if err != nil {
		panic(err)
	}
	fmt.Println("row", row)
	return otp

}

var mySigningKey = []byte("dukaanlogintoken")

func CreateToken(mobile int) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = mobile
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	fmt.Println(tokenString)
	return tokenString, nil
}
