package ramukaka

import (
	"crypto/sha1"
	"encoding/base64"
)

func findIndex(value byte) int {
	whereAt := 0
	if value >= 97 && value <= 122 {
		whereAt = int(value) - 97
	} else if value >= 48 && value <= 57 {
		whereAt = int(value) - 48 + 26
	}
	return whereAt
}

func hash(policys string) string {
	hashman := sha1.New()
	policy := []byte(policys)
	hashman.Write(policy)
	hash := base64.URLEncoding.EncodeToString(hashman.Sum(nil))
	return hash
}
