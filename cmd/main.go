package main

import (
	"context"
	"github.com/julienschmidt/httprouter"
	"go-library/configs"
	"go-library/internal/composite"
	"go-library/pkg/logging"
	"net/http"
)

func main() {
	cfg := config.GetConfig()
	log := logger.GetLogger()

	mysqlComposite, err := composite.NewMysqlComposite(
		context.Background(),
		cfg.MySQLDb.Host,
		cfg.MySQLDb.Port,
		cfg.MySQLDb.User,
		cfg.MySQLDb.Password,
		cfg.MySQLDb.DBName)

	if err != nil {
		log.Error.Println("mysql composite failed")
	}

	router := httprouter.New()

	authorComposite, err := composite.NewAuthorComposite(mysqlComposite)
	if err != nil {
		log.Error.Println("author composite failed")
	}
	authorComposite.Handler.Register(router)

	bookComposite, err := composite.NewBookComposite(mysqlComposite)
	if err != nil {
		log.Error.Println("author composite failed")
	}
	bookComposite.Handler.Register(router)

	readerComposite, err := composite.NewReaderComposite(mysqlComposite)
	if err != nil {
		log.Error.Println("reader composite failed")
	}
	readerComposite.Handler.Register(router)

	err = http.ListenAndServe(":8088", router)
	if err != nil {
		return
	}
}
