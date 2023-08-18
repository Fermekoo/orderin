package utils

import (
	crypto_rand "crypto/rand"
	"math/big"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

/** RandomInt generate a random integer beetwen min and max
 * @param min int
 * @param max int
 * @return int64
 */
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

/** RandomString generate a random string of length n
 * @param n int
 * @return string
 */
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

/** RandomOwner generate a random owner name
 * @return string
 */
func RandomOwner() string {
	return RandomString(int(RandomInt(1, 50)))
}

/** RandomMonye generate a random amount of money
 * @return int64
 */
func RandomMoney() int64 {
	return RandomInt(10, 10000)
}

func RandomEmail() string {
	name := RandomOwner()
	domain_list := []string{"@gmail.com", "@yahoo.com"}
	n := len(domain_list)
	domain := domain_list[rand.Intn(n)]

	return name + domain
}

func RandomBillKey() string {
	max := big.NewInt(999999999999)
	randomNumber, _ := crypto_rand.Int(crypto_rand.Reader, max)

	return randomNumber.String()
}
