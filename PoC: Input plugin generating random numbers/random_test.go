package random

import (
	"math"
	"testing"

	"github.com/influxdata/telegraf/testutil"
)

type Random struct {
	x         float64
	Amplitude float64
}

func TestRandom(t *testing.T) {
	s := &Random{
		Amplitude: 10.0,
	}

	for i := 0.0; i < 10.0; i++ {

		var acc testutil.Accumulator

		randint := rand.Float64()

		s.Gather(&acc)

		fields := make(map[string]interface{})
		fields["random"] = randint

		acc.AssertContainsFields(t, "random", fields)
	}
}
