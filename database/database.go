package database

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2/bson"

	mgo "gopkg.in/mgo.v2"
)

const hostURL = "mongodb://localhost:27017"
const dbName = "social_net"
const collName = "tweets"

var db *mgo.Database
var coll *mgo.Collection
var session *mgo.Session

//Tweet is the main structure for the type of tweet in the database.
type Tweet struct {
	ID       int    `bson:"id"`
	Polarity int    `bson:"polarity"`
	Date     string `bson:"date"`
	Query    string `bson:"query"`
	User     string `bson:"user"`
	Text     string `bson:"text"`
}

func init() {
	var err error
	session, err = GetSession()
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(dbName)
	coll = db.C(collName)

}

//GetSession returns the mgo MongoDB session.
func GetSession() (*mgo.Session, error) {
	var err error
	session, err = mgo.Dial(hostURL)
	if err != nil {
		log.Fatal(err)
	}

	return session, err
}

//GetUserCount returns the amount of users in the database.
func GetUserCount() int {
	var result []string
	err := coll.Find(nil).Distinct("user", &result)
	if err != nil {
		log.Fatal(err)
	}

	return len(result)

}

//GetTopTaggers returns a list of the users who has tagged the most people in their tweets.
func GetTopTaggers() {
	//var result []Tweet
	var result []bson.M
	//var result []string

	//err := coll.Find(bson.M{"text": bson.M{"$regex": bson.RegEx{`@\w`, ""}}})
	//pipeline := []bson.M{
	//	{"$match": bson.M{"text": bson.M{"$regex": bson.RegEx{`@\w`, ""}}}

	pipeline := []bson.M{
		{"$match": bson.M{"text": bson.M{"$regex": bson.RegEx{`@\w`, ""}}}},
		{"$group": bson.M{"_id": "$user",
			"matches": bson.M{"$sum": 1},
		},
		},
		{"$sort": bson.M{"matches": -1}}, //1: Ascending, -1: Descending
		{"$limit": 5},
	}

	/*
		pipeline := []bson.M{
			{"$match": bson.M{"text": bson.M{"$regex": bson.RegEx{`/@\w+\/`, ""}}}},
			{"$group": bson.M{"_id": "null",
				"text": bson.M{"$push": "$text"},
			},
			},
			{"$sort": bson.M{"user": 1}}, //1: Ascending, -1: Descending

		}
	*/

	err := coll.Pipe(pipeline).All(&result)
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range result {
		fmt.Println(user)

	}
}
