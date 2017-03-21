package main

import (
	"testing"
	assert "github.com/pilu/miniassert"
)

func TestPluralizeBottles(t *testing.T) {
	assert.Equal(t, "Нет бутылок", pluralizeBottles(0))
	assert.Equal(t, "15 бутылок", pluralizeBottles(15))
	assert.Equal(t, "100 бутылок", pluralizeBottles(100))
	assert.Equal(t, "1 бутылка", pluralizeBottles(1))
	assert.Equal(t, "2 бутылки", pluralizeBottles(2))
	assert.Equal(t, "23 бутылки", pluralizeBottles(23))
	assert.Equal(t, "104 бутылки", pluralizeBottles(104))
	assert.Equal(t, "112 бутылок", pluralizeBottles(112))
	assert.Equal(t, "110 бутылок", pluralizeBottles(110))
}
