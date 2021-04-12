package mongodb

import (
	"fmt"
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func find(t *testing.T, filter bson.D, doc interface{}, opts ...*options.FindOptions) {
	// todo: bson.D 等类型的作用

	// 通过查看源码, bson.D 是 bson.E 的数组形式, 而 bson.E 是个 struct.
	// 在构造 bson.D 时, 可以参考字符数组:
	//     字符数组: []string{"abc", "def"}
	//     bson.D: bson.D{{Key: "abc", Value: "def"}, {Key: "123", Value: "456"}}
	cursor, err := coll.Find(nil, filter, opts...)
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

	docType := reflect.TypeOf(doc)
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
	find(t, filter, Doc2{})
}

func TestFindEqualDoc2(t *testing.T) {
	filter := bson.D{{Key: "status", Value: "D"}}
	find(t, filter, Doc2{})
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

	find(t, filter, Doc2{})
}

// TestFindAnd 测试 and 查询
func TestFindAndDoc2(t *testing.T) {
	filter := bson.D{
		{"status", "A"},
		{"qty", bson.D{{"$lt", 30}}},
	}
	find(t, filter, Doc2{})
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
	find(t, filter, Doc2{})
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
	find(t, filter, Doc2{})
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

	find(t, bson.D{}, Doc3{}, opts)
}

func TestProjectionDoc3EmbeddedField(t *testing.T) {
	opts := &options.FindOptions{}
	opts.Projection = bson.D{
		{"item", 1},
		{"status", 1},
		{"size.uom", 1},
	}

	filter := bson.D{
		{"status", "A"},
	}

	find(t, filter, Doc3{}, opts)
}

func TestProjectionDoc3ArrayElem(t *testing.T) {
	filter := bson.D{
		{"status", "A"},
	}

	opts := &options.FindOptions{}
	opts.Projection = bson.D{
		{"item", 1},
		{"status", 1},
		{"instock.qty", 1},
	}

	find(t, filter, Doc3{}, opts)
}

// TestProjectionDoc3ArrayElemSlice test $slice cmd
func TestProjectionDoc3ArrayElemSlice(t *testing.T) {
	filter := bson.D{
		{"status", "A"},
	}

	opts := &options.FindOptions{}
	opts.Projection = bson.D{
		{"item", 1},
		{"status", 1},
		// 对比
		//{"instock", 1},
		{
			"instock",
			bson.D{{"$slice", -1}},
		},
	}

	find(t, filter, Doc3{}, opts)
}

func TestFindDoc4ArrayElem(t *testing.T) {
	filter := bson.D{
		{
			"instock",
			// 注意字段顺序, 顺序不一致是无法匹配的
			bson.D{
				{"warehouse", "A"},
				{"qty", 5},
			},
		},
	}

	find(t, filter, Doc4{})
}

func TestFindDoc4ArrayElemField(t *testing.T) {
	filter := bson.D{
		{
			"instock.qty",
			// 这种方式 (没有使用 $elemMatch) 表示数组中的元素的字段只要满足其中一个条件就会返回该 doc
			bson.D{{"$lte", 20}},
		},
	}

	find(t, filter, Doc4{})
}

func TestFindDoc4ArrayIndexField(t *testing.T) {
	filter := bson.D{
		{
			"instock.0.qty",
			bson.D{{"$lte", 20}},
		},
	}

	find(t, filter, Doc4{})
}

func TestFindDoc4ArrayElemMatch(t *testing.T) {
	filter := bson.D{
		{
			"instock",
			bson.D{
				{
					"$elemMatch",
					// 同时满足指定条件
					bson.D{
						{"qty", 5},
						{"warehouse", "A"},
					},
				},
			},
		},
	}

	find(t, filter, Doc4{})
	fmt.Println("---------")

	filter2 := bson.D{
		{
			"instock",
			bson.D{
				{
					"$elemMatch",
					bson.D{
						{
							"qty",
							bson.D{
								{"$gt", 10},
								{"$lte", 20},
							},
						},
					},
				},
			},
		},
	}
	find(t, filter2, Doc4{})
	fmt.Println("------")

	// 与 $elemMatch 对比
	filter3 := bson.D{
		{
			"instock.qty",
			bson.D{
				// 猜测, 先对员集合 $gt, 再进行 $lte, 两个结果做交集
				{"$gt", 10},
				{"$lte", 20},
			},
		},
	}
	find(t, filter3, Doc4{})
}

func TestFindDoc5Array(t *testing.T) {
	// 精确匹配
	filter := bson.D{
		{"tags", []string{"red", "blank"}},
	}
	fmt.Println("精确匹配:")
	find(t, filter, Doc5{})

	// 只要包含就可以
	filter = bson.D{
		{
			"tags",
			bson.D{
				{"$all", []string{"red", "blank"}},
			},
		},
	}
	fmt.Println("只要包含就可以:")
	find(t, filter, Doc5{})

}
