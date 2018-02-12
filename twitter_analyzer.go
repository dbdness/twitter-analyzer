package main

import (
	"fmt"
	"log"
	"os"

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

	userCommand := os.Args[1]

	switch userCommand {
	case "--CountUsers":
		//1.
		fmt.Println("Counting users...")
		fmt.Println(database.GetUserCount())
	case "--TopTaggers":
		//2.
		fmt.Println("Counting top taggers...")
		database.GetTopTaggers()
	case "--MostTagged":
		//3.
		fmt.Println("Counting most tagged users...")
		database.GetMostTagged()
	case "--MostActive":
		//4.
		fmt.Println("Counting most active users...")
		database.GetMostActive()
	case "--Grumpiest":
		//5
		fmt.Println("Getting grumpiest users...")
		database.GetGrumpiest()
	case "--Happiest":
		//6.
		fmt.Println("Getting happiest users...")
		database.GetHappiest()
	default:
		fmt.Println("Unknown command: " + userCommand)
		printUsage()

	}

}

func printUsage() {
	fmt.Println("usage:\n--CountUsers\n--TopTaggers\n--MostTagged\n--MostActive\n--Grumpiest\n--Happiest")
}
