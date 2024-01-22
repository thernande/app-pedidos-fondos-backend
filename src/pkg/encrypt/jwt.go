package encrypt

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/thernande/app-pedidos-fondos-backend/internal/config/enviroment"
	"github.com/thernande/app-pedidos-fondos-backend/pkg/appLogs"
	"github.com/thernande/app-pedidos-fondos-backend/pkg/errores"
)

var env = enviroment.New()

func getSecretKey() []byte {
	env.Load()
	return []byte(env.AccessTokenSecret)
}

var logger = &appLogs.Logger{}

func GenerateJWTLoginUsuario(id uint, Document string, tokenString *string, chann chan errores.ChannelErrors) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Header["typ"] = "JWT"
	token.Header["alg"] = "HS256"
	claims := token.Claims.(jwt.MapClaims)
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Duration(env.AccessTokenExpiryHour) * time.Hour).Unix()
	claims["Id"] = id
	claims["Document"] = Document

	TS, err := token.SignedString(getSecretKey())
	if err != nil {
		logger.ErrorLogPrint(err.Error())
		chann <- errores.ChannelErrors{Condition: true, Error: err.Error()}
		(*tokenString) = ""
	}
	(*tokenString) = TS
	chann <- errores.ChannelErrors{Condition: false, Error: "todo bien"}
}

func JwtAuthMiddleware(secret string, page string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := strings.Replace(t[1], "\"", "", 2)
			authorized, err := IsAuthorized(authToken, secret)
			if err != nil {
				logger.ErrorLogPrint(err.Error())
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msj": err.Error()})
				return
			}
			if authorized {
				userID, err := ExtractIDFromToken(authToken, secret)
				if err != nil {
					logger.ErrorLogPrint(err.Error())
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msj": err.Error()})
					return
				}
				c.Set("x-user-id", userID)
				c.Next()
				return
			}
			logger.ErrorLogPrint(err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msj": err.Error()})
			return
		}
		logger.InfoLogPrint("no lleva cabecera de autorizacion")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msj": "no autorizado"})
	}
}

func IsAuthorized(requestToken string, secret string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			logger.WarnLogPrint(fmt.Errorf("unexpected signing method: %v", token.Header["alg"]))
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		claims := token.Claims.(jwt.MapClaims)
		if claims["exp"].(float64) < float64(time.Now().Unix()) {
			logger.InfoLogPrint("token Expirado")
			return false, fmt.Errorf("token Expirado")
		}
		return []byte(secret), nil
	})

	if err != nil {
		logger.ErrorLogPrint(err.Error())
		return false, err
	}
	return true, nil
}

func ExtractIDFromToken(requestToken string, secret string) (string, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			logger.ErrorLogPrint(fmt.Errorf("unexpected signing method: %v", token.Header["alg"]))
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		logger.ErrorLogPrint(err.Error())
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		logger.ErrorLogPrint(err.Error())
		return "", fmt.Errorf("invalid Token")
	}

	return claims["Id"].(string), nil
}
