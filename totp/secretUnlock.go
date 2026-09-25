package totp

import "github.com/xlzd/gotp"

func GetCode(key string) string {
	t := gotp.NewDefaultTOTP(key)

	// 3. Generate a token (e.g., to display to a user or test)

	return t.Now()
}
