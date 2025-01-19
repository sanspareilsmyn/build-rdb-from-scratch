package main

import (
	"fmt"
	"github.com/sanspareilsmyn/build-rdb-from-scratch/01-implementing-bplus-tree/bplustree"
)

func main() {
	tree := bplustree.NewTree()
	// Insert
	tree.Insert(1, []byte("data1"))
	tree.Insert(4, []byte("data4"))
	tree.Insert(2, []byte("data2"))
	tree.Insert(3, []byte("data3"))
	tree.Insert(5, []byte("data5"))
	tree.Insert(6, []byte("data6"))

	// Get
	r, err := tree.Find(3, false)
	if err == nil {
		fmt.Println("key 3:", string(r.Value))
	}

	if _, err := tree.Find(7, false); err != nil {
		fmt.Println("key 7 not found")
	}

	// Delete
	tree.Delete(3)
	tree.Delete(1)

	if _, err := tree.Find(3, false); err != nil {
		fmt.Println("key 3 is deleted")
	}

	if _, err := tree.Find(1, false); err != nil {
		fmt.Println("key 1 is deleted")
	}

	r, err = tree.Find(2, false)
	if err == nil {
		fmt.Println("key 2:", string(r.Value))
	}

	r, err = tree.Find(4, false)
	if err == nil {
		fmt.Println("key 4:", string(r.Value))
	}

	r, err = tree.Find(5, false)
	if err == nil {
		fmt.Println("key 5:", string(r.Value))
	}
	r, err = tree.Find(6, false)
	if err == nil {
		fmt.Println("key 6:", string(r.Value))
	}
}
