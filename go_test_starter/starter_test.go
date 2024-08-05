package starter_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	starter "github.com/williaminfante/go_test_starter"
)

func TestSayHello(t *testing.T) {
	greeting := starter.SayHello("William")
	assert.Equal(t, "Testando pela primeira vez", greeting)
}
