package base

type User struct {
	UserId   string `bson:"user_id,omitempty" json:"user_id, omitempty"`
	UserName string `json:"user_name,omitempty" bson:"user_name,omitempty"`
}
