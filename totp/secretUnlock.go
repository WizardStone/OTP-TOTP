package totp

import (
	"github.com/xlzd/gotp"
)

func GetCode(Key) {
	totp := gotp.NewDefaultTOTP(Key)

	// 3. Generate a token (e.g., to display to a user or test)
	token, err := totp.Now()
	if err != nil {
		panic(err)
	}
	return token
}
