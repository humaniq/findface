package findface

import (
	"context"
	"path"
)

type GalleriesDeleteResponse struct {
	FindFaceResponse
}

// Delete function deletes the gallery and removed all the faces from it.
func (s *GalleriesService) Delete(ctx context.Context, name string) (*GalleriesDeleteResponse, error) {
	if err := s.validateName(name); err != nil {
		return nil, err
	}

	deletePath := path.Join("/galleries", name)

	req, err := s.client.NewRequest("DELETE", deletePath, nil)
	if err != nil {
		return nil, err
	}

	var result = &GalleriesDeleteResponse{}
	resp, err := s.client.Do(ctx, req, result)
	result.Response = resp
	return result, err
}