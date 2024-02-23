package page_assembler

func translatePlanetCase(planet string) string {
	switch planet {
	case "sun":
		return "Солнце"
	case "moon":
		return "Луна"
	case "mercury":
		return "Меркурий"
	case "venus":
		return "Венера"
	case "mars":
		return "Марс"
	case "jupiter":
		return "Юпитер"
	case "saturn":
		return "Сатурн"
	case "uranus":
		return "Уран"
	case "neptune":
		return "Нептун"
	case "pluto":
		return "Плутон"
	case "charon":
		return "Хирон"
	}
	return planet
}
