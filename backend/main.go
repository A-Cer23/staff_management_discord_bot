package main

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/A-Cer23/adminbot-backend/db"
	"github.com/A-Cer23/adminbot-backend/routes"
)

const (
	defaultPortNumber  = "9090"
	defaultDatabaseURL = "postgres://user:password@:5432/adminstrationbotdb"
)

func main() {

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	configureMultiLogger()

	portNumber := os.Getenv("PORT_NUMBER")
	if len(portNumber) == 0 {
		// err := errors.New(`"PORT_NUMBER" missing from .env file`)
		// log.Fatal().Err(err).Send()
		portNumber = defaultPortNumber
	}

	dbUrl := os.Getenv("DATABASE_URL")
	if len(dbUrl) == 0 {
		// err := errors.New(`"DATABASE_URL" missing from .env file`)
		// log.Fatal().Err(err).Send()
		dbUrl = defaultDatabaseURL
	}

	log.Info().Msg("Attempting to connect to database")

	dbPool, err := db.ConnectDB(dbUrl)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	if err := dbPool.Ping(context.Background()); err != nil {
		log.Fatal().Err(err).Send()
	} else {
		log.Info().Msg("Database connection established")
	}

	router := gin.New()
	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Output: log.Logger,
	}))
	router.SetTrustedProxies(nil)

	log.Info().Msg("Setting up routes")

	routes.SetupRoutes(router, dbPool)

	log.Info().Msg("Routes are setup")

	log.Info().Msgf("Listening on port: %v", portNumber)

	router.Run(":" + portNumber)

}

func configureMultiLogger() {

	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}

	logFile, _ := os.OpenFile(
		"admin_backend.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)

	fileWriter := zerolog.ConsoleWriter{Out: logFile}

	multiWriter := zerolog.MultiLevelWriter(consoleWriter, fileWriter)

	log.Logger = zerolog.New(multiWriter).With().Caller().Timestamp().Logger()
}
