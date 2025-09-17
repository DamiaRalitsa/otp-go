package presenters

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateOTP() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	digits := []rune("0123456789")

	r.Shuffle(len(digits), func(i, j int) {
		digits[i], digits[j] = digits[j], digits[i]
	})

	return string(digits[:6])
}

func StringToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return n
}
