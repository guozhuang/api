package editor

import (
	"api/base"
	"encoding/json"
	"fmt"
)

//
type User base.User

func (user *User) Save() {
	jsonStr, err := json.Marshal(user)
	if err != nil {
		fmt.Println("json user is error")
	}
	baseProvider.User.RedisProvider.SetUserInfo(user.UserId, string(jsonStr))

	//进行mongo存储
}

func GetUserEditor(userId string) *User {
	user := new(User)

	editorData := baseProvider.User.RedisProvider.GetUserInfo(userId)

	json.Unmarshal([]byte(editorData), user)
	return user
}

func GetUserEditors(userIds []string) []*User {
	result := make([]*User, 0)

	return result
}
