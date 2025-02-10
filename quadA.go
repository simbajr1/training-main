package main

import "github.com/01-edu/z01" //modules to print to the console

func QuadA(x, y int) { //function to print a rectangle
	for row := 1; row <= y; row++ { //loop to print the rectangle
		if row == 1 {
			z01.PrintRune('o')                       //print the first row
			for column := 2; column <= x; column++ { //column loop, prints column of the first row with a star
				z01.PrintRune('*')
			}
			z01.PrintRune('o') //print the last character of the first row
		} else if row == y { //print the last row with o
			z01.PrintRune('o')
			for column := 2; column <= x; column++ { //print the column of the last row with a star
				z01.PrintRune('*')
			}
			z01.PrintRune('o') //print the last character of the last row
		} else {
			z01.PrintRune('|') //print the column of the rectangle the height of the rectangle of the middle row
			for column := 2; column <= x; column++ {
				z01.PrintRune(' ') //print the space in the rectangle of the middle row
			}
			z01.PrintRune('|') //print the last character of the column in the rectangle of the middle row
		}
		z01.PrintRune('\n')
	}
}

//basically the code prints the rectangle with the given width and height on all quads, just a matter of changing the position of the characters

//function to print the rectangle, given the width and height
func main() {
	QuadA(5, 3)
}
