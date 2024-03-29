package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockStructure struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type mockStructureNew mockStructure

func Test_ConvertEntity(t *testing.T) {

	t.Parallel()

	t.Run("When an entity of a specific type is delivered and another of another type returns with the same structure ok", func(t *testing.T) {
		e1 := mockStructure{
			Id:   1,
			Name: "name",
		}

		e2 := new(mockStructureNew)

		err := ConvertEntity(e1, e2)

		require.NoError(t, err)
		assert.Equal(t, EntityToJson(e1), EntityToJson(e2))
	})
	t.Run("When an first params is invalid, return error with first params", func(t *testing.T) {
		err := ConvertEntity(make(chan int), "<body>")
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrWithFirstParam)
	})
	t.Run("When an entity is null, return a json empty", func(t *testing.T) {
		json := EntityToJson(nil)
		assert.Equal(t, json, "{}")
	})

	t.Run("When an entity of a specific type is delivered and another of another type returns with a distillate structure and returns nil", func(t *testing.T) {
		e1 := mockStructure{
			Id:   1,
			Name: "name",
		}

		e2 := new(string)

		err := ConvertEntity(e1, e2)

		assert.Error(t, err)

	})
}

func Test_EntityToJson(t *testing.T) {

	t.Parallel()

	t.Run("When an entity is delivered and json returns with values", func(t *testing.T) {

		jsonExpected := "{\"id\":1,\"name\":\"name\"}"
		e1 := mockStructure{
			Id:   1,
			Name: "name",
		}

		result := EntityToJson(e1)

		assert.Equal(t, jsonExpected, result)

	})

	t.Run("When an invalid entity is delivered and json returns empty", func(t *testing.T) {

		jsonExpected := "{}"

		result := EntityToJson(make(chan int))

		assert.Equal(t, jsonExpected, result)

	})
}

func Test_EntityToJsonEscape(t *testing.T) {

	t.Parallel()

	t.Run("When an entity is delivered and json returns with values", func(t *testing.T) {

		jsonExpected := "{\"id\":1,\"name\":\"name\"}"
		e1 := mockStructure{
			Id:   1,
			Name: "name",
		}

		result := EntityToJsonEscape(e1)

		assert.Equal(t, jsonExpected, result)

	})

	t.Run("When an invalid entity is delivered and json returns empty", func(t *testing.T) {

		jsonExpected := "{}"

		result := EntityToJsonEscape(make(chan int))

		assert.Equal(t, jsonExpected, result)

	})
}

func Test_JsonToEntity(t *testing.T) {
	t.Parallel()
	t.Run("When json is sent valid and return entity with data", func(t *testing.T) {

		json := "{\"id\":1,\"name\":\"name\"}"

		entity := new(mockStructure)
		JsonToEntity(json, entity)
		assert.Equal(t, 1, entity.Id)
		assert.Equal(t, "name", entity.Name)
	})

	t.Run("When json is sent valid and returns empty entity", func(t *testing.T) {
		json := "$%&/()/"
		entityExpected := new(mockStructure)
		entity := new(mockStructure)

		JsonToEntity(json, entity)
		assert.Equal(t, entityExpected, entity)
	})
}

func Test_ValidateIntValue(t *testing.T) {
	t.Parallel()
	t.Run("when String value, is a int value to convert", func(t *testing.T) {
		var intValue = "1"
		var err = ValidateIntValue(intValue)
		assert.NoError(t, err)
	})
	t.Run("when String value, not is a int value to convert", func(t *testing.T) {
		var intValue = "Hello"
		var err = ValidateIntValue(intValue)
		assert.Error(t, err)
	})
}
