package findface

import (
	"context"
	"fmt"
	"path"
	"strconv"
)

type FacesDeleteResponse struct {
	FindFaceResponse
}

func (s *FacesService) Delete(ctx context.Context, faceID int) (*FacesDeleteResponse, error) {
	if faceID <= 0 {
		return nil, fmt.Errorf("faceID shuld be > 0, but was: %d", faceID)
	}

	urlPath := path.Join("/faces/id", strconv.Itoa(faceID))

	req, err := s.client.NewRequest("DELETE", urlPath, nil)
	if err != nil {
		return nil, err
	}

	var result = &FacesDeleteResponse{}
	resp, err := s.client.Do(ctx, req, result)
	result.Response = resp
	return result, err
}
