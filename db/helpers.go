package db

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func preparerPipeline(field, by, name string) mongo.Pipeline {
	stages := make([]bson.D, 0, 5)
	if by != "" {
		stages = append(stages, bson.D{{"$match", bson.D{{by, name}}}})
	}

	stages = append(stages, bson.D{
		{"$group", bson.D{
			{"_id", fmt.Sprintf("$%s", field)},
			{"count", bson.D{{"$sum", 1}}},
		}}})
	stages = append(stages, bson.D{{"$sort", bson.D{{"count", -1}}}})
	stages = append(stages, bson.D{{"$limit", 100}})
	return stages
}
