package krakendowinaaaauth

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type IdentityClaims struct {
	Claims []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"claims"`
}

func validateOwin(_ *gin.Context, accessToken string) Claims {
	client := &http.Client{}
	var identityClaims IdentityClaims
	req, _ := http.NewRequest("GET", "https://id-qa.dhhmena.com/v1/identity/claims", nil)
	req.Header.Add("Authorization", "Bearer "+accessToken)

	resp, _ := client.Do(req)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	err := json.Unmarshal([]byte(body), &identityClaims)
	if err != nil {
		fmt.Println(err)
		return Claims{}
	}
	if identityClaims.Claims != nil && &identityClaims.Claims[0] != nil {
		var userClaims = Claims{}
		for index, value := range identityClaims.Claims {
			if identityClaims.Claims[index].Key == "http://schemas.xmlsoap.org/ws/2005/05/identity/claims/nameidentifier" {
				userClaims.UserID = value.Value
			}
			if identityClaims.Claims[index].Key == "Email" {
				userClaims.Email = value.Value
			}
		}
		return userClaims
	}
	return Claims{}
}
