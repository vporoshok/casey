package casey_test

import (
	"testing"

	"github.com/vporoshok/casey"

	"github.com/stretchr/testify/assert"
)

func TestCamel(t *testing.T) {
	cases := [...]struct {
		name string
		src  string
		res  casey.Tokens
	}{
		{
			name: "camelCase",
			src:  "someCamelCase",
			res:  casey.Tokens{"some", "Camel", "Case"},
		},
		{
			name: "CapitalCase",
			src:  "SomeCapitalCase",
			res:  casey.Tokens{"Some", "Capital", "Case"},
		},
		{
			name: "abbr",
			src:  "JSONFileAndSOMEMore",
			res:  casey.Tokens{"JSON", "File", "And", "SOME", "More"},
		},
		{
			name: "numbers",
			src:  "JSON42File42An42dSO42MEMore",
			res:  casey.Tokens{"JSON42", "File42", "An42d", "SO42ME", "More"},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			assert.EqualValues(t, c.res, casey.Camel(c.src))
		})
	}
}

func TestTokensCapital(t *testing.T) {
	cases := [...]struct {
		name string
		src  casey.Tokens
		res  string
	}{
		{
			name: "camelCase",
			src:  casey.Tokens{"some", "Camel", "Case"},
			res:  "SomeCamelCase",
		},
		{
			name: "CapitalCase",
			src:  casey.Tokens{"Some", "Capital", "Case"},
			res:  "SomeCapitalCase",
		},
		{
			name: "abbr",
			src:  casey.Tokens{"JSON", "File", "And", "SOME", "More"},
			res:  "JSONFileAndSOMEMore",
		},
		{
			name: "numbers",
			src:  casey.Tokens{"JSON42", "File42", "An42d", "SO42ME", "More"},
			res:  "JSON42File42An42dSO42MEMore",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			assert.EqualValues(t, c.res, c.src.Capital())
		})
	}
}

func TestTokensCamel(t *testing.T) {
	cases := [...]struct {
		name string
		src  casey.Tokens
		res  string
	}{
		{
			name: "camelCase",
			src:  casey.Tokens{"some", "Camel", "Case"},
			res:  "someCamelCase",
		},
		{
			name: "CapitalCase",
			src:  casey.Tokens{"Some", "Capital", "Case"},
			res:  "someCapitalCase",
		},
		{
			name: "abbr",
			src:  casey.Tokens{"JSON", "File", "And", "SOME", "More"},
			res:  "jsonFileAndSOMEMore",
		},
		{
			name: "numbers",
			src:  casey.Tokens{"JSON42", "File42", "An42d", "SO42ME", "More"},
			res:  "json42File42An42dSO42MEMore",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			assert.EqualValues(t, c.res, c.src.Camel())
		})
	}
}

func TestTokensSNAKE(t *testing.T) {
	cases := [...]struct {
		name string
		src  casey.Tokens
		res  string
	}{
		{
			name: "camelCase",
			src:  casey.Tokens{"some", "Camel", "Case"},
			res:  "SOME_CAMEL_CASE",
		},
		{
			name: "CapitalCase",
			src:  casey.Tokens{"Some", "Capital", "Case"},
			res:  "SOME_CAPITAL_CASE",
		},
		{
			name: "abbr",
			src:  casey.Tokens{"JSON", "File", "And", "SOME", "More"},
			res:  "JSON_FILE_AND_SOME_MORE",
		},
		{
			name: "numbers",
			src:  casey.Tokens{"JSON42", "File42", "An42d", "SO42ME", "More"},
			res:  "JSON42_FILE42_AN42D_SO42ME_MORE",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			assert.EqualValues(t, c.res, c.src.SNAKE())
		})
	}
}

func TestTokensSnake(t *testing.T) {
	cases := [...]struct {
		name string
		src  casey.Tokens
		res  string
	}{
		{
			name: "camelCase",
			src:  casey.Tokens{"some", "Camel", "Case"},
			res:  "some_camel_case",
		},
		{
			name: "CapitalCase",
			src:  casey.Tokens{"Some", "Capital", "Case"},
			res:  "some_capital_case",
		},
		{
			name: "abbr",
			src:  casey.Tokens{"JSON", "File", "And", "SOME", "More"},
			res:  "json_file_and_some_more",
		},
		{
			name: "numbers",
			src:  casey.Tokens{"JSON42", "File42", "An42d", "SO42ME", "More"},
			res:  "json42_file42_an42d_so42me_more",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			assert.EqualValues(t, c.res, c.src.Snake())
		})
	}
}

func TestTokensKebab(t *testing.T) {
	cases := [...]struct {
		name string
		src  casey.Tokens
		res  string
	}{
		{
			name: "camelCase",
			src:  casey.Tokens{"some", "Camel", "Case"},
			res:  "some-Camel-Case",
		},
		{
			name: "CapitalCase",
			src:  casey.Tokens{"Some", "Capital", "Case"},
			res:  "Some-Capital-Case",
		},
		{
			name: "abbr",
			src:  casey.Tokens{"JSON", "File", "And", "SOME", "More"},
			res:  "JSON-File-And-SOME-More",
		},
		{
			name: "numbers",
			src:  casey.Tokens{"JSON42", "File42", "An42d", "SO42ME", "More"},
			res:  "JSON42-File42-An42d-SO42ME-More",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			assert.EqualValues(t, c.res, c.src.Kebab())
		})
	}
}
