package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Student struct {
	Name string
	Age  int
}

var s1 = Student{"小红", 12}
var s2 = Student{"小兰", 10}
var s3 = Student{"小黄", 11}
var filter = bson.D{{"name", "小兰"}}

func main() {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	// 关闭
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Connection to MongoDB closed.")

	// db, err := ConnectToDB("mongodb://localhost:27017", "client", 10000, 10)

	// 指定获取要操作的数据
	c := client.Database("test").Collection("student")
	insertOne(c)
	insertMany(c)
	update(c)
	queryOne(c)
	queryMany(c)
	delete(c)
}
func insertOne(collection *mongo.Collection) {
	// 插入单条操作
	insertResult, err := collection.InsertOne(context.TODO(), s1)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}
func insertMany(collection *mongo.Collection) {
	// 插入多条操作
	students := []interface{}{s2, s3}
	insertManyResult, err := collection.InsertMany(context.TODO(), students)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
}
func update(collection *mongo.Collection) {
	// 更新文档
	var _update = bson.D{{"$inc", bson.D{
		{"age", 1},
	}}}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, _update)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}
func queryOne(collection *mongo.Collection) {
	// 查找文档
	var result Student
	filter := bson.D{{"name", "小兰"}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("Found a single document: %+v\n", result)
}
func queryMany(collection *mongo.Collection) {
	// 查询多个
	// 将选项传递给Find()
	findOptions := options.Find()
	findOptions.SetLimit(2)

	// 定义一个切片用来存储查询结果
	var results []*Student

	// 把bson.D{{}}作为一个filter来匹配所有文档
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 查找多个文档返回一个光标
	// 遍历游标允许我们一次解码一个文档
	for cur.Next(context.TODO()) {
		// 创建一个值，将单个文档解码为该值
		var elem Student
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// 完成后关闭游标
	cur.Close(context.TODO())
	fmt.Printf("Found multiple documents (array of pointers): %#v\n", results)
}
func delete(collection *mongo.Collection) {
	// 删除名字是小黄的那个
	deleteResult1, err := collection.DeleteOne(context.TODO(), bson.D{{"name", "小黄"}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult1.DeletedCount)
	// 删除所有
	deleteResult2, err := collection.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult2.DeletedCount)
}
func ConnectToDB(uri, name string, timeout time.Duration, num uint64) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	o := options.Client().ApplyURI(uri)
	o.SetMaxPoolSize(num)
	client, err := mongo.Connect(ctx, o)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return client.Database(name), nil
}
