package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		setSHA256()
	} else {
		checkFlag(os.Args[1])
	}
}

func checkFlag(flag string) {
	if flag == "256" {
		setSHA256()
	} else if flag == "384" {
		setSHA384()
	} else if flag == "512" {
		setSHA512()
	} else {
		setSHA256()
	}
}

func setSHA256() {
	c1 := sha256.Sum256(getArray1())
	c2 := sha256.Sum256(getArray2())

	fmt.Println("Printing SHA-256")
	Print32(c1)
	Print32(c2)
}

func setSHA384() {
	c1 := sha512.Sum384(getArray1())
	c2 := sha512.Sum384(getArray1())

	fmt.Println("Printing SHA-384")
	Print48(c1)
	Print48(c2)
}

func setSHA512() {
	c1 := sha512.Sum512(getArray1())
	c2 := sha512.Sum512(getArray2())

	fmt.Println("Printing SHA-512")
	Print64(c1)
	Print64(c2)
}

func getArray1() []byte {
	return []byte("x")
}

func getArray2() []byte {
	return []byte("X")
}

func Print32(data [32]byte) {
	fmt.Printf("%x\n", data)
}

func Print48(data [48]byte) {
	fmt.Printf("%x\n", data)
}

func Print64(data [64]byte) {
	fmt.Printf("%x\n", data)
}
