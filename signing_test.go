package imageboss

import (
	"encoding/hex"
	"testing"
)

func TestSignPath(t *testing.T) {
	path := "/mysecureimages/width/500/01.jpg"
	secret := "mysecret"
	token := signPath(secret, path)
	if len(token) != 64 {
		t.Errorf("signPath: token length = %d; want 64 (hex SHA-256)", len(token))
	}
	if _, err := hex.DecodeString(token); err != nil {
		t.Errorf("signPath: token is not hex: %v", err)
	}
	// Deterministic
	if signPath(secret, path) != token {
		t.Error("signPath: same inputs should yield same token")
	}
	// Different path → different token
	if signPath(secret, "/other/path.jpg") == token {
		t.Error("signPath: different path should yield different token")
	}
	// Different secret → different token
	if signPath("othersecret", path) == token {
		t.Error("signPath: different secret should yield different token")
	}
}
