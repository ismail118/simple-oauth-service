package database

import (
	"database/sql"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"simple-oauth-service/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:Colonelgila123@tcp(localhost:3306)/auth?parseTime=true")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return client
}
