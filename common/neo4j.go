package common

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Neo4jConfig struct {
	Host string `env:"HOST" value-name:"HOST" long:"host" description:"Neo4j host name" default:"localhost"`
	Port int    `env:"PORT" value-name:"PORT" long:"port" description:"Neo4j port" default:"7687"`
}

type Neo4jZerologBoltLogger struct {
	Logger *zerolog.Logger
}

func (nl *Neo4jZerologBoltLogger) LogClientMessage(id, msg string, args ...interface{}) {
	nl.Logger.Debug().
		Str("type", "client").
		Str("id", id).
		Msgf(msg, args...)
}

func (nl *Neo4jZerologBoltLogger) LogServerMessage(id, msg string, args ...interface{}) {
	nl.Logger.Debug().
		Str("type", "server").
		Str("id", id).
		Msgf(msg, args...)
}

func GetNeo4jSession(ctx context.Context, driver neo4j.DriverWithContext) neo4j.SessionWithContext {
	return driver.NewSession(ctx, neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: "greenstar",
		BoltLogger:   &Neo4jZerologBoltLogger{Logger: log.Ctx(ctx)},
	})
}
