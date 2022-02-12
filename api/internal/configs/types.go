package configs

type (
	DB struct {
		Host string
		Port int
		User string
		Password string
		Name string
	}

	Logger struct {
	}

	App struct {
		Port int
		Enviroment string
		Production bool
		JWTSecret string
 		DB *DB
		Logger *Logger
	}
)