package verify

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"sort"
	"strings"
)

func Verify(c *gin.Context) {

}

func checkSignature(signature, token,timestamp, nonce string) bool {
	strList := []string{token, timestamp, nonce}
	sort.Strings(strList)
	str := strings.Join(strList, "")
	tmpStr := sha1s(str)
	if tmpStr == signature {
		return true
	} else {
		return false
	}
}

func sha1s(s string) string {
	r := sha1.Sum([]byte(s))
	return hex.EncodeToString(r[:])
}
