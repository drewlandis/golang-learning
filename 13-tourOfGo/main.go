package main

import "fmt"

/*
https://tour.golang.org/moretypes/18

Exercise: Slices
Implement Pic. It should return a slice of length dy, each element of which is a
  slice of dx 8-bit unsigned integers. When you run the program, it will display
  your picture, interpreting the integers as grayscale (well, bluescale) values.
The choice of image is up to you. Interesting functions include (x+y)/2, x*y,
  and x^y.
(You need to use a loop to allocate each []uint8 inside the [][]uint8.)
(Use uint8(intValue) to convert between types.)
*/

func Pic(dx, dy int) [][]uint8 {
	var pic [][]uint8

	for y := 0; y < dy; y++ {
		ySlice := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			//ySlice[x] = uint8(x + y)
			ySlice[x] = uint8(x + 2*y)
			//ySlice[x] = uint8(x ^ y)
		}
		pic = append(pic, ySlice)
	}

	return pic
}

func main() {
	//pic.Show(Pic)
	/* */
	pic := Pic(4, 4)
	for i := range pic {
		fmt.Println(pic[len(pic)-i-1])
	}
	/* */
}
