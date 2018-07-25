package shared

import "github.com/ti/mdb"

func Connect(url string) *mdb.Database {
	db, err := mdb.Dial(url)
	if err != nil {
		panic(err)
	}
	return db
}
