package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHome(t *testing.T) {
	_, err := Home()

	assert.Nil(t, err, "in Home, error should be nil")
}

func TestSaveData(t *testing.T) {
	_, err := SaveData()

	assert.Nil(t, err, "in SaveData, error should be nil")
}
