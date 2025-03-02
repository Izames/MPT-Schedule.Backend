package main

import (
	"MPT-Schedule/DataBase"
	"MPT-Schedule/Roaming"
)

func main() {
	DataBase.Connect()
	DataBase.SetDefault()
	go DataBase.ClearPincodes()
	Roaming.Routes()
}
