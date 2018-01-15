package main

import (
	"gopkg.in/resty.v1"
	"fmt"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RequestSICD (c *gin.Context) {

	var sicd Sicd
	var url string
	if err := c.ShouldBindJSON(&sicd); err == nil {
		url = "http://api.briconnect.bri.co.id/sid/sicd/" + sicd.Name + "/" + sicd.Birth_date + "/" + sicd.Personal_number + "/" + sicd.Branch_code
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
	}

	READ_WS_SICD: resp, err := resty.R().
		SetHeader("Authorization", auth).
		Get(url)
	if err==nil {
		if resp.Status() == "200 OK" {	//Auth token is valid
			c.String(http.StatusOK, resp.String()) 	//Send SICD data as response in JSON format
		} else {
			var resp_body map[string]interface{}
			json.Unmarshal(resp.Body(), &resp_body)
			if resp_body["error"] == "invalid_token" || resp_body["error"] == "missing_token" {
				//fmt.Println("Token "+ auth +" is not valid") // Auth token is not valid or expired 
				auth,_ = RequestToken()
				goto READ_WS_SICD
			} else {	//Other error
				//fmt.Printf("Error: %v\n",resp_body["message"])
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid url or JSON format"})
			}
		}
	} else {
		// Request Error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to connect to prescreening service"})
	}
}

func RequestDHN (c *gin.Context) {

	var dh Dh
	var url string
	if err := c.ShouldBindJSON(&dh); err == nil {
		url = "http://api.briconnect.bri.co.id/sid/dhn/"+ dh.Name + "/" + dh.Birth_date
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
	}

	READ_WS_DHN: resp, err := resty.R().
		SetHeader("Authorization", auth).
		Get(url)
	if err==nil {
		if resp.Status() == "200 OK" {	//Auth token is valid
			c.String(http.StatusOK, resp.String()) 	//Send SICD data as response in JSON format
		} else {
			var resp_body map[string]interface{}
			json.Unmarshal(resp.Body(), &resp_body)
			if resp_body["error"] == "invalid_token" || resp_body["error"] == "missing_token" {
				//fmt.Println("Token "+ auth +" is not valid") // Auth token is not valid or expired 
				auth,_ = RequestToken()
				goto READ_WS_DHN
			} else {	//Other error
				//fmt.Printf("Error: %v\n",resp_body["message"])
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid url or JSON format"})
			}
		}
	} else {
		// Request Error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to connect to prescreening service"})
	}
}

func RequestKEMENDAGRI (c *gin.Context) {

	var nik Nik
	var url string

	if err := c.ShouldBindJSON(&nik); err == nil {
		url = "http://api.briconnect.bri.co.id/sid/kemendagri/"+ nik.NIK
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
	}

	READ_WS_KEMENDAGRI: resp, err := resty.R().
		SetHeader("Authorization", auth).
		Get(url)
	if err==nil {
		if resp.Status() == "200 OK" {	//Auth token is valid
			c.String(http.StatusOK, resp.String()) 	//Send SICD data as response in JSON format
		} else {
			var resp_body map[string]interface{}
			json.Unmarshal(resp.Body(), &resp_body)
			fmt.Printf("erro: %v\n", resp_body["error"])
			if resp_body["error"] == "invalid_token" || resp_body["error"] == "missing_token" {
				//fmt.Println("Token "+ auth +" is not valid") // Auth token is not valid or expired 
				auth,_ = RequestToken()
				goto READ_WS_KEMENDAGRI
			} else {	//Other error
				//fmt.Printf("Error: %v\n",resp_body["message"])
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid url or JSON format"})
			}
		}
	} else {
		// Request Error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to connect to prescreening service"})
	}
}