// Package pluginkeycloak main package
package pluginkeycloak

import (
	"context"
	"net/http"

	"github.com/Nerzal/gocloak"
)

// Config main config
type Config struct {
	URL   string `json:"url"`
	Token string `json:"token"`
}

// CreateConfig make a new config
func CreateConfig() *Config {
	return &Config{}
}

type keycloak struct {
	name   string
	client *gocloak.Client
	next   http.Handler
	config *Config
}

// New makes a plugin instance
func New(_ context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &keycloak{
		name:   name,
		next:   next,
		config: config,
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
	//TODO Verify the token
	k.next.ServeHTTP(res, req)
}
