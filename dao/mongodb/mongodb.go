package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbName      = "study_mongodb"
	cInventory  = "inventory"
	cCharacters = "characters"
	cStores     = "stores"
)

var (
	client     *mongo.Client
	db         *mongo.Database
	inventory  *mongo.Collection
	characters *mongo.Collection
	stores     *mongo.Collection
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
	db = client.Database(dbName)
	inventory = db.Collection(cInventory)
	characters = db.Collection(cCharacters)
	stores = db.Collection(cStores)
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

// Doc6 Doc7 用于测试空值

type Doc6 struct {
	ID   primitive.ObjectID `bson:"_id"`
	Item *int
}

type Doc7 struct {
	ID primitive.ObjectID `bson:"_id"`
}

type Doc8 struct {
	ID    primitive.ObjectID `bson:"_id"`
	Char  string
	Class string
	LVL   int
}

type Doc9 struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string
	Description string
}

type Doc10 struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string
	Description string
	Score       float32
}
