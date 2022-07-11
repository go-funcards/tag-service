package tag

import (
	"github.com/go-funcards/tag-service/proto/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Tag struct {
	TagID     string    `json:"tag_id" bson:"_id,omitempty"`
	OwnerID   string    `json:"owner_id" bson:"owner_id,omitempty"`
	BoardID   string    `json:"board_id" bson:"board_id,omitempty"`
	Name      string    `json:"name" bson:"name,omitempty"`
	Color     string    `json:"color" bson:"color,omitempty"`
	CreatedAt time.Time `json:"created_at" bson:"created_at,omitempty"`
}

type Filter struct {
	TagIDs   []string `json:"tag_ids,omitempty"`
	OwnerIDs []string `json:"owner_ids,omitempty"`
	BoardIDs []string `json:"board_ids,omitempty"`
}

func (t Tag) toResponse() *v1.TagsResponse_Tag {
	return &v1.TagsResponse_Tag{
		TagId:     t.TagID,
		OwnerId:   t.OwnerID,
		BoardId:   t.BoardID,
		Name:      t.Name,
		Color:     t.Color,
		CreatedAt: timestamppb.New(t.CreatedAt),
	}
}

func CreateTag(in *v1.CreateTagRequest) Tag {
	return Tag{
		TagID:     in.GetTagId(),
		OwnerID:   in.GetOwnerId(),
		BoardID:   in.GetBoardId(),
		Name:      in.GetName(),
		Color:     in.GetColor(),
		CreatedAt: time.Now().UTC(),
	}
}

func UpdateTag(in *v1.UpdateTagRequest) Tag {
	return Tag{
		TagID:   in.GetTagId(),
		BoardID: in.GetBoardId(),
		Name:    in.GetName(),
		Color:   in.GetColor(),
	}
}

func CreateFilter(in *v1.TagsRequest) Filter {
	return Filter{
		TagIDs:   in.GetTagIds(),
		OwnerIDs: in.GetOwnerIds(),
		BoardIDs: in.GetBoardIds(),
	}
}
