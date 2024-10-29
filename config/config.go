package config

import (
	"test/handlers"
	"test/infra"
	"test/repo/db"
	"test/usecase"

	"test/domain/general"

	"github.com/spf13/viper"
)

func GetCoreConfig() (*general.SectionService, error) {
	viper.SetConfigName("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	data := &general.SectionService{
		App: general.AppAccount{
			Name:         viper.GetString("APP.NAME"),
			Environtment: viper.GetString("APP.ENV"),
			URL:          viper.GetString("APP.URL"),
			Port:         viper.GetString("APP.PORT"),
			SecretKey:    viper.GetString("APP.KEY"),
			Endpoint:     viper.GetString("APP.ENDPOINT"),
		},
		Database: general.DatabaseAccount{
			Read: general.DBDetailAccount{
				Username:     viper.GetString("DATABASE.READ.USERNAME"),
				Password:     viper.GetString("DATABASE.READ.PASSWORD"),
				URL:          viper.GetString("DATABASE.READ.URL"),
				Port:         viper.GetString("DATABASE.READ.PORT"),
				DBName:       viper.GetString("DATABASE.READ.NAME"),
				MaxIdleConns: viper.GetInt("DATABASE.READ.MAXIDLECONNS"),
				MaxOpenConns: viper.GetInt("DATABASE.READ.MAXOPENCONNS"),
				MaxLifeTime:  viper.GetInt("DATABASE.READ.MAXLIFETIME"),
				Timeout:      viper.GetString("DATABASE.READ.TIMEOUT"),
				SSLMode:      viper.GetString("DATABASE.READ.SSL_MODE"),
			},
			Write: general.DBDetailAccount{
				Username:     viper.GetString("DATABASE.WRITE.USERNAME"),
				Password:     viper.GetString("DATABASE.WRITE.PASSWORD"),
				URL:          viper.GetString("DATABASE.WRITE.URL"),
				Port:         viper.GetString("DATABASE.WRITE.PORT"),
				DBName:       viper.GetString("DATABASE.WRITE.NAME"),
				MaxIdleConns: viper.GetInt("DATABASE.WRITE.MAXIDLECONNS"),
				MaxOpenConns: viper.GetInt("DATABASE.WRITE.MAXOPENCONNS"),
				MaxLifeTime:  viper.GetInt("DATABASE.WRITE.MAXLIFETIME"),
				Timeout:      viper.GetString("DATABASE.WRITE.TIMEOUT"),
				SSLMode:      viper.GetString("DATABASE.READ.SSL_MODE"),
			},
		},
		Authorization: general.AuthAccount{
			JWT: general.JWTCredential{
				IsActive:              viper.GetBool("AUTHORIZATION.JWT.IS_ACTIVE"),
				AccessTokenSecretKey:  viper.GetString("AUTHORIZATION.JWT.ACCESS_TOKEN_SECRET_KEY"),
				AccessTokenDuration:   viper.GetInt("AUTHORIZATION.JWT.ACCESS_TOKEN_DURATION"),
				RefreshTokenSecretKey: viper.GetString("AUTHORIZATION.JWT.REFRESH_TOKEN_SECRET_KEY"),
				RefreshTokenDuration:  viper.GetInt("AUTHORIZATION.JWT.REFRESH_TOKEN_DURATION"),
			},
		},
	}

	return data, nil
}

func NewRepoContext(conf *general.SectionService) (handlers.Handler, error) {
	var handler handlers.Handler

	// Init DB Read Connection.
	dbRead := infra.NewDB()
	dbRead.ConnectDB(&conf.Database.Read)
	if dbRead.Err != nil {
		return handler, dbRead.Err
	}

	// Init DB Write Connection.
	dbWrite := infra.NewDB()
	dbWrite.ConnectDB(&conf.Database.Write)
	if dbWrite.Err != nil {
		return handler, dbWrite.Err
	}

	dbList := &infra.DatabaseList{
		Backend: infra.DatabaseType{
			Read:  &dbRead,
			Write: &dbWrite,
		},
	}

	repo := db.NewDBRepo(dbList)
	usecase := usecase.NewUsecase(conf, repo)
	handler = handlers.NewHandler(usecase, conf)

	return handler, nil
}
