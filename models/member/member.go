package member

// Member model
type Member struct {
	FirstName string `bson:"firstName" json:"firstName"`
	LastName  string `bson:"lastName" json:"lastName"`
	Username  string `bson:"username" json:"username"`
	Password  string `bson:"password" json:"password"`
}
