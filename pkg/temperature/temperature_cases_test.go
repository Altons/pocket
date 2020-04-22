package temperature_test

var temperatureTestCases = []struct {
	value       float64
	unitIn      string
	unitOut     string
	want        float64
	expectError bool
}{
	{ // empty strands
		0.0,
		"",
		"",
		0.0,
		true,
	},
	{ // missing unitIn
		0.0,
		"",
		"k",
		0.0,
		true,
	},
	{ // missing unitOut
		0.0,
		"c",
		"",
		0.0,
		true,
	},
	{ // same units and diff cases
		5.5,
		"c",
		"C",
		5.5,
		false,
	},
	{ // celsius to Fahrenheit
		0,
		"c",
		"f",
		32,
		false,
	},
	{ // kelvin to Fahrenheit
		200.0,
		"k",
		"f",
		-99.67,
		false,
	},
	{ // Fahrenheit to celsius
		10,
		"f",
		"c",
		-12.222,
		false,
	},
	{ // Celsius to Kelvin
		10,
		"c",
		"k",
		283.15,
		false,
	},
	{ // Kelvin to Celsius to
		100,
		"k",
		"c",
		-173.15,
		false,
	},
	{ // Fahrenheit to Kelvin
		97,
		"f",
		"k",
		309.26,
		false,
	},
}
