package pkg

type ClerkUserCreated struct {
	Object string `json:"object"`
	Type   string `json:"type"`
	Data   struct {
		Id             string         `json:"id"`
		FirstName      string         `json:"first_name"`
		LastName       string         `json:"last_name"`
		EmailAddresses []EmailAddress `json:"email_addresses"`
	} `json:"data"`
}

type EmailAddress struct {
	EmailAddress string `json:"email_address"`
}
