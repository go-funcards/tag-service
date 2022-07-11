package tag

import (
	"context"
	"github.com/go-funcards/slice"
	"github.com/go-funcards/tag-service/proto/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ v1.TagServer = (*tagService)(nil)

type tagService struct {
	v1.UnimplementedTagServer
	storage Storage
}

func NewTagService(storage Storage) *tagService {
	return &tagService{storage: storage}
}

func (s *tagService) CreateTag(ctx context.Context, in *v1.CreateTagRequest) (*emptypb.Empty, error) {
	if err := s.storage.Save(ctx, CreateTag(in)); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *tagService) UpdateTag(ctx context.Context, in *v1.UpdateTagRequest) (*emptypb.Empty, error) {
	if err := s.storage.Save(ctx, UpdateTag(in)); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *tagService) DeleteTag(ctx context.Context, in *v1.DeleteTagRequest) (*emptypb.Empty, error) {
	if err := s.storage.Delete(ctx, in.GetTagId()); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *tagService) GetTags(ctx context.Context, in *v1.TagsRequest) (*v1.TagsResponse, error) {
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
			return item.toResponse()
		}),
		Total: total,
	}, nil
}
