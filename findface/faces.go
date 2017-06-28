package findface

import (
	"context"
	"strings"
)

type FacesService service

type BoundingBox struct {
	X1 int `json:"x1"`
	Y1 int `json:"y1"`
	X2 int `json:"x2"`
	Y2 int `json:"y2"`
}

type Face struct {
	BoundingBox

	ID int `json:"id"`

	// Metadata string that you can use to store any information associated with the face.
	Meta string `json:"meta"`

	// Age
	Age int `json:"age"`

	// List of emotions
	Emotions []string `json:"emotions"`
	// Gender
	Gender string `json:"gender"`

	// Url of the photo
	Photo     string `json:"photo"`
	PhotoHash string `json:"photo_hash"`
	Thumbnail string `json:"thumbnail"`
	Timestamp string `json:"timestamp"`

	// List of gallery names to add face(s) to.
	Galleries []string `json:"galleries"`
}

type FaceListOptions struct {
	// Gallery name
	GalleryName string `json:"gallery"`

	// Pagination parameters
	MinID, MaxID int
}

type FaceListResult struct {
	FindFaceResponse
	Faces    []*Face `json:"results"`
	NextPage string  `json:"next_page"`
}

// Returns the list of all faces stored in gallery or account.
func (s *FacesService) List(ctx context.Context, opt *FaceListOptions) (*FaceListResult, error) {
	var path = "/faces"
	if opt.GalleryName != "" {
		path = strings.Join([]string{path, "gallery", opt.GalleryName}, "/")
	}

	req, err := s.client.NewRequest("GET", path, opt)
	if err != nil {
		return nil, err
	}

	var result *FaceListResult
	resp, err := s.client.Do(ctx, req, &result)
	result.Response = resp
	return result, err
}
