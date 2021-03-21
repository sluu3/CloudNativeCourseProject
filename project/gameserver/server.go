package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"


	"google.golang.org/grpc"
)

const (
	port = ":50051"
)


type userID struct {
	userName     string
	monsterNames []monsterID
}

type monsterID struct {
	monsterName string
	attackMoves []string
	healthPoint int32
	elementType string 
}

type gameID struct {
	userFirst     userID
	userSecond    userID
	monsterFirst  monsterID
	monsterSecond monsterID
	whoseTurn     string
}

type database struct {
	users    []userID
	games    []gameID
	monsters []monsterID
}

type server struct {
	
}



func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}