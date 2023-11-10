package mongox

type Config struct {
	Scheme string `yaml:"scheme" validate:"required"`
	URI    string `yaml:"uri" validate:"required,uri"`
	Auth   Auth   `yaml:"auth"`
}

type Auth struct {
	AuthMechanism           string            `yaml:"authMechanism"`
	AuthMechanismProperties map[string]string `yaml:"authMechanismProperties"`
	AuthSource              string            `yaml:"authSource"`
	Username                string            `yaml:"username"`
	Password                string            `yaml:"password"`
	PasswordSet             bool              `yaml:"passwordSet"`
}
