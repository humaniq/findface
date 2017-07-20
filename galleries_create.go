package findface

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
)

type GalleriesCreateResponse struct {
	FindFaceResponse
	Error *FindFaceError
}

// GalleriesCreateOptions valdation
func (s *GalleriesService) validateName(name string) error {
	if name == "" {
		return fmt.Errorf("Name should be present")
	}

	match, err := regexp.MatchString(galleryNameRegexp, name)
	if err != nil {
		return err
	}

	if !match {
		return fmt.Errorf("Name \"%s\" invalid", name)
	}

	return nil
}

// Creates new gallery with the specified name.
// Gallery name may contain English letters, numbers, underscore and minus sign.
func (s *GalleriesService) Create(ctx context.Context, name string) (*GalleriesCreateResponse, error) {
	if err := s.validateName(name); err != nil {
		return nil, err
	}

	opt := struct {
		Name string `json:"name"`
	}{Name: name}

	req, err := s.client.NewRequest("POST", "galleries", &opt)
	if err != nil {
		return nil, err
	}

	var result = GalleriesCreateResponse{}
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
