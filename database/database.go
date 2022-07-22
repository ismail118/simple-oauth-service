package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"simple-oauth-service/config"
	"simple-oauth-service/helper"
	"time"
)

func NewDB(config *config.Config) *sql.DB {
	// ex root:Colonelgila123@tcp(localhost:3306)/auth?parseTime=true
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		config.Database.MySql.User,
		config.Database.MySql.Password,
		config.Database.MySql.Addr,
		config.Database.MySql.DatabaseName,
	)

	db, err := sql.Open("mysql", dataSourceName)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	if err2 := db.Ping(); err2 != nil {
		panic(errors.New("mysql database connection failed"))
	}
	return db
}

func NewRedisClient(config *config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Database.Redis.Addr,
		Password: config.Database.Redis.Password,
		DB:       config.Database.Redis.DB,
	})

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		panic(errors.New("redis database connection failed"))
	}

	return client
}
