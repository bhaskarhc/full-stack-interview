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

type Shop struct {
	ID          int    `json:"id"`
	AccountID   int    `json:"accountID"`
	ItemGroupID string `json:"itemGroupID"`
	Location    string `json:"location"`
	Rating      int    `json:"rating"`
	Status      string `json:"status"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
}
type Shops struct {
	Shops []Shop `json:"shops"`
}
