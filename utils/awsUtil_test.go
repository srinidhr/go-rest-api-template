package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAwsSesSession(t *testing.T) {
	sesClient, err := GetAwsSesSession("us-east-1")

	assert.Nil(t, err)
	assert.NotNil(t, sesClient)
}
