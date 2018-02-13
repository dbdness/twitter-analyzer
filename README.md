# Assignment 2: Tweet Analyzer 

*Made by Danny Nielsen* 

Details on this assignment can be found [here](https://github.com/datsoftlyngby/soft2018spring-databases-teaching-material/blob/master/lecture_notes/02-Intro_to_MongoDB.ipynb).

This project is written in [Go](golang.org).

## My answers

1. How many Twitter users are in the database?

   ```bash
   659621
   ```

2. Which Twitter users link the most to other Twitter users? (Provide the top ten.)

   ```bash
   1 lost_dog has tagged others: 549 times
   2 tweetpet has tagged others: 310 times
   3 VioletsCRUK has tagged others: 251 times
   4 what_bugs_u has tagged others: 246 times
   5 tsarnick has tagged others: 245 times
   6 SallytheShizzle has tagged others: 229 times
   7 mcraddictal has tagged others: 217 times
   8 Karen230683 has tagged others: 216 times
   9 keza34 has tagged others: 211 times
   10 TraceyHewins has tagged others: 202 times
   ```

3. Who are the most mentioned Twitter users? (Provide the top five.)

   ```bash
   1 @workformeonline
   2 @mileycyrus
   3 @taylorswift13
   4 @ddlovato
   5 @tommcfly
   ```

4. Who are the most active Twitter users (top ten)?

   ```bash
   1 lost_dog has made: 549 tweets
   2 webwoke has made: 345 tweets
   3 tweetpet has made: 310 tweets
   4 SallytheShizzle has made: 281 tweets
   5 VioletsCRUK has made: 279 tweets
   6 mcraddictal has made: 276 tweets
   7 tsarnick has made: 248 tweets
   8 what_bugs_u has made: 246 tweets
   9 Karen230683 has made: 238 tweets
   10 DarkPiano has made: 236 tweets
   ```

5. Who are the five most grumpy (most negative tweets) and the most happy (most positive tweets)? (Provide five users for each group)

   Grumpy:

   ```bash
   1 webwoke - 37 negative tweets.
   2 Spidersamm - 29 negative tweets.
   3 SallytheShizzle - 29 negative tweets.
   4 mr_apollo - 26 negative tweets.
   5 D_AMAZIN - 24 negative tweets.
   ```

   Happy:

   ```bash
   1 caldjr - 61 positive tweets.
   2 VioletsCRUK - 58 positive tweets.
   3 iHomeTech - 57 positive tweets.
   4 Jeff_Hardyfan - 55 positive tweets.
   5 sierrabardot - 50 positive tweets.
   ```

## Using the program

### Prerequisites

1. You need to have Go installed on your system in order to run this program. It's also possible to run it on a virtual machine or a container that has Go installed.

2. You also need to have a MongoDB instance with all the required twitter data running on your localhost, otherwise you will see no results. A guide to do that can also be found in the [project requirements](https://github.com/datsoftlyngby/soft2018spring-databases-teaching-material/blob/master/lecture_notes/02-Intro_to_MongoDB.ipynb).

3. I am using [mgo](https://labix.org/mgo) as the MongoDB driver for Go. It's necessary to install the required dependency before you can run the program:

   ```bash
   $ go get gopkg.in/mgo.v2
   ```

### Building it

1. Open your favorite terminal, and clone this repository to a location you prefer:

   ```bash
   $ git clone https://github.com/dbdness/twitter-analyzer.git
   ```

2. Navigate to the project folder, and build the project like so:

   ```bash
   $ go build .
   ```

3. Use any one of the commands mentioned in the section below to make the program analyze the tweets. Example:

   ```bash
   $ ./twitter_analyzer --Happiest
   ```

The results will be printed to your terminal.

### Commands

```bash
--CountUsers
--TopTaggers
--MostTagged
--MostActive
--Grumpiest
--Happiest
```