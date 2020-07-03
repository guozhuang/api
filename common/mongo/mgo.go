package mongo

import (
	"api/base"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//没问题，应该将这里的方法统一下沉到provider中，业务要使用，直接调用provider的封装代码，使得对数据源的操作封装实现
//因为本身这里操作数据源就是关联了相应的数据源的结构体。
//这样的实现和redis和mysql的golang实现是一致的【而且也是体现了封装这一golang的编程哲学】

func (mongo *Client) searchAll(query bson.M, results []interface{}) {
	//如何将相应的结果统一处理？
	//var results []User
	expr := func(c *mgo.Collection) error {
		return c.Find(query).All(&results)
	}

	err := mongo.WithMongoCollection(expr)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (mongo *Client) GetAll(query map[string]string, results []interface{}) {
	bsonQuery := sel(query)
	mongo.searchAll(bsonQuery, results)
}

//似乎只能按照需求来实现相应的方法【all，那这样公共类内其实就没必要实现相应的标准方法，因为标准方法已经实现到对应的mgo中】
func (mongo *Client) FindOne(query map[string]string, result *base.User) {
	bsonQuery := sel(query)

	expr := func(c *mgo.Collection) error {
		return c.Find(bsonQuery).One(&result)
	}

	fmt.Println(result)

	err := mongo.WithMongoCollection(expr)
	if err != nil {
		fmt.Println(err.Error())
	}
}

//统一query转为bson
func sel(query map[string]string) bson.M {
	result := make(bson.M, len(query))

	for k, v := range query {
		result[k] = v
	}
	return result
}

/**
 * 执行查询，此方法可拆分做为公共方法
 * [SearchPerson description]
 * @param {[type]} collectionName string [description]
 * @param {[type]} query          bson.M [description]
 * @param {[type]} sort           bson.M [description]
 * @param {[type]} fields         bson.M [description]
 * @param {[type]} skip           int    [description]
 * @param {[type]} limit          int)   (results      []interface{}, err error [description]
 */
/*func SearchPerson(collectionName string, query bson.M, sort string, fields bson.M, skip int, limit int) (results []interface{}, err error) {
	exop := func(c *mgo.Collection) error {
		return c.Find(query).Sort(sort).Select(fields).Skip(skip).Limit(limit).All(&results)
	}
	err = witchCollection(collectionName, exop)
	return
}*/

//todo:参考上面用例来进行mongo的基本操作【同样需要设置接口来进行mock】
//标准构造操作方法
