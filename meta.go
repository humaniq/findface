package findface

import (
	"context"
	"encoding/json"
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

	var result *MetaListResponse
	resp, rawResp, err := s.client.Do(ctx, req)
	result.Response = resp
	result.RawResponseBody = rawResp
	unErr := json.Unmarshal(rawResp, &result.Results)
	if unErr != nil {
		return nil, unErr
	}
	return result, err
}
