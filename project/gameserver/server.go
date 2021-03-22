package main

import (
	"context"
	"errors"
	"log"
	"net"

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
	userFirst     userID
	userSecond    userID
	monsterFirst  monsterID
	monsterSecond monsterID
	whoseTurn     string
}

type database struct {
	users    map[string]monsterID
	games    []gameID
	monsters map[string]monsterStats
}

type server struct {
	pokmonapi.UnimplementedPokmonInfoServer
}

var monsterNamesDB []string = []string{"Bulbasaur", "Charmander", "Squirtle", "Chikorita", "Cyndaquil", "Totodile", "Treecko", "Torchic", "Mudkip", "Turtwig", "Chimchar", "Piplup"}
var monsterAttackDB [][]string = [][]string{{"Leaf Blade", "Energy Ball", "Apple Acid"}, {"Flamethrower", "Blaze Kick", "Searing Shot"}, {"Hydro Cannon", "Surf", "Water Shuriken"}}
var monsterHealthDB []int32 = []int32{90, 78, 88, 90, 78, 88, 80, 90, 100, 110, 88, 106} // twice the amount they had in pokemon
var monsterElementDB []string = []string{"Grass", "Fire", "Water"}

var pokmonDB database = database{}

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