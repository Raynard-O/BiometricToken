package main

import "BiometricToken/router"

func main() {
	e := router.New()
	e.Logger.Fatal(e.Start(":1111"))
}
