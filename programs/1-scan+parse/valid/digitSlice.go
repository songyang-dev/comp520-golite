/*Takes a number as input and creates a slice
with all the digits of that number*/
package main

func intToSlice(int num){
	numSlice = []rune;
	for num>0 {
		append(numSlice, rune(num%10));
		num /= 10;
	}
	return numSlice
}