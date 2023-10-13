package main

func Checker(cardNumber int) bool {
	var cardNumberArray []int //will store card num as slice

	//-----int to slice converting-----//
	for cardNumber > 0 {
		number := cardNumber % 10
		// cardNumberArray = append(cardNumberArray, number)
		cardNumberArray = append([]int{number}, cardNumberArray...)
		cardNumber /= 10
	}

	//----------LUHN Algorithm Starts here----------//
	oddSum := 0
	evneSum := 0

	for i := len(cardNumberArray) - 1; i >= 0; i-- {
		if i%2 != 0 {
			oddSum += cardNumberArray[i]
		} else {
			n := cardNumberArray[i] * 2
			if n >= 10 {
				evneSum += ((n % 10) + (n / 10))
			} else {
				evneSum += n
			}
		}
	}

	totalSum := oddSum + evneSum
	if totalSum%10 == 0 {
		return true
	} else {
		return false
	}
}

/*
Their is an logical error, I should have limited the card number within 16 digits(cz real life cards have only 16 digits)
*/