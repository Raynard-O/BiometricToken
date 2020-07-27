package main

import (
	"BiometricToken/db"
	"BiometricToken/router"
)

func main() {
	db.DbInit()
	e := router.New()
	e.Logger.Fatal(e.Start(":1111"))
}
