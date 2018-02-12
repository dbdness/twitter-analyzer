package database

import (
	"fmt"
	"log"
	"regexp"

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

//GetTopTaggers returns a top-10 of the users who has tagged the most people in their tweets.
//It matches and grabs all tweets that start with a '@'. Afterwards it groups the matches together by unique user,
//and puts a 'sum' value on each user. It iterates that sum value each time a user matches, and has tagged another user.
//Finally, it sorts descending and limits the query to 5 results.
func GetTopTaggers() {
	var result []bson.M

	pipeline := []bson.M{
		{"$match": bson.M{"text": bson.M{"$regex": bson.RegEx{`@\w`, ""}}}},
		{"$group": bson.M{"_id": "$user",
			"matches": bson.M{"$sum": 1},
		},
		},
		{"$sort": bson.M{"matches": -1}}, //1: Ascending, -1: Descending
		{"$limit": 10},
	}

	err := coll.Pipe(pipeline).All(&result)
	if err != nil {
		log.Fatal(err)
	}

	for i, user := range result {
		fmt.Println(i+1, user["_id"], "has tagged others:", user["matches"], "times")

	}
}

//GetMostTagged returns a top-5 of the most tagged users.
func GetMostTagged() {
	var result []bson.M

	pipeline := []bson.M{
		{"$match": bson.M{"text": bson.M{"$regex": bson.RegEx{`@\w+`, ""}}}},
		{"$group": bson.M{"_id": "$text",
			"matches": bson.M{"$sum": 1},
		},
		},
		{"$sort": bson.M{"matches": -1}}, //1: Ascending, -1: Descending
		{"$limit": 5},
	}

	regEx, _ := regexp.Compile(`@\w+`)

	err := coll.Pipe(pipeline).AllowDiskUse().All(&result)
	if err != nil {
		log.Fatal(err)
	}

	for i, user := range result {
		fmt.Println(i+1, regEx.FindString(user["_id"].(string)))

	}

}

//GetMostActive returns 10 the most active twitter users based on numbers of tweets.
//It simply counts every tweet by each unique username.
func GetMostActive() {
	var result []bson.M

	pipeline := []bson.M{
		{"$match": bson.M{"user": bson.M{"$regex": bson.RegEx{`.*`, ""}}}}, //This RegEx grabs everything in the 'user'-field.
		{"$group": bson.M{"_id": "$user",
			"matches": bson.M{"$sum": 1},
		},
		},
		{"$sort": bson.M{"matches": -1}}, //1: Ascending, -1: Descending
		{"$limit": 10},
	}

	err := coll.Pipe(pipeline).All(&result)
	if err != nil {
		log.Fatal(err)
	}

	for i, user := range result {
		fmt.Println(i+1, user["_id"], "has made:", user["matches"], "tweets")

	}

	fmt.Println(result)
}

//GetGrumpiest returns the five most grumpy/angry/sad/negative users, based on their wording.
func GetGrumpiest() {
	var result []bson.M

	const negativeWords = "(shit|fuck|damn|bitch|crap|piss|dick|darn|asshole|bastard|douche|sad|angry|stupid)"

	pipeline := []bson.M{
		{"$match": bson.M{"text": bson.M{"$regex": bson.RegEx{negativeWords, ""}}}},
		{"$group": bson.M{"_id": "$user",
			"matches": bson.M{"$sum": 1},
		},
		},
		{"$sort": bson.M{"matches": -1}}, //1: Ascending, -1: Descending
		{"$limit": 5},
	}

	err := coll.Pipe(pipeline).All(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Most negative tweeters:")
	for i, user := range result {
		fmt.Println(i+1, user["_id"], "-", user["matches"], "negative tweets.")
	}

}

//GetHappiest returns the five most happy/glad/positive users, based on their wording.
func GetHappiest() {
	var result []bson.M

	const positiveWords = "(love|happy|amazing|beautiful|yay|joy|pleasure|smile|win|winning|smiling|healthy|delight|paradise|positive|fantastic|blessed|splendid|sweetheart|great|funny)"

	pipeline := []bson.M{
		{"$match": bson.M{"text": bson.M{"$regex": bson.RegEx{positiveWords, ""}}}},
		{"$group": bson.M{"_id": "$user",
			"matches": bson.M{"$sum": 1},
		},
		},
		{"$sort": bson.M{"matches": -1}}, //1: Ascending, -1: Descending
		{"$limit": 5},
	}

	err := coll.Pipe(pipeline).All(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Most positive tweeters:")
	for i, user := range result {
		fmt.Println(i+1, user["_id"], "-", user["matches"], "positive tweets.")
	}
}
