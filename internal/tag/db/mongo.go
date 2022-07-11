package db

import (
	"context"
	"github.com/go-funcards/mongodb"
	"github.com/go-funcards/tag-service/internal/tag"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"time"
)

var _ tag.Storage = (*storage)(nil)

const (
	timeout    = 5 * time.Second
	collection = "tags"
)

type storage struct {
	c mongodb.Collection[tag.Tag]
}

func NewStorage(ctx context.Context, db *mongo.Database, logger *zap.Logger) (*storage, error) {
	s := &storage{c: mongodb.Collection[tag.Tag]{
		Inner: db.Collection(collection),
		Log:   logger,
	}}

	if err := s.indexes(ctx); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *storage) indexes(ctx context.Context) error {
	name, err := s.c.Inner.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{
			{"owner_id", 1},
			{"board_id", 1},
			{"created_at", 1},
		},
	})
	if err == nil {
		s.c.Log.Info("indexes created", zap.String("collection", collection), zap.String("name", name))
	}

	return err
}

func (s *storage) Save(ctx context.Context, model tag.Tag) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	data, err := s.c.ToM(model)
	if err != nil {
		return err
	}

	delete(data, "_id")
	delete(data, "owner_id")
	delete(data, "created_at")

	return s.c.UpdateOne(
		ctx,
		bson.M{"_id": model.TagID},
		bson.M{
			"$set": data,
			"$setOnInsert": bson.M{
				"owner_id":   model.OwnerID,
				"created_at": model.CreatedAt,
			},
		},
		options.Update().SetUpsert(true),
	)
}

func (s *storage) Delete(ctx context.Context, id string) error {
	return s.c.DeleteOne(ctx, bson.M{"_id": id})
}

func (s *storage) Find(ctx context.Context, filter tag.Filter, index uint64, size uint32) ([]tag.Tag, error) {
	return s.c.Find(ctx, s.filter(filter), s.c.FindOptions(index, size).SetSort(bson.D{{"created_at", -1}}))
}

func (s *storage) Count(ctx context.Context, filter tag.Filter) (uint64, error) {
	return s.c.CountDocuments(ctx, s.filter(filter))
}

func (s *storage) filter(filter tag.Filter) bson.M {
	f := make(bson.M)
	if len(filter.TagIDs) > 0 {
		f["_id"] = bson.M{"$in": filter.TagIDs}
	}
	if len(filter.OwnerIDs) > 0 {
		f["owner_id"] = bson.M{"$in": filter.OwnerIDs}
	}
	if len(filter.BoardIDs) > 0 {
		f["board_id"] = bson.M{"$in": filter.BoardIDs}
	}
	return f
}
