package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8{
	ret := make([][]uint8,dy)
	for i := 0 ; i < dy ;i++{
		ret[i] = make([]uint8 ,dx)
		for j := 0 ; j < dx ; j++{
			ret[i][j] = uint8(i*j)
		}
	}
	return ret
}

func Pic1(dx, dy int) [][]uint8{
	ret := make([][]uint8,dy)
	for i,_ := range ret{
		ret[i] = make([]uint8,dx)
		for j,_:= range ret[i]{
			ret[i][j] = uint8(i*j)
		}
	}
	return ret
}

func main() {
	pic.Show(Pic)
	pic.Show(Pic1)
}
