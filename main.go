package main

import (
	"fmt"
	"os"
	"polyadip/db"
	"polyadip/models"
	"polyadip/router"
)

func main() {
	e := router.Init()
	conn, err := db.Connect()
	if err != nil {
		fmt.Println(err)
		os.Exit(4)
	}
	models.ConnectDB(conn)
	port := fmt.Sprintf(":%v", "1400")
	e.Logger.Fatal(e.Start(port))
}
