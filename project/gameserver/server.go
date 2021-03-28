package main

import (
	"context"
	"errors"
	"log"
	"net"
	"strings"
	"math/rand"
	"time"

	"project/pokmonapi"

	"google.golang.org/grpc"
	"github.com/google/uuid"
)

const (
	port = ":50051"
)

type monsterStats struct {
	attackMoves []string
	healthPoint int32
	elementType string 
}

type monsterID struct {
	monsterName string
	attributes  monsterStats
}

type userID struct {
	userName string
	monster monsterID
}

type gameID struct {
	users                [2]userID
	currentMonsterHealth [2]int32
	whoseTurn            string
	gameUuid             string
}

type queueID struct {
	userNames []string
	length    int
}

type database struct {
	users    map[string]monsterID
	games    map[string]gameID
	monsters map[string]monsterStats
}

type server struct {
	pokmonapi.UnimplementedPokmonInfoServer
}

var monsterNamesDB       []string   = []string{"Bulbasaur", "Charmander", "Squirtle", "Chikorita", "Cyndaquil", "Totodile", "Treecko", "Torchic", "Mudkip", "Turtwig", "Chimchar", "Piplup"}
var monsterAttackDB      [][]string = [][]string{{"Leaf Blade", "Energy Ball", "Apple Acid", "tackle"}, {"Flamethrower", "Blaze Kick", "Searing Shot", "tackle"}, {"Hydro Cannon", "Surf", "Water Shuriken", "tackle"}}
var attackPowerDB 		 map[string]int32
var monsterHealthDB      []int32    = []int32{90, 78, 88, 90, 78, 88, 80, 90, 100, 110, 88, 106} // twice the amount they had in pokemon
var monsterElementDB     []string   = []string{"Grass", "Fire", "Water"}
var attackpower          []int32    = []int32{40, 40, 40, 30}

var pokmonDB database = database{}
var queue queueID = queueID{}

func (s *server) SetUserName(ctx context.Context, in *pokmonapi.UserName) (*pokmonapi.Status, error) {
	name := in.GetName()
	status := &pokmonapi.Status{}

	if value, ok := pokmonDB.users[name]; ok {
		if value.monsterName == "new"{
			status.Code = "Username in system. Enter monster"

			return status, errors.New("Username already in system. No monster")
		} else {
			status.Code = "Username in system. Does not need to enter monster"

			return status, errors.New("Username already in system. Has mosnter already")
		}
	} else {
		tempAttributes := monsterStats{attackMoves: []string{"none"}, healthPoint: 0, elementType: "none"}
		tempMonsterID  := monsterID{monsterName: "new", attributes: tempAttributes}
		pokmonDB.users[name] = tempMonsterID

		status.Code = "Set User name. Need monster name"

		return status, nil
	}
}

func (s *server) GetMonsterInfo(ctx context.Context, in *pokmonapi.MonsterName) (*pokmonapi.MonsterNames, error) {
	monsterNames := &pokmonapi.MonsterNames{}

	// var tempStr string
	// for _, s := range monsterNamesDB {
	// 	tempStr += " " + s
	// }

	monsterNames.Monsters = monsterNamesDB

	return monsterNames, nil
}

func (s *server) SetMonsterInfo(ctx context.Context, in *pokmonapi.UserAndName) (*pokmonapi.Status, error) {
	name := in.GetName()
	monster := in.GetMonster()
	status := &pokmonapi.Status{}

	if _, ok := pokmonDB.monsters[monster]; ok {
		if value, ok := pokmonDB.users[name]; ok {
			if value.monsterName == "new"{
				tempMonsterID := monsterID{monsterName: monster,attributes: pokmonDB.monsters[monster]}
				pokmonDB.users[name] = tempMonsterID

				status.Code = "Added monster to your team"

				return status, nil
			} else {
				status.Code = "Username already has monster"

				return status, errors.New("Username already has mosnter")
			}
		} else {
			status.Code = "Username not in database"

			return status, errors.New("Unable to add mosnter. Username not in database")
		}
	} else {
		status.Code = "Monster not in database"

		return status, errors.New("Unable to add mosnter. Monster not in database")
	}
}

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

func (s *server) GetGameInfo(ctx context.Context, in *pokmonapi.RequestInfo) (*pokmonapi.GameStatus, error) {
	name := in.GetName()
	gameStatus := &pokmonapi.GameStatus{}

	for _, value := range pokmonDB.games {
		if name == pokmonDB.games[value.gameUuid].users[0].userName || name == pokmonDB.games[value.gameUuid].users[1].userName {
			gameStatus.Code = "Game created"

			if name != pokmonDB.games[value.gameUuid].users[0].userName {
				gameStatus.OpponentName = pokmonDB.games[value.gameUuid].users[0].userName
				gameStatus.OpponentMonster = pokmonDB.users[pokmonDB.games[value.gameUuid].users[0].userName].monsterName
				gameStatus.OpponentHealth = pokmonDB.games[value.gameUuid].currentMonsterHealth[0]
				gameStatus.MyHealth = pokmonDB.games[value.gameUuid].currentMonsterHealth[1]
			} else {
				gameStatus.OpponentName = pokmonDB.games[value.gameUuid].users[0].userName
				gameStatus.OpponentMonster = pokmonDB.users[pokmonDB.games[value.gameUuid].users[0].userName].monsterName
				gameStatus.OpponentHealth = pokmonDB.games[value.gameUuid].currentMonsterHealth[1]
				gameStatus.MyHealth = pokmonDB.games[value.gameUuid].currentMonsterHealth[0]
			}

			gameStatus.WhoseTurn = pokmonDB.games[value.gameUuid].users[0].userName
			gameStatus.Uuid = value.gameUuid

			return gameStatus, nil
		}
	}

	for queue.length < 2 {
		// wait for the queue lengh to increase to atleast 2
	}		

	if queue.length >= 2 {
		gameStatus.Code = "Game created"

		// this is game identifier
		uuidWithHyphen := uuid.New()
   		gameUUID := strings.Replace(uuidWithHyphen.String(), "-", "", -1)

		tempGame := gameID{}

		tempGame.users[0] = userID{userName: queue.userNames[0], monster: pokmonDB.users[queue.userNames[0]]}
		tempGame.users[1] = userID{userName: queue.userNames[1], monster: pokmonDB.users[queue.userNames[1]]}
		tempGame.currentMonsterHealth[0] = tempGame.users[0].monster.attributes.healthPoint
		tempGame.currentMonsterHealth[1] = tempGame.users[1].monster.attributes.healthPoint
		tempGame.whoseTurn = queue.userNames[0]
		tempGame.gameUuid = gameUUID

		pokmonDB.games[gameUUID] = tempGame

		if name != queue.userNames[0] {
			gameStatus.OpponentName = queue.userNames[0]
			gameStatus.OpponentMonster = pokmonDB.users[queue.userNames[0]].monsterName
			gameStatus.OpponentHealth = tempGame.currentMonsterHealth[0]
			gameStatus.MyHealth = tempGame.currentMonsterHealth[1]
		} else {
			gameStatus.OpponentName = queue.userNames[1]
			gameStatus.OpponentMonster = pokmonDB.users[queue.userNames[1]].monsterName
			gameStatus.OpponentHealth = tempGame.currentMonsterHealth[1]
			gameStatus.MyHealth = tempGame.currentMonsterHealth[0]
		}

		gameStatus.WhoseTurn = queue.userNames[0]
		gameStatus.Uuid = gameUUID

		remove(queue.userNames, 0)
		queue.length = queue.length - 1
		remove(queue.userNames, 0)
		queue.length = queue.length - 1 

		return gameStatus, nil
	} else { 
		gameStatus.Code = "Error creating the game. Too few players"
		return gameStatus, errors.New("Error creating the game. Too few players")
	}

}

func (s *server) GetHealthPoints(ctx context.Context, in *pokmonapi.HealthRequest) (*pokmonapi.HealthPoints, error) {
	name := in.GetName()
	game := in.GetUuid()
	healthPoints := &pokmonapi.HealthPoints{}

	if value, ok := pokmonDB.games[game]; ok {
		for pokmonDB.games[game].whoseTurn != name{
			// wait for the opponent to make its move
			time.Sleep(2 * time.Second)
		}

		if name != value.users[0].userName {
			healthPoints.Health = value.currentMonsterHealth[1]
		} else {
			healthPoints.Health = value.currentMonsterHealth[0]
		}

		healthPoints.WhoseTurn = value.whoseTurn

		return healthPoints, nil
	} else {
		return healthPoints, errors.New("Error getting Health and whose turn")
	}
}

func (s *server) MonsterAttack(ctx context.Context, in *pokmonapi.MonsterAction) (*pokmonapi.HealthPoints, error) {
	name := in.GetName()
	action := in.GetAction()
	game := in.GetUuid()
	healthPoints := &pokmonapi.HealthPoints{}

	if value, ok := pokmonDB.games[game]; ok {
		randMax := int(attackPowerDB[action])
		healthLoss := int32(rand.Intn(randMax))
		healthLoss = int32(rand.Intn(randMax))
		healthLoss = int32(rand.Intn(randMax))

		if name != value.users[0].userName {
			value.currentMonsterHealth[0] = value.currentMonsterHealth[0] - healthLoss
			healthPoints.Health = value.currentMonsterHealth[0]
			value.whoseTurn = value.users[0].userName
		} else {
			value.currentMonsterHealth[1] = value.currentMonsterHealth[1] - healthLoss
			healthPoints.Health = value.currentMonsterHealth[1]
			value.whoseTurn = value.users[1].userName
		}

		healthPoints.WhoseTurn = value.whoseTurn

		pokmonDB.games[game] = value

		return healthPoints, nil
	} else {
		return healthPoints, errors.New("Error getting Health and whose turn")
	}
}

func (s *server) GetActionInfo(ctx context.Context, in *pokmonapi.RequestInfo ) (*pokmonapi.AttackActions, error) {
	name := in.GetName()
	actions := &pokmonapi.AttackActions{}

	actions.Actions = pokmonDB.monsters[pokmonDB.users[name].monsterName].attackMoves

	return actions, nil
}

func (s *server) GetOpponentInfo(ctx context.Context, in *pokmonapi.RequestInfo) (*pokmonapi.OpponentStatus, error) {

	opponentStatus := &pokmonapi.OpponentStatus{}

	return opponentStatus, nil
}

func remove(slice []string, index int) []string {
    return append(slice[:index], slice[index+1:]...)
}

func main() {
	// initialize the database maps
	pokmonDB.users = make(map[string]monsterID)
	pokmonDB.monsters = make(map[string]monsterStats)
	pokmonDB.games =make(map[string]gameID)
	attackPowerDB = make(map[string]int32)

	// initialize the monster database
	var tempStats monsterStats
	for i := 0; i < 12; i++ {
		if (i % 3) == 0 { // grass monsters
			tempStats.attackMoves = monsterAttackDB[0]
			tempStats.healthPoint = monsterHealthDB[i]
			tempStats.elementType = monsterElementDB[0]
		} else if (i % 3) == 1 { // fire monsters
			tempStats.attackMoves = monsterAttackDB[1]
			tempStats.healthPoint = monsterHealthDB[i]
			tempStats.elementType = monsterElementDB[1]
		} else if (i % 3) == 2 { // water monsters
			tempStats.attackMoves = monsterAttackDB[2]
			tempStats.healthPoint = monsterHealthDB[i]
			tempStats.elementType = monsterElementDB[2]
		}

		pokmonDB.monsters[monsterNamesDB[i]] = tempStats
	}

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

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}