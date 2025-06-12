package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var list [27]int
	for i := 0; i < len(list); i++ {
		list[i] = 50
		if i == 10 {
			list[i] = 25
		}
	}
	rand.Shuffle(len(list), func(i, j int) {
		list[i], list[j] = list[j], list[i]
	})

	// fmt.Println(list)
	var group1 []int
	var group2 []int
	var total1 int
	var total2 int

	group1 = append(group1, list[0:9]...)
	group2 = append(group2, list[9:18]...)

	for i := 0; i < len(group1); i++ {
		total1 = total1 + group1[i]
	}
	for i := 0; i < len(group2); i++ {
		total2 = total2 + group2[i]
	}
	fmt.Println("Total del grupo 1:", total1)
	fmt.Println("Total del grupo 2:", total2)
	if total1 < total2 {
		clear(group1)
		clear(group2)
		group1 = append(group1, list[0:3]...)
		group2 = append(group2, list[3:6]...)
		totalcaso21 := 0
		totalcaso22 := 0
		for i := 0; i < len(group1); i++ {
			totalcaso21 = totalcaso21 + group1[i]
		}
		for i := 0; i < len(group2); i++ {
			totalcaso22 = totalcaso22 + group2[i]
		}
		if totalcaso21 < totalcaso22 {
			clear(group1)
			clear(group2)
			group1 = append(group1, list[0])
			group2 = append(group2, list[1])
			total1 = 0
			total2 = 0
			for i := 0; i < len(group1); i++ {
				total1 = total1 + group1[i]
			}
			for i := 0; i < len(group2); i++ {
				total2 = total2 + group2[i]
			}
			if total1 < total2 {
				fmt.Println("La bola mas ligera está en la posición 0 con un peso de: ", list[0])
			}
		}

	} else if total1 > total2 {
		clear(group1)
		clear(group2)
		group1 = append(group1, list[9:12]...)
		group2 = append(group2, list[12:15]...)
		total1 = 0
		total2 = 0
		for i := 0; i < len(group1); i++ {
			total1 = total1 + group1[i]
		}
		for i := 0; i < len(group2); i++ {
			total2 = total2 + group2[i]
		}
	} else {
		clear(group1)
		clear(group2)
		group1 = append(group1, list[18:21]...)
		group2 = append(group2, list[21:24]...)
		total1 = 0
		total2 = 0
		for i := 0; i < len(group1); i++ {
			total1 = total1 + group1[i]
		}
		for i := 0; i < len(group2); i++ {
			total2 = total2 + group2[i]
		}

	}

}
