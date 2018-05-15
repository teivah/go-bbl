package main

func main() {
	var sum int // Zeroed (sum := 0)

	// For
	for i := 0; i < 5; i++ {
		sum += i
	}

	j := 0
	for ; j < 5; {
		j++
	}

	// While
	for sum < 30 {
		sum += sum
	}

	// Forever loop
	for {
		sum++
	}
}
