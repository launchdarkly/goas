package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_genSchemaObjectID(t *testing.T) {
	t.Run("check none", func(t *testing.T) {
		emptyMap := make(map[string]*string)
		result := genSchemaObjectID("", "sample", emptyMap)

		require.Equal(t, "sample", string(result))
	})
	t.Run("check single", func(t *testing.T) {
		emptyMap := make(map[string]*string)
		result := genSchemaObjectID("sample", "sample", emptyMap)

		require.Equal(t, "sample.sample", string(result))
	})
	t.Run("check multiple", func(t *testing.T) {
		emptyMap := make(map[string]*string)
		result := genSchemaObjectID("test.sample", "sample", emptyMap)

		require.Equal(t, "test.sample.sample", string(result))
	})
}

func Test_getAliasedTypeName(t *testing.T) {
	t.Run("check no change", func(t *testing.T) {
		emptyMap := make(map[string]*string)
		result := getAliasedTypeName("mypackage.test", emptyMap)

		require.Equal(t, "mypackage.test", string(result))
	})

	t.Run("check empty rename", func(t *testing.T) {
		emptyMap := make(map[string]*string)
		emptyString := ""
		emptyMap["mypackage"] = &emptyString
		result := getAliasedTypeName("mypackage.test", emptyMap)

		require.Equal(t, "test", string(result))
	})

	t.Run("check alias rename", func(t *testing.T) {
		emptyMap := make(map[string]*string)
		emptyString := "newpackage"
		emptyMap["mypackage"] = &emptyString
		result := getAliasedTypeName("mypackage.test", emptyMap)

		require.Equal(t, "newpackage.test", string(result))
	})
}

func Test_getAliasedPackageName(t *testing.T) {
	pkgName := "example/go/mypackage"
	t.Run("check no change", func(t *testing.T) {
		emptyMap := make(map[string]*string)
		result := getAliasedPackageName(pkgName, emptyMap)

		require.Equal(t, "example/go/mypackage", string(result))
	})

	t.Run("check empty rename", func(t *testing.T) {
		emptyMap := make(map[string]*string)
		emptyString := ""
		emptyMap["mypackage"] = &emptyString
		result := getAliasedPackageName(pkgName, emptyMap)

		require.Equal(t, "", string(result))
	})

	t.Run("check rename alias package", func(t *testing.T) {
		emptyMap := make(map[string]*string)
		emptyString := "newpackage"
		emptyMap["mypackage"] = &emptyString
		result := getAliasedPackageName(pkgName, emptyMap)

		require.Equal(t, "newpackage", string(result))
	})
}
