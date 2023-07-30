package database

import (
	"github.com/yuttasakcom/go-kafka-simple/src/core/config"
)

type Store struct {
	*GormDB
	// *MongoDB
}

func NewStore(config config.DB) *Store {
	return &Store{&GormDB{(GormConnect(config.Pg))}}
	// return &Store{&MongoDB{(MongoConnect(config.Mg))}}
}
