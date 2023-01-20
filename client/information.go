package main

import (
	"context"
	"log"

	pb "github.com/MatthewSerre/car/information/proto"

	"github.com/MatthewSerre/car/helper"
)

func GetVehicleInfo(c pb.InformationServiceClient, auth helper.Auth) (helper.Vehicle, error) {
	log.Println("GetVehicleInfo was invoked")

	res, err := c.GetVehicleInfo(context.Background(), &pb.VehicleInfoRequest{
		Username: auth.Username,
		Pin: auth.PIN,
		JwtToken: auth.JWT_Token,
		JwtExpiry: auth.JTW_Expiry,
	})

	if err != nil {
		return helper.Vehicle{}, err
	}

	return helper.Vehicle{ RegistrationID: res.RegistrationId, VIN: res.Vin, Generation: res.Generation, Mileage: res.Mileage }, nil
}