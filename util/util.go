package util

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"fmt"
	"errors"
	"encoding/json"
)

func Hash(text string) string {
	h := sha256.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}

func GetChainFromRemote(host string, target *FullChainResponse) error {
	response, error := http.Get(fmt.Sprintf("http://%s/chain", host))
	if error != nil {
		return error
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		return json.NewDecoder(response.Body).Decode(target)
	} else {
		return errors.New(fmt.Sprintf("Connection failed: %d", response.StatusCode))
	}
}