package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	pb "github.com/MatthewSerre/car/proto"
	"github.com/golang-jwt/jwt/v4"
)

const Base_URL = "https://owners.hyundaiusa.com"

type Auth struct {
	Username   string
	PIN        string
	JWT_Token  string
	JTW_Expiry int64
}

func (s *Server) Login(context context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	log.Printf("Login function was invoked for user %v\n", request.Username)

	getCSFRToken()

	auth, err := login(request.Username, request.Password, request.Pin)

	if err != nil {
		log.Println("error loggin in: ", err)
	}

	return &pb.LoginResponse{
		Username: auth.Username,
		Pin: auth.PIN,
		JwtToken: auth.JWT_Token,
		JwtExpiry: auth.JTW_Expiry,
	}, nil
}

func getCSFRToken() (string, error) {
	// Generate a new request to obtain a cross-site forgery request (CSFR) token
	req, err := http.NewRequest("GET", Base_URL + "/etc/designs/ownercommon/us/token.json", nil)

	if err != nil {
		log.Println("error generating CSRF token req: ", err)
		return "", err
	}

	// Call the request
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println("error calling CSRF token req: ", err)
		return "", err
	}

	// Read the token from the response body and print it
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("error reading CSRF token: ", err)
		return "", err
	}

	var result map[string]interface{}

	json.Unmarshal([]byte(body), &result)

	csrf := result["jwt_token"].(string)

	log.Println("CSFR token: ", csrf)

	// Generate a new request to validate the token

	req, err = http.NewRequest("GET", Base_URL + "/libs/granite/csrf/token.json", nil)

	if err != nil {
		log.Println("error generating csrf_token validation req: ", err)
		return "", err
	}

	// Add the token to the request header

	req.Header.Add("csrf_token", csrf)

	// Send a request to validate the token

	resp, err = http.DefaultClient.Do(req)

	if err != nil {
		log.Println("error sending csrf validation request: ", err)
		return "", err
	}

	// Validate the token

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Println("error could not validate csrf ", resp.Status)
		return "", err
	}

	return csrf, nil
}

func login(username, password, pin string) (Auth, error) {
	// Obtain CSFR token

	csrf, err := getCSFRToken()

	if err != nil {
		log.Println("error obtaining CSFR token: ", err)
		return Auth{}, err
	}

	// Generate a new request to obtain a JSON web token

	req, err := http.NewRequest("POST", Base_URL + "/bin/common/connectCar", nil)

	if err != nil {
		log.Println("Error getting csrf_token req: ", err)
		return Auth{}, err
	}

	// Add query parameters to the request

	q := req.URL.Query()
	q.Add(":cq_csrf_token", csrf)
	q.Add("url", Base_URL+"/us/en/index.html")
	q.Add("username", username)
	q.Add("password", password)
	req.URL.RawQuery = q.Encode()
	
	// Check the response status

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println("error obtaining JSON web token: ", err)
		return Auth{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Println("error logging in: ", resp.Status)
		return Auth{}, err
	}

	// Print the response body as JSON

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("error reading login response: ", err)
		return Auth{}, err
	}

	var login_result map[string]interface{}

	json.Unmarshal([]byte(body), &login_result)

	jwtID := login_result["RESPONSE_STRING"].(map[string]interface{})["jwtID"].(string)

	log.Println("jwtID: ", jwtID)

	// Remove the first 4 characters from jwtID if it begins with "JWT-"

	var jwtID_truncated string

	jwtID_truncated = jwtID

	if strings.HasPrefix(jwtID, "JWT-") {
		jwtID_truncated = jwtID[4:]
	}

	// Decode the JWT and obtain the expiration date from the "exp" field
	
	token, _ := jwt.Parse(jwtID_truncated, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	expires_at := int64(token.Claims.(jwt.MapClaims)["exp"].(float64) / 1000)

	log.Println("Raw expiration date: ", expires_at)

	auth := Auth{ Username: username, PIN: pin, JWT_Token: jwtID, JTW_Expiry: expires_at }

	return auth, nil
}