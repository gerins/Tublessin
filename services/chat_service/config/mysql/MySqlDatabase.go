package mysql

import (
	"database/sql"
	"fmt"
	"tublessin/services/chat_service/config"

	log "github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDatabase() *sql.DB {
	db, err := sql.Open(config.DbDriver, config.DbUser+":"+config.DbPass+"@tcp("+config.DbHost+":"+config.DbPort+")/"+config.DbName)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Print(err)
		fmt.Scanln()
		log.Fatal(err)
	}
	log.Println("DataBase Successfully Connected")
	return db
}
