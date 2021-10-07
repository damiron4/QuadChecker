package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func QuadA(x, y int) string {
	var str string
	for i := 0; i <= y-1; i++ {
		if x <= 0 || y <= 0 {
			break
		}
		for j := 0; j <= x-1; j++ {
			if (i == 0 && j == 0) || (i == y-1 && j == 0) || (i == 0 && j == x-1) || (i == y-1 && j == x-1) {
				str += "o"
			} else if i == 0 || i == y-1 {
				str += "-"
			} else if j == 0 || j == x-1 {
				str += "|"
			} else {
				str += " "
			}
		}
		if i == y-1 {
			break
		}
		str += "\n"
	}
	return str
}

func QuadB(x, y int) string {
	var str string
	for i := 0; i <= y-1; i++ {
		if y <= 0 || x <= 0 {
			break
		}
		for j := 0; j <= x-1; j++ {
			if i == 0 && j == x-1 && x != 1 {
				str += "\\"
			} else if j == 0 && i == y-1 && y != 1 {
				str += "\\"
			} else if i == 0 && j == 0 || i == y-1 && j == x-1 {
				str += "/"
			} else if i == y-1 && j == 0 || i == 0 && j == x-1 {
				str += "\\"
			} else if i == 0 || j == 0 || j == x-1 || i == y-1 {
				str += "*"
			} else {
				str += " "
			}
		}
		if i == y-1 {
			break
		}
		str += "\n"
	}
	return str
}

func QuadC(x,y int) string {
	var str string
	for i := 0; i < y; i++ {
		if x <= 0 || y <= 0 {
			break
		}
		for j := 0; j < x; j++ {
			if j == 0 || j == x-1 {
				if i == 0 {
					str += "A"
					continue
				} else if i > 0 && i <= y-2 {
					str += "B"
					continue
				} else if i == y-1 {
					str += "C"
					continue
				}
			}
			if j > 0 && j < x-1 && i > 0 && i < y-1 {
				str += " "
				continue
			}
			str += "B"
		}
		if i == y-1 {
			break
		}
		str += "\n"
	}
	return str
}

func QuadD(x, y int) string {
	var str string
	for i := 0; i <= y-1; i++ {
		if x <= 0 || y <= 0 {
			break
		}
		for j := 0; j <= x-1; j++ {
			if (i == 0 && j == 0) || (i == y-1 && j == 0) {
				str += "A"
			} else if (i == 0 && j == x-1) || (i == y-1 && j == x-1) {
				str += "C"
			} else if i == 0 || i == y-1 || j == 0 || j == x-1 {
				str += "B"
			} else {
				str += " "
			}
		}
		if i == y-1 {
			break
		}
		str += "\n"
	}
	return str
}

func QuadE(x,y int) string {
	var str string
	for i := 0; i < y; i++ {
		if x <= 0 || y <= 0 {
			break
		}
		for j := 0; j < x; j++ {
			if j == 0 || j == x-1 {
				if i == 0 {
					if j == x-1 && x != 1 {
						str += "C"
						continue
					}
					str += "A"
					continue
				} else if i > 0 && i <= y-2 {
					str += "B"
					continue
				} else if i == y-1 {
					if j == x-1  && x != 1 {
						str += "A"
						continue
					}
					str += "C"
					continue
				}
			}
			if j > 0 && j < x-1 && i > 0 && i < y-1 {
				str += " "
				continue
			}
			str += "B"
		}
		if i == y-1 {
			break
		}
		str += "\n"
	}
	return str
}

func main() {
	var temp2 string = ""
	var temp []byte
	var s string
	var count int
	temp, err := ioutil.ReadAll(os.Stdin)
    if err != nil {
        panic(err)
    }
	var x, y int
	for _, word := range temp {
		if word == 10 {
			y++
		} else {
			x++
		}
 	}
	if x == 1 && y == 0 {
		y = 1
	}
	x -= y
	x /= y
	for i := 0; i < len(temp); i++ {
		if temp[i] == 13 || (temp[i] == 10 && i == len(temp)-1) {
			continue
		}
		s += string(temp[i])
	}
	for i := 1; i <= x; i++ {
		for j := 1; j <= y; j++ {
			if QuadA(i, j) == s {
				temp2 = "s"
				fmt.Printf("[quadA] [%v] [%v]", i, j)
			}
		}
	}
	for i := 1; i <= x; i++ {
		for j := 1; j <= y; j++ {
			if QuadB(i, j) == s {
				temp2 = "s"
				fmt.Printf("[quadB] [%v] [%v]", i, j)
			}
		}
	}
	for i := 1; i <= x; i++ {
		for j := 1; j <= y; j++ {
			if QuadC(i, j) == s {
				count++
			}
		}
	}
	for i := 1; i <= x; i++ {
		for j := 1; j <= y; j++ {
			if QuadD(i, j) == s {
				count++
			}
		}
	}
	for i := 1; i <= x; i++ {
		for j := 1; j <= y; j++ {
			if QuadE(i, j) == s {
				count++
			}
		}
	}
	for i := 1; i <= x; i++ {
		for j := 1; j <= y; j++ {
			if QuadC(i, j) == s {
				temp2 = "s"
				fmt.Printf("[quadC] [%v] [%v]", i, j)
			}
			if count != 1 && QuadC(i, j) == s {
				fmt.Printf(" || ")
				count--
			}
			if QuadD(i, j) == s {
				temp2 = "s"
				fmt.Printf("[quadD] [%v] [%v]", i, j)
			}
			if count != 1 && QuadD(i, j) == s{
				fmt.Printf(" || ")
				count--
			}
			if QuadE(i, j) == s {
				temp2 = "s"
				fmt.Printf("[quadE] [%v] [%v]", i, j)
			}
		}
	}
	if len(temp2) == 0 {
		fmt.Print("Not a Raid function")
	}
	fmt.Print("\n")
}
