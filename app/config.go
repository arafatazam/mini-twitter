package app

type Config struct {
	Port string 		`env:"PORT" envDefault:"3000"`
	Secret string		`env:"SECRET" envDefault:"secret"`
	DBHost string		`env:"DB_HOST" envDefault:"localhost"`
	DBPort string		`env:"DB_PORT" envDefault:"5432"`
	DBName string		`env:"DB_NAME" envDefault:"twitter"`
	DBUser string		`env:"DB_USER" envDefault:"twitter"`
	DBPass string		`env:"DB_PASS" envDefault:"twitter"`
}