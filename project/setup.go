package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongodbEndpoint = "mongodb://192.168.0.174:30953" // Find this from the Mongo container
)

type userDatabase struct {
	ID          primitive.ObjectID `bson:"_id"`
	User        string             `bson:"user"`
	Monster     string             `bson:"monster"`
	AttackMoves []string           `bson:"attack_moves"`
	Health      int                `bson:"health"`
	Element     string             `bson:"element"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

type monsterDatabase struct {
	ID          primitive.ObjectID `bson:"_id"`
	Monster     string             `bson:"monster"`
	AttackMoves []string           `bson:"attack_moves"`
	Health      int                `bson:"health"`
	Element     string             `bson:"element"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

var monsterNamesDB []string = []string{"Bulbasaur", "Charmander", "Squirtle", "Chikorita", "Cyndaquil", "Totodile", "Treecko", "Torchic", "Mudkip", "Turtwig", "Chimchar", "Piplup"}
var monsterAttackDB [][]string = [][]string{{"Leaf blade", "Energy ball", "Apple acid", "Tackle"}, {"Flamethrower", "Blaze kick", "Searing shot", "Tackle"}, {"Hydro cannon", "Surf", "Water ball", "Tackle"}}
var monsterHealthDB []int = []int{90, 78, 88, 90, 78, 88, 80, 90, 100, 110, 88, 106} // twice the amount they had in pokemon
var monsterElementDB []string = []string{"Grass", "Fire", "Water"}
var attackpower []int = []int{40, 40, 40, 30}
var attackPowerDB map[string]int

var client *mongo.Client

func main() {
	// create a mongo client
	clientInstance, err := mongo.NewClient(
		options.Client().ApplyURI(mongodbEndpoint),
	)

	// Connect to mongo
	err = clientInstance.Connect(context.TODO())
	if err == nil {
		// do something here
	} else {
		// do something here
	}

	// Disconnect
	defer clientInstance.Disconnect(context.TODO())

	client = clientInstance

	// select collection from database
	colMonsters := client.Database("Pokmon").Collection("monsters")
	colUsers := client.Database("Pokmon").Collection("users")
	colGames := client.Database("Pokmon").Collection("games")

	input := bufio.NewScanner(os.Stdin)
	var number string

	// connecting to the server with a certain username
	fmt.Println("1000: setup monster database\n2: delete users and games\n45: delete games\n67: show monsters and moves\n1234567: delete everything\n33: delete monsters\n25: show users in database")
	input.Scan()
	number = input.Text()

	switch number {
	case "1000":
		// initialize the monster database
		for i := 0; i < 12; i++ {
			// Insert one
			colMonsters.InsertOne(context.TODO(), &monsterDatabase{
				ID:          primitive.NewObjectID(),
				Monster:     monsterNamesDB[i],
				AttackMoves: monsterAttackDB[i%3],
				Health:      monsterHealthDB[i],
				Element:     monsterElementDB[i%3],
				CreatedAt:   time.Now(),
			})
		}

		foundMonsters, err := colMonsters.Find(context.TODO(), bson.D{})
		if err != nil {
			fmt.Printf("Error listing monsters")
		}
		//Map result to varaible
		for foundMonsters.Next(context.TODO()) {
			monsters := monsterDatabase{}
			err := foundMonsters.Decode(&monsters)
			if err != nil {
				fmt.Printf("Error listing monster")
			} else {
				fmt.Printf("Monster: %q\tAttacks: %q\n", monsters.Monster, monsters.AttackMoves)
			}
		}
	case "2":
		// Delete everything in the users and games database
		colUsers.DeleteMany(context.TODO(), bson.D{})
		colGames.DeleteMany(context.TODO(), bson.D{})
	case "45":
		// Delete everything in the games database
		colGames.DeleteMany(context.TODO(), bson.D{})
	case "67":
		foundMonsters, err := colMonsters.Find(context.TODO(), bson.D{})
		if err != nil {
			fmt.Printf("Error listing monsters")
		}
		//Map result to varaible
		for foundMonsters.Next(context.TODO()) {
			monsters := monsterDatabase{}
			err := foundMonsters.Decode(&monsters)
			if err != nil {
				fmt.Printf("Error listing monster")
			} else {
				fmt.Printf("Monster: %q\tAttacks: %q\n", monsters.Monster, monsters.AttackMoves)
			}
		}
	case "1234567":
		// Delete everything in the users and games and monsters database
		colUsers.DeleteMany(context.TODO(), bson.D{})
		colGames.DeleteMany(context.TODO(), bson.D{})
		colMonsters.DeleteMany(context.TODO(), bson.D{})
	case "33":
		// Delete everything in the monsters database
		colMonsters.DeleteMany(context.TODO(), bson.D{})
	case "25":
		foundUsers, err := colUsers.Find(context.TODO(), bson.D{})
		if err != nil {
			fmt.Printf("Error listing monsters")
		}
		//Map result to varaible
		for foundUsers.Next(context.TODO()) {
			users := userDatabase{}
			err := foundUsers.Decode(&users)
			if err != nil {
				fmt.Printf("Error listing monster")
			} else {
				fmt.Printf("User name: %q\tMonsteer: %q\n", users.User, users.Monster)
			}
		}
	}

}
