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
	var fErr *FindFaceError
	resp, rawResp, err := s.client.Do(ctx, req)
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
	}
	result.Response = resp
	result.RawResponseBody = rawResp
	return result, err
}
