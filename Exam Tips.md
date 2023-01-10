#   How to hack the z01.PrintRune constraint

##  Rewrite the function PrintStr on your coding space
```go
func PrintStr(str string) {
	for _, char := range str {
		z01.PrintRune(char)
	}
}
PrintStr("Cool! I can print a string quickly")
```

##  Add line break instruction if needed
```go
func PrintStr(str string) {
	for _, char := range str {
		z01.PrintRune(char)
	}
    z01.PrintRune('\n')
}
PrintStr("Cool! a can print a string quickly")
```

##  Concatenate the useful variables directly with your string before printing the whole thing
```go
name := "Serigne Saliou Mbaye"
str := "My name is "+name
PrintStr(str)
```

<div class="page"></div>

##  Write an Itoa function to convert integers to string if you need to print a number
```go
func Itoa(nb int) string {
	runes := []rune(nil)
	for nb != 0 {
		char := (nb % 10) + 48
		runes = append(runes, rune(char))
		nb /= 10
	}
	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(runes)-i-1] = runes[len(runes)-i-1], runes[i]
	}
	return string(runes)
}

nb1 := 12
nb2 := 13
addResult := nb1 + nb2

str := "The result of adding "+Itoa(nb1)+" to "+Itoa(nb2)+" is :"+Itoa(addResult)
PrintStr(str)
```

##   BONUS: You can get an array of runes directly from a string with the following method. This is useful when must apply some operation on the string
```go
// when you must reverse the string for example
str := "aBcDeFgHiJklMnOpQrStUvWxYz"
runes := []rune(str)
for i := 0; i < len(runes)/2; i++ {
	runes[i], runes[len(runes)-i-1] = runes[len(runes)-i-1], runes[i]
}
str = string(runes)
PrintStr(str)
```

<div class="page"></div>

#   Array for dummy

##  Get the first and last element of an array
```go
number := []int{1,2,3,4,5}
first := number[0]
last := number[len(number)-1]
```

## Get a part of an array
```go
number := []int{10, 32, 6, 97, 45, 13}
low := 1
high := 4
// Keep in mind that the element at 'high' position is not to retrieve
part1 := number[low:high]
part2 := number[2:] // Retrieve all elements from the third element to the end
part3 := number[:2] // Retrieve all elements for the first to the third element (the third exclude)
```

## Delete an element at a specific position in an array
```go
number := []int{10,32,6,97,45}
x := 3 // If you want delete the 97
// Create a new array with all element before the position x : number[:x]
// Append all elements after the position x : number[x+1:] (Don't forgot the ...)
number = append(number[:x], number[x+1:]...)
```

#   ASCII the essentials to know

##	The ASCII code for `a` is : **97**

##	The ASCII code for `A` is : **65**

##	The ASCII code for `0` is : **48**

##	Get in mind the previous and for finding the ASCII code for the others add the value of his normal position (for example : `z` is the 26th letter of the alphabet to find their ASCII code you add 25 to th `a` ASCII code : **122**)

##  Is lower case condition
```go
char >= 'a' && char <= 'z'
```

##  Is not lower case condition
```go
char < 'a' || char > 'z'
```

##  Is upper case condition
```go
char >= 'A' && char <= 'Z'
```

##  Is not upper case condition
```go
char < 'A' || char > 'Z'
```

##  Is digit condition
```go
char >= '0' && char <= '9'
```

##  Is not digit condition
```go
char < '0' || char > '9'
```


#   Bits
    +   activeBits, BinaryAddition, reverseBit, swapbits
#   Array
    +   Sort
    +   Sum
#   Ascii
    +   Is Alpha
#   Loop
    +   Count
#   PrintRune
    +   Atoi, TrimAtoi, AtoiBase, BasicAtoi, BasicAtoi2
    +   Itoa
#   Function
    +   Function of function
#   Mathematiques
    +   base
    +   gcd
    +   lcm
    +   power
    +   median
    +   negative
    +   prime