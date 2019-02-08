package permissions

import (
	"github.com/globalsign/mgo/bson"
	"github.com/ti/mdb"
)

type MongoRepository struct {
	collection *mdb.Collection
}

func NewMongoRepository(database *mdb.Database) *MongoRepository {
	return &MongoRepository{collection: database.C("permissions")}
}


func (self *MongoRepository) GetPermissions(userId string) []string {
	var result map[string]interface{}
	query := bson.M{"user_id": bson.ObjectIdHex(userId)}
	err := self.collection.Find(query).One(&result)
	if err != nil{
		_, err := self.collection.Upsert(query, bson.M{"_id": bson.NewObjectId(), "permissions": []string{},
			"user_id": bson.ObjectIdHex(userId)})
		if err != nil{
			panic(err)
		}
		return []string{}
	}
	return result["permissions"].([]string)
}

