package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

//RandomInt generates a random integer between min and max
func RandomInt(min, max int32) int32 {
	return min + rand.Int31n(max-min+1) //returns 0 => max-min
}

//RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

//RandomUserId generates a random string user_id
//func RandomUserId() string {
//	return RandomString(9)
//}

//RandomFirstName generates a random first name
func RandomFirstName() string {
	return RandomString(5)

}

//RandomLastName generates a random last name
func RandomLastName() string {
	return RandomString(4)
}

//RandomEmail generates a random email for a user
func RandomEmail() string {
	return RandomFirstName() + "@" + RandomLastName() + ".com"
}

//RandomPass generates a random password
func RandomPass() string {
	return RandomString(6)
}
