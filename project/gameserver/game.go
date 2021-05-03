package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	"project/gameapi"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
	//mongodbEndpoint = "mongodb://172.17.0.2:27017" // Find this from the Mongo container
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

type gameDatabase struct {
	ID         primitive.ObjectID `bson:"_id"`
	User1      string             `bson:"user1"`
	User2      string             `bson:"user2"`
	Health1    int                `bson:"health1"`
	Health2    int                `bson:"health2"`
	WhoseTurn  string             `bson:"whose_turn"`
	LastAttack string             `bson:"last_attack"`
	GamePort   int                `bson:"game_port"`
	CreatedAt  time.Time          `bson:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at"`
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

type server struct {
	gameapi.UnimplementedGameInfoServer
}

var monsterAttackDB [][]string = [][]string{{"Leaf blade", "Energy ball", "Apple acid", "Tackle"}, {"Flamethrower", "Blaze kick", "Searing shot", "Tackle"}, {"Hydro cannon", "Surf", "Water ball", "Tackle"}}
var attackpower []int = []int{40, 40, 40, 30}
var attackPowerDB map[string]int

var client *mongo.Client

func (s *server) GetHealthPoints(ctx context.Context, in *gameapi.HealthRequest) (*gameapi.HealthPoints, error) {
	name := in.GetName()
	gameID := in.GetGameID()
	healthPoints := &gameapi.HealthPoints{}

	objectID, _ := primitive.ObjectIDFromHex(gameID)

	// select collection from database
	colGames := client.Database("Pokmon").Collection("games")

	// filters for the databases
	gameFilter := bson.M{"_id": bson.M{"$eq": objectID}}

	// find one user and one monster
	var filterGame gameDatabase
	var tempHealth int

	errG := colGames.FindOne(context.TODO(), gameFilter).Decode(&filterGame)
	if errG == nil {
		// found game with user in it
		if name != filterGame.User1 {
			tempHealth = filterGame.Health2
		} else {
			tempHealth = filterGame.Health1
		}

		var check bool = true
		for check == true {
			// wait for the opponent to make its move
			time.Sleep(2 * time.Second)
			colGames.FindOne(context.TODO(), gameFilter).Decode(&filterGame)
			if filterGame.WhoseTurn == name {
				break
			}
		}
		colGames.FindOne(context.TODO(), gameFilter).Decode(&filterGame)
		if name != filterGame.User1 {
			healthPoints.Health = int32(filterGame.Health2)
			tempHealth = tempHealth - filterGame.Health2
		} else {
			healthPoints.Health = int32(filterGame.Health1)
			tempHealth = tempHealth - filterGame.Health1
		}

		healthPoints.WhoseTurn = filterGame.WhoseTurn
		healthPoints.LastAttack = filterGame.LastAttack
		healthPoints.Damage = int32(tempHealth)

		// delete the game if someone has hit zero HP
		if filterGame.Health2 == 0 || filterGame.Health1 == 0 {
			//Perform DeleteOne operation & validate against the error.
			_, err := colGames.DeleteOne(context.TODO(), gameFilter)
			if err == nil {
				// successfully deleted the game from database
				fmt.Println("Someone won the game, Game deleted")
			} else {
				// error happened while deleting the game from the database
				fmt.Println("Someone won the game, Game not deleted")
			}
		}

		return healthPoints, nil
	} else {
		// could not find the game with user in it
		return healthPoints, errors.New("Error getting Health and whose turn")
	}
}

func (s *server) MonsterAttack(ctx context.Context, in *gameapi.MonsterAction) (*gameapi.HealthPoints, error) {
	name := in.GetName()
	action := in.GetAction()
	gameID := in.GetGameID()
	healthPoints := &gameapi.HealthPoints{}

	fmt.Println(name, "  ", action)

	objectID, _ := primitive.ObjectIDFromHex(gameID)

	// select collection from database
	colUsers := client.Database("Pokmon").Collection("users")
	colGames := client.Database("Pokmon").Collection("games")

	// filters for the databases
	user1Filter := bson.M{"user": bson.M{"$eq": name}}
	gameFilter := bson.M{"_id": bson.M{"$eq": objectID}}

	// find one user and one monster
	var filterUser1 userDatabase
	var filterUser2 userDatabase
	var filterGame gameDatabase

	rand.Seed(time.Now().UnixNano())

	errG := colGames.FindOne(context.TODO(), gameFilter).Decode(&filterGame)
	if errG == nil {
		if name != filterGame.User1 {
			user2Filter := bson.M{"user": bson.M{"$eq": filterGame.User1}}
			if errU := colUsers.FindOne(context.TODO(), user1Filter).Decode(&filterUser1); errU == nil {
				if errU := colUsers.FindOne(context.TODO(), user2Filter).Decode(&filterUser2); errU == nil {
					var randMax int

					randMax = elementDamage(filterUser1.Element, filterUser2.Element, action)

					healthLoss := rand.Intn(randMax)
					healthLoss = rand.Intn(randMax)
					healthLoss = rand.Intn(randMax)

					var tempHealth int = filterGame.Health1 - healthLoss

					if tempHealth < 0 {
						tempHealth = 0
					}

					healthPoints.Health = int32(tempHealth)
					// upadter for specified item and price
					updater := bson.M{"$set": bson.M{"health1": tempHealth, "health2": filterGame.Health2, "whose_turn": filterGame.User1, "last_attack": action, "updated_at": time.Now()}}

					//Perform UpdateOne operation & validate against the error.
					_, err := colGames.UpdateOne(context.TODO(), gameFilter, updater)
					colGames.FindOne(context.TODO(), gameFilter).Decode(&filterGame)

					if err == nil {
						healthPoints.WhoseTurn = filterGame.User1
						healthPoints.Health = int32(filterGame.Health1)
						healthPoints.LastAttack = action
						healthPoints.Damage = int32(healthLoss)

						return healthPoints, nil
					} else {
						// error updating the game
					}
				}
			}
		} else {
			user2Filter := bson.M{"user": bson.M{"$eq": filterGame.User2}}
			if errU := colUsers.FindOne(context.TODO(), user1Filter).Decode(&filterUser1); errU == nil {
				if errU := colUsers.FindOne(context.TODO(), user2Filter).Decode(&filterUser2); errU == nil {
					var randMax int

					randMax = elementDamage(filterUser1.Element, filterUser2.Element, action)

					healthLoss := rand.Intn(randMax)
					healthLoss = rand.Intn(randMax)
					healthLoss = rand.Intn(randMax)

					var tempHealth int = filterGame.Health2 - healthLoss

					if tempHealth < 0 {
						tempHealth = 0
					}

					// upadter for specified item and price
					updater := bson.M{"$set": bson.M{"health1": filterGame.Health1, "health2": tempHealth, "whose_turn": filterGame.User2, "last_attack": action, "updated_at": time.Now()}}

					//Perform UpdateOne operation & validate against the error.
					_, err := colGames.UpdateOne(context.TODO(), gameFilter, updater)
					colGames.FindOne(context.TODO(), gameFilter).Decode(&filterGame)

					if err == nil {
						healthPoints.WhoseTurn = filterGame.User2
						healthPoints.Health = int32(filterGame.Health2)
						healthPoints.LastAttack = action
						healthPoints.Damage = int32(healthLoss)

						return healthPoints, nil
					} else {
						// error updating the game
					}
				}
			}
		}
	} else {
		return healthPoints, errors.New("Error getting Health and whose turn")
	}

	return healthPoints, errors.New("Error getting Health and whose turn")
}

func elementDamage(myType string, oppType string, action string) int {
	switch myType {
	case "Grass":
		switch oppType {
		case "Grass":
			return int(attackPowerDB[action])
		case "Fire":
			return int(attackPowerDB[action]) / 2
		case "Water":
			return int(attackPowerDB[action]) * 2
		}
	case "Fire":
		switch oppType {
		case "Grass":
			return int(attackPowerDB[action]) * 2
		case "Fire":
			return int(attackPowerDB[action])
		case "Water":
			return int(attackPowerDB[action]) / 2
		}
	case "Water":
		switch oppType {
		case "Grass":
			return int(attackPowerDB[action]) / 2
		case "Fire":
			return int(attackPowerDB[action]) * 2
		case "Water":
			return int(attackPowerDB[action])
		}
	default:
		return 0
	}

	return 0
}

func main() {
	fmt.Println("Starting Server")

	// create a mongo client
	clientInstance, err := mongo.NewClient(
		options.Client().ApplyURI(mongodbEndpoint),
	)

	// Connect to mongo
	err = clientInstance.Connect(context.TODO())

	// Disconnect
	defer clientInstance.Disconnect(context.TODO())

	client = clientInstance

	// creates map for attack moves and their power level
	attackPowerDB = make(map[string]int)
	var counter int = 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			attackPowerDB[monsterAttackDB[i][j]] = attackpower[counter]
			counter++
		}
		counter = 0
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	gameapi.RegisterGameInfoServer(s, &server{})

	fmt.Println("Server successfully started")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
