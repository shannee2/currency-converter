package dao

import (
	"context"
	"log"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var (
	db     *dynamodb.Client
	dbOnce sync.Once
)

func InitDB() *dynamodb.Client {
	dbOnce.Do(func() {
		cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithEndpointResolver(aws.EndpointResolverFunc(
			func(service, region string) (aws.Endpoint, error) {
				return aws.Endpoint{URL: "http://localhost:8000"}, nil // Use for local DynamoDB
			},
		)))
		if err != nil {
			log.Fatalf("Failed to load AWS config: %v", err)
		}

		db = dynamodb.NewFromConfig(cfg)
		log.Println("DynamoDB connected successfully!")
	})

	return db
}

func GetDB() *dynamodb.Client {
	return db
}

//package dao
//
//import (
//	"database/sql"
//	"fmt"
//	"log"
//	"sync"
//
//	_ "github.com/lib/pq"
//)
//
//const (
//	host     = "localhost"
//	port     = 5432
//	user     = "shannee_ahirwar_ftc"
//	password = "postgres"
//	dbname   = "currency"
//)
//
//var (
//	db     *sql.DB
//	dbOnce sync.Once
//)
//
//func InitDB() (*sql.DB, error) {
//	var err error
//	dbOnce.Do(func() {
//		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
//			"password=%s dbname=%s sslmode=disable",
//			host, port, user, password, dbname)
//
//		db, err = sql.Open("postgres", psqlInfo)
//		if err != nil {
//			log.Fatalf("Failed to connect to database: %v", err)
//			return
//		}
//
//		if err = db.Ping(); err != nil {
//			log.Fatalf("Failed to ping database: %v", err)
//			return
//		}
//
//		log.Println("Database connected successfully!")
//	})
//
//	return db, err
//}
//
//func GetDB() *sql.DB {
//	return db
//}
