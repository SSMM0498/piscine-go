#   Quest 02
    Go is a language that I will discover in this piscine. I have heard on it in the pass. Some classmates told me that it is a very performant language but actually, I will experiment that by myself.

##  Print alphabet
>   [Here it is](../printalphabet/main.go)

First I learn that go need a package declaration at top of each files and at least one main func in a go project for a starting point.
We initialize a go module with the command :

```sh
go mod init
```

We can also import both external and internal package by using the `import` statements. Notice that for external package we need to download it by running:

```sh
go get <package-name>
```

Particularly for this task, I learn that go can loop through characters and make comparaison between them.

⚠️FOR THIS QUEST, WE WILL USE A PACKAGE DEVELOP BY ZONE01 WHICH CONTAIN A FUNCTION NAMED PRINTRUNE. WE ARE NOT ALLOW TO USE THE BASIC FMT PACKAGE OF GOLANG TO PRINT TEXT IN THIS EXERCICE.⚠️

⚠️I have lost a lot of time because the code you will push must be correctly formated before being executed.⚠️

We resolve this task by using a for loop with each we print each alphabet letter with the `PrintRune`.

##  Print reverse alphabet
>   [Here it is](../printreversealphabet/main.go)

No comment, this task don't really have challenge.

##  Print digits
>   [Here it is](../printdigits/main.go)

We use the rune function which we pass as parameters ASCII code of each number between 0 and 9. (Their ASCII code is between the range 48 and 58)

##  Is negative
>   [Here it is](../isnegative.go)

This is the first task where we need to handle several package. One for the function we are supposed to write and another for testing it. Indeed we will import the first package into the second one for making the test.

##  Print combination
##  Print combination 2
##  Print nbr

You need making search about [the maximum and minimum value of the int types in go](https://gosamples.dev/int-min-max/#:~:text=%F0%9F%8D%B0%20The%20maximum%20and%20minimum%20value%20of%20the%20int%20types%20in%20Go&text=For%20example%2C%20to%20get%20the,MaxInt64%20.)
##  Print combination n