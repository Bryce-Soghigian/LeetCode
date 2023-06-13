/*
Given a positive integer that is less than 4000, return a string containing that value in Roman numeral representation. 
*/ 

package main


/*
Let us first consider the relationship between the input and the output of our presented function. 

What are we transforming? 


We start with a positive integer, then we want to represent that int with characters because more memory allocated means more money we spend on our app must be good. 


Lets outline the rules for Roman numerals: 
 1. I can be placed before V and X to make 4 units (IV) and 9 units (IX) respectively 
 2. X can be placed before L and C to make 40 (XL) and 90 (XC) respectively
 3. C can be placed before D and M to make 400 (CD) and 900 (CM) according to the same pattern 
 4. Symbols are placed from left to right in order of value, starting with the largest.


 So here, we need to look at our number, and get the biggest denomination represented first. 

 0-9 = I, II, III, IV, V, VI, VII, VIII, IX, X 
 10-99 = X, XX, XXX, XL, L, LX, LXX, LXXX, XC, C 
 100-999 = C, CC, CCC, CD, D, DC, DCC, DCCC, CM, M 
 1000-3999 = M, MM, MMM, MMMCMXCIX 

 So we can see that we need to get the largest denomination first, and then work our way down. 

 We need to look at the number of significant digits in our number. 
 Then based on that value we should get some value from our map. 



 Lets write up some test cases just to test our assumptions. 

 39 = XXX + IX = XXXIX. We can represent this as 30 + 9. so since we are below 40, we must use the next lowest denomination to represent our number, then after we represent 30 we still have 9 to represent, 1X represents 9 



 2421 = MM + CD + XX + I = MMCDXXI. We can represent this as 2000 + 400 + 20 + 1. So we need to get the largest denomination first, then work our way down. 


*/


func romanNumeralForm(num int) string {
	digits := []int{}
	mappings := map[int]string{ 
		1: "I",
		4: "IV",
		5: "V",
		9: "IX",
		10: "X",
		40: "XL",
		50: "L",
		90: "XC",
		100: "C",
		400: "CD",
		500: "D",
		900: "CM",
		1000: "M",
	}
	for val, roman := range mappings { 
		if num == 0 {
			break
		}
		quotient, remainder := divmod(num)
		digits = append(digits, ConcatNTimes(roman, quotient))

	}
}

// divmod is a function in python that returns the quotient and remainder of a division operation. Yoinking it into go
func divmod(num int) (int, int) {
	return num / 10, num % 10
}

func ConcatNTimes(s string, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += s
	}
	return result
}


romanNumeralForm(39) // XXXIX 
romanNumeralForm(2421) // MMCDXXI
romanNumeralForm(3999) // MMMCMXCIX
romanNumeralForm(1) // I 
romanNumeralForm(4) // IV 
