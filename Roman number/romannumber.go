package main

func romanNumber(number int) string {
	decimalValue := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanNumeral := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	romanized := ""

	for index := 0; index < len(decimalValue); index++ {
		for decimalValue[index] <= number {
			romanized += romanNumeral[index]
			number -= decimalValue[index]
		}
	}

	return romanized
}

func main() {

}
