package exercises

type TemperatureUnit string

const (
	Kelvin     TemperatureUnit = "kelvin"
	Celsius    TemperatureUnit = "celsius"
	Fahrenheit TemperatureUnit = "fahrenheit"
)

func TemperatureUnitConversion(temperature float32, from TemperatureUnit, to TemperatureUnit) float32 {

	switch from {
	case "celsius":
		switch to {
		case "fahrenheit":
			return (temperature * 1.8) + 32
		case "kelvin":
			return temperature + 273.15
		}
	case "fahrenheit":
		switch to {
		case "celsius":
			return (temperature - 32) / 1.8
		case "kelvin":
			return ((temperature - 32) * 5 / 9) + 273.15
		}
	case "kelvin":
		switch to {
		case "celsius":
			return temperature - 273.15
		case "fahrenheit":
			return ((temperature - 273.15) * 9 / 5) + 32
		}
	}

	return 0
}
