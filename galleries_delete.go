package findface

import (
	"context"
	"encoding/json"
	"fmt"
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

	var result = GalleriesDeleteResponse{}
	resp, rawResp, dErr := s.client.Do(ctx, req)
	var fErr *FindFaceError
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
	result.Response = resp
	result.RawResponseBody = rawResp
	result.Error = fErr
	return &result, dErr
}
