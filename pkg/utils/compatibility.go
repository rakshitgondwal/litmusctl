package utils

var (
	// CompatibilityMatrix to store the compatible versions of litmusctl and ChaosCenter
	CompatibilityMatrix = map[string][]string{
		"0.6.0":  {"2.2.0", "2.3.0"},
		"0.7.0":  {"2.4.0", "2.5.0", "2.6.0", "2.7.0", "2.8.0"},
		"0.8.0":  {"2.4.0", "2.5.0", "2.6.0", "2.7.0", "2.8.0"},
		"0.9.0":  {"2.4.0", "2.5.0", "2.6.0", "2.7.0", "2.8.0"},
		"0.10.0": {"2.9.0", "2.10.0", "2.11.0", "2.12.0", "2.13.0"},
		"0.11.0": {"2.9.0", "2.10.0", "2.11.0", "2.12.0", "2.13.0"},
		"0.12.0": {"2.9.0", "2.10.0", "2.11.0", "2.12.0", "2.13.0"},
		"0.13.0": {"2.9.0", "2.10.0", "2.11.0", "2.12.0", "2.13.0"},
	}
)
