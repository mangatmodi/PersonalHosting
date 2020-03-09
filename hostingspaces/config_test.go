package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewConfig(t *testing.T) {
	c, err := NewConfig("conf.yml")
	assert.Nil(t, err)
	assert.Equal(t, uint16(8000), c.Port)
	assert.Equal(t, 2, len(c.Routes))
	assert.Equal(t, "resume.pdf", c.Routes[0].ToRoute)
}
