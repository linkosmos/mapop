package mapop

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var splitTests = []struct {
	input          map[string]interface{}
	expectedKeys   []string
	expectedValues []interface{}
}{
	{
		input: map[string]interface{}{
			"key1": 2,
			"key3": 2,
		},
		expectedKeys:   []string{"key1", "key3"},
		expectedValues: []interface{}{2, 2},
	},
	{
		input: map[string]interface{}{
			"key1": 1,
			"key2": 2,
			"key3": 292929,
			"key4": 4,
			"key5": nil,
		},
		expectedKeys:   []string{"key1", "key2", "key3", "key4", "key5"},
		expectedValues: []interface{}{1, 2, 292929, 4, nil},
	},
}

func TestSplit(t *testing.T) {
	for _, test := range splitTests {
		keysGot, valuesGot := Split(test.input)

		// Assert size
		assert.Equal(t, len(keysGot), len(valuesGot))

		// Assert keys maps to values
		assert.Len(t, keysGot, len(test.input))
		assert.Len(t, valuesGot, len(test.input))

		for _, keyGot := range keysGot {
			assert.Contains(t, test.expectedKeys, keyGot)
		}

		for _, valueGot := range valuesGot {
			assert.Contains(t, test.expectedValues, valueGot)
		}

		// Assert order mapped is correct
		for keyGotIndex, keyGot := range keysGot {
			inputValue, ok := test.input[keyGot]
			assert.True(t, ok)

			gotValue := valuesGot[keyGotIndex]
			assert.Equal(t, inputValue, gotValue)
		}
	}
}

var selectTests = []struct {
	input        map[string]interface{}
	selectedKeys []string
	expectedKeys []string
}{
	{
		input: map[string]interface{}{
			"key1": 2,
			"key3": 2,
		},
		selectedKeys: []string{"key1", "key3"},
		expectedKeys: []string{"key1", "key3"},
	},
	{
		input: map[string]interface{}{
			"key1": 1,
			"key2": 2,
			"key3": 292929,
			"key4": 4,
			"key5": nil,
		},
		selectedKeys: []string{"key1", "key5"},
		expectedKeys: []string{"key1", "key5"},
	},
	{
		input: map[string]interface{}{
			"key1": 2,
			"key3": 2,
		},
		selectedKeys: []string{},
		expectedKeys: []string{"key1", "key3"}, // If select empty return same input
	},
	{
		input: map[string]interface{}{
			"key1": 2,
			"key3": 2,
		},
		selectedKeys: []string{"noKey"},
		expectedKeys: []string{},
	},
	{
		input: map[string]interface{}{
			"key1": 2,
			"key3": 2,
		},
		selectedKeys: []string{"noKey", "no", "nokey2", "nokey3"},
		expectedKeys: []string{},
	},
	{
		input:        map[string]interface{}{},
		selectedKeys: []string{"noKey", "no", "nokey2", "nokey3"},
		expectedKeys: []string{},
	},
	{
		input:        map[string]interface{}{},
		selectedKeys: []string{},
		expectedKeys: []string{},
	},
}

func TestSelect(t *testing.T) {
	for _, test := range selectTests {
		got := Select(test.input, test.selectedKeys...)
		keysGot, _ := Split(got)

		assert.Equal(t, len(keysGot), len(test.expectedKeys))

		for _, keyGot := range keysGot {
			assert.Contains(t, test.expectedKeys, keyGot)
		}
	}
}

var rejectTests = []struct {
	input        map[string]interface{}
	rejectedKeys []string
	expectedKeys []string
}{
	{
		input: map[string]interface{}{
			"key1": 2,
			"key3": 2,
		},
		rejectedKeys: []string{"key3"},
		expectedKeys: []string{"key1"},
	},
	{
		input: map[string]interface{}{
			"key1": 1,
			"key2": 2,
			"key3": 292929,
			"key4": 4,
			"key5": nil,
		},
		rejectedKeys: []string{"key5"},
		expectedKeys: []string{"key1", "key2", "key3", "key4"},
	},
	{
		input: map[string]interface{}{
			"key1": 2,
			"key3": 2,
		},
		rejectedKeys: []string{},
		expectedKeys: []string{"key1", "key3"}, // If reject empty return same input
	},
	{
		input: map[string]interface{}{
			"k3": 2,
			"k1": 2,
		},
		rejectedKeys: []string{"noKey"},
		expectedKeys: []string{"k1", "k3"},
	},
	{
		input: map[string]interface{}{
			"key1": 2,
			"key3": 2,
		},
		rejectedKeys: []string{"noKey", "no", "nokey2", "nokey3"},
		expectedKeys: []string{"key1", "key3"},
	},
	{
		input:        map[string]interface{}{},
		rejectedKeys: []string{"noKey", "no", "nokey2", "nokey3"},
		expectedKeys: []string{},
	},
	{
		input:        map[string]interface{}{},
		rejectedKeys: []string{},
		expectedKeys: []string{},
	},
}

func TestReject(t *testing.T) {
	for _, test := range rejectTests {
		got := Reject(test.input, test.rejectedKeys...)
		keysGot, _ := Split(got)

		assert.Equal(t, len(keysGot), len(test.expectedKeys))

		for _, keyGot := range keysGot {
			assert.Contains(t, test.expectedKeys, keyGot)
		}
	}
}
