package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const todoCollection = "todos"

func OpenConnection(conString string, dbName string) *DBHandle {
	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI(conString)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	return &DBHandle{monDB: client.Database(dbName)}
}

func (db *DBHandle) GetTodos(userID string) ([]Todo, error) {
	var todos []Todo
	cursor, err := db.monDB.Collection(todoCollection).Find(context.TODO(), bson.M{"userId": userID})

	if err != nil {
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var todo Todo
		err := cursor.Decode(&todo)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (db *DBHandle) DeleteTodo(todoId string) error {
	_, err := db.monDB.Collection(todoCollection).DeleteOne(context.TODO(), bson.M{"_id": todoId})
	return err
}

func (db *DBHandle) UpdateTodo(userID string, todo Todo) error {
	res := db.monDB.Collection(todoCollection).FindOneAndReplace(context.TODO(), bson.M{"_id": todo.ID}, todo)
	if res.Err() != nil {
		return res.Err()
	}
	return nil
}

func (db *DBHandle) AddTodo(userID string, todo Todo) error {
	_, err := db.monDB.Collection(todoCollection).InsertOne(context.TODO(), todo)
	return err
}

func (db *DBHandle) GetTodoById(todoId string) (Todo, error) {
	var todo Todo
	err := db.monDB.Collection(todoCollection).FindOne(context.TODO(), bson.M{"_id": todoId}).Decode(&todo)
	return todo, err
}
