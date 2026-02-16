package imageboss

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// signPath returns the HMAC SHA-256 hex digest of path using secret, as required by
// ImageBoss signed URLs. See https://imageboss.me/docs/security.
func signPath(secret, path string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(path))
	return hex.EncodeToString(mac.Sum(nil))
}
