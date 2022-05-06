package configs

type (
	DB struct {
		Host string
		Port int
		User string
		Password string
		Name string
		ForceTLS bool
	}

	Logger struct {
	}

	App struct {
		Port int
		Enviroment string
		Production bool
		JWTSecret string
		SecureCookieHashKey string
		SecureCookieBlockKey string
 		DB *DB
		Logger *Logger
	}
)