func intToRoman(num int) string {
	oneToNine := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	tenToNinety := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	hundredToNineHundred := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	thousandToThreeThousand := []string{"", "M", "MM", "MMM"}

	return thousandToThreeThousand[num/1000] + hundredToNineHundred[(num%1000)/100] + tenToNinety[(num%100)/10] + oneToNine[num%10]

}
