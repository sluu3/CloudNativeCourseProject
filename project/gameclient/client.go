// Package main imlements a client for movieinfo service
package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"project/pokmonapi"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
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
	var userName string
	var monster string
	var readyCheck string

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
				if err != nil {
					fmt.Println(gameStatus)
				}

				for err != nil {
					// create game after joining the queue
					gameStatus, err = c.GetGameInfo(context.TODO(), &pokmonapi.RequestInfo{Name: userName})
				}

				fmt.Println(gameStatus)


			}
		}
	}
}