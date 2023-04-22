package domain

type Seller struct {
	UserId string `json:"userId"`

	User User `gorm:"references:ID"`
}
