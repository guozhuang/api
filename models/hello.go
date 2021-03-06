package models

import (
	"api/base"
	"api/editor"
)

type HelloModel struct {
	//
}

func (hello *HelloModel) GetHelloData(id string) string {
	/*id += "aaaa"

	//model需要各自层面实现对应provider的实例化，在当前包内有效还是全局有效，都需要进行判断
	//以及连接池的初始化
	id = baseProvider.User.RedisProvider.GetUserName(id)

	id = baseProvider.Test.RedisProvider.GetUserName(id)*/

	result := &base.User{}
	baseProvider.User.MongoProvider.GetMongoName(id, result)

	return result.UserName
}

//对editor的使用
func (hello *HelloModel) GetUserName(id string) string {
	userEditor := editor.GetUserEditor(id)
	userEditor.UserName = "save"
	userEditor.Save()

	userEditor = editor.GetUserEditor(id)
	return userEditor.UserName
}

//test mongo
func (hello *HelloModel) GetMongoData(id string) string {
	/*info := baseProvider.User.MongoProvider.GetMongoName(id)
	fmt.Println(info)*/
	return id
}
