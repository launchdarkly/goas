package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReffableString(t *testing.T) {
	t.Run("marshals a string", func(t *testing.T) {
		result, err := json.Marshal(&ReffableString{Value: "Foobarbazington Esq."})

		require.NoError(t, err)
		require.Equal(t, ([]byte)("\"Foobarbazington Esq.\""), result)
	})

	t.Run("marshals an object", func(t *testing.T) {
		result, err := json.Marshal(&ReffableString{Value: "$ref:foo/bar/baz"})

		require.NoError(t, err)
		require.Equal(t, ([]byte)("{\"$ref\":\"foo/bar/baz\"}"), result)
	})

	t.Run("missing url", func(t *testing.T) {
		_, err := json.Marshal(&ReffableString{Value: "$ref:"})

		require.Error(t, err)
	})
}
