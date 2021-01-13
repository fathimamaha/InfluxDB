package random

import (
	"math/rand"
	"fmt"
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
)

type Random struct {
	x         float64
	Amplitude float64
}

var RandomConfig = `
  ## Function used is float64()
  float value between 0 and 1
`

func (s *Random) SampleConfig() string {
	return RandomConfig
}

func (s *Random) Description() string {
	return "Generates a random value"
}

func (s *Random) Gather(acc telegraf.Accumulator) error {
	randint := rand.Float64()

	fields := make(map[string]interface{})
	fields["random"] = randint

	tags := make(map[string]string)

	acc.AddFields("random", fields, tags)

	fmt.Println("Random integer:",randint)
	return nil
}

func init() {
	inputs.Add("Random", func() telegraf.Input { return &Random{x: 0.0} })
}
