package findface

import (
	"context"
	"strings"
)

type MetaService service

type MetaData struct {
	Count int    `json:"count"`
	Face  *Face  `json:"face"`
	Meta  string `json:"meta"`
}

type MetaListResponse struct {
	FindFaceResponse
	Error   *FindFaceError
	Results []*MetaData `json:"results"`
}

func (s *MetaService) List(ctx context.Context, galleryName string) (*MetaListResponse, error) {
	var path = "/meta"
	if galleryName != "" {
		path = strings.Join([]string{path, "gallery", galleryName}, "/")
	}

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	result := MetaListResponse{}

	err = s.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
