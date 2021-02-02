package types

type User struct {
	Index    int    `json:"index"`
	UserID   string `json:userID`
	Mobile   int    `json:mobile`
	Token    string `json:token`
	Name     string `json:name`
	Location string `json:location`
	ShopType string `json:shopType`
	OTP      int    `json:"otp"`
}
