package editor

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
)

//作为nosql的数据源的实例转换使用
type User struct {
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
}

func (user *User) Save() {
	jsonStr, err := json.Marshal(user)
	if err != nil {
		fmt.Println("json user is error")
	}
	baseProvider.User.RedisProvider.SetUserInfo(user.UserId, string(jsonStr))
}

func GetUserEditor(userId string) *User {
	user := new(User)
	user.UserId = userId
	editorData := baseProvider.User.RedisProvider.GetUserInfo(userId)
	if editorData != "" {
		//进行json解析
		user.UserName = gjson.Get(editorData, "user_name").String()
	}

	return user
}

func GetUserEditors(userIds []string) []*User {
	result := make([]*User, 0)

	return result
}
