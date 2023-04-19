package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/database"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/logger"

	// repo
	_announcementsTeamsRepo "gitlab.com/chaihanij/evat/app/layers/repositories/announcements"
	_omiseRepo "gitlab.com/chaihanij/evat/app/layers/repositories/omise"

	_assignmentTeamsRepo "gitlab.com/chaihanij/evat/app/layers/repositories/assignment_teams"
	_assignmentsRepo "gitlab.com/chaihanij/evat/app/layers/repositories/assignments"
	_configRepo "gitlab.com/chaihanij/evat/app/layers/repositories/config"
	_considerationRepo "gitlab.com/chaihanij/evat/app/layers/repositories/consideration"
	_emailRepo "gitlab.com/chaihanij/evat/app/layers/repositories/email"
	_emailcontact "gitlab.com/chaihanij/evat/app/layers/repositories/emailcontact"
	_fildracteamsRepo "gitlab.com/chaihanij/evat/app/layers/repositories/field_race_teams"
	_fieldracesRepo "gitlab.com/chaihanij/evat/app/layers/repositories/field_races"
	_filesRepo "gitlab.com/chaihanij/evat/app/layers/repositories/files"
	_membersRepo "gitlab.com/chaihanij/evat/app/layers/repositories/members"
	_teamsRepo "gitlab.com/chaihanij/evat/app/layers/repositories/teams"
	_userRepo "gitlab.com/chaihanij/evat/app/layers/repositories/users"
	_visitRepo "gitlab.com/chaihanij/evat/app/layers/repositories/visit"

	// use case
	_announcementsUseCase "gitlab.com/chaihanij/evat/app/layers/usecase/announcements"
	_assignmentTeamUseCase "gitlab.com/chaihanij/evat/app/layers/usecase/assignment_teams"
	_assignmentsUseCase "gitlab.com/chaihanij/evat/app/layers/usecase/assignments"
	_configUseCase "gitlab.com/chaihanij/evat/app/layers/usecase/config"
	_considerationUseCase "gitlab.com/chaihanij/evat/app/layers/usecase/consideration"
	_emailUseCase "gitlab.com/chaihanij/evat/app/layers/usecase/email"
	_emailcontactUseCase "gitlab.com/chaihanij/evat/app/layers/usecase/emailcontact"
	_fildracteamsUseCas "gitlab.com/chaihanij/evat/app/layers/usecase/field_race_teams"
	_fieldracesUseCase "gitlab.com/chaihanij/evat/app/layers/usecase/field_races"
	_filesUseCase "gitlab.com/chaihanij/evat/app/layers/usecase/files"
	_memberUseCase "gitlab.com/chaihanij/evat/app/layers/usecase/members"
	_teamsUseCase "gitlab.com/chaihanij/evat/app/layers/usecase/teams"
	_usersUseCase "gitlab.com/chaihanij/evat/app/layers/usecase/users"
	_visitUseCase "gitlab.com/chaihanij/evat/app/layers/usecase/visit"

	// Deliveries
	_announcementsHttp "gitlab.com/chaihanij/evat/app/layers/deliveries/http/announcements"
	_assignmentTeamHttp "gitlab.com/chaihanij/evat/app/layers/deliveries/http/assignment_teams"
	_assignmentsHttp "gitlab.com/chaihanij/evat/app/layers/deliveries/http/assignments"
	_configHttp "gitlab.com/chaihanij/evat/app/layers/deliveries/http/config"
	_considerationHttp "gitlab.com/chaihanij/evat/app/layers/deliveries/http/consideration"
	_emailHttp "gitlab.com/chaihanij/evat/app/layers/deliveries/http/email"
	_emailcontactHttp "gitlab.com/chaihanij/evat/app/layers/deliveries/http/emailcontact"
	__fildracteamsHttp "gitlab.com/chaihanij/evat/app/layers/deliveries/http/field_race_teams"
	_fieldracesHttp "gitlab.com/chaihanij/evat/app/layers/deliveries/http/field_races"
	_filesHttp "gitlab.com/chaihanij/evat/app/layers/deliveries/http/files"
	_healthCheck "gitlab.com/chaihanij/evat/app/layers/deliveries/http/health_check"
	_membersHttp "gitlab.com/chaihanij/evat/app/layers/deliveries/http/members"
	_omiseHttp "gitlab.com/chaihanij/evat/app/layers/deliveries/http/omise"
	_teamsHttp "gitlab.com/chaihanij/evat/app/layers/deliveries/http/teams"
	_usersHttp "gitlab.com/chaihanij/evat/app/layers/deliveries/http/users"
	_visitHttp "gitlab.com/chaihanij/evat/app/layers/deliveries/http/visit"

	middlewares "gitlab.com/chaihanij/evat/app/layers/deliveries/http/middlewares"

	// swagger
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	docs "gitlab.com/chaihanij/evat/app/docs/swagger"
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
	os.MkdirAll(filepath.Join(env.DataPath, "teams", "slips"), os.ModePerm)
	os.MkdirAll(filepath.Join(env.DataPath, "members", "images"), os.ModePerm)
	os.MkdirAll(filepath.Join(env.DataPath, "members", "documents"), os.ModePerm)
	os.MkdirAll(filepath.Join(env.DataPath, "assignments", "images"), os.ModePerm)
	os.MkdirAll(filepath.Join(env.DataPath, "assignments", "documents"), os.ModePerm)
	os.MkdirAll(filepath.Join(env.DataPath, "assignments", "template"), os.ModePerm)

	db := database.ConnectMongoDB()

	// init repo
	assignmentsRepo := _assignmentsRepo.InitRepo(db)
	filesRepo := _filesRepo.InitRepo(db)
	membersRepo := _membersRepo.InitRepo(db)
	teamsRepo := _teamsRepo.InitRepo(db)
	userRepo := _userRepo.InitRepo(db)
	assignmentTeamsRepo := _assignmentTeamsRepo.InitRepo(db)
	announcementsTeamsRepo := _announcementsTeamsRepo.InitRepo(db)
	fildracteamsRepo := _fildracteamsRepo.InitRepo(db)
	considerationRepo := _considerationRepo.InitRepo(db)
	fieldraces := _fieldracesRepo.InitRepo(db)
	visitRepo := _visitRepo.InitRepo(db)
	emailRepo := _emailRepo.InitRepo(db)
	emailcontactRepo := _emailcontact.InitRepo(db)
	omiseRepo := _omiseRepo.InitRepo(env.OMISEPublicKey, env.OMISEPrivateKey)
	configRepo := _configRepo.InitRepo(db)
	// config repo
	assignmentsRepo.Config()
	filesRepo.Config()
	membersRepo.Config()
	teamsRepo.Config()
	userRepo.Config()
	assignmentTeamsRepo.Config()
	announcementsTeamsRepo.Config()
	fildracteamsRepo.Config()
	considerationRepo.Config()
	fieldraces.Config()
	visitRepo.Config()
	emailRepo.Config()
	emailcontactRepo.Config()
	configRepo.Config()

	// usecase
	assignmentsUseCase := _assignmentsUseCase.InitUseCase(assignmentsRepo, filesRepo)
	userUseCase := _usersUseCase.InitUseCase(userRepo)
	teamsUseCase := _teamsUseCase.InitUseCase(teamsRepo, userRepo, membersRepo, filesRepo, assignmentTeamsRepo, omiseRepo)
	memberUseCase := _memberUseCase.InitUseCase(membersRepo, filesRepo)
	filesUseCase := _filesUseCase.InitUseCase(filesRepo)
	announcementsUseCase := _announcementsUseCase.InitUseCase(announcementsTeamsRepo)
	fildracteamsUseCas := _fildracteamsUseCas.InitUseCase(fildracteamsRepo)
	considerationUseCase := _considerationUseCase.InitUseCase(considerationRepo)
	fieldracesUseCase := _fieldracesUseCase.InitUseCase(fieldraces)
	assignmentTeamUseCase := _assignmentTeamUseCase.InitUseCase(assignmentTeamsRepo)
	visitUseCase := _visitUseCase.InitUseCase(visitRepo)
	emailUseCase := _emailUseCase.InitUseCase(emailRepo)
	emailcontactUseCase := _emailcontactUseCase.InitUseCase(emailcontactRepo)
	configUseCase := _configUseCase.InitUseCase(configRepo)

	//
	ginEngine := gin.New()
	ginEngine.Use(gin.Logger())
	ginEngine.Use(gin.Recovery())
	ginEngine.Use(middlewares.CORSMiddleware())
	// swagerr
	docs.SwaggerInfo.Host = env.Host
	docs.SwaggerInfo.BasePath = env.BasePath
	ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

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
	_announcementsHttp.NewEndpointHttpHandler(ginEngine, authMiddleware, announcementsUseCase)
	__fildracteamsHttp.NewEndpointHttpHandler(ginEngine, authMiddleware, fildracteamsUseCas)
	_considerationHttp.NewEndpointHttpHandler(ginEngine, authMiddleware, considerationUseCase)
	_fieldracesHttp.NewEndpointHttpHandler(ginEngine, authMiddleware, fieldracesUseCase)
	_assignmentTeamHttp.NewEndpointHttpHandler(ginEngine, authMiddleware, assignmentTeamUseCase)
	_visitHttp.NewEndpointHttpHandler(ginEngine, authMiddleware, visitUseCase)
	_emailHttp.NewEndpointHttpHandler(ginEngine, emailUseCase)
	_emailcontactHttp.NewEndpointHttpHandler(ginEngine, emailcontactUseCase)
	_omiseHttp.NewEndpointHttpHandler(ginEngine, teamsUseCase)
	_configHttp.NewEndpointHttpHandler(ginEngine, authMiddleware, configUseCase)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: ginEngine,
	}
	err := srv.ListenAndServe()
	if err != http.ErrServerClosed {
		log.Fatal("error :", err)
	}
}
