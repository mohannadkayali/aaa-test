package krakendowinaaaauth

import (
	"context"
	sts "github.com/deliveryhero/pd-sts-go-sdk"
	"github.com/deliveryhero/pd-sts-go-sdk/agent"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

var aaaCtx context.Context
var stsAgent sts.IntrospectionService

func loadAaaConf() {
	privateKey := os.Getenv("AAA_PRIVATE_KEY")
	keyID := os.Getenv("AAA_KEY_ID")
	host := os.Getenv("AAA_HOST_URL")
	clientID := os.Getenv("AAA_CLIENT_ID")

	// create config for agent to connect to the control place
	cfg := &agent.Config{
		Host:       host,               // STS host url
		KeyID:      keyID,              // id provided by STS
		ClientID:   clientID,           // your client id registered with STS
		Timeout:    5 * time.Second,    // response header timeout with STS
		PrivateKey: string(privateKey), // private key
	}

	// initialize the introspection agent
	stsAgent = sts.NewIntrospectionAgent(cfg)
}

func validateAaa(ginCtx *gin.Context, accessToken string) Claims {
	headers := http.Header{
		"X-Global-Entity-ID": []string{"TB_KW"},
	}
	aaaCtx = sts.WithHttpHeaders(ginCtx, headers)

	var claims Claims
	// get claims from token, provided token is valid
	token, err := stsAgent.GetMetadata(aaaCtx, accessToken)
	if err == nil {
		claims.UserID = token.Subject
	}
	return claims
}
