package adapters

import (
	"github.com/gin-gonic/gin"
	"github.com/intwone/ddd-golang/internal/constants"
	"github.com/intwone/ddd-golang/internal/infra/cryptography"
	er "github.com/intwone/ddd-golang/internal/presentation/errors"
)

func GinGonicMiddlewareAdapter(crypto cryptography.CryptographyInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		userID, err := crypto.Decrypt(token)

		if err != nil {
			switch err.Error() {
			case constants.InvalidTokenError, constants.TokenClaimsError, constants.FieldNotFoundOnToken:
				restErr := er.NewUnauthorizedError(constants.InvalidTokenError)
				c.JSON(restErr.Code, restErr)
				c.Abort()
				return
			}
		}

		c.Set("userID", *userID)
	}
}
