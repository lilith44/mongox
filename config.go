package mongox

import "go.mongodb.org/mongo-driver/mongo/options"

type Config struct {
	Scheme string             `yaml:"scheme" validate:"required"`
	URI    string             `yaml:"uri" validate:"required,uri"`
	Auth   options.Credential `yaml:"auth"`
}
