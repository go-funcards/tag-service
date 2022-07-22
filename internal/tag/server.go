package tag

import (
	"context"
	"github.com/go-funcards/slice"
	"github.com/go-funcards/tag-service/proto/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ v1.TagServer = (*server)(nil)

type server struct {
	v1.UnimplementedTagServer
	storage Storage
}

func NewTagServer(storage Storage) *server {
	return &server{storage: storage}
}

func (s *server) CreateTag(ctx context.Context, in *v1.CreateTagRequest) (*emptypb.Empty, error) {
	err := s.storage.Save(ctx, CreateTag(in))

	return s.empty(err)
}

func (s *server) UpdateTag(ctx context.Context, in *v1.UpdateTagRequest) (*emptypb.Empty, error) {
	err := s.storage.Save(ctx, UpdateTag(in))

	return s.empty(err)
}

func (s *server) DeleteTag(ctx context.Context, in *v1.DeleteTagRequest) (*emptypb.Empty, error) {
	err := s.storage.Delete(ctx, in.GetTagId())

	return s.empty(err)
}

func (s *server) GetTags(ctx context.Context, in *v1.TagsRequest) (*v1.TagsResponse, error) {
	filter := CreateFilter(in)

	data, err := s.storage.Find(ctx, filter, in.GetPageIndex(), in.GetPageSize())
	if err != nil {
		return nil, err
	}

	total := uint64(len(data))
	if len(in.GetTagIds()) == 0 && uint64(in.GetPageSize()) == total {
		if total, err = s.storage.Count(ctx, filter); err != nil {
			return nil, err
		}
	}

	return &v1.TagsResponse{
		Tags: slice.Map(data, func(item Tag) *v1.TagsResponse_Tag {
			return item.toProto()
		}),
		Total: total,
	}, nil
}

func (s *server) empty(err error) (*emptypb.Empty, error) {
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
