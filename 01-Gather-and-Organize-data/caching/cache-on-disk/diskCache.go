package main

import (
	"fmt"
	"os"

	"github.com/boltdb/bolt"
)

func main() {
	// Open an embedded.db data file in current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("embedded.db", 0600, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	// Create a "bucket" in the boltdb file for our data.
	err = db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucket([]byte("MyBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		return nil
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Put the map keys and values into the BoltDB file.
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		err = b.Put([]byte("Petros"), []byte("Trak"))
		err = b.Put([]byte("Eirini"), []byte("Tour"))
		err = b.Put([]byte("George"), []byte("Trak"))
		err = b.Put([]byte("Maria"), []byte("Geo"))
		err = b.Put([]byte("Maggie"), []byte("Trak"))
		err = b.Put([]byte("Elias"), []byte("Tour"))
		return err
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Output the keys and values in the embedded db file to
	// stdOut.
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key: %s, value: %s\n", k, v)
		}

		return nil
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
