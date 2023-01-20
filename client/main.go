package main

import (
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/MatthewSerre/car/authentication/proto"
	"github.com/MatthewSerre/car/helper"
	bp "github.com/MatthewSerre/car/information/proto"
)

func main() {
	log.Println("Welcome to the unofficial Hyundai Bluelink CLI!")

	log.Println("Establishing connection to the authentication service...")

	authConn, err := grpc.Dial(helper.Authentication_Address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println("failed to connect to the authentication service:", err)
		os.Exit(1)
	}

	defer authConn.Close()

	c := pb.NewAuthenticationServiceClient(authConn)

	log.Println("Connection established!")

	log.Println("Authenticating!")

	auth, err := Login(c)

	if err != nil {
		log.Println("authentication failed with error:", err)
		os.Exit(1)
	}

	if (helper.Auth{}) == auth {
		log.Println("authentication failed")
		os.Exit(1)
	}

	log.Println("Authentication successful!")

	//

	log.Println("Establishing connection to the information service...")

	infoConn, err := grpc.Dial(helper.Information_Address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println("failed to connect to the information service:", err)
		os.Exit(1)
	}

	defer infoConn.Close()

	d := bp.NewInformationServiceClient(infoConn)

	log.Println("Connection established!")

	log.Println("Obtaining vehicle information...")

	info, err := GetVehicleInfo(d, auth)

	if err != nil {
		log.Println("vehicle information request failed with error:", err)
		os.Exit(1)
	}

	log.Println("Vehicle information:")
	log.Println("Registration ID:", info.RegistrationID)
	log.Println("VIN:", info.VIN)
	log.Println("Mileage:", info.Mileage)
}