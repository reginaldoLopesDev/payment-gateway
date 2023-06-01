package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreditCardNumber(t *testing.T) {
	_, err := NewCreditCard("0123456789123456", "José da Silva", 12, 2024, 123)
	assert.Equal(t, "invalid credit card number", err.Error())

	_, err = NewCreditCard("5159227242857123", "José da Silva", 12, 2024, 123)
	assert.Nil(t, err)
}

func TestCreditCardExpirationMonth(t *testing.T) {
	_, err := NewCreditCard("5159227242857123", "José da Silva", 13, 2024, 123)
	assert.Equal(t, "invalid expiration month", err.Error())

	_, err = NewCreditCard("5159227242857123", "José da Silva", 0, 2024, 123)
	assert.Equal(t, "invalid expiration month", err.Error())

	_, err = NewCreditCard("5159227242857123", "José da Silva", 4, 2024, 123)
	assert.Nil(t, err)
}
func TestCreditCardExpirationYear(t *testing.T) {
	_, err := NewCreditCard("5159227242857123", "José da Silva", 4, 2022, 123)
	assert.Equal(t, "invalid expiration year", err.Error())

	_, err = NewCreditCard("5159227242857123", "José da Silva", 4, 2023, 123)
	assert.Nil(t, err)
}

func TestCreditCardCVV(t *testing.T) {
	_, err := NewCreditCard("5159227242857123", "José da Silva", 4, 2023, 12345)
	assert.Equal(t, "invalid cvv", err.Error())

	_, err = NewCreditCard("5159227242857123", "José da Silva", 4, 2023, 12)
	assert.Equal(t, "invalid cvv", err.Error())

	_, err = NewCreditCard("5159227242857123", "José da Silva", 4, 2023, 123)
	assert.Nil(t, err)

}
