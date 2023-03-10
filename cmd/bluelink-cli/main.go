package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	authv1 "github.com/MatthewSerre/car/gen/go/protos/authentication/v1"
	infov1 "github.com/MatthewSerre/car/gen/go/protos/information/v1"
	"github.com/joho/godotenv"
)

type Auth struct {
	Username   string
	PIN        string
	JWT_Token  string
	JWT_Expiry int64
}

type Vehicle struct {
	RegistrationID string
	VIN string
	Generation string
	Mileage string
}

const Authentication_Address = "localhost:50051"
const Information_Address = "localhost:50052"

func main() {
	log.Println("Welcome to the unofficial Hyundai Bluelink CLI!")

	log.Println("Establishing connection to the authentication service...")

	authConn, err := grpc.Dial(Authentication_Address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println("failed to connect to the authentication service:", err)
		os.Exit(1)
	}

	defer authConn.Close()

	c := authv1.NewAuthenticationServiceClient(authConn)

	log.Println("Connection established!")

	log.Println("Authenticating!")

	auth, err := authenticate(c)

	if err != nil {
		log.Println("authentication failed with error:", err)
		os.Exit(1)
	}

	if (Auth{}) == auth {
		log.Println("authentication failed")
		os.Exit(1)
	}

	log.Println("Authentication successful!")

	//

	log.Println("Establishing connection to the information service...")

	infoConn, err := grpc.Dial(Information_Address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println("failed to connect to the information service:", err)
		os.Exit(1)
	}

	defer infoConn.Close()

	d := infov1.NewInformationServiceClient(infoConn)

	log.Println("Connection established!")

	log.Println("Obtaining vehicle information...")

	info, err := getVehicleInfo(d, auth)

	if err != nil {
		log.Println("vehicle information request failed with error:", err)
		os.Exit(1)
	}

	log.Println("Vehicle information:")
	log.Println("Registration ID:", info.RegistrationID)
	log.Println("VIN:", info.VIN)
	log.Println("Mileage:", info.Mileage)
}


func authenticate(c authv1.AuthenticationServiceClient) (Auth, error) {


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
	
	res, err := c.Authenticate(context.Background(), &authv1.AuthenticationRequest{
		Username: username,
		Password: password,
		Pin: pin,
	})

	if err != nil {
		return Auth{}, err
	}

	return Auth{ Username: res.Username, PIN: res.Pin, JWT_Token: res.JwtToken, JWT_Expiry: res.JwtExpiry }, nil
}

func getInput(message string) (input string) {
	log.Println(message)
	input_scanner := bufio.NewScanner(os.Stdin)
	input_scanner.Scan()
	return input_scanner.Text()
}

func getVehicleInfo(c infov1.InformationServiceClient, auth Auth) (Vehicle, error) {
	log.Println("GetVehicleInfo was invoked")

	res, err := c.GetVehicleInfo(context.Background(), &infov1.VehicleInfoRequest{
		Username: auth.Username,
		Pin: auth.PIN,
		JwtToken: auth.JWT_Token,
		JwtExpiry: auth.JWT_Expiry,
	})

	if err != nil {
		return Vehicle{}, err
	}

	return Vehicle{ RegistrationID: res.RegistrationId, VIN: res.Vin, Generation: res.Generation, Mileage: res.Mileage }, nil
}