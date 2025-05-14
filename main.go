package main

import (
	"MPT-Schedule/DataBase"
	"MPT-Schedule/Roaming"
	"log"
)

func main() {
	DataBase.Connect()
	if !DataBase.DB.Migrator().HasTable("users") {
		DataBase.RunMigrations()
		log.Println("tables created")
	}
	DataBase.SetDefault()
	go DataBase.ClearPincodes()
	Roaming.Routes()
}
