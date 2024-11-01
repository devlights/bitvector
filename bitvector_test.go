package bitvector_test

import (
	"fmt"
	"log"

	"github.com/devlights/bitvector"
)

func ExampleBitVector() {
	const (
		bitSize = 32
	)

	log.SetFlags(0)

	//
	// ビットの設定
	//
	var (
		bv  = bitvector.New(bitSize)
		err error
	)

	for _, i := range []int{1, 3, 23} {
		if err = bv.Set(i, true); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("ORIG : %s\n", bv)

	//
	// ビットの取得
	//
	var (
		bits []bool
		bit  bool
	)

	for _, i := range []int{1, 2, 3, 4} {
		if bit, err = bv.Get(i); err != nil {
			log.Fatal(err)
		}

		bits = append(bits, bit)
	}
	fmt.Printf("BITS : %v\n", bits)

	//
	// 別のビットベクタとの演算
	//
	var (
		bv2   = bitvector.New(bitSize)
		bvAnd *bitvector.BitVector
		bvOr  *bitvector.BitVector
		bvXor *bitvector.BitVector
	)

	for _, i := range []int{2, 23, 24} {
		if err = bv2.Set(i, true); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("OTHER: %s\n", bv2)

	// AND
	if bvAnd, err = bv.And(bv2); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("AND  : %s\n", bvAnd)

	// OR
	if bvOr, err = bv.Or(bv2); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("OR   : %s\n", bvOr)

	// XOR
	if bvXor, err = bv.Xor(bv2); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("XOR  : %s\n", bvXor)

	// Output:
	// ORIG : 01010000000000000000000100000000
	// BITS : [true false true false]
	// OTHER: 00100000000000000000000110000000
	// AND  : 00000000000000000000000100000000
	// OR   : 01110000000000000000000110000000
	// XOR  : 01110000000000000000000010000000
}
