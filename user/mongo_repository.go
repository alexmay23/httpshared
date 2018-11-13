package user

import (
	"github.com/ti/mdb"
	"github.com/globalsign/mgo/bson"
	"github.com/alexmay23/httputils"
)

type MongoRepository struct {
	collection *mdb.Collection
}

func (self *MongoRepository) GetByIdList(idList []string) []Model {
	var objectIdList []bson.ObjectId
	for _, value := range idList{
		objectIdList = append(objectIdList, bson.ObjectIdHex(value))
	}
	return self.manyBy(bson.M{"_id": bson.M{"$in": objectIdList}}, nil, nil).Objects
}

type MongoModel struct {
	ID     bson.ObjectId `bson:"_id"`
	Phone  string
	Avatar string
	FBId   string        `bson:"fb_id"`
	Name   string
	Code   int
	Secret string
}

func (self *MongoRepository) manyBy(query bson.M, skip, limit *int) *List {
	list := []MongoModel{}
	qry := self.collection.Find(query)
	total, err := qry.Count()
	err = httputils.ApplySkipLimit(qry, skip, limit).All(&list)
	if err != nil {
		panic(err)
	}
	count := len(list)
	var updated = make([]Model, count)
	for idx, value := range list {
		updated[idx] = ModelFromMongo(&value)
	}
	r := List{Objects: updated[:], Total: total}
	return &r
}

func ModelFromMongo(model *MongoModel) Model {
	return Model{ID: model.ID.Hex(), Phone: model.Phone, Avatar:model.Avatar, Name: model.Name, Code: model.Code, Secret: model.Secret}
}



func NewMongoRepository(database *mdb.Database) *MongoRepository {
	return &MongoRepository{collection: database.C("user")}
}

func (self *MongoRepository) GetByPhone(phone string) *Model {
	return self.oneBy(bson.M{"phone": phone})
}

func (self *MongoRepository) GetByFBId(id string) *Model {
	return self.oneBy(bson.M{"fb_id":id})
}


func (self *MongoRepository) GetById(id string) *Model {
	return self.oneBy(bson.M{"_id": bson.ObjectIdHex(id)})
}

func (self *MongoRepository) oneBy(query bson.M) *Model {
	m := &MongoModel{}
	err := self.collection.Find(query).One(m)
	if err != nil {
		return nil
	}
	o := Model{ID: m.ID.Hex(), Phone: m.Phone, Code: m.Code, Name:m.Name, Avatar:m.Avatar, FBId:m.FBId, Secret:m.Secret}
	return &o
}

func (self *MongoRepository) Update(model *Model) {
	m := MongoModel{bson.ObjectIdHex(model.ID), model.Phone, model.Avatar, model.FBId, model.Name,model.Code, model.Secret}
	err := self.collection.UpdateId(m.ID, m)
	if err != nil {
		panic(err)
	}
}

func (self *MongoRepository) CreateWithPhone(phone string, code int) *Model {
	m := MongoModel{Phone: phone, Code: code, Secret: httputils.RandStringBytes(8)}
	info, err := self.collection.Upsert(bson.M{"phone": phone}, &m)
	if err != nil {
		return nil
	}
	v := Model{info.UpsertedId.(bson.ObjectId).Hex(), m.Phone, "", "","",m.Code, m.Secret}
	return &v
}

func (self *MongoRepository)CreateWithFB(id string, name string, avatar string, phone *string)*Model{

	m := MongoModel{FBId:id, Name:name, Avatar:avatar}
	query := bson.M{"fb_id": id}
	if phone != nil{
		m.Phone = *phone;
		query = bson.M{"$or":[]bson.M{query, {"phone": *phone}}}
	}
	var current MongoModel
	err := self.collection.Find(query).One(&current)
	if err == nil{
		m.ID = current.ID
		m.Secret = current.Secret
	}else{
		m.Secret = httputils.RandStringBytes(8)
		m.ID = bson.NewObjectId()
	}
	_, err = self.collection.Upsert(query, &m)
	if err != nil {
		panic(err)
	}
	v := &Model{m.ID.Hex(), m.Phone, m.Name, m.Avatar, m.FBId,0, m.Secret}
	return v
}
