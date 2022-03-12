package entity

import "go-api/db"

type EnvConfig struct {
	Host  string
	Port  int
	Mongo db.MongoConfig
}
