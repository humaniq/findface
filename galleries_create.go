package findface

import (
	"context"
	"fmt"
	"regexp"
)

type GalleriesCreateResponse struct {
	FindFaceResponse
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

	result := GalleriesCreateResponse{}

	err = s.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
