package mongodb

import (
	"fmt"
	"testing"
)

// insert 系列用于插入测试用例

func TestInsertOneDoc1(t *testing.T) {
	// 不使用指针类型, 会不会有性能问题?
	d := Doc1{"canvas", 100, []string{"cotton"}, Size{28, 35.5, "cm"}}
	res, err := coll.InsertOne(nil, d)
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
	res, err := coll.InsertMany(nil, data)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%#v\n", res)
}

func TestInsertManyDoc3(t *testing.T) {
	data := []interface{}{
		Doc3{"journal", "A", Size{14, 21, "cm"}, []Instock{{"A", 5}}},
		Doc3{"notebook", "A", Size{8.5, 11, "in"}, []Instock{{"C", 5}}},
		Doc3{"paper", "D", Size{8.5, 11, "in"}, []Instock{{"A", 60}}},
		Doc3{"planner", "D", Size{22.85, 30, "cm"}, []Instock{{"A", 40}}},
		Doc3{"postcard", "A", Size{10, 15.25, "cm"}, []Instock{{"B", 15}, {"C", 35}}},
	}
	res, err := coll.InsertMany(nil, data)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%#v\n", res)
}
