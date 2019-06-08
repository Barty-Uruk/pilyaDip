package db

import (
	"fmt"
	"polyadip/models"

	pg "github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/sirupsen/logrus"
)

var (
	modelsList   []interface{}
	User         = "barman"
	Password     = "ba4man80"
	DatabaseName = "polyadb"
)

func init() {
	modelsList = append(modelsList, (*models.User)(nil))
	modelsList = append(modelsList, (*models.Ad)(nil))
}

// Connect return DB connection
func Connect() (*pg.DB, error) {
	var (
		conn *pg.DB
		n    int
	)
	addr := fmt.Sprintf("%s:%v", "192.168.0.103", 6001)
	conn = pg.Connect(&pg.Options{
		Addr:     addr,
		User:     User,
		Password: Password,
		Database: DatabaseName,
	})
	_, err := conn.QueryOne(pg.Scan(&n), "SELECT 1")
	if err != nil {
		return conn, fmt.Errorf("Error conecting to DB. addr: %s, user: %s, db: %s,%v", addr, User, DatabaseName, err)
	}

	if err := createSchema(conn); err != nil {
		return conn, fmt.Errorf("Error creating DB schemas. %v", err)
	}
	return conn, nil
}

// CloseDbConnection closing connection for defer in main
func CloseDbConnection(db *pg.DB) {
	db.Close()
}

func createSchema(db *pg.DB) error {
	logrus.Info("Creatind tables if not exist...")
	for _, m := range modelsList {
		err := db.CreateTable(m, &orm.CreateTableOptions{
			IfNotExists:   true,
			FKConstraints: false,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
