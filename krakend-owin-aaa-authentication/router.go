package krakendowinaaaauth

import (
	"github.com/gin-gonic/gin"
	"github.com/luraproject/lura/v2/config"
	"github.com/luraproject/lura/v2/logging"
	"github.com/luraproject/lura/v2/proxy"
	router "github.com/luraproject/lura/v2/router/gin"
	"strings"
)

func NewHandlerFactory(
	logger logging.Logger,
	factory router.HandlerFactory,
) router.HandlerFactory {
	return func(config *config.EndpointConfig, proxy proxy.Proxy) gin.HandlerFunc {
		// TODO: Check if middleware is enabled for endpoint.
		// TODO: Read configuration of middleware.
		handler := factory(config, proxy)
		loadAaaConf()
		return func(ginCtx *gin.Context) {
			userClaims := GetClaimsBasedOnTokenPolicy(ginCtx)
			ginCtx.Params = append(
				ginCtx.Params, gin.Param{
					Key:   "JWT.user_id",
					Value: userClaims.UserID,
				},
			)
			ginCtx.Params = append(
				ginCtx.Params, gin.Param{
					Key:   "JWT.email",
					Value: userClaims.Email,
				},
			)
			handler(ginCtx)
		}
	}
}

func GetClaimsBasedOnTokenPolicy(ginCtx *gin.Context) Claims {
	var userClaims Claims
	accessToken := strings.Replace(ginCtx.Request.Header.Get("Authorization"), "Bearer ", "", 1)
	if len(accessToken) > 0 {
		if strings.Contains(accessToken, ".") {
			userClaims = validateAaa(ginCtx, accessToken)
		} else {
			userClaims = validateOwin(ginCtx, accessToken)
		}
	}

	return userClaims
}
