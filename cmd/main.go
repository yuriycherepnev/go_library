/* TODO
в гайде по чистой архитектуре:
https://www.youtube.com/watch?v=PqQyCFygiZg&t=16s

мы подошли к моменту подключения базы данных (10 минута)
в видео используется монго дб,
мой код надо будет адаптировать под mysql
*/

package main

import (
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	config "go-library/configs"
	"go-library/internal/composite"
)

func main() {
	cfg := config.GetConfig()
	mysqlComposite, err := composite.NewMysqlComposite(
		context.Background(),
		cfg.MySQLDb.Host,
		cfg.MySQLDb.Port,
		cfg.MySQLDb.User,
		cfg.MySQLDb.Password,
		cfg.MySQLDb.DBName)

	if err != nil {
		fmt.Println("aloha")
	}

	router := httprouter.New()

	authorComposite, err := composite.NewAuthorComposite(mysqlComposite)

	if err != nil {
		fmt.Println("aloha")
	}

	bookComposite, err := composite.NewBookComposite(mysqlComposite)

}
