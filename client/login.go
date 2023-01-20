package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/MatthewSerre/car/authentication/proto"
	"github.com/joho/godotenv"

	"github.com/MatthewSerre/car/helper"
)

func Login(c pb.AuthenticationServiceClient) (helper.Auth, error) {


	var username, password, pin string

	exit := false;
	for !exit {
		var command string
		log.Println("Enter 1 to input your credentials or 2 to have them read from the environment.")
		fmt.Scan(&command)
		switch command {
		case "1":
			username = getInput("Enter your username\n")
			password = getInput("Enter your password\n")
			pin = getInput("Enter your PIN\n")
			exit = true
		case "2":
			envFile, _ := godotenv.Read(".env")
			username = envFile["USERNAME"]
			password = envFile["PASSWORD"]
			pin = envFile["PIN"]
			exit = true
		default:
			continue
		}
	}
	
	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Username: username,
		Password: password,
		Pin: pin,
	})

	if err != nil {
		return helper.Auth{}, err
	}

	return helper.Auth{ Username: res.Username, PIN: res.Pin, JWT_Token: res.JwtToken, JTW_Expiry: res.JwtExpiry }, nil
}

func getInput(message string) (input string) {
	log.Println(message)
	input_scanner := bufio.NewScanner(os.Stdin)
	input_scanner.Scan()
	return input_scanner.Text()
}