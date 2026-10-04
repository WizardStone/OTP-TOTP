package storage

import (
	"errors"
	"fmt"

	"github.com/zalando/go-keyring"
)

func addKey(username string, secret string, appID string) error {
	if username == "" || secret == "" {
		return errors.New("Username and Password cannot be empty")
	}
	KeyringError := keyring.Set(appID, username, secret)
	if KeyringError != nil {
		return fmt.Errorf("Failed to add keys to password manager: %w", KeyringError)
	}
	return nil
}

func getKey(username string, appID string) (string, error) {
	// Check if username is empty
	if username == "" {
		return "", errors.New("Username cannot be empty")
	}
	if appID == "" {
		fmt.Println("Internal Error, contact devlopers on github: https://github.com/wizardStone/OTP-TOTP and open an issue stating the problem (appID is empty storage api)")
		return "", errors.New("Internal Error, contact devlopers on github: https://github.com/wizardStone/OTP-TOTP")
	}
	secret, ReadError :=  keyring.Get(appID, username)

	if ReadError != nil {
		return "", fmt.Errorf("Storage read error: %w", ReadError)
	}
	return secret, ReadError
}