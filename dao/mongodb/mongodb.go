package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbName = "study_mongodb"
	coName = "test"
)

var (
	client *mongo.Client
	coll   *mongo.Collection
)

func Init(url string) error {
	opt := options.Client().ApplyURI(url)
	c, err := mongo.Connect(context.TODO(), opt)
	if err != nil {
		return err
	}

	err = c.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}
	client = c
	coll = client.Database(dbName).Collection(coName)
	return nil
}

// 为了使数据清晰, 所以定义了多种文档类型

type Size struct {
	H   float32
	W   float32
	UOM string
}

type Doc1 struct {
	Item string
	QTY  int
	Tags []string
	Size Size
}

type Doc2 struct {
	Item   string
	QTY    int
	Size   Size
	Status string
}

type Instock struct {
	Warehouse string
	QTY       int
}

type Doc3 struct {
	Item    string
	Status  string
	Size    Size
	Instock []Instock
}

type Doc4 struct {
	Item    string
	Instock []Instock
}

type Doc5 struct {
	Item  string
	QTY   int
	Tags  []string
	DimCm []float32
}
