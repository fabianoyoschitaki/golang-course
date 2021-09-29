package addresses

import "strings"

// TypeOfAddress verifies if address has a valid type and returns it
func TypeOfAddress(address string) string {
	// slice of valid address types
	validTypes := []string{"street", "avenue", "road", "highway"}

	// Avenue X Y = ["Avenue", "X", "Y"]
	firstAddressWord := strings.Split(strings.ToLower(address), " ")[0]

	addressHasAValidType := false
	// iterate over the valid types
	for _, validType := range validTypes {
		if firstAddressWord == validType {
			addressHasAValidType = true
		}
	}

	// if it has a valid type, return it
	if addressHasAValidType {
		return strings.Title(firstAddressWord)
	}
	return "Invalid Type"
}
