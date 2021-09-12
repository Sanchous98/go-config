package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var bag = DotNotationBag{bag: map[string]interface{}{
	"another": map[string]interface{}{
		"test": "value",
	},
}}

type TestStruct struct {
	TestField string `config:"another.test"`
}

func TestDotNotationBag_Has(t *testing.T) {
	assert.True(t, bag.Has("another"))
	assert.True(t, bag.Has("another.test"))
}

func TestDotNotationBag_Set(t *testing.T) {
	bag.Set("another.test", "value2")
	value, err := bag.Get("another.test")
	assert.Nil(t, err)
	assert.Equal(t, "value2", value)
}

func TestConfig_Configure(t *testing.T) {
	testStruct := &TestStruct{}
	err := Unmarshall(&Config{bag}, testStruct)
	assert.Nil(t, err)
	assert.NotNil(t, testStruct.TestField)
}
