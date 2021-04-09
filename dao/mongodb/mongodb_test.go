package mongodb

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.M) {
	url := "mongodb://root:toor@127.0.0.1:27017"
	err := Init(url)
	if err != nil {
		fmt.Printf("Init: %s\n", err)
		return
	}
	t.Run()
}
