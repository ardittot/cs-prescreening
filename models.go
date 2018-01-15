package main

import (
	"gopkg.in/resty.v1"
	"fmt"
	"encoding/json"
)

const FILE = "authtoken.gob"
const GRANT_TYPE = "client_credentials"
const CLIENT_ID = "9a8f6e1662c07e0df4f19c3852779a6ec6e14c22"
const CLIENT_SECRET = "770c62343f579a90fe43c1000b57344492ef1ba8"
const USERNAME = "ivojulistira@gmail.com"
const PASSWORD = "092215Conand"

var auth string

type Sicd struct {
	Name string
	Birth_date string
	Personal_number string
	Branch_code string
}

type Dh struct {
	Name string
	Birth_date string
}

type Token struct {
	Access_token string
	Expires_in uint64
	Token_type string
	Scope string
}

func RequestToken() (auth string, success bool) {
	var auth_token = new (Token)
	body := `{"grant_type":"` + GRANT_TYPE + `", "client_id":"` + CLIENT_ID + `", "client_secret":"` + CLIENT_SECRET + `", "username":"` + USERNAME + `", "password":"` + PASSWORD + `"}`
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post("http://sandbox.bri.co.id:82/oauth/token")
	if err==nil {
		if resp.Status() == "200 OK" {
			json.Unmarshal(resp.Body(), &auth_token)
			Save(FILE, auth_token)
			auth = `Bearer ` + auth_token.Access_token
			success = true
			return
		} else {
			// Error
			success = false
			return
		}
	} else {
		// Error
		success = false
		return
	}
}

func InitToken() (auth string, success bool) {
	var auth_token = new (Token)
	err := Load(FILE, auth_token)
	if err==nil {
		fmt.Println("Take from file")
		auth = `Bearer ` + auth_token.Access_token
		success = true
		return
	} else {
		fmt.Println("Take from new request")
		auth,success = RequestToken()
		return
	}
}
