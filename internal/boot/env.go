package boot

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	// Load environment variables. If STAGE is prod, do not load .env file
	if os.Getenv("STAGE") != "prod" {
		err := godotenv.Load()
		if err != nil {
			return fmt.Errorf("error loading .env file")
		}
	}
	return nil
}
