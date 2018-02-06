package findface

import (
	"context"
)

const (
	galleryNameRegexp = "[a-zA-Z0-9_-]+"
)

type GalleriesService service

type GalleriesListResponse struct {
	FindFaceResponse
	Error     *FindFaceError
	Galleries []string `json:"results"`
}

// Returns list of all galleries.
func (s *GalleriesService) List(ctx context.Context) (*GalleriesListResponse, error) {
	req, err := s.client.NewRequest("GET", "/galleries", nil)
	if err != nil {
		return nil, err
	}

	result := GalleriesListResponse{}

	err = s.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
