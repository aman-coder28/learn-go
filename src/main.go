package main

func main() {
	// scanner := bufio.NewScanner(os.Stdin)

	// RunCalculator(scanner)
	//
	if file, err := LoadFile("../.env"); err != nil {
		println(err.Error())
	} else {
		println(file)
	}

}
