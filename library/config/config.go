package config

import (
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"goecho-boilerplate/internal/constant"
	"log"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	HttpPort    string
	CORSOrigins []string

	// Database
	DBMaster *DB
	DBSlave  *DB

	// JWT Token
	JWTRS256PrivateKey *rsa.PrivateKey
	JWTRS256PubKey     *rsa.PublicKey
	JWTExpiration      int
}

// DB config model
type DB struct {
	ConnectionString string
	MaxIdle          int
	MaxOpen          int
}

// DatabaseConfig stores database configurations.
type configEnv struct {
	Port        string   `envconfig:"APP_PORT" default:"8080"`
	CORSOrigins []string `envconfig:"CORS_ORIGINS"`

	// Database
	DBMasterMaxIdle int    `envconfig:"DBMASTERMAXIDLECONN"`
	DBMasterMaxOpen int    `envconfig:"DBMASTERMAXOPENCONN"`
	DBSlaveMaxIdle  int    `envconfig:"DBSLAVEMAXIDLECONN"`
	DBSlaveMaxOpen  int    `envconfig:"DBSLAVEMAXOPENCONN"`
	DBMasterUser    string `envconfig:"DBMASTERUSER"`
	DBMasterPass    string `envconfig:"DBMASTERPASS"`
	DBMasterHost    string `envconfig:"DBMASTERHOST"`
	DBMasterPort    string `envconfig:"DBMASTERPORT"`
	DBMasterName    string `envconfig:"DBMASTERNAME"`
	DBSlaveUser     string `envconfig:"DBSLAVEUSER"`
	DBSlavePass     string `envconfig:"DBSLAVEPASS"`
	DBSlaveHost     string `envconfig:"DBSLAVEHOST"`
	DBSlavePort     string `envconfig:"DBSLAVEPORT"`
	DBSlaveName     string `envconfig:"DBSLAVENAME"`

	//auth
	JWTRS256PrivateKey string `envconfig:"JWT_RS256_PRIVATE_KEY" required:"true"`
	JWTRS256PubKey     string `envconfig:"JWT_RS256_PUBLIC_KEY" required:"true"`
	JWTExpiration      int    `envconfig:"JWT_EXPIRATION" required:"true"`
}

var appConfig *Config

// ReadConfig populates configurations from environment variables.
func Init() {
	_ = godotenv.Overload()
	var cfg configEnv
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("[Init] failed to map config, %+v\n", err)
	}

	appConfig = &Config{}
	appConfig.HttpPort = cfg.Port
	appConfig.CORSOrigins = cfg.CORSOrigins
	initDB(&cfg)

	jwtPrivateKey, jwtPubKey := getJWTConfig(cfg)

	appConfig.JWTRS256PrivateKey = jwtPrivateKey
	appConfig.JWTRS256PubKey = jwtPubKey
	appConfig.JWTExpiration = cfg.JWTExpiration
}

func initDB(c *configEnv) {
	appConfig.DBMaster = &DB{
		ConnectionString: fmt.Sprintf(
			constant.DBStringConnection,
			c.DBMasterUser,
			c.DBMasterPass,
			c.DBMasterHost,
			c.DBMasterPort,
			c.DBMasterName,
		),
		MaxIdle: c.DBMasterMaxIdle,
		MaxOpen: c.DBMasterMaxOpen,
	}
	appConfig.DBSlave = &DB{
		ConnectionString: fmt.Sprintf(
			constant.DBStringConnection,
			c.DBSlaveUser,
			c.DBSlavePass,
			c.DBSlaveHost,
			c.DBSlavePort,
			c.DBSlaveName,
		),
		MaxIdle: c.DBMasterMaxIdle,
		MaxOpen: c.DBMasterMaxOpen,
	}
}

func getJWTConfig(cfg configEnv) (*rsa.PrivateKey, *rsa.PublicKey) {
	// Decode base64 RS256 JWT Secret
	jwtPrivateKeyPEM, err := base64.StdEncoding.DecodeString(cfg.JWTRS256PrivateKey)
	if err != nil {
		log.Fatalf("Failed to load jwt private key, %+v\n", err)
	}

	jwtPrivateKey, err := jwt.ParseRSAPrivateKeyFromPEM(jwtPrivateKeyPEM)
	if err != nil {
		log.Fatalf("Failed to load jwt private key, %+v\n", err)
	}

	jwtPubKeyPEM, err := base64.StdEncoding.DecodeString(cfg.JWTRS256PubKey)
	if err != nil {
		log.Fatalf("Failed to load jwt public key, %+v\n", err)
	}

	jwtPubKey, err := jwt.ParseRSAPublicKeyFromPEM(jwtPubKeyPEM)

	return jwtPrivateKey, jwtPubKey
}

// Get private instance config
func Get() *Config {
	return appConfig
}
