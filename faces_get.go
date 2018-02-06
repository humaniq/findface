package findface

import (
	"context"
	"encoding/json"
	"fmt"
	"path"
	"strconv"
)

type FaceGetResponse struct {
	FindFaceResponse
	Faces []*Face
}

func (r *FaceGetResponse) UnmarshalJSON(data []byte) error {
	faces := []*Face{}
	if err := json.Unmarshal(data, &faces); err != nil {
		return err
	}
	r.Faces = faces
	return nil
}

func (s *FacesService) Get(ctx context.Context, faceID int) (*FaceGetResponse, error) {
	if faceID <= 0 {
		return nil, fmt.Errorf("faceID should be > 0, but was: %d", faceID)
	}

	urlPath := path.Join("/faces/id", strconv.Itoa(faceID))

	req, err := s.client.NewRequest("GET", urlPath, nil)
	if err != nil {
		return nil, err
	}

	result := FaceGetResponse{}

	err = s.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
