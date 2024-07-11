package clout

import (
	"context"
	"fmt"
	"log"

	dg "github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type (
	Repository struct {
		mongo *mongo.Client
	}

	RelationDAO struct {
		From string `bson:"from"`
		To string `bson:"to"`
		Count uint64 `bson:"count"`
	}
)

func (r *Repository) getRelationshipsCollection() *mongo.Collection {
	return r.mongo.Database("clout").Collection("relationships")
}

func NewRepository() *Repository {
	clientOptions := options.Client().ApplyURI("mongodb://cloutbot:sw4g_cloutbot@mongo:27017")
	ctx := context.Background()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	repository := &Repository{client}
	return repository
}

func bulkWriteRequest(from *dg.User, to []*dg.User) []mongo.WriteModel {
	request := make([]mongo.WriteModel, 0)
	for _, i := range to {
		filter := &bson.M{
			"from": from.ID,
			"to": i.ID,
		}
		update := &bson.M{
			"$setOnInsert": bson.M{ 
				"from": from.ID,
				"to": i.ID,
			},
			"$inc": bson.M{ "count": 1 },
		}
		model := mongo.NewUpdateOneModel().SetUpsert(true).SetFilter(filter).SetUpdate(update)
		request = append(request, model)
	}
	return request
}

func (r *Repository) Upsert(from *dg.User, to []*dg.User) {
	relationships := r.getRelationshipsCollection()
	request := bulkWriteRequest(from, to)
	result, err := relationships.BulkWrite(context.TODO(), request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func (r *Repository) GetMany(user *dg.User) []RelationDAO {
	relationships := r.getRelationshipsCollection()
	filter := &bson.M{
		"to": user.ID,
	}
	cursor, err := relationships.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
		return nil
	} else {
		fmt.Println(cursor)
	}
	var relations []RelationDAO
	if err = cursor.All(context.TODO(), &relations); err != nil {
		fmt.Println(err)
		return nil
	}
	return relations
}