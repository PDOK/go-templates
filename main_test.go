package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)
func Test_addKvpArgumentsToMap(t *testing.T) {
	testMap := createTestMap()
    arguments := []string{"a.a2=A2new", "c.c1=C1"}

    addKvpArgumentsToMap(arguments, testMap)

	aMap := testMap["a"].(map[string]interface{})
	require.Equal(t, aMap["a1"], "A1")
	require.Equal(t, aMap["a2"], "A2new")
	cMap := testMap["c"].(map[string]interface{})
	require.Equal(t, cMap["c1"], "C1")
}

func createTestMap () map[string]interface{} {
	testMap := make(map[string]interface{})
	aMap := make(map[string]interface{})
	aMap["a1"]="A1"
	aMap["a2"]="A2"
	testMap["a"]= aMap
	testMap["b"]="B"
	testMap["x"]="C"
	return testMap
}
