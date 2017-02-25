package util

import "fmt"

// Env is the environment
type Env string

// To string
func (e Env) String() string {
	return string(e)
}

const (
	// EnvDev development environment
	EnvDev Env = "dev"
	// EnvStage staging environment
	EnvStage Env = "stage"
	// EnvProd production environment
	EnvProd Env = "prod"
)

// Config tells us the port environment and the mongo url
type Config struct {
	Port        int    `envconfig:"MYWIKI_PORT"`
	Environment string `envconfig:"MYWIKI_ENVIRONMENT"`
	MongoURL    string `envconfig:"MYWIKI_MONGO_URL"`
}

// Env gets the environment
func (c Config) Env() (Env, error) {
	switch c.Environment {
	case EnvDev.String():
		return EnvDev, nil
	case EnvStage.String():
		return EnvStage, nil
	case EnvProd.String():
		return EnvProd, nil
	default:
		return Env(""), fmt.Errorf("invalid environment %s", c.Environment)
	}
}
