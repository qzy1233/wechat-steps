package utils

import (
	"fmt"
	"io"
	"math/rand"
	"net/url"
	"strings"
	"time"
)

var r *rand.Rand

func init() {
	source := rand.NewSource(time.Now().UnixNano())
	r = rand.New(source)
}
func FakeIP() string {
	// 随便找的国内IP段：223.64.0.0 - 223.117.255.255
	return fmt.Sprintf("223.%d.%d.%d", r.Intn(54)+64, r.Intn(254), r.Intn(254))
}

func GetFromData(data map[string]string) io.Reader {
	formData := url.Values{}
	for k, v := range data {
		formData.Add(k, v)
	}
	return strings.NewReader(formData.Encode())
}

func GetRandInt64(n int64) int64 {
	return r.Int63n(n)
}
