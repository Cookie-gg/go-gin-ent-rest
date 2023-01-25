package infra

import (
	"context"
	"go-gin-ent-rest/ent"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

func InitDB() *ent.Client {
	entOptions := []ent.Option{}
	entOptions = append(entOptions, ent.Debug())

	mc := mysql.Config{
		User:      os.Getenv("MYSQL_USER"),
		Passwd:    os.Getenv("MYSQL_PASSWORD"),
		Net:       "tcp",
		Addr:      os.Getenv("MYSQL_HOST") + ":" + os.Getenv("MYSQL_PORT"),
		DBName:    os.Getenv("MYSQL_DATABASE"),
		ParseTime: true,
		Loc:       time.Local,
	}

	client, err := ent.Open("mysql", mc.FormatDSN(), entOptions...)

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	log.Println("db connected!")

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}