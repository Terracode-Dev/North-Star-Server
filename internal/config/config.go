package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config struct holds all configuration settings
type Config struct {
	AppName    string
	Port       string
	DBString   string
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
	// AWSRegion  string
	MainBranchId int
	JwtExpHour   int
}

// LoadConfig loads environment variables from a `.env` file
func LoadConfig() *Config {
	// Load environment variables from .env file (only for development)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Convert DB_PORT to integer
	dbPort, err := strconv.Atoi(getEnv("DB_PORT", "5432"))
	if err != nil {
		log.Fatalf("Invalid DB_PORT: %v", err)
	}

	mainbranchid, err := strconv.Atoi(getEnv("MAINBRANCHID", "1"))
	if err != nil {
		log.Fatalf("Invalid DB_PORT: %v", err)
	}

	jwtexthour, err := strconv.Atoi(getEnv("JWT_EXP_HOUR", "24"))
	if err != nil {
		log.Fatalf("Invalid DB_PORT: %v", err)
	}

	// awsregion := getEnv("AWS_REGION", "")
	// if awsregion == "" {
	// 	log.Println("aws region not found in the env")
	// }

	return &Config{
		AppName:    getEnv("APP_NAME", "EchoApp"),
		Port:       getEnv("PORT", ":8080"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBString:   getEnv("DB_STR", "user:0000@tcp(127.0.0.1:3306)/NS_db"),
		DBPort:     dbPort,
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "mydb"),
		JWTSecret:  getEnv("JWT_SECRET", "supersecret"),
		// AWSRegion:  awsregion,
		MainBranchId: mainbranchid,
		JwtExpHour:   jwtexthour,
	}
}

// Helper function to get environment variables with a fallback value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
