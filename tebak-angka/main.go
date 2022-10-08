package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 10
	randNumber := rand.Intn(max-min+1) + min

	for chance := 3; chance > 0; chance-- {
		var inputNumber int
		fmt.Println("Tebak angka dari 1 sampai 10")
		fmt.Print("Masukkan angka : ")
		_, err := fmt.Scan(&inputNumber)
		if err != nil {
			panic(err)
		}

		fmt.Println()
		if inputNumber == randNumber {
			fmt.Println("selamat tebakan anda benar!")
			return
		} else if inputNumber > randNumber {
			fmt.Println("angka terlalu tinggi")
		} else {
			fmt.Println("angka terlalu rendah")
		}
	}

	fmt.Println("kesempatan habis! kamu kalah!")
	time.Sleep(10 * time.Second)
}
