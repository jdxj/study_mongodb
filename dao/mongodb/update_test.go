package mongodb

import (
	"fmt"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// resetUpdate 重置为初始数据
func resetUpdate(t *testing.T) {
	err := inventory.Drop(nil)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Println("reset update:")
	TestInsertManyDoc2Update(t)
}

func updateOne(t *testing.T, filter bson.D, update bson.D, opts ...*options.UpdateOptions) {
	res, err := inventory.UpdateOne(nil, filter, update, opts...)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%#v\n", res)
}

func replaceOne(t *testing.T, filter bson.D, replacement bson.D, opts ...*options.ReplaceOptions) {
	res, err := inventory.ReplaceOne(nil, filter, replacement, opts...)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%#v\n", res)
}

func TestUpdateDoc2(t *testing.T) {
	resetUpdate(t)

	filter := bson.D{{"item", "paper"}}
	update := bson.D{
		{"$set", bson.D{{"size.uom", "cm"}, {"status", "P"}}},
		// 会添加 lastModified 字段
		{"$currentDate", bson.D{{"lastModified", true}}},
	}
	fmt.Println("更新前:")
	find(t, inventory, filter, Doc2{})
	fmt.Println("更新后:")
	updateOne(t, filter, update)
	find(t, inventory, filter, Doc2{})
}

func TestReplaceOne(t *testing.T) {
	resetUpdate(t)

	filter := bson.D{{"item", "paper"}}
	replace := bson.D{
		{"item", "paper"},
		{
			"instock",
			[]Instock{{"A", 60}, {"B", 40}},
		},
	}
	fmt.Println("替换前:")
	find(t, inventory, filter, Doc2{})
	fmt.Println("替换后:")
	replaceOne(t, filter, replace)
	find(t, inventory, filter, Doc4{})
}
