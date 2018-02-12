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

	//1.
	//fmt.Println("Counting users...")
	//fmt.Println(database.GetUserCount())

	//2.
	//fmt.Println("Counting top taggers...")
	//database.GetTopTaggers()

	//3.
	//fmt.Println("Counting most tagged users...")
	//database.GetMostTagged()

	//4.
	//fmt.Println("Counting most active users...")
	//database.GetMostActive()

	//5 & 6.
	fmt.Println("Getting grumpiest users...")
	database.GetGrumpiest()
	//fmt.Println("Getting happiest users...")
	//database.GetHappiest()
}
