package mongo

import (
	"api/base"
	"api/common/mongo"
	"api/config"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

var UserMongoClient = &mongo.Client{}

const (
	UserCollection = "user"
)

func init() {
	//进行mongo基础结构初始化
	UserMongoClient.Collection = UserCollection
	UserMongoClient.IsCluster = false
	UserMongoClient.ServerAddr = config.UserMongoServerSetting.MongodbAddr
}

//mongo基础加载
type Provider struct {
	//
}

//todo：当前包用来隔离指定editor针对mongo处理

func (provider *Provider) GetMongoName(userId string, result *base.User) {
	//进行mongo包处理
	//userInfo :=
	query := map[string]string{
		"user_id": userId,
	}

	UserMongoClient.FindOne(query, result)

	userName := ""

	fmt.Println(userName)
	fmt.Println(result)
}

func Bson2Odj(val interface{}, obj interface{}) error {
	data, err := bson.Marshal(val)
	if err != nil {
		return err
	}
	bson.Unmarshal(data, obj)
	return nil
}
