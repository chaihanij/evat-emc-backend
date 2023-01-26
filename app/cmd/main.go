package main

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pseidemann/finish"
	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/database"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/logger"

	// repo
	_userRepo "gitlab.com/chaihanij/evat/app/layers/repositories/users"

	// use case
	_userUseCase "gitlab.com/chaihanij/evat/app/layers/usecase/users"

	// Deliveries
	_healthCheck "gitlab.com/chaihanij/evat/app/layers/deliveries/http/health_check"
	//

	_usersHttp "gitlab.com/chaihanij/evat/app/layers/deliveries/http/users"

	middlewares "gitlab.com/chaihanij/evat/app/layers/deliveries/http/middlewares"
)

// @title EVAT Service
// @version 1.0.0
// @description Evat Service.
// @contact.name chaihanij@gmail.com
// @BasePath {{evat-service}}

func main() {
	os.Setenv("TZ", "Asia/Bangkok")

	env.Init()
	logger.Init()
	log.WithFields(log.Fields{
		"DEBUG":                   env.Debug,
		"MONGODB_URL":             env.MongoDBUrl,
		"MONGODB_NAME":            env.MongoDBName,
		"MONGODB_USER":            env.MongoDBUser,
		"MONGODB_PASS":            env.MongoDBPass,
		"MONGODB_REQUEST_TIMEOUT": env.MongoDBRequestTimeout,
	}).Info("main")

	log.WithFields(log.Fields{
		"JWT_TOKEN_LIFE": env.JwtTokenLife,
		"ENCRYPT_KEY":    env.EncryptKey,
		"PUBLIC_KEY":     env.RsaPublicKey,
		"PRIVAT_KEY":     env.RsaPrivateKey,
	}).Info("main")

	db := database.ConnectMongoDB()

	// init repo
	userRepo := _userRepo.InitRepo(db)

	// config repo
	userRepo.Config()

	// usecase
	userUseCase := _userUseCase.InitUseCase(userRepo)
	ginEngine := gin.New()
	ginEngine.Use(gin.Recovery())
	ginEngine.Use(middlewares.CORSMiddleware())

	// middlewares
	authMiddleware := middlewares.InitAuthMiddleware(userUseCase)
	// diliveries
	// Health Check
	_healthCheck.NewEndpointHTTPHandler(ginEngine)

	// diliveries
	_usersHttp.NewEndpointHttpHandler(ginEngine, authMiddleware, userUseCase)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: ginEngine,
	}

	timeOut, err := strconv.Atoi(os.Getenv("GRACEFUL_TIMEOUT"))
	if err != nil {
		timeOut = 30 // second
	}

	graceful := &finish.Finisher{Timeout: time.Duration(timeOut) * time.Second}
	graceful.Add(srv)

	go func() {
		err := srv.ListenAndServe()
		if err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	graceful.Wait()

}
