// Package main imlements a client for movieinfo service
package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
	"strconv"

	"project/pokmonapi"
	"project/gameapi"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"      //local
	//address = "10.152.183.178:50051" //kubernetes
)

type gameID struct {
	users                [2]string
	monsters			 [2]string
	totalMonsterHealth   [2]int
	currentMonsterHealth [2]int
	whoseTurn            string
	gameId               string
	gamePort             int
	lastAttack           string
	damage               int
}

func main() {
	var userName string
	var monster string
	var readyCheck string
	var game gameID = gameID{}
	var displayType string
	var action string
	var checkSpelling bool = false

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	server_grpc := pokmonapi.NewPokmonInfoClient(conn)

	// Timeout if server doesn't respond
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()

	input := bufio.NewScanner(os.Stdin)
	//reader := bufio.NewReader(os.Stdin)

	// connecting to the server with a certain username
	fmt.Println("Enter your username:")
	input.Scan()
	userName = input.Text()

	status, err := server_grpc.SetUserName(ctx, &pokmonapi.UserName{Name: userName})
	fmt.Println(status)

	if err == nil {
		if status.GetCode() == "Username in system. Does not need to enter monster" {
			// does nothing 
		} else {
			// printing the monsters the users can choose from
			monsterNames, err := server_grpc.GetMonsterInfo(ctx, &pokmonapi.MonsterName{Monster: "none"})

			if err == nil {
				fmt.Printf("\nMonsters to choose from: %v\n", monsterNames.GetMonsters())

				checkSpelling = false
				for checkSpelling != true {
					fmt.Println("Enter monster you want:")
					input.Scan()
					monster = input.Text()

					// check to see if user spelled monster name correctly
					for _, checkName := range monsterNames.GetMonsters() {
						if monster == checkName { // spelled correctly
							checkSpelling = true
							break
						}
					}
				}

				// set the user's monster to the userName
				status, err = server_grpc.SetMonsterInfo(ctx, &pokmonapi.UserAndName{Name: userName, Monster: monster})
				fmt.Println(status)	
			}	
		}
			
		if err == nil {
			fmt.Printf("\nYou are ready for battle! Enter 'Ready' if you want to join the Queue.\n")

			for {
				input.Scan()
				readyCheck = input.Text()
				//readyCheck, _ = reader.ReadString('\n')

				if readyCheck == "Ready" || readyCheck == "ready" {
					status, err = server_grpc.JoinQueue(ctx, &pokmonapi.UserName{Name: userName})
					fmt.Println(status)
					break
				}

				fmt.Println("Please enter 'Ready' if you want to join the queue")
			}

			// client game
			for {
				// create game after joining the queue
				gameStatus, err := server_grpc.GetGameInfo(context.TODO(), &pokmonapi.RequestInfo{Name: userName})

				fmt.Println(gameStatus)

				// set users of the game
				game.users[0] = gameStatus.GetOpponentName()
				game.users[1] = userName

				// set monsters of the game
				game.monsters[0] = gameStatus.GetOpponentMonster()
				game.monsters[1] = gameStatus.GetMyMonster()

				// set monsters health
				game.currentMonsterHealth[0] = int(gameStatus.GetOpponentHealth())
				game.currentMonsterHealth[1] = int(gameStatus.GetMyHealth())
				game.totalMonsterHealth[0] = int(gameStatus.GetOpponentHealth())
				game.totalMonsterHealth[1] = int(gameStatus.GetMyHealth())

				// set whose turn
				game.whoseTurn = gameStatus.GetWhoseTurn()

				// set game ID
				game.gameId = gameStatus.GetGameID()

				// set game port
				game.gamePort = int(gameStatus.GetGamePort())
				fmt.Println(game.gamePort, "\n")

				// set display to setup 
				displayType = "setup"

				err = displayGame(displayType, game)
				if err != nil {
					log.Fatalf("did not display game: %v", err)
				}

				// get the available actions for the user's monster
				attackActions, err := server_grpc.GetActionInfo(context.TODO(), &pokmonapi.RequestInfo{Name: userName})

				conn.Close()
				newAddress := "localhost:" + strconv.Itoa(game.gamePort)
				// Set up a connection to the game.
				conn, err := grpc.Dial(newAddress, grpc.WithInsecure(), grpc.WithBlock())
				if err != nil {
					log.Fatalf("did not connect: %v", err)
				}
				game_grpc := gameapi.NewGameInfoClient(conn)


				// pokmon battle happens here as long as no player has 0 or fewer HP
				for { 
					// check to see if you start first
					if game.whoseTurn == userName {
						fmt.Printf("Attack 1: %v \t\tAttack 2: %v\nAttack 3: %v \t\tAttack 4: %v\n\n", attackActions.GetActions()[0], attackActions.GetActions()[1], attackActions.GetActions()[2], attackActions.GetActions()[3])

						checkSpelling = false
						for checkSpelling != true {
							// get action from user
							fmt.Println("Enter the monster's Action Number (1, 2, 3, 4):")
							input.Scan()
							action = input.Text()
							switch action {
							case "1":
								action = attackActions.GetActions()[0]
								checkSpelling = true
							case "2":
								action = attackActions.GetActions()[1]
								checkSpelling = true
							case "3":
								action = attackActions.GetActions()[2]
								checkSpelling = true
							case "4":
								action = attackActions.GetActions()[3]
								checkSpelling = true
							default:
								break
							}
						}

						// send action returns opponents Health Points and turn change
						gameInfo, err := game_grpc.MonsterAttack(context.TODO(), &gameapi.MonsterAction{Name: userName, Action: action, GameID: game.gameId})

						if err == nil{
							game.currentMonsterHealth[0] = int(gameInfo.GetHealth())
							game.whoseTurn = gameInfo.GetWhoseTurn()
							game.lastAttack = gameInfo.GetLastAttack()
							game.damage = int(gameInfo.GetDamage())
						}

						err = displayGame("battle-attacker", game)
						if err != nil {
							log.Fatalf("did not display game: %v", err)
						}
						
						if game.currentMonsterHealth[0] <= 0 {
							fmt.Println("Opponent's monster reached 0 Hp.\nYou have Won! ")
							break
						}
					} else {
						fmt.Println("Waiting for opponent to attack")
					
						// check my Health Points returns my Health Points and turn change
						gameInfo, err := game_grpc.GetHealthPoints(context.TODO(), &gameapi.HealthRequest{Name: userName, GameID: game.gameId})

						if err == nil{
							game.currentMonsterHealth[1] = int(gameInfo.GetHealth())
							game.whoseTurn = gameInfo.GetWhoseTurn()
							game.lastAttack = gameInfo.GetLastAttack()
							game.damage = int(gameInfo.GetDamage())
						}

						err = displayGame("battle-defender", game)
						if err != nil {
							log.Fatalf("did not display game: %v", err)
						}

						if game.currentMonsterHealth[1] <= 0 {
							fmt.Println("Your monster reached 0 Hp.\nYou have lost! ")
							break
						}
					}
				}

				fmt.Println("\nWould you like to play again? \nEnter 'Ready' if you want to rejoin the queue: ")

				conn.Close()

				for {
					input.Scan()
					readyCheck = input.Text()
					//readyCheck, _ = reader.ReadString('\n')

					if readyCheck == "Ready" || readyCheck == "ready" {
						
						// Set up a connection to the game.
						conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
						if err != nil {
							log.Fatalf("did not connect: %v", err)
						}

						server_grpc = pokmonapi.NewPokmonInfoClient(conn)

						status, err = server_grpc.JoinQueue(context.TODO(), &pokmonapi.UserName{Name: userName})
						fmt.Println(status)
						if err == nil {
							break
						} else {
							fmt.Println("something went really wrong, you should be scared")
						}
					} else if readyCheck == "Quit" || readyCheck == "quit" {
						return
					}

					fmt.Println("Please enter 'Ready' if you want to join the queue")
				}
			}
		}
	}
}

func displayGame(display string, game gameID) error {
	switch display{
	case "setup":
		// clear the terminal
		cmd := exec.Command("clear") //Linux example, its tested
        cmd.Stdout = os.Stdout
		cmd.Run()
		
		// display opponent information
		fmt.Printf("Opponent's Name: %s\n", game.users[0])
		fmt.Printf("Opponent's Monster: %s\n", game.monsters[0])
		fmt.Printf("Monster's Health: %d / %d\n", game.currentMonsterHealth[0], game.totalMonsterHealth[0])

		fmt.Printf("\n\n\n\n")

		// display self information
		fmt.Printf("Your Name: %s\n", game.users[1])
		fmt.Printf("Your Monster: %s\n", game.monsters[1])
		fmt.Printf("Monster's Health: %d / %d\n\n", game.currentMonsterHealth[1], game.totalMonsterHealth[1])

		fmt.Println("----------------------------------------------------\n\n")

		return nil
	case "battle-attacker":
		// something
		// clear the terminal
		cmd := exec.Command("clear") //Linux example, its tested
        cmd.Stdout = os.Stdout
		cmd.Run()
		
		// display opponent information
		fmt.Printf("Opponent's Name: %s\n", game.users[0])
		fmt.Printf("Opponent's Monster: %s\n", game.monsters[0])
		fmt.Printf("Monster's Health: %d / %d\n", game.currentMonsterHealth[0], game.totalMonsterHealth[0])

		fmt.Printf("\n\n\n\n")

		// display self information
		fmt.Printf("Your Name: %s\n", game.users[1])
		fmt.Printf("Your Monster: %s\n", game.monsters[1])
		fmt.Printf("Monster's Health: %d / %d\n\n", game.currentMonsterHealth[1], game.totalMonsterHealth[1])

		fmt.Printf("------------------------------------------------------------\n")
		fmt.Printf("%s's %s used %s and did %d damage            \n", game.users[1], game.monsters[1], game.lastAttack, game.damage)
		fmt.Printf("------------------------------------------------------------\n\n")

		return nil	
	case "battle-defender":
		// something
		// clear the terminal
		cmd := exec.Command("clear") //Linux example, its tested
        cmd.Stdout = os.Stdout
		cmd.Run()
		
		// display opponent information
		fmt.Printf("Opponent's Name: %s\n", game.users[0])
		fmt.Printf("Opponent's Monster: %s\n", game.monsters[0])
		fmt.Printf("Monster's Health: %d / %d\n", game.currentMonsterHealth[0], game.totalMonsterHealth[0])

		fmt.Printf("\n\n\n\n")

		// display self information
		fmt.Printf("Your Name: %s\n", game.users[1])
		fmt.Printf("Your Monster: %s\n", game.monsters[1])
		fmt.Printf("Monster's Health: %d / %d\n\n", game.currentMonsterHealth[1], game.totalMonsterHealth[1])

		fmt.Printf("------------------------------------------------------------\n")
		fmt.Printf("%s's %s used %s and did %d damage            \n", game.users[0], game.monsters[0], game.lastAttack, game.damage)
		fmt.Printf("------------------------------------------------------------\n\n")

		return nil	
	}

	return errors.New("Could not dsisplay the game")
}