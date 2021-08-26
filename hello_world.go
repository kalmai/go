package main

import(
	f "fmt" // you can alias imports like this
	"time"
)

// i can comment like this

/* or
like 
this */

const exampleConst = 1 + 1 // this is an example of how to write a constant in go

/*
go datatypes: https://content.codecademy.com/courses/updated_images/Go%2Bnumeric%2Btypes_Updated_1-01.svg
int8 = -128..128
int16 -32768..32768
int32 -1bil osrs..1bil osrs
int64 -bigNum..bigNum
using 'u' in front of the above only allows positive numbers starting at 0 (just double the upper range for the corresponding type)
a type with u is called unsigned or starting with 0 and signed is - to +
float32
float64 are the types for floats
*/

func main(){
	var helloVar string = "helllllllooo from var" // this is how you assign variables, var {varName} {type}
	outsideOfMainVar := false // you can have go infer the datatype when declaring something, however ints will only default to 32 and 64 bits, when declaring a variable like this you do not use 'var', when using the := operator the variable will take on the datatype it has been assigned to
	var insideOfMainVar = true // this is another syntax you can use to infer the datatype 
	var emptyInt int8 // go recommends that you do not specify the side of an int unless you're tying to save memory or if you know the size of the variable beforehand
	var iKnowHowBigThisIntIs int8 = 1
	var iDontKnowHowBigThisIntIs int = 1234123111
	f.Println("hello world")
	f.Println(helloVar + " with concatenation tm")
	f.Println(helloVar, emptyInt, outsideOfMainVar, insideOfMainVar, iKnowHowBigThisIntIs, iDontKnowHowBigThisIntIs) // go lets you print variables without converting them to strings it seems
	f.Println(time.Now())
	f.Println(exampleConst)
	f.Println(3i) // complex number example that evaluates to (0 + 3i) ????



	// you can declare multiple variables with the following syntax:
	var part1, part2 string // when doing this they must both the same datatype
	part1 = "hello"
	part2 = "world!"
	// can also be done on one like like so: var part1, part2 string = "asdf", ";lkj"
	f.Println(part1,part2)
	// you can also do it without having matching datatypes with the := operator:
	str, falsey := "sup", false
	f.Println(str, falsey)
}
