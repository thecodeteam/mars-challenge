package sensorsuite

// Default values
const (
	InitSolarFlare = false

	InitTemp      = -53.5
	MinTemp       = -142.00
	MaxTemp       = 35.00
	VariationTemp = 5.0

	InitRadiation      = 500
	MinRadiation       = 0
	MaxRadiation       = 1000
	VariationRadiation = 20

	MinTrendSec = 5
	MaxTrendSec = 20

	DefaultFlareWSAddr = ":9000"
	DefaultTempWSAddr  = ":9001"
	DefaultRadWSAddr   = ":9002"
	DefaultPubWSAddr   = ":9003"
	DefaultAggWSAddr   = ":9004"
)
