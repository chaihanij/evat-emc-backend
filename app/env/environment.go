package env

import (
	"log"
	"os"
	"strconv"
	"time"
)

var (
	Debug                 bool
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
	}

}
