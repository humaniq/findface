package findface

import (
	"context"
	"path"
)

type GalleriesDeleteResponse struct {
	FindFaceResponse
	Error *FindFaceError
}

// Delete function deletes the gallery and removed all the faces from it.
func (s *GalleriesService) Delete(ctx context.Context, name string) (*GalleriesDeleteResponse, error) {
	if err := s.validateName(name); err != nil {
		return nil, err
	}

	deletePath := path.Join("galleries", name)

	req, err := s.client.NewRequest("DELETE", deletePath, nil)
	if err != nil {
		return nil, err
	}

	result := GalleriesDeleteResponse{}

	err = s.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
