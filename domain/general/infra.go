package general

type SectionService struct {
	App           AppAccount      `json:",omitempty"`
	Database      DatabaseAccount `json:",omitempty"`
	Authorization AuthAccount     `json:",omitempty"`
}

type AppAccount struct {
	Name         string `json:",omitempty"`
	Environtment string `json:",omitempty"`
	URL          string `json:",omitempty"`
	Port         string `json:",omitempty"`
	SecretKey    string `json:",omitempty"`
	Endpoint     string `json:",omitempty"`
}

type DatabaseAccount struct {
	Read  DBDetailAccount `json:",omitempty"`
	Write DBDetailAccount `json:",omitempty"`
}
type DBDetailAccount struct {
	Username     string `json:",omitempty"`
	Password     string `json:",omitempty"`
	URL          string `json:",omitempty"`
	Port         string `json:",omitempty"`
	DBName       string `json:",omitempty"`
	MaxIdleConns int    `json:",omitempty"`
	MaxOpenConns int    `json:",omitempty"`
	MaxLifeTime  int    `json:",omitempty"`
	Timeout      string `json:",omitempty"`
	SSLMode      string `json:",omitempty"`
}

type AuthAccount struct {
	JWT JWTCredential `json:",omitempty"`
}

type JWTCredential struct {
	IsActive              bool   `json:",omitempty"`
	AccessTokenSecretKey  string `json:",omitempty"`
	AccessTokenDuration   int    `json:",omitempty"`
	RefreshTokenSecretKey string `json:",omitempty"`
	RefreshTokenDuration  int    `json:",omitempty"`
}
