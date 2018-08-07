package shared

import "gopkg.in/mgo.v2/bson"

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
