package mid_test

import (
	"github.com/MealeyAU/utils-go/pkg/mid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIDString(t *testing.T) {
	id := mid.ID{
		Timestamp:   1257894000,
		Entity:      "test",
		ServerIdent: "a0fd197d69f3",
		Entropy:     "0a50df0ab835",
	}
	expected := "1257894000-test-a0fd197d69f3-0a50df0ab835"
	assert.Equal(t, expected, id.String())
}

func TestParse_Valid(t *testing.T) {
	expected := mid.ID{
		Timestamp:   1257894000,
		ServerIdent: "a0fd197d69f3",
		Entity:      "test",
		Entropy:     "0a50df0ab835",
	}

	out, err := mid.Parse("1257894000-test-a0fd197d69f3-0a50df0ab835")
	assert.Nil(t, err)
	assert.EqualValues(t, expected, out)
}

func TestParse_Error_Malformed(t *testing.T) {
	_, err := mid.Parse("")
	assert.NotNil(t, err)
	assert.Equal(t, "invalid number of parts, have 1 need 4", err.Error())
}

func TestParse_Error_NotIntTimestamp(t *testing.T) {
	_, err := mid.Parse("cooked-test-a0fd197d69f3-0a50df0ab835")
	assert.NotNil(t, err)
	assert.Equal(t,
		"failed to parse timestamp component: strconv.ParseInt: parsing \"cooked\": invalid syntax",
		err.Error())
}