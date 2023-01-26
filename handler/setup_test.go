package handler_test

import (
	"go-gin-ent-rest/ent"
	"go-gin-ent-rest/infra"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
)

var (
	clientMock *ent.Client
	sqlMock    sqlmock.Sqlmock
	err        error
)

func init() {
	clientMock, sqlMock, err = infra.InitDBMock()
	if err != nil {
		log.Fatal(err)
	}
}
