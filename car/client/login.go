package main

import (
	"bufio"
	"context"
	"log"
	"os"

	pb "github.com/MatthewSerre/car/proto"
)

func Login(c pb.CarServiceClient) {
	log.Println("Login was invoked")

	username := getInput("Enter your username\n")
	password := getInput("Enter your password\n")
	pin := getInput("Enter your PIN\n")


	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Username: username,
		Password: password,
		Pin: pin,
	})

	if err != nil {
		log.Fatalf("Could not login: %v\n", err)
	}

	log.Print("LoginResponse:\n")
	log.Printf("Username: %v\n", res.Username)
	log.Printf("PIN: %v\n", res.Pin)
	log.Printf("JWT Token: %v\n", res.JwtToken)
	log.Printf("JWT Expiry: %v\n", res.JwtExpiry)
}

func getInput(message string) (input string) {
	log.Println(message)
	input_scanner := bufio.NewScanner(os.Stdin)
	input_scanner.Scan()
	return input_scanner.Text()
}