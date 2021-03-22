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

	reader := bufio.NewReader(os.Stdin)
	var userName string
	var monster string

	// connecting to the server with a certain username
	fmt.Println("Enter your username:")
	userName, _ = reader.ReadString('\n')
	status, err := c.SetUserName(ctx, &pokmonapi.UserName{Name: userName})
	fmt.Println(status)

	if err == nil {
		monsterNames, err := c.GetMonsterInfo(ctx, &pokmonapi.MonsterName{Monster: "none"})
		if err == nil {
			fmt.Printf("\nMonsters to choose from: %v\n", monsterNames.GetMonsters())
			fmt.Println("Enter monster you want:")
			monster, _ = reader.ReadString('\n')
			status, err := c.SetMonsterInfo(ctx, &pokmonapi.UserAndName{Name: userName, Monster: monster})
			fmt.Println(status)

			if err == nil {
				fmt.Printf("\nYou are ready for battle! Enter 'Ready' if you want to join the Queue.\n")
			}
		}
	}
}