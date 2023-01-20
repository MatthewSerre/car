package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/MatthewSerre/car/helper"
	pb "github.com/MatthewSerre/car/information/proto"
)

func (s *Server) GetVehicleInfo(context context.Context, request *pb.VehicleInfoRequest) (*pb.VehicleInfoResponse, error) {
	log.Printf("GetVehicleInfo function was invoked for user %v\n", request.Username)

	vehicleInfo, err := getVehicleInfo(helper.Auth{ Username: request.Username, PIN: request.Pin, JWT_Token: request.JwtToken, JTW_Expiry: request.JwtExpiry })

	if err != nil {
		log.Println("error getting vehicle info:", err)
	}

	return &pb.VehicleInfoResponse{
		RegistrationId: vehicleInfo.RegistrationID,
		Vin: vehicleInfo.VIN,
		Generation: vehicleInfo.Generation,
		Mileage: vehicleInfo.Mileage,
	}, nil
}

// func (s *Server) Login(context context.Context, request *pb.VehicleStatusRequest) (*pb.VehicleStatusResponse, error) {
// 	log.Printf("VehicleStatus function was invoked for user %v\n", request.Username)

// 	return &pb.VehicleStatusResponse{

// 	}, nil
// }

func getVehicleInfo(auth helper.Auth) (helper.Vehicle, error) {
	// Generate a request to obtain owner information

	req, err := http.NewRequest("POST", helper.Base_URL + "/bin/common/MyAccountServlet", nil)

	if err != nil {
		log.Println("error getting owner info req:", err)
		return helper.Vehicle{}, err
	}

	// Set the request headers using a helper method

	helper.SetReqHeaders(req, auth)

	// Add query parameters to the request

	q := req.URL.Query()
	q.Add("username", auth.Username)
	q.Add("token", auth.JWT_Token)
	q.Add("service", "getOwnerInfoService")
	q.Add("url", helper.Base_URL + "/us/en/page/dashboard.html")
	req.URL.RawQuery = q.Encode()

	// Check the response status

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println("error obtaining vehicle information:", err)
		return helper.Vehicle{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Println("error obtaining vehicle information:", resp.Status)
		return helper.Vehicle{}, err
	}

	// Print the response body as JSON

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println("error reading owner info response:", err)
		return helper.Vehicle{}, err
	}

	var ownerInfo helper.OwnerInfoService

	json.Unmarshal([]byte(body), &ownerInfo)

	vehicles := ownerInfo.ResponseString.OwnersVehiclesInfo

	vehicle := helper.Vehicle{ RegistrationID: vehicles[0].RegistrationID, VIN: vehicles[0].VinNumber, Generation: "", Mileage: vehicles[0].Mileage }

	return vehicle, nil
}