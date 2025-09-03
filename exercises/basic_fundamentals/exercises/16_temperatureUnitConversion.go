package exercises

type TemperatureUnit string

const (
	Kelvin     TemperatureUnit = "kelvin"
	Celsius    TemperatureUnit = "celsius"
	Fahrenheit TemperatureUnit = "fahrenheit"
)

func TemperatureUnitConversion(temperature float32, from TemperatureUnit, to TemperatureUnit) float32 {

	switch from == "celsius" {
	case to == "fahrenheit":

	}
	return 1.1
}
