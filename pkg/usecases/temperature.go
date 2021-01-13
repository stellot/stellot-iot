package usecases

import (
	"math"
	"math/rand"
	"strconv"
)

var tempBase = 10.0

var temps = []float64{4, 3, 2, 1, 2, 3, 6, 10, 12, 14, 15, 18,
	20, 22, 22, 20, 18, 19, 16, 14, 12, 10, 8, 6}

func RandomTemperature(value int) [32]byte {
	timeDay := value % len(temps)
	temp := temps[timeDay] + tempBase + rand.Float64()*temps[timeDay]
	multipled := temp * 10
	result := int(math.Round(multipled))
	var output [32]byte
	copy(output[:], strconv.Itoa(result))
	return output
}
