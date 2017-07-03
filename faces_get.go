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
	Error *FindFaceError
	Faces []*Face
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

	var result []*Face
	response := &FaceGetResponse{}

	resp, rawResp, err := s.client.Do(ctx, req)
	fErr := &FindFaceError{}
	switch resp.StatusCode {
	case 200:
		unErr := json.Unmarshal(rawResp, &result)
		if unErr != nil {
			return nil, unErr
		}
	case 400:
		unErr := json.Unmarshal(rawResp, &fErr)
		if unErr != nil {
			return nil, unErr
		}
	default:
		err = fmt.Errorf("FindFace returned an unhandled status: %s, body: %s", resp.Status, string(rawResp))
	}
	response.Response = resp
	response.Faces = result
	response.Error = fErr
	response.RawResponseBody = rawResp
	return response, err
}
