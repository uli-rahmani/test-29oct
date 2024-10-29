package general

type RenewToken struct {
	Token        string `json:"token"`
	TokenExpired string `json:"token_expired"`
}

type JWTAccess struct {
	AccessToken        string `json:"access"`
	AccessTokenExpired int64  `json:"access_expired"`
	RenewToken         string `json:"renew"`
	RenewTokenExpired  int64  `json:"renew_expired"`
}

type AdditionalDataJWT struct {
	Role string `json:"role"`
}
