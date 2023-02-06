package env

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var (
	Debug                 bool
	BaseUrl               string
	MongoDBUrl            string
	MongoDBName           string
	MongoDBUser           string
	MongoDBPass           string
	MongoDBRequestTimeout time.Duration
	RsaPublicKey          string
	RsaPrivateKey         string
	JwtTokenLife          string
	EncryptKey            string
	HttpClientTimeOut     time.Duration
	RetryAttempts         int
	DataPath              string
	LogPath               string
)

func Init() {
	DebugStr := os.Getenv("DEBUG")
	if DebugStr != "" {
		debug, err := strconv.ParseBool(DebugStr)
		if err != nil {
			log.Fatal(err)
		}
		Debug = debug
	}
	BaseUrl = os.Getenv("BASE_URL")
	if BaseUrl == "" {
		BaseUrl = "http://localhost:8080"
	}
	MongoDBUrl = os.Getenv("MONGODB_URL")
	MongoDBName = os.Getenv("MONGODB_NAME")
	MongoDBRequestTimeoutStr := os.Getenv("MONGODB_REQUEST_TIMEOUT")
	if MongoDBRequestTimeoutStr != "" {
		d, err := strconv.ParseUint(MongoDBRequestTimeoutStr, 10, 64)
		if err != nil {
			panic(err)
		}
		MongoDBRequestTimeout = time.Duration(d) * time.Second
	}
	MongoDBUser = os.Getenv("MONGODB_USER")
	MongoDBPass = os.Getenv("MONGODB_PASS")
	RsaPublicKey = os.Getenv("RSA_PUBLIC_KEY")
	RsaPrivateKey = os.Getenv("RSA_PRIVATE_KEY")
	JwtTokenLife = os.Getenv("JWT_TOKEN_LIFE")
	EncryptKey = os.Getenv("ENCRYPT_KEY")

	HttpClientTimeOutStr := os.Getenv("HTTP_CLIENT_TIMEOUT")
	if HttpClientTimeOutStr != "" {
		d, err := strconv.ParseUint(HttpClientTimeOutStr, 10, 64)
		if err != nil {
			panic(err)
		}
		HttpClientTimeOut = time.Duration(d) * time.Second
	}
	RetryAttemptsStr := os.Getenv("RETRY_ATTEMPTS")
	if RetryAttemptsStr != "" {
		i, err := strconv.Atoi(RetryAttemptsStr)
		if err != nil {
			panic(err)
		}
		RetryAttempts = i
	} else {
		RetryAttempts = 1
	}

	DataPath = os.Getenv("DATA_PATH")
	if DataPath == "" {
		current, _ := os.Getwd()
		DataPath = filepath.Join(current, "data")
	}

	LogPath = os.Getenv("LOG_PATH")
	if LogPath == "" {
		current, _ := os.Getwd()
		LogPath = filepath.Join(current, "logs")
	}

}
