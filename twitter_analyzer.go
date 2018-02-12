package main

import (
	"fmt"
	"log"

	database "./database"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	session, err := database.GetSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	//Switching the session to monotonic behavior. Not necessary.
	session.SetMode(mgo.Monotonic, true)

	//fmt.Println("Counting...")
	//fmt.Println(database.GetUserCount())
	fmt.Println("Checking...")
	database.GetTopTaggers()

}
