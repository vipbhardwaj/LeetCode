// Function to calculate the greatest common divisor
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// Function to parse a fraction and return numerator and denominator
func parseFraction(fraction string) (int, int) {
	parts := strings.Split(fraction, "/")
	numerator, _ := strconv.Atoi(parts[0])
	denominator, _ := strconv.Atoi(parts[1])
	return numerator, denominator
}

// Function to add or subtract fractions from the expression
func fractionAddition(expression string) string {
	// Initialize numerator and denominator
	numerator, denominator := 0, 1
	// Split the expression into fractions based on the signs
	i := 0
	for i < len(expression) {
		start := i
		if expression[i] == '+' || expression[i] == '-' {
			i++
		}
		for i < len(expression) && (expression[i] >= '0' && expression[i] <= '9' || expression[i] == '/') {
			i++
		}
		// Parse the current fraction
		num, denom := parseFraction(expression[start:i])
		// Calculate the new numerator and denominator
		numerator = numerator*denom + num*denominator
		denominator *= denom
		// Simplify the result using gcd
		g := gcd(abs(numerator), denominator)
		numerator /= g
		denominator /= g
	}
	return fmt.Sprintf("%d/%d", numerator, denominator)
}

// Helper function to get absolute value
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}