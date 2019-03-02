package shared

import (
	"github.com/alexmay23/httputils"
	"github.com/globalsign/mgo/bson"
	"github.com/ti/mdb"
	"net/http"
)

func IsEqual(lhs *bson.ObjectId, rhs *bson.ObjectId) bool {
	if lhs == rhs {
		return true
	} else {
		if lhs == nil || rhs == nil {
			return false
		} else {
			return lhs.Hex() == rhs.Hex()
		}
	}
}


func ObjectIdOrError(r *http.Request, key string)(string, error){
	id := httputils.GetValueFromURLInRequest(r, key)
	if id == nil || !bson.IsObjectIdHex(*id){
		return "", httputils.HTTP400()
	}
	return *id, nil
}

func FindMany(collection *mdb.Collection, parameters bson.M, skip, limit *int)(map[string]interface{}, error){
	var values []bson.M
	query := collection.Find(parameters)
	count, err  := query.Count()
	if err != nil{
		return nil, err
	}
	query = httputils.ApplySkipLimit(query, skip, limit)
	err = query.All(&values)
	if err != nil{
		return nil, err
	}

	result := map[string]interface{}{
		"total": count,
		"objects": values,
	}
	return result, nil
}


func IsEqualStrings(lhs *string, rhs *string) bool {
	if lhs == rhs {
		return true
	} else {
		if lhs == nil || rhs == nil {
			return false
		} else {
			return *lhs == *rhs
		}
	}
}

func SOI(value *bson.ObjectId) *string {
	if value == nil {
		return nil
	}
	r := value.Hex()
	return &r
}

func OIS(value *string) *bson.ObjectId {
	if value == nil {
		return nil
	}
	if !bson.IsObjectIdHex(*value) {
		return nil
	}
	r := bson.ObjectIdHex(*value)
	return &r
}

func ToObjectIdList(slice []string) []bson.ObjectId {
	var items []bson.ObjectId
	for _, item := range slice {
		items = append(items, bson.ObjectIdHex(item))
	}
	return items
}
