package mongox

import (
	"context"

	"github.com/lilith44/mongox/timer"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	database *mongo.Database

	logger      timer.Logger
	collections map[string]*collection
}

func New(c Config, logger timer.Logger) (*Mongo, error) {
	auth := options.Credential{
		AuthMechanism:           c.Auth.AuthMechanism,
		AuthMechanismProperties: c.Auth.AuthMechanismProperties,
		AuthSource:              c.Auth.AuthSource,
		Username:                c.Auth.Username,
		Password:                c.Auth.Password,
		PasswordSet:             c.Auth.PasswordSet,
	}

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(c.URI).SetAuth(auth))
	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return &Mongo{
		database:    client.Database(c.Scheme),
		logger:      logger,
		collections: make(map[string]*collection),
	}, nil
}
