package mid_test

import (
	"github.com/MealeyAU/utils-go/pkg/clock"
	"github.com/MealeyAU/utils-go/pkg/mid"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDefaultGenerator_Generate(t *testing.T) {
	clk := &clock.Static{}
	clk.Set(time.Date(1995, 8, 25, 8, 20, 35, 100, &time.Location{}))

	dg := mid.DefaultGenerator{
		Clock: clk,
		ServerIdent: uuid.MustParse("090817e5-2630-49fb-b640-0a50df0ab835"),
	}

	id, err := dg.Generate("example")
	assert.Nil(t, err)
	assert.Equal(t, int64(809338835), id.Timestamp)
	assert.Equal(t, "example", id.Entity)
	assert.Equal(t, "0a50df0ab835", id.ServerIdent)
	assert.NotEmpty(t, id.Entropy)
}

func TestStaticGenerator_Generate_Success(t *testing.T) {
	clk := &clock.Static{}
	clk.Set(time.Date(1995, 8, 25, 8, 20, 35, 100, &time.Location{}))

	dg := mid.DefaultGenerator{
		Clock: clk,
		ServerIdent: uuid.MustParse("090817e5-2630-49fb-b640-0a50df0ab835"),
	}
	sg := mid.StaticGenerator{
		DefaultGenerator: dg,
		IDQueue: []uuid.UUID{uuid.MustParse("01e4d5a8-31d6-49ce-ae7d-43c13a15d91c")},
	}

	id, err := sg.Generate("example")
	assert.Nil(t, err)
	assert.Equal(t, int64(809338835), id.Timestamp)
	assert.Equal(t, "example", id.Entity)
	assert.Equal(t, "0a50df0ab835", id.ServerIdent)
	assert.Equal(t, "43c13a15d91c", id.Entropy)
}

func TestStaticGenerator_Generate_Error(t *testing.T) {
	clk := &clock.Static{}
	clk.Set(time.Date(1995, 8, 25, 8, 20, 35, 100, &time.Location{}))

	dg := mid.DefaultGenerator{
		Clock: clk,
		ServerIdent: uuid.MustParse("090817e5-2630-49fb-b640-0a50df0ab835"),
	}
	sg := mid.StaticGenerator{
		DefaultGenerator: dg,
		IDQueue: []uuid.UUID{},
	}

	_, err := sg.Generate("example")
	assert.NotNil(t, err)
	assert.Equal(t,
		"failed to generate entropy bits: no uuids remaining to be read",
		err.Error())
}