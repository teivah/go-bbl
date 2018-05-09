package main

func main() {
	var sum int // Zeroed (sum := 0)

	// For
	for i := 0; i < 5; i++ {
		sum += i
	}

	// While (complex form)
	j := 0
	for ; j<5; {
		j++
	}

	// While (easy form)
	for sum < 30 {
		sum += sum
	}

	// Forever loop
	for {
		sum++
	}
}
