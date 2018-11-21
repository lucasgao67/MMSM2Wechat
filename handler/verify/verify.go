package verify

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"sort"
	"strings"
)

func Verify(c *gin.Context) {
	var r struct {
		Signature string `json:"signature"`
		Timestamp string `json:"timestamp"`
		Nonce     string `json:"nonce"`
		Echostr   string `json:"echostr"`
	}
	if err := c.Bind(&r); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	logrus.Debugf("signature is %s, timestamp is %s, Nonce is %s, echostr is %s", r.Signature, r.Timestamp, r.Nonce, r.Echostr)
	isValid := checkSignature(r.Signature, viper.GetString("verify.token"), r.Timestamp, r.Nonce)
	if isValid {
		c.String(http.StatusOK, r.Echostr)
	} else {
		c.Status(http.StatusUnauthorized)
	}
}

func checkSignature(signature, token, timestamp, nonce string) bool {
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
