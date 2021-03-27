package main

import (
	"context"
	"errors"
	"log"
	"net"
	"fmt"

	"project/pokmonapi"

	"google.golang.org/grpc"
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
}

type queueID struct {
	userNames []string
	length    int
}

type database struct {
	users    map[string]monsterID
	games    []gameID
	monsters map[string]monsterStats
}

type server struct {
	pokmonapi.UnimplementedPokmonInfoServer
}

var monsterNamesDB   []string   = []string{"B", "Charmander", "Squirtle", "Chikorita", "Cyndaquil", "Totodile", "Treecko", "Torchic", "Mudkip", "Turtwig", "Chimchar", "Piplup"}
var monsterAttackDB  [][]string = [][]string{{"Leaf Blade", "Energy Ball", "Apple Acid", "tackle"}, {"Flamethrower", "Blaze Kick", "Searing Shot", "tackle"}, {"Hydro Cannon", "Surf", "Water Shuriken", "tackle"}}
var monsterHealthDB  []int32    = []int32{90, 78, 88, 90, 78, 88, 80, 90, 100, 110, 88, 106} // twice the amount they had in pokemon
var monsterElementDB []string   = []string{"Grass", "Fire", "Water"}

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

	var tempStr string
	for _, s := range monsterNamesDB {
		tempStr += " " + s
	}

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

func (s *server) GetHealthPoints(ctx context.Context, in *pokmonapi.HealthRequest) (*pokmonapi.HealthPoints, error) {

	healthPoints := &pokmonapi.HealthPoints{}

	return healthPoints, nil
}

func (s *server) GetGameInfo(ctx context.Context, in *pokmonapi.RequestInfo) (*pokmonapi.GameStatus, error) {
	gameStatus := &pokmonapi.GameStatus{}

	if queue.length >= 2 {
		gameStatus.Code = "Game created"

		tempGame := gameID{}
		tempGame.users[0] = pokmonDB.users[queue.userNames[0]]
		tempGame.users[1] = pokmonDB.users[queue.userNames[1]]
		tempGame.currentMonsterHealth[0] = tempGame.users[0].monster.attributes.healthPoint
		tempGame.currentMonsterHealth[1] = tempGame.users[1].monster.attributes.healthPoint
		tempGame.whoseTurn = queue.userNames[0]

		pokmonDB.games = append(pokmonDB.games, tempGame)

		remove(queue.users, 0)
		queue.length = queue.length - 1
		remove(queue.users, 0)
		queue.length = queue.length - 1

		fmt.Printf("queue Length: %d", queue.length)

		return gameStatus, nil
	} else { 
		gameStatus.Code = "Game not created. Waiting on opponent"
		return gameStatus, errors.New("Game not created. Waiting on opponent")
	}
}

func (s *server) GetOpponentInfo(ctx context.Context, in *pokmonapi.RequestInfo) (*pokmonapi.OpponentStatus, error) {

	opponentStatus := &pokmonapi.OpponentStatus{}

	return opponentStatus, nil
}

func (s *server) MonsterAttack(ctx context.Context, in *pokmonapi.MonsterAction) (*pokmonapi.Status, error) {

	status := &pokmonapi.Status{}

	return status, nil
}

func remove(slice []int, index int) []int {
    return append(slice[:index], slice[index+1:]...)
}

func main() {
	// initialize the database maps
	pokmonDB.users = make(map[string]monsterID)
	pokmonDB.monsters = make(map[string]monsterStats)

	// initialize the monster database
	var tempStats monsterStats
	for i := 0; i < 12; i++ {
		if (i % 3) == 1 { // grass monsters
			tempStats.attackMoves = monsterAttackDB[0]
			tempStats.healthPoint = monsterHealthDB[i]
			tempStats.elementType = monsterElementDB[0]
		} else if (i % 3) == 2 { // fire monsters
			tempStats.attackMoves = monsterAttackDB[1]
			tempStats.healthPoint = monsterHealthDB[i]
			tempStats.elementType = monsterElementDB[1]
		} else if (i % 3) == 0 { // water monsters
			tempStats.attackMoves = monsterAttackDB[2]
			tempStats.healthPoint = monsterHealthDB[i]
			tempStats.elementType = monsterElementDB[2]
		}

		pokmonDB.monsters[monsterNamesDB[i]] = tempStats
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