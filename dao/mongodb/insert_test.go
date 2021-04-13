package mongodb

import (
	"context"
	"fmt"
	"testing"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// insert 系列用于插入测试用例, 测试 find_test.go 中的函数前需要清除不同的 Doc.

func insert(t *testing.T, coll *mongo.Collection, data []interface{}) {
	res, err := coll.InsertMany(nil, data)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%#v\n", res)
}

func TestInsertOneDoc1(t *testing.T) {
	// 不使用指针类型, 会不会有性能问题?
	d := Doc1{"canvas", 100, []string{"cotton"}, Size{28, 35.5, "cm"}}
	res, err := inventory.InsertOne(nil, d)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%#v\n", res)
}

func TestInsertManyDoc2(t *testing.T) {
	data := []interface{}{
		Doc2{"journal", 25, Size{14, 21, "cm"}, "A"},
		Doc2{"notebook", 50, Size{8.5, 11, "in"}, "A"},
		Doc2{"paper", 100, Size{8.5, 11, "in"}, "D"},
		Doc2{"planner", 75, Size{22.85, 30, "cm"}, "D"},
		Doc2{"postcard", 45, Size{10, 15.25, "cm"}, "A"},
	}

	insert(t, inventory, data)
}

// TestInsertManyDoc2Update 用于测试 update 相关 api
func TestInsertManyDoc2Update(t *testing.T) {
	data := []interface{}{
		Doc2{"canvas", 100, Size{28, 35.5, "cm"}, "A"},
		Doc2{"journal", 25, Size{14, 21, "cm"}, "A"},
		Doc2{"mat", 85, Size{27.9, 35.5, "cm"}, "A"},
		Doc2{"mousepad", 25, Size{19, 22.85, "cm"}, "P"},
		Doc2{"notebook", 50, Size{8.5, 11, "in"}, "P"},
		Doc2{"paper", 100, Size{8.5, 11, "in"}, "D"},
		Doc2{"planner", 75, Size{22.85, 30, "cm"}, "D"},
		Doc2{"postcard", 45, Size{10, 15.25, "cm"}, "A"},
		Doc2{"sketchbook", 80, Size{14, 21, "cm"}, "A"},
		Doc2{"sketch pad", 95, Size{22.85, 30.5, "cm"}, "A"},
	}
	insert(t, inventory, data)
}

func TestInsertManyDoc3(t *testing.T) {
	data := []interface{}{
		Doc3{"journal", "A", Size{14, 21, "cm"}, []Instock{{"A", 5}}},
		Doc3{"notebook", "A", Size{8.5, 11, "in"}, []Instock{{"C", 5}}},
		Doc3{"paper", "D", Size{8.5, 11, "in"}, []Instock{{"A", 60}}},
		Doc3{"planner", "D", Size{22.85, 30, "cm"}, []Instock{{"A", 40}}},
		Doc3{"postcard", "A", Size{10, 15.25, "cm"}, []Instock{{"B", 15}, {"C", 35}}},
	}

	insert(t, inventory, data)
}

func TestInsertManyDoc4(t *testing.T) {
	data := []interface{}{
		Doc4{"journal", []Instock{{"A", 5}, {"C", 15}}},
		Doc4{"notebook", []Instock{{"C", 5}}},
		Doc4{"paper", []Instock{{"A", 60}, {"B", 15}}},
		Doc4{"planner", []Instock{{"A", 40}, {"B", 5}}},
		Doc4{"postcard", []Instock{{"B", 15}, {"C", 35}}},
	}

	insert(t, inventory, data)
}

func TestInsertManyDoc5(t *testing.T) {
	data := []interface{}{
		Doc5{"journal", 25, []string{"blank", "red"}, []float32{14, 21}},
		Doc5{"notebook", 50, []string{"red", "blank"}, []float32{14, 21}},
		Doc5{"paper", 100, []string{"red", "blank", "plain"}, []float32{14, 21}},
		Doc5{"planner", 75, []string{"blank", "red"}, []float32{22.85, 30}},
		Doc5{"postcard", 45, []string{"blue"}, []float32{10, 15.25}},
	}

	insert(t, inventory, data)
}

func TestInsertManyDoc6(t *testing.T) {
	data := []interface{}{
		Doc6{ID: primitive.NewObjectID(), Item: nil},
		Doc7{ID: primitive.NewObjectID()},
	}
	insert(t, inventory, data)
}

func TestInsertDoc8(t *testing.T) {
	data := []interface{}{
		Doc8{[12]byte{11: 1}, "Brisbane", "monk", 4},
		Doc8{[12]byte{11: 2}, "Eldon", "alchemist", 3},
		Doc8{[12]byte{11: 3}, "Meldane", "ranger", 3},
	}

	insert(t, characters, data)
}

func TestInsertManyDoc9(t *testing.T) {
	data := []interface{}{
		Doc9{[12]byte{11: 1}, "Java Hut", "Coffee and cakes"},
		Doc9{[12]byte{11: 2}, "Burger Buns", "Gourmet hamburgers"},
		Doc9{[12]byte{11: 3}, "Coffee Shop", "Just coffee"},
		Doc9{[12]byte{11: 4}, "Clothes Clothes Clothes", "Discount clothing"},
		Doc9{[12]byte{11: 5}, "Java Shopping", "Indonesian goods"},
	}

	insert(t, stores, data)

	// 用于测试文本查询所建立的索引
	im := mongo.IndexModel{
		Keys: bson.D{{"name", "text"}, {"description", "text"}},
	}

	indexName, err := stores.Indexes().CreateOne(context.TODO(), im)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("index name: %s\n", indexName)
}
