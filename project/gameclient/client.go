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

	"project/pokmonapi"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

type gameID struct {
	users                [2]string
	monsters			 [2]string
	totalMonsterHealth   [2]int32
	currentMonsterHealth [2]int32
	whoseTurn            string
	gameID             string
}

func main() {
	var userName string
	var monster string
	var readyCheck string
	var game gameID = gameID{}
	var displayType string
	var action string

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pokmonapi.NewPokmonInfoClient(conn)

	// Timeout if server doesn't respond
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	input := bufio.NewScanner(os.Stdin)
	//reader := bufio.NewReader(os.Stdin)

	// connecting to the server with a certain username
	fmt.Println("Enter your username:")
	input.Scan()
	userName = input.Text()
	//userName, _ = reader.ReadString('\n')
	status, err := c.SetUserName(ctx, &pokmonapi.UserName{Name: userName})
	fmt.Println(status)

	if err == nil {
		// printing the monsters the users can choose from
		monsterNames, err := c.GetMonsterInfo(ctx, &pokmonapi.MonsterName{Monster: "none"})
		if err == nil {
			fmt.Printf("\nMonsters to choose from: %v\n", monsterNames.GetMonsters())

			checkSpelling := false
			for checkSpelling != true {
				fmt.Println("Enter monster you want:")
				input.Scan()
				monster = input.Text()
				//monster, _ = reader.ReadString('\n')

				// check to see if user spelled monster name correctly
				for _, checkName := range monsterNames.GetMonsters() {
					if monster == checkName { // spelled correctly
						checkSpelling = true
						break;
					}
				}
			}

			// set the user's monster to the userName
			status, err = c.SetMonsterInfo(ctx, &pokmonapi.UserAndName{Name: userName, Monster: monster})
			fmt.Println(status)		

			if err == nil {
				fmt.Printf("\nYou are ready for battle! Enter 'Ready' if you want to join the Queue.\n")

				for {
					input.Scan()
					readyCheck = input.Text()
					//readyCheck, _ = reader.ReadString('\n')

					if readyCheck == "Ready" || readyCheck == "ready" {
						status, err = c.JoinQueue(ctx, &pokmonapi.UserName{Name: userName})
						fmt.Println(status)
						break;
					}

					fmt.Println("Please enter 'Ready' if you want to join the queue")
				}

				// create game after joining the queue
				gameStatus, err := c.GetGameInfo(context.TODO(), &pokmonapi.RequestInfo{Name: userName})

				fmt.Println(gameStatus)

				// set users of the game
				game.users[0] = gameStatus.GetOpponentName()
				game.users[1] = userName

				// set monsters of the game
				game.monsters[0] = gameStatus.GetOpponentMonster()
				game.monsters[1] = monster

				// set monsters health
				game.currentMonsterHealth[0] = gameStatus.GetOpponentHealth()
				game.currentMonsterHealth[1] = gameStatus.GetMyHealth()
				game.totalMonsterHealth[0] = gameStatus.GetOpponentHealth()
				game.totalMonsterHealth[1] = gameStatus.GetMyHealth()

				// set whose turn
				game.whoseTurn = gameStatus.GetWhoseTurn()

				// set playerID
				game.gameID = gameStatus.GetUuid()


				// set display to start 
				displayType = "start"

				err = displayGame(displayType, game)
				if err != nil {
					log.Fatalf("did not display game: %v", err)
				}

				// get the available actions for the user's monster
				attackActions, err := c.GetActionInfo(ctx, &pokmonapi.RequestInfo{Name: userName})

				// game happens here as long as no player has 0 or fewer HP
				for game.currentMonsterHealth[0] <= 0 {
					// check to see if you start first
					if game.whoseTurn == userName {
						fmt.Println(attackActions)

						// get action from user
						fmt.Println("Enter the monster's Action:")
						input.Scan()
						action = input.Text()

						// send action returns opponents Health Points and turn change
						gameInfo, err := c.MonsterAttack(context.TODO(), &pokmonapi.MonsterAction{Name: userName, Action: action, Uuid: game.gameID})

						game.currentMonsterHealth[0] = gameInfo.GetHealth()
						game.whoseTurn = gameInfo.GetWhoseTurn()

						err = displayGame(displayType, game)
						if err != nil {
							log.Fatalf("did not display game: %v", err)
						}				
					} else {
						fmt.Println("Waiting for opponent to attack")
					
						// check my Health Points returns my Health Points and turn change
						gameInfo, err := c.GetHealthPoints(context.TODO(), &pokmonapi.HealthRequest{Name: userName, Uuid: game.gameID})

						fmt.Println("opponent has attacked")

						if err == nil{
							game.currentMonsterHealth[1] = gameInfo.GetHealth()
							game.whoseTurn = gameInfo.GetWhoseTurn()
						}

						err = displayGame(displayType, game)
						if err != nil {
							log.Fatalf("did not display game: %v", err)
						}
					}
				}
			}
		}
	}
}

func displayGame(display string, game gameID) error {
	switch display{
	case "start":
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
		fmt.Printf("Monster's Health: %d / %d\n", game.currentMonsterHealth[1], game.totalMonsterHealth[1])

		fmt.Println("----------------------------------------------------\n\n")

		return nil
	case "battle":
		// something

		return nil	
	}

	return errors.New("Could not dsisplay the game")
}