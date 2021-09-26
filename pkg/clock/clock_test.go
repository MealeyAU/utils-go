package clock_test

import (
	"github.com/MealeyAU/utils-go/pkg/clock"
	"reflect"
	"testing"
	"time"
)

func TestRealtime_Now(t *testing.T) {
	rc := clock.Realtime{}
	generated := rc.Now()
	empty := time.Time{}

	if reflect.DeepEqual(generated, empty) {
		t.Errorf("generated and empty times are the same")
	}
}

func TestStatic(t *testing.T) {
	sc := clock.Static{}

	empty := time.Time{}

	t1 := sc.Now()
	if !reflect.DeepEqual(t1, empty) {
		t.Errorf("t1 and empty times are not the same")
	}

	sc.Set(time.Date(2020, 10, 10, 12, 30, 00, 0, &time.Location{}))
	t2 := sc.Now()
	if reflect.DeepEqual(t2, empty) {
		t.Errorf("t2 and empty values are the same, set didn't work")
	}
}