package entity

import (
	"log"
	"os"
	"strconv"

	"github.com/runetale/runetale-handshake-server/utility"
)

type TurnConfig struct {
	URL                  string
	Username             string
	Password             string
	Secret               string
	CredentialsTTL       utility.Duration
	TimeBasedCredentials bool
}

type StunConfig struct {
	URL      string
	Username string
	Password string
}

type Env struct {
	// STUN/TURN ServerConfig
	StunConfig *StunConfig
	TurnConfig *TurnConfig

	// Application Port
	Port string

	// Log
	LogLevel string
	LogPath  string
	LogFmt   string

	IsDev bool
}

func NewEnv() *Env {
	stunURL := os.Getenv("STUN_URL")
	stunUsername := os.Getenv("STUN_USERNAME")
	stunPassword := os.Getenv("STUN_PASSWORD")

	turnURL := os.Getenv("TURN_URL")
	turnUsername := os.Getenv("TURN_USERNAME")
	turnPassword := os.Getenv("TURN_PASSWORD")

	port := os.Getenv("PORT")

	logLevel := os.Getenv("LOG_LEVEL")
	logPath := os.Getenv("LOG_PATH")
	logFmt := os.Getenv("LOG_FORMAT")

	isdevstr := os.Getenv("IS_DEV")
	isDev, err := strconv.ParseBool(isdevstr)
	if err != nil {
		isDev = false
	}

	return &Env{
		// STUN/TURN ServerConfig
		StunConfig: &StunConfig{
			URL:      stunURL,
			Username: stunUsername,
			Password: stunPassword,
		},
		TurnConfig: &TurnConfig{
			URL:      turnURL,
			Username: turnUsername,
			Password: turnPassword,
			// TODO: fix later CredentialsTTL & TimeBasedCredentials
			CredentialsTTL:       utility.Duration{Duration: 0},
			TimeBasedCredentials: false,
		},

		Port: port,

		// Log
		LogLevel: logLevel,
		LogPath:  logPath,
		LogFmt:   logFmt,

		// Redis
		IsDev: isDev,
	}
}

func (e *Env) GetLogFile() *os.File {
	f, err := os.Create(e.LogPath)
	if err != nil {
		log.Fatalf("failed to create log file, %v", err)
		log.Printf("please set the specified log format.\n`/dev/stderr` or `/dev/stdout` or `your favorite path`")
	}
	return f
}
