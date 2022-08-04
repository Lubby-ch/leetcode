package main

func intToRoman(num int) string {
	var m = map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}
	var example = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	var ans string
	var index int
	for num != 0 {
		for index=0;index<len(example);index++ {
			if num-example[index] >= 0 {
				num -= example[index]
				ans += m[example[index]]
				break
			}
		}
	}
	return ans
}
