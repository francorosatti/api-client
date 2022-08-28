package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccountData_WithID(t *testing.T) {
	// Arrange
	ad := &AccountData{}
	expectedID := "id"

	// Act
	ad.WithID(expectedID)

	// Assert
	assert.Equal(t, ad.ID, expectedID)
}

// TODO: Add same test for all fields
