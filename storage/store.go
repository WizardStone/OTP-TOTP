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