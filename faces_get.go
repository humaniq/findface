package findface

import (
	"context"
	"fmt"
	"path"
	"strconv"
)

type FaceGetResponse struct {
	FindFaceResponse
	Faces []*Face
}

func (s *FacesService) Get(ctx context.Context, faceID int) (*FaceGetResponse, error) {
	if faceID <= 0 {
		return nil, fmt.Errorf("faceID shuld be > 0, but was: %d", faceID)
	}

	urlPath := path.Join("/faces/id", strconv.Itoa(faceID))

	req, err := s.client.NewRequest("GET", urlPath, nil)
	if err != nil {
		return nil, err
	}

	var result []*Face
	response := &FaceGetResponse{}

	resp, err := s.client.Do(ctx, req, &result)
	response.Response = resp
	if err != nil {
		return response, err
	}
	response.Faces = result
	return response, err
}
