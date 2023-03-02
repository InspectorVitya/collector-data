package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateUserDeviceInfo(ctx context.Context, info InfoDevice) error {
	_, err := mongodb.collection.InsertOne(ctx, info)
	return err
}

func GetUserInfoDeviceById(ctx context.Context, id uint32) (InfoDevice, error) {
	filter := bson.D{{"id", id}}
	result := InfoDevice{}
	err := mongodb.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return InfoDevice{}, err
	}
	return result, nil
}

func GetTop100(ctx context.Context, field, by, name string) ([]Top, error) {

	cursor, err := mongodb.collection.Aggregate(ctx, preparerPipeline(field, by, name))
	if err != nil {
		panic(err)
	}
	result := make([]Top, 0)
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		row := Top{}
		if err = cursor.Decode(&row); err != nil {
			return nil, err
		}
		if row.Name == "" {
			continue
		}
		result = append(result, row)
	}

	return result, err
}
