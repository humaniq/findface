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
	Galleries []string `json:"results"`
}

// Returns list of all galleries.
func (s *GalleriesService) List(ctx context.Context) (*GalleriesListResponse, error) {
	req, err := s.client.NewRequest("GET", "/galleries", nil)
	if err != nil {
		return nil, err
	}

	var result *GalleriesListResponse
	resp, err := s.client.Do(ctx, req, &result)
	result.Response = resp
	return result, err
}
