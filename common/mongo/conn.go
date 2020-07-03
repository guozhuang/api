package mongo

import (
	"api/config"
	"errors"
	"fmt"
	"gopkg.in/mgo.v2"
	"time"
)

type Client struct {
	Session    *mgo.Session
	Collection string
	ServerAddr string
	IsCluster  bool
}

func (mongo *Client) GetMongoSession() {
	if mongo.Session == nil {
		var err error

		dialInfo := &mgo.DialInfo{
			//todo：后续追加到结构体内进行传递
			Addrs:     []string{config.UserMongoServerSetting.MongodbAddr}, //此处显然可以进行副本集数据的替换
			Direct:    false,
			Timeout:   time.Second * 30,
			Database:  config.UserMongoServerSetting.MongodbName,
			Source:    config.UserMongoServerSetting.MongodbName, //此处的source和db一致
			Username:  config.UserMongoServerSetting.MongodbUser,
			Password:  config.UserMongoServerSetting.MongodbPassword,
			PoolLimit: 4096, // Session.SetPoolLimit
		}

		mongo.Session, err = mgo.DialWithInfo(dialInfo)

		if err != nil {
			fmt.Println(err)
		}
	}
}

func (mongo *Client) WithMongoCollection(s func(*mgo.Collection) error) error {
	mongo.GetMongoSession()
	session := mongo.Session.Clone()
	if session == nil {
		return errors.New("获取mongodb连接失败")
	}
	defer session.Close()

	c := session.DB(config.UserMongoServerSetting.MongodbName).C(mongo.Collection)
	return s(c)
}
