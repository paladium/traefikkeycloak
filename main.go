// Package traefikkeycloak main package
package traefikkeycloak

import (
	"context"
	"errors"
	"net/http"

	"github.com/Nerzal/gocloak"
)

// Config main config
type Config struct {
	URL         string `json:"url"`
	Token       string `json:"token"`
	Realm       string `json:"realm"`
	ParsedToken string `json:"parsedToken"`
}

// CreateConfig make a new config
func CreateConfig() *Config {
	return &Config{
		URL:         "",
		Token:       "",
		Realm:       "",
		ParsedToken: "",
	}
}

type keycloak struct {
	name   string
	client gocloak.GoCloak
	next   http.Handler
	config *Config
}

// New makes a plugin instance
func New(_ context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if config.URL == "" {
		return nil, errors.New("The URL is required")
	}
	client := gocloak.NewClient(config.URL)
	return &keycloak{
		name:   name,
		next:   next,
		config: config,
		client: client,
	}, nil
}

func (k *keycloak) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	cookies := req.Cookies()
	var token *string
	for _, cookie := range cookies {
		if cookie.Name == k.config.Token {
			token = &cookie.Value
			break
		}
	}
	if token == nil {
		res.WriteHeader(http.StatusForbidden)
		return
	}
	// _, claims, err := k.client.DecodeAccessToken(*token, k.config.Realm)
	// if err != nil {
	// 	res.WriteHeader(http.StatusForbidden)
	// 	return
	// }
	//Pass the claims as a header encoded in base64
	// jsonClaims, err := json.Marshal(claims)
	// if err != nil {
	// 	res.WriteHeader(http.StatusForbidden)
	// 	return
	// }
	// encodedClaims := base64.StdEncoding.EncodeToString([]byte(jsonClaims))
	// res.Header().Add(k.config.ParsedToken, encodedClaims)
	res.Header().Add(k.config.ParsedToken, "test")
	k.next.ServeHTTP(res, req)
}
