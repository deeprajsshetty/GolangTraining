// Alignment is about placing fields on address alignment boundaries
// for more efficient reads and writes to memory.

// Sample program to show how struct types align on boundaries.
package main

import (
	"fmt"
	"unsafe"
)

// No byte padding.
type nbp struct {
	a bool //	1 byte		sizeof 1
	b bool //	1 byte		sizeof 2
	c bool //	1 byte		sizeof 3	-	Aligned on 1 byte
}

// Single byte padding.
type sbp struct {
	a bool //	1 byte					sizeof 1
	//					1 byte padding	sizeof 2
	b int16 //	2 bytes					sizeof 4	-	Aligned on 2 bytes
}

// Two byte padding.
type tbp struct {
	a bool //	1 byte					sizeof 1
	//					3 byte padding	sizeof 4
	b int32 //	4 byte					sizeof 8	-	Aligned on 2 bytes
}

// Four byte padding.
type fbp struct {
	a bool //	1 byte					sizeof 1
	//					7 bytes padding	sizeof 8
	b int64 //	8 bytes					sizeof 16	-	Aligned on 8 bytes
}

// Eight byte padding on 64bit Arch. Word size is 8 bytes.
type ebp64 struct {
	a string //	16 bytes							sizeof 16
	b int32  //	4 bytes								sizeof 20
	//							4 bytes padding		sizeof 24
	c string //	16 bytes							sizeof 40
	d int32  //	4 bytes								sizeof 44
	//							4 bytes padding		sizeof 48	-	Aligned on 8 bytes
}

// No padding on 32bit Arch. Word size is 4 bytes.
// To see this build as 32 bit: GOARCH=386 go build
type ebp32 struct {
	a string // 8 bytes							sizeof 8
	b int32  // 4 bytes							sizeof 12
	c string // 8 bytes							sizeof 20
	d int32  // 4 bytes							sizeof 24	- Aligned on 4 bytes
}

// No padding
type np struct {
	a string // 16 bytes						sizeof 16
	b string // 16 bytes						sizeof 32
	c int32  // 8 bytes							sizeof 40
	d int32  // 8 bytes							sizeof 48
}

func main() {

	// No byte padding.
	fmt.Println("******No byte padding.*******************************************************************************************")
	var nbp nbp
	size := unsafe.Sizeof(nbp)
	fmt.Printf("SizeOf[%d][%p %p %p]\n", size, &nbp.a, &nbp.b, &nbp.c)

	// Single byte padding.
	fmt.Println("******Single byte padding.****************************************************************************************")
	var sbp sbp
	size = unsafe.Sizeof(sbp)
	fmt.Printf("SizeOf[%d][%p %p]\n", size, &sbp.a, &sbp.b)

	// Two byte padding.
	fmt.Println("******Two byte padding.*******************************************************************************************")
	var tbp tbp
	size = unsafe.Sizeof(tbp)
	fmt.Printf("SizeOf[%d][%p %p]\n", size, &tbp.a, &tbp.b)

	// Four byte padding.
	fmt.Println("******Four byte padding.*******************************************************************************************")
	var fbp fbp
	size = unsafe.Sizeof(fbp)
	fmt.Printf("SizeOf[%d][%p %p]\n", size, &fbp.a, &fbp.b)

	// Eight byte padding on 64bit Arch. Word size is 8 bytes.
	fmt.Println("******Eight byte padding on 64bit Arch. Word size is 8 bytes.*******************************************************")
	var ebp64 ebp64
	size = unsafe.Sizeof(ebp64)
	fmt.Printf("SizeOf[%d][%p %p %p %p]\n", size, &ebp64.a, &ebp64.b, &ebp64.c, &ebp64.d)

	// Eight byte padding on 32bit Arch. Word size is 8 bytes.
	fmt.Println("******Eight byte padding on 32bit Arch. Word size is 8 bytes.*******************************************************")
	var ebp32 ebp32
	size = unsafe.Sizeof(ebp32)
	fmt.Printf("SizeOf[%d][%p %p %p %p]\n", size, &ebp32.a, &ebp32.b, &ebp32.c, &ebp32.d)

	// No Padding
	fmt.Println("******No Padding.***************************************************************************************************")
	var np np
	size = unsafe.Sizeof(np)
	fmt.Printf("SizeOf[%d][%p %p %p %p]\n", size, &np.a, &np.b, &np.c, &np.d)
}

/*Output
******No byte padding.*******************************************************************************************
SizeOf[3][0xc00001a0b8 0xc00001a0b9 0xc00001a0ba]
******Single byte padding.****************************************************************************************
SizeOf[4][0xc00001a0bc 0xc00001a0be]
******Two byte padding.*******************************************************************************************
SizeOf[8][0xc00001a0e0 0xc00001a0e4]
******Four byte padding.*******************************************************************************************
SizeOf[16][0xc00001a0f0 0xc00001a0f8]
******Eight byte padding on 64bit Arch. Word size is 8 bytes.*******************************************************
SizeOf[48][0xc00007c4b0 0xc00007c4c0 0xc00007c4c8 0xc00007c4d8]
******Eight byte padding on 32bit Arch. Word size is 8 bytes.*******************************************************
SizeOf[48][0xc00007c4e0 0xc00007c4f0 0xc00007c4f8 0xc00007c508]
******No Padding.***************************************************************************************************
SizeOf[40][0xc00007c510 0xc00007c520 0xc00007c530 0xc00007c534]
*/
