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

	var result *MetaListResponse
	resp, err := s.client.Do(ctx, req, &result)
	result.Response = resp
	return result, err
}
