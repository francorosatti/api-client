package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_AccountAttributes_WithAccountClassification(t *testing.T) {
	// Arrange
	aa := &AccountAttributes{}
	expectedAccountClassification := "account_classification"

	// Act
	aa.WithAccountClassification(expectedAccountClassification)

	// Assert
	assert.Equal(t, *aa.AccountClassification, expectedAccountClassification)
}

// TODO: Add same test for all fields
