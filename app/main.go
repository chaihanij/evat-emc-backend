package main

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pseidemann/finish"
	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/database"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/logger"

	// repo
	_assignmentsRepo "gitlab.com/chaihanij/evat/app/layers/repositories/assignments"
	_filesRepo "gitlab.com/chaihanij/evat/app/layers/repositories/files"
	_membersRepo "gitlab.com/chaihanij/evat/app/layers/repositories/members"
	_teamsRepo "gitlab.com/chaihanij/evat/app/layers/repositories/teams"
	_userRepo "gitlab.com/chaihanij/evat/app/layers/repositories/users"

	// use case
	_assignmentsUseCase "gitlab.com/chaihanij/evat/app/layers/usecase/assignments"
	_filesUseCase "gitlab.com/chaihanij/evat/app/layers/usecase/files"
	_memberUseCase "gitlab.com/chaihanij/evat/app/layers/usecase/members"
	_teamsUseCase "gitlab.com/chaihanij/evat/app/layers/usecase/teams"
	_usersUseCase "gitlab.com/chaihanij/evat/app/layers/usecase/users"

	// Deliveries
	_healthCheck "gitlab.com/chaihanij/evat/app/layers/deliveries/http/health_check"

	_assignmentsHttp "gitlab.com/chaihanij/evat/app/layers/deliveries/http/assignments"
	_filesHttp "gitlab.com/chaihanij/evat/app/layers/deliveries/http/files"
	_membersHttp "gitlab.com/chaihanij/evat/app/layers/deliveries/http/members"
	_teamsHttp "gitlab.com/chaihanij/evat/app/layers/deliveries/http/teams"
	_usersHttp "gitlab.com/chaihanij/evat/app/layers/deliveries/http/users"

	middlewares "gitlab.com/chaihanij/evat/app/layers/deliveries/http/middlewares"
)

// @title EVAT eMCS Service
// @version 1.0.0
// @description EVAT eMCS Service.
// @contact.name chaihanij@gmail.com
// @BasePath {{evat-service}}
func main() {
	os.Setenv("TZ", "Asia/Bangkok")
	env.Init()
	logger.Init()
	log.WithFields(log.Fields{
		"BASE_URL":                env.BaseUrl,
		"DEBUG":                   env.Debug,
		"MONGODB_URL":             env.MongoDBUrl,
		"MONGODB_NAME":            env.MongoDBName,
		"MONGODB_USER":            env.MongoDBUser,
		"MONGODB_PASS":            env.MongoDBPass,
		"MONGODB_REQUEST_TIMEOUT": env.MongoDBRequestTimeout,
	}).Debugln("main")

	log.WithFields(log.Fields{
		"JWT_TOKEN_LIFE": env.JwtTokenLife,
		"ENCRYPT_KEY":    env.EncryptKey,
		"PUBLIC_KEY":     env.RsaPublicKey,
		"PRIVAT_KEY":     env.RsaPrivateKey,
	}).Debugln("main")
	// Create data dir
	os.MkdirAll(filepath.Join(env.DataPath, "members", "images"), os.ModePerm)
	os.MkdirAll(filepath.Join(env.DataPath, "members", "documents"), os.ModePerm)
	os.MkdirAll(filepath.Join(env.DataPath, "assignments", "images"), os.ModePerm)
	os.MkdirAll(filepath.Join(env.DataPath, "assignments", "documents"), os.ModePerm)

	db := database.ConnectMongoDB()

	// init repo
	assignmentsRepo := _assignmentsRepo.InitRepo(db)
	filesRepo := _filesRepo.InitRepo(db)
	membersRepo := _membersRepo.InitRepo(db)
	teamsRepo := _teamsRepo.InitRepo(db)
	userRepo := _userRepo.InitRepo(db)

	// config repo
	assignmentsRepo.Config()
	filesRepo.Config()
	membersRepo.Config()
	teamsRepo.Config()
	userRepo.Config()

	// usecase
	assignmentsUseCase := _assignmentsUseCase.InitUseCase(assignmentsRepo, filesRepo)
	userUseCase := _usersUseCase.InitUseCase(userRepo)
	teamsUseCase := _teamsUseCase.InitUseCase(teamsRepo, membersRepo, filesRepo)
	memberUseCase := _memberUseCase.InitUseCase(membersRepo, filesRepo)
	filesUseCase := _filesUseCase.InitUseCase(filesRepo)

	//
	ginEngine := gin.New()
	ginEngine.Use(gin.Logger())
	ginEngine.Use(gin.Recovery())
	ginEngine.Use(middlewares.CORSMiddleware())

	// middlewares
	authMiddleware := middlewares.InitAuthMiddleware(userUseCase)
	// diliveries
	// Health Check
	_healthCheck.NewEndpointHTTPHandler(ginEngine)

	// diliveries
	_assignmentsHttp.NewEndpointHttpHandler(ginEngine, authMiddleware, assignmentsUseCase)
	_usersHttp.NewEndpointHttpHandler(ginEngine, authMiddleware, userUseCase)
	_teamsHttp.NewEndpointHttpHandler(ginEngine, authMiddleware, teamsUseCase)
	_membersHttp.NewEndpointHttpHandler(ginEngine, authMiddleware, memberUseCase)
	_filesHttp.NewEndpointHttpHandler(ginEngine, filesUseCase)
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
