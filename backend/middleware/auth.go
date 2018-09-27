package middleware

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"log"
	"strings"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"github.com/ja1984/cogCMS/backend/config"
	"github.com/ja1984/cogCMS/backend/services"
)

func VerifyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			c.Next()
			return
		}

		if *config.Environment == "local" && *config.DisableAuthentication {
			c.Next()
			return
		}

		tokenHeader := c.GetHeader("Authorization")

		if tokenHeader == "" {
			respondWithError(401, "API token required", c)
			return
		}

		token, err := fromAuthHeader(tokenHeader)

		if err != nil {
			log.Printf("error fromAuthHeader: %v\n", err)
			respondWithError(401, "Invalid API token", c)
			return
		}

		if *config.Environment != "local" {
			idToken, err := services.FirebaseClient.VerifyIDToken(context.Background(), token)

			if err != nil {
				log.Printf("error VerifyIDToken: %v\n", err)
				respondWithError(401, "Invalid VerifyIDToken token", c)
				return
			}
			c.Set("token", *idToken)
			c.Next()
			return
		}

		idToken, err := getUnverifiedDevToken(token)
		if err != nil {
			log.Printf("error getUnverifiedDevToken: %v\n", err)
			respondWithError(401, "Invalid getUnverifiedDevToken token", c)
			return
		}
		c.Set("token", *idToken)
		c.Next()
	}
}

func fromAuthHeader(authHeader string) (string, error) {
	if authHeader == "" {
		return "", nil // No error, just no token
	}

	// TODO: Make this a bit more robust, parsing-wise
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", errors.New("Authorization header format must be Bearer {token}")
	}

	return authHeaderParts[1], nil
}

func respondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}

	c.JSON(code, resp)
	c.Abort()
}

func getUnverifiedDevToken(idToken string) (*auth.Token, error) {
	segments := strings.Split(idToken, ".")

	var (
		header  jwtHeader
		payload auth.Token
		claims  map[string]interface{}
	)
	if err := decode(segments[0], &header); err != nil {
		return nil, err
	}
	if err := decode(segments[1], &payload); err != nil {
		return nil, err
	}
	if err := decode(segments[1], &claims); err != nil {
		return nil, err
	}
	// Delete standard claims from the custom claims maps.
	for _, r := range []string{"iss", "aud", "exp", "iat", "sub", "uid"} {
		delete(claims, r)
	}
	payload.Claims = claims

	payload.UID = payload.Subject
	return &payload, nil
}

func decode(segment string, i interface{}) error {
	decoded, err := base64.RawURLEncoding.DecodeString(segment)
	if err != nil {
		return err
	}
	return json.NewDecoder(bytes.NewBuffer(decoded)).Decode(i)
}

type jwtHeader struct {
	Algorithm string `json:"alg"`
	Type      string `json:"typ"`
	KeyID     string `json:"kid,omitempty"`
}
