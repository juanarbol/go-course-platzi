package maps

// GetMap deveulve un mapita
func GetMap(name string) int {
	// maptest := make(map[string]int)
	maptest := map[string]int{
		"Juan":    10,
		"Luz":     100,
		"Natalia": 1000,
	}

	return maptest[name]
}
