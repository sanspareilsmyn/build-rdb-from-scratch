package main

import (
	"fmt"
	"github.com/collinglass/bptree"
	"testing"
)

func TestBPTree(t *testing.T) {
	tree := bptree.NewTree()

	keys := []int{1, 4, 2, 3, 5, 6, 8, 9, 7, 10}
	values := []string{"data1", "data4", "data2", "data3", "data5", "data6", "data8", "data9", "data7", "data10"}

	for i, key := range keys {
		err := tree.Insert(key, []byte(values[i]))
		if err != nil {
			t.Fatalf("Error inserting key %d: %v", key, err)
		}
	}

	fmt.Println("Find key 3:")
	tree.FindAndPrint(3, true)

	fmt.Println("\nFind key 7:")
	tree.FindAndPrint(7, true)

	fmt.Println("\nFind key 11:")
	tree.FindAndPrint(11, true)

	fmt.Println("\nRange search from 3 to 7:")
	tree.FindAndPrintRange(3, 7, true)

	fmt.Println("\nDelete key 3:")
	tree.Delete(3)
	fmt.Println("Find key 3 after deletion:")
	tree.FindAndPrint(3, true)

	fmt.Println("\nDelete key 1:")
	tree.Delete(1)
	fmt.Println("Find key 1 after deletion:")
	tree.FindAndPrint(1, true)

	fmt.Println("\nFind key 2 after delete 1:")
	tree.FindAndPrint(2, true)

	fmt.Println("\nFind key 4 after delete 1 and 3:")
	tree.FindAndPrint(4, true)

	fmt.Println("\nFind key 5 after delete 1 and 3:")
	tree.FindAndPrint(5, true)

	fmt.Println("\nFind key 6 after delete 1 and 3:")
	tree.FindAndPrint(6, true)

	fmt.Println("\nFind key 8 after delete 1 and 3:")
	tree.FindAndPrint(8, true)

	fmt.Println("\nFind key 9 after delete 1 and 3:")
	tree.FindAndPrint(9, true)

	fmt.Println("\nFind key 7 after delete 1 and 3:")
	tree.FindAndPrint(7, true)

	fmt.Println("\nFind key 10 after delete 1 and 3:")
	tree.FindAndPrint(10, true)

}
