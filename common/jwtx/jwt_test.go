package jwtx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetToken(t *testing.T) {
	token, _ := GetToken("123", 1643255473, 3600, 123456)
	assert.Equal(t, token, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDMyNTkwNzMsImlhdCI6MTY0MzI1NTQ3MywidWlkIjoxMjM0NTZ9.M6kTaZbR0ZtgTlH_-RZh5vkTGeh4nXyBaxBp1bPfEZ0")
}
