package main

import (
	"context"
	"errors"
	"log"
	"net"
	"time"
	"fmt"
	"strconv"
	"os/exec"

	"project/pokmonapi"

	"google.golang.org/grpc"
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

const (
	port = ":50051"
	mongodbEndpoint = "mongodb://172.17.0.2:27017" // Find this from the Mongo container
)

type queueID struct {
	userNames []string
	length    int
}

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
	ID            primitive.ObjectID `bson:"_id"`
	User1         string             `bson:"user1"`
	User2         string             `bson:"user2"`
	Health1       int                `bson:"health1"`
	Health2       int                `bson:"health2"`
	WhoseTurn     string             `bson:"whose_turn"`
	LastAttack    string             `bson:"last_attack"`
	GamePort      int             `bson:"game_port"`
	CreatedAt     time.Time          `bson:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at"`
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
	pokmonapi.UnimplementedPokmonInfoServer
}

var monsterAttackDB      [][]string = [][]string{{"Leaf blade", "Energy ball", "Apple acid", "Tackle"}, {"Flamethrower", "Blaze kick", "Searing shot", "Tackle"}, {"Hydro cannon", "Surf", "Water ball", "Tackle"}}
var attackpower          []int      = []int{40, 40, 40, 30}
var attackPowerDB 		 map[string]int
var availablePorts       []int = []int{8080, 8081, 8082, 8083, 8084, 8085, 8086, 8087, 8088, 80089, 8090, 8091}

var queue queueID = queueID{}

var client *mongo.Client

func (s *server) SetUserName(ctx context.Context, in *pokmonapi.UserName) (*pokmonapi.Status, error) {
	name := in.GetName()
	status := &pokmonapi.Status{}

	// select collection from database
	col := client.Database("Pokmon").Collection("users")

	// filter user tagged as specified name
	filter := bson.M{"user": bson.M{"$eq": name}}

	// find one user
	var filterUser userDatabase
	err := col.FindOne(context.TODO(), filter).Decode(&filterUser)

	if err != nil { // user not found in system
    	// Insert one user
		col.InsertOne(context.TODO(), &userDatabase{
			ID:          primitive.NewObjectID(),
			User:        name,
			Monster:     "new",
			CreatedAt:   time.Now(),
		})

		status.Code = "Set User name. Need monster name"

		return status, nil
	} else { // user found in system
		if filterUser.Monster == "new"{ // never added monster to their team
			status.Code = "Username in system. Enter monster"

			return status, nil
		} else { // added monster to team already
			status.Code = "Username in system. Does not need to enter monster"

			return status, nil
		}
	}
}

func (s *server) GetMonsterInfo(ctx context.Context, in *pokmonapi.MonsterName) (*pokmonapi.MonsterNames, error) {
	monsterNames := &pokmonapi.MonsterNames{}

	var monsterNameSlice []string

	// select collection from database
	col := client.Database("Pokmon").Collection("monsters")

	foundMonsters, err := col.Find(context.TODO(), bson.D{})
	if err != nil {

		return monsterNames, errors.New("Unable to display the monsters")
	}

	// Map found monters to an array of monster names 
	for foundMonsters.Next(context.TODO()) {
		monsters := monsterDatabase{}
		err := foundMonsters.Decode(&monsters)
		if err != nil {
			// no monster was found
		} else {
			// save mosnter names to array here
			monsterNameSlice = append(monsterNameSlice, monsters.Monster)
		}
	}

	monsterNames.Monsters = monsterNameSlice

	return monsterNames, nil
}

func (s *server) SetMonsterInfo(ctx context.Context, in *pokmonapi.UserAndName) (*pokmonapi.Status, error) {
	name := in.GetName()
	monster := in.GetMonster()
	status := &pokmonapi.Status{}

	// select collection from database
	colUsers := client.Database("Pokmon").Collection("users")
	colMonsters := client.Database("Pokmon").Collection("monsters")

	// filters for the databases
	userFilter := bson.M{"user": bson.M{"$eq": name}}
	monsterFilter := bson.M{"monster": bson.M{"$eq": monster}}

	// find one user and one monster
	var filterUser userDatabase
	var filterMonster monsterDatabase
	
	if errM := colMonsters.FindOne(context.TODO(), monsterFilter).Decode(&filterMonster); errM == nil {
		// the mosnter is in the database
		if errU := colUsers.FindOne(context.TODO(), userFilter).Decode(&filterUser); errU == nil {
			// user is in database
			if filterUser.Monster == "new" {
				// new user needs to add mosnter
				// upadter for specified item and price
				updater := bson.M{"$set": bson.M{"monster": filterMonster.Monster, "attack_moves": filterMonster.AttackMoves, "health": filterMonster.Health, "element": filterMonster.Element, "updated_at": time.Now()}}

				//Perform UpdateOne operation & validate against the error.
				_, err := colUsers.UpdateOne(context.TODO(), userFilter, updater)

				if err != nil {
					// unable to update item, could not find the user
					status.Code = "Unable to add monster to your team"

					return status, errors.New("Unable to add monster to your team")
				} else {
					// successfully added user monster attributes to the user's database
					status.Code = "Added monster to your team"

					return status, nil
				}
			} else {
				// user already has a monster
				status.Code = "Username already has monster"

				return status, nil
			}
		} else {
			// user not in database
			status.Code = "Username not in database"

			return status, errors.New("Unable to add mosnter. Username not in database")
		}
	} else {
		// no monster in database
		status.Code = "Monster not in database"

		return status, errors.New("Unable to add mosnter. Monster not in database")
	}
}

// mutex? race condition could happen here
func (s *server) JoinQueue(ctx context.Context, in *pokmonapi.UserName) (*pokmonapi.Status, error) {
	name := in.GetName()
	status := &pokmonapi.Status{}

	for i := 0; i < queue.length; i++ {
		if queue.userNames[i] == name {
			status.Code = "Username is already in the Queue"

			return status, errors.New("Username is already in the Queue")
		}
	}
	
	queue.userNames = append(queue.userNames, name)
	queue.length = queue.length + 1

	status.Code = "Added to the Queue"

	return status, nil
}


func (s *server) GetActionInfo(ctx context.Context, in *pokmonapi.RequestInfo ) (*pokmonapi.AttackActions, error) {
	name := in.GetName()
	actions := &pokmonapi.AttackActions{}

	// select collection from database
	colUsers := client.Database("Pokmon").Collection("users")
	colMonsters := client.Database("Pokmon").Collection("monsters")

	// filters for the databases
	userFilter := bson.M{"user": bson.M{"$eq": name}}

	// find one user and one monster
	var filterUser userDatabase
	var filterMonster monsterDatabase

	if errM := colUsers.FindOne(context.TODO(), userFilter).Decode(&filterUser); errM == nil {
		monsterFilter := bson.M{"monster": bson.M{"$eq": filterUser.Monster}}
		if errU := colMonsters.FindOne(context.TODO(), monsterFilter).Decode(&filterMonster); errU == nil {
			// found monster in database
			var tempActionSlice []string
			for _, actions := range filterMonster.AttackMoves {
				tempActionSlice = append(tempActionSlice, actions)
			}
			actions.Actions = tempActionSlice
			actions.Actions = filterMonster.AttackMoves
			return actions, nil
		} else {
			// could not find the monster in the database\
			return actions, errors.New("Error getting actions of monster")
		}
	} else {
		// could not find the user in the database
		return actions, errors.New("Error getting actions of monster")
	}

}

func (s *server) GetGameInfo(ctx context.Context, in *pokmonapi.RequestInfo) (*pokmonapi.GameStatus, error) {
	name := in.GetName()
	gameStatus := &pokmonapi.GameStatus{}

	// select collection from database
	colUsers := client.Database("Pokmon").Collection("users")
	colGames := client.Database("Pokmon").Collection("games")

	// filters for the databases
	user1Filter := bson.M{"user": bson.M{"$eq": name}}
	gameFilter1 := bson.M{"user1": bson.M{"$eq": name}}
	gameFilter2 := bson.M{"user2": bson.M{"$eq": name}}

	// find one user and one monster
	var filterUser1 userDatabase
	var filterUser2 userDatabase
	var filterGame  gameDatabase

	time.Sleep(1 * time.Second)

	fmt.Println(name, " trying to get game info")

	// checking to see if the user is already in a game
	if errG := colGames.FindOne(context.TODO(), gameFilter1).Decode(&filterGame); errG == nil {
		// user is in a game, find user information
		if name != filterGame.User1 {
			user2Filter := bson.M{"user": bson.M{"$eq": filterGame.User1}}
			errU := colUsers.FindOne(context.TODO(), user2Filter).Decode(&filterUser2)
			errU = colUsers.FindOne(context.TODO(), user1Filter).Decode(&filterUser1)
			if errU == nil {
				// found user 2
				gameStatus.OpponentName = filterGame.User1
				gameStatus.OpponentMonster = filterUser2.Monster
				gameStatus.OpponentHealth = int32(filterGame.Health1)
				gameStatus.MyHealth = int32(filterGame.Health2)
				gameStatus.MyMonster = filterUser1.Monster
				gameStatus.WhoseTurn = filterGame.WhoseTurn
				gameStatus.GameID = filterGame.ID.Hex()
				gameStatus.GamePort = int32(filterGame.GamePort)

				gameStatus.Code = "Game created"

				return gameStatus, nil
			}else{
				// error finding user 2
				gameStatus.Code = "User 2 not found, so game did not create"

				return gameStatus, errors.New("User 2 not found, so game didnt create")
			}
		} else {
			user2Filter := bson.M{"user": bson.M{"$eq": filterGame.User2}}
			errU := colUsers.FindOne(context.TODO(), user2Filter).Decode(&filterUser2)
			errU = colUsers.FindOne(context.TODO(), user1Filter).Decode(&filterUser1)
			if errU == nil {
				// found user 2
				gameStatus.OpponentName = filterGame.User2
				gameStatus.OpponentMonster = filterUser2.Monster
				gameStatus.OpponentHealth = int32(filterGame.Health2)
				gameStatus.MyHealth = int32(filterGame.Health1)
				gameStatus.MyMonster = filterUser1.Monster
				gameStatus.WhoseTurn = filterGame.WhoseTurn
				gameStatus.GameID = filterGame.ID.Hex()
				gameStatus.GamePort = int32(filterGame.GamePort)

				gameStatus.Code = "Game created"

				return gameStatus, nil
			}else{
				// error finding user 2
				gameStatus.Code = "User 2 not found, so game did not create"

				return gameStatus, errors.New("User 2 not found, so game didnt create")
			}
		}
	} else if errG := colGames.FindOne(context.TODO(), gameFilter2).Decode(&filterGame); errG == nil {
		// user is in a game, find user information
		if name != filterGame.User1 {
			user2Filter := bson.M{"user": bson.M{"$eq": filterGame.User1}}
			errU := colUsers.FindOne(context.TODO(), user2Filter).Decode(&filterUser2)
			errU = colUsers.FindOne(context.TODO(), user1Filter).Decode(&filterUser1)
			if errU == nil {
				// found user 2
				gameStatus.OpponentName = filterGame.User1
				gameStatus.OpponentMonster = filterUser2.Monster
				gameStatus.OpponentHealth = int32(filterGame.Health1)
				gameStatus.MyHealth = int32(filterGame.Health2)
				gameStatus.MyMonster = filterUser1.Monster
				gameStatus.WhoseTurn = filterGame.WhoseTurn
				gameStatus.GameID = filterGame.ID.Hex()
				gameStatus.GamePort = int32(filterGame.GamePort)

				gameStatus.Code = "Game created"

				return gameStatus, nil
			}else{
				// error finding user 2
				gameStatus.Code = "User 2 not found, so game did not create"

				return gameStatus, errors.New("User 2 not found, so game didnt create")
			}
		} else {
			user2Filter := bson.M{"user": bson.M{"$eq": filterGame.User2}}
			errU := colUsers.FindOne(context.TODO(), user2Filter).Decode(&filterUser2)
			errU = colUsers.FindOne(context.TODO(), user1Filter).Decode(&filterUser1)
			if errU == nil {
				// found user 2
				gameStatus.OpponentName = filterGame.User2
				gameStatus.OpponentMonster = filterUser2.Monster
				gameStatus.OpponentHealth = int32(filterGame.Health2)
				gameStatus.MyHealth = int32(filterGame.Health1)
				gameStatus.MyMonster = filterUser1.Monster
				gameStatus.WhoseTurn = filterGame.WhoseTurn
				gameStatus.GameID = filterGame.ID.Hex()
				gameStatus.GamePort = int32(filterGame.GamePort)

				gameStatus.Code = "Game created"

				return gameStatus, nil
			}else{
				// error finding user 2
				gameStatus.Code = "User 2 not found, so game did not create"

				return gameStatus, errors.New("User 2 not found, so game didnt create")
			}
		}
	} else {
		// user was not found in any game, so we need to create a new game

		for queue.length < 2 {
			// wait for the queue lengh to increase to atleast 2 before creating the game
		}	
		fmt.Println(name, " is done wainting for antoher person in queue")	
	
		if queue.length >= 2 {
			if name != queue.userNames[0] {
				user2Filter := bson.M{"user": bson.M{"$eq": queue.userNames[0]}}
				if errU := colUsers.FindOne(context.TODO(), user1Filter).Decode(&filterUser1); errU == nil {
					if errU := colUsers.FindOne(context.TODO(), user2Filter).Decode(&filterUser2); errU == nil {
						var tempPort int = 12345
						for i := 0; i < len(availablePorts); i++ {
							portFilter := bson.M{"game_port": bson.M{"$eq": availablePorts[i]}}

							var filterPort gameDatabase

							if errG := colGames.FindOne(context.TODO(), portFilter).Decode(&filterPort); errG != nil {
								tempPort = 	availablePorts[i]
								break
							}
						}

						// Insert one game
						res, _ := colGames.InsertOne(context.TODO(), &gameDatabase{
							ID:            primitive.NewObjectID(),
							User1:         filterUser1.User, 
							User2:  	   filterUser2.User,
						    Health1:       filterUser1.Health, 
							Health2:       filterUser2.Health,
							WhoseTurn:     filterUser2.User,
							GamePort:      tempPort,
							CreatedAt:     time.Now(),
						})

						gameStatus.OpponentName = filterUser2.User
						gameStatus.OpponentMonster = filterUser2.Monster
						gameStatus.OpponentHealth = int32(filterUser2.Health)
						gameStatus.MyHealth = int32(filterUser1.Health)
						gameStatus.MyMonster = filterUser1.Monster
						gameStatus.WhoseTurn = filterUser2.User
						gameStatus.GameID = res.InsertedID.(primitive.ObjectID).Hex()
						gameStatus.GamePort = int32(tempPort)

						port := strconv.Itoa(tempPort) + ":8080"
						cmd2 := exec.Command("sudo", "docker", "container", "run", "-p", port, "pokmongame")
						if err := cmd2.Start(); err != nil {
							log.Fatal(err)
						}
					}
				}
			} else {
				user2Filter := bson.M{"user": bson.M{"$eq": queue.userNames[1]}}
				if errU := colUsers.FindOne(context.TODO(), user1Filter).Decode(&filterUser1); errU == nil {
					if errU := colUsers.FindOne(context.TODO(), user2Filter).Decode(&filterUser2); errU == nil {
						var tempPort int = 12345
						for i := 0; i < len(availablePorts); i++ {
							portFilter := bson.M{"game_port": bson.M{"$eq": availablePorts[i]}}

							var filterPort gameDatabase

							if errG := colGames.FindOne(context.TODO(), portFilter).Decode(&filterPort); errG != nil {
								tempPort = 	availablePorts[i]
								break
							}
						}

						// Insert one game
						res, _ := colGames.InsertOne(context.TODO(), &gameDatabase{
							ID:            primitive.NewObjectID(),
							User1:         filterUser1.User, 
							User2:  	   filterUser2.User,
						    Health1:       filterUser1.Health, 
							Health2:       filterUser2.Health,
							WhoseTurn:     filterUser1.User,
							GamePort:      tempPort,
							CreatedAt:     time.Now(),
						})

						gameStatus.OpponentName = filterUser2.User
						gameStatus.OpponentMonster = filterUser2.Monster
						gameStatus.OpponentHealth = int32(filterUser2.Health)
						gameStatus.MyHealth = int32(filterUser1.Health)
						gameStatus.MyMonster = filterUser1.Monster
						gameStatus.WhoseTurn = filterUser1.User
						gameStatus.GameID = res.InsertedID.(primitive.ObjectID).Hex()
						gameStatus.GamePort = int32(tempPort)

						port :=  strconv.Itoa(tempPort) + ":8080"
						cmd2 := exec.Command("sudo", "docker", "container", "run", "-p", port, "pokmongame")
						if err := cmd2.Start(); err != nil {
							log.Fatal(err)
						}
					}
				}
			}
	
			// remove the first user from the queue
			if queue.length > 1 {
				queue.userNames = queue.userNames[1:queue.length]
				queue.length = queue.length - 1
			}
			// remove the next person from the queue, check to see if there is more than one perosn left 
			if queue.length > 1 {
				queue.userNames = queue.userNames[1:queue.length]
				queue.length = queue.length - 1
			} else {
				queue.userNames = make([]string, 0)
				queue.length = 0
			}
	
			gameStatus.Code = "Game created"
	
			return gameStatus, nil
		} else { 
			gameStatus.Code = "Error creating the game. Too few players"
	
			return gameStatus, errors.New("Error creating the game. Too few players")
		}
	}
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
	pokmonapi.RegisterPokmonInfoServer(s, &server{})

	fmt.Println("Server successfully started")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}