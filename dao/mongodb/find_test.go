package mongodb

import (
	"fmt"
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func find(t *testing.T, filter bson.D, docType reflect.Type) {
	// todo: bson.D 等类型的作用

	// 通过查看源码, bson.D 是 bson.E 的数组形式, 而 bson.E 是个 struct.
	// 在构造 bson.D 时, 可以参考字符数组:
	//     字符数组: []string{"abc", "def"}
	//     bson.D: bson.D{{Key: "abc", Value: "def"}, {Key: "123", Value: "456"}}
	cursor, err := coll.Find(nil, filter)
	if err != nil {
		t.Fatalf("Find: %s\n", err)
	}
	defer func() {
		// 参考 https://mongoing.com/archives/27257 例子,
		// 先判断 Err, 再 Close.
		if err := cursor.Err(); err != nil {
			t.Logf("Err: %s\n", err)
		}
		if err := cursor.Close(nil); err != nil {
			t.Fatalf("Close: %s\n", err)
		}
	}()

	var docs []interface{}
	for cursor.Next(nil) {
		doc := reflect.New(docType).Interface()
		err := cursor.Decode(doc)
		if err != nil {
			t.Fatalf("Decode: %s\n", err)
		}
		docs = append(docs, doc)
	}

	for _, doc := range docs {
		fmt.Printf("%#v\n", doc)
	}
}

func TestFindAllDoc2(t *testing.T) {
	filter := bson.D{}
	find(t, filter, reflect.TypeOf(Doc2{}))
}

func TestFindEqualDoc2(t *testing.T) {
	filter := bson.D{{Key: "status", Value: "D"}}
	find(t, filter, reflect.TypeOf(Doc2{}))
}

func TestFindInDoc2(t *testing.T) {
	filter := bson.D{{
		"status",
		bson.D{{
			"$in",
			// bson.A 的底层类型是 []interface{}
			bson.A{"A", "D"},
		}},
	}}

	find(t, filter, reflect.TypeOf(Doc2{}))
}

// TestFindAnd 测试 and 查询
func TestFindAndDoc2(t *testing.T) {
	filter := bson.D{
		{"status", "A"},
		{"qty", bson.D{{"$lt", 30}}},
	}
	find(t, filter, reflect.TypeOf(Doc2{}))
}

// TestFindOr 测试 or 查询, 这里也进行了 lt 查询, 可以看到查询写起来比较复杂.
func TestFindOrDoc2(t *testing.T) {
	filter := bson.D{
		{
			"$or",
			bson.A{
				bson.D{{"status", "A"}},
				bson.D{{"qty", bson.D{{"$lt", 30}}}},
			},
		},
	}
	// 查询结果不重复
	find(t, filter, reflect.TypeOf(Doc2{}))
}

func TestFindAndOrDoc2(t *testing.T) {
	filter := bson.D{
		{"status", "A"},
		{
			"$or",
			bson.A{
				bson.D{{"qty", bson.D{{"$lt", 30}}}},
				bson.D{
					{
						"item",
						primitive.Regex{Pattern: "^p", Options: "i"},
					},
				},
			},
		},
	}
	find(t, filter, reflect.TypeOf(Doc2{}))
}

// todo
func TestFindRegex(t *testing.T) {

}

func TestFindProjectionDoc3(t *testing.T) {
	opts := &options.FindOptions{}
	opts.Projection = bson.D{
		{"item", 1},
		{"status", 1},
	}

	cursor, err := coll.Find(nil, bson.D{}, opts)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	defer func() {
		if err := cursor.Err(); err != nil {
			t.Logf("Err: %s\n", err)
		}
		if err := cursor.Close(nil); err != nil {
			t.Fatalf("Close: %s\n", err)
		}
	}()

	var docs []*Doc3
	for cursor.Next(nil) {
		doc := &Doc3{}
		err := cursor.Decode(doc)
		if err != nil {
			t.Fatalf("%s\n", err)
		}
		docs = append(docs, doc)
	}
	for _, doc := range docs {
		fmt.Printf("%#v\n", doc)
	}
}
