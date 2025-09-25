package pkg

import (
	"log"
	"os"

	"github.com/clerk/clerk-sdk-go/v2"
)

func InitClerk() {
	key := os.Getenv("CLERK_API_KEY")
	if key == "" {
		log.Fatal("Invalid clerk key")
	}

	clerk.SetKey(key)
}
