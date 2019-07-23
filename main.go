package main

import "go-hero/route"

func main() {
	router := route.Init()
	router.Logger.Fatal(router.Start(":8888"))
}
