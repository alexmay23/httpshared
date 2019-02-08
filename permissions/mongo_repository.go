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


type mongoObject struct {
	UserId bson.ObjectId `bson:"user_id"`
	Permissions []string `bson:"permissions"`
	ID bson.ObjectId `bson:"_id"`
}


func (self *MongoRepository) GetPermissions(userId string) []string {
	var result mongoObject
	query := bson.M{"user_id": bson.ObjectIdHex(userId)}
	err := self.collection.Find(query).One(&result)
	if err != nil{
		_, err := self.collection.Upsert(query, mongoObject{UserId:bson.ObjectIdHex(userId), ID:bson.NewObjectId(), Permissions:[]string{}})
		if err != nil{
			panic(err)
		}
		return []string{}
	}
	return result.Permissions
}

