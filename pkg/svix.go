package pkg

import (
	"log"
	"os"

	svix "github.com/svix/svix-webhooks/go"
)

func InitSvix() *svix.Webhook {
	secret := os.Getenv("SVIX_SECRET")

	wh, err := svix.NewWebhook(secret)
	if err != nil {
		log.Fatal("Error: initlisizing svix")
	}

	return wh
}
