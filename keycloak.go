// Package pluginkeycloak main package
package pluginkeycloak

import (
	"context"
	"net/http"

	"github.com/Nerzal/gocloak"
)

// Config main config
type Config struct {
	URL   string `json:"url,omitempty"`
	Token string `json:"token,omitempty"`
}

// CreateConfig make a new config
func CreateConfig() *Config {
	return &Config{
		URL:   "",
		Token: "",
	}
}

// Keycloak plugin
type Keycloak struct {
	name   string
	client *gocloak.Client
	next   http.Handler
	config *Config
}

// New makes a plugin instance
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &Keycloak{
		name:   name,
		next:   next,
		config: config,
	}, nil
}

func (k *Keycloak) ServeHTTP(res http.ResponseWriter, req *http.Request) {
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
