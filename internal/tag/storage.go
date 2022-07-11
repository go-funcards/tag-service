package tag

import "context"

type Storage interface {
	Save(ctx context.Context, model Tag) error
	Delete(ctx context.Context, id string) error
	Find(ctx context.Context, filter Filter, index uint64, size uint32) ([]Tag, error)
	Count(ctx context.Context, filter Filter) (uint64, error)
}
