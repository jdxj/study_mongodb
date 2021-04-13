package mongodb

import (
	"fmt"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

func TestBulkWrite(t *testing.T) {
	iom1 := &mongo.InsertOneModel{
		Doc8{[12]byte{11: 1}, "Brisbane", "monk", 4},
	}
	iom2 := &mongo.InsertOneModel{
		Doc8{[12]byte{11: 2}, "Eldon", "alchemist", 3},
	}
	iom3 := &mongo.InsertOneModel{
		Doc8{[12]byte{11: 3}, "Meldane", "ranger", 3},
	}

	ioms := []mongo.WriteModel{iom1, iom2, iom3}

	res, err := characters.BulkWrite(nil, ioms)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%#v\n", res)
}
