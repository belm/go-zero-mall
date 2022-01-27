package cryptx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPasswordEncrypt(t *testing.T) {
	pwd := PasswordEncrypt("123", "123")
	assert.Equal(t, pwd, "cd2b18ac1e2b647af8e4bf2f3ca934afd080fa558fae09167839890e64ca6e15")
}