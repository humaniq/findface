package findface

import (
	"context"
	"path"
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
	Age float64 `json:"age"`

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
	GalleryName string

	// Pagination parameters
	MinID, MaxID int

	// Metadata string to filter faces by.
	Meta string
}

func (o *FaceListOptions) Path() (string, error) {
	var urlPath string
	switch {
	case o.GalleryName != "" && o.Meta != "":
		urlPath = path.Join("/faces/gallery/", o.GalleryName, "meta", o.Meta)
	case o.Meta != "":
		urlPath = path.Join("/faces/meta/", o.Meta)
	case o.GalleryName != "":
		urlPath = path.Join("/faces/gallery/", o.GalleryName)
	default:
		urlPath = "/faces"
	}

	return urlPath, nil
}

type FaceListResult struct {
	FindFaceResponse
	Error    *FindFaceError
	Faces    []*Face `json:"results"`
	NextPage string  `json:"next_page"`
}

// Returns the list of all faces stored in gallery or account.
func (s *FacesService) List(ctx context.Context, opt *FaceListOptions) (*FaceListResult, error) {
	path, err := opt.Path()
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	result := &FaceListResult{}
	resp, rawResp, err := s.client.Do(ctx, req)
	result.Response = resp
	result.RawResponseBody = rawResp
	return result, err
}
