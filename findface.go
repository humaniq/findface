package findface

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	libraryVersion = "0.1.0"
	defaultBaseURL = "https://api.findface.pro/v1/"
	userAgent      = "go-findface/" + libraryVersion

	headerTokenAuth = "Authorization"
)

type service struct {
	client *Client
}

type Client struct {
	client *http.Client // HTTP client used to communicate with the API.

	// Base URL for API requests. BaseURL should
	// always be specified with a trailing slash.
	BaseURL *url.URL

	// User agent used when communicating with the FindFace API.
	UserAgent string

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Services used for talking to different parts of the FindFace
	Face      *FacesService
	Meta      *MetaService
	Galleries *GalleriesService
}

// NewClient returns a new FindFace API client with Authentication header.
// If a nil httpClient is provided, http.Client with TokenAuthTransport will be used.
func NewClient(token string, httpClient *http.Client, endpoint *url.URL) *Client {
	if httpClient == nil {
		tp := &TokenAuthTransport{Token: token}
		httpClient = tp.Client()
	}

	c := &Client{
		client:    httpClient,
		UserAgent: userAgent,
	}

	if endpoint != nil {
		c.BaseURL = endpoint
	} else {
		c.BaseURL, _ = url.Parse(defaultBaseURL)
	}

	c.common.client = c
	c.Face = (*FacesService)(&c.common)
	c.Meta = (*MetaService)(&c.common)
	c.Galleries = (*GalleriesService)(&c.common)

	return c
}

// NewRequest creates an API request. A relative URL can be provided in urlPath,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash.
func (c *Client) NewRequest(method, urlPath string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlPath)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	var b []byte
	if body != nil {
		bb, mErr := json.Marshal(body)
		if mErr != nil {
			return nil, mErr
		}
		b = bb
	}

	req, err := http.NewRequest(method, u.String(), bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	return req, nil
}

// Do sends an API request and set response into `result`.
// The provided ctx must be non-nil. If it is canceled or times out,
// ctx.Err() will be returned.
func (c *Client) Do(ctx context.Context, req *http.Request, result responser) error {
	response, body, err := c.makeRequest(ctx, req)
	if err != nil {
		return err
	}

	var findFaceError *FindFaceError
	switch response.StatusCode {
	case 200, 201, 204:
		if len(body) == 0 {
			break
		}

		if err := json.Unmarshal(body, result); err != nil {
			return err
		}
	case 400, 500:
		if err := json.Unmarshal(body, &findFaceError); err != nil {
			return err
		}
	default:
		return fmt.Errorf("FindFace returned an unhandled status: %s, body: %s", response.Status, string(body))
	}

	result.
		setResponse(response).
		setRawResponseBody(body).
		setError(findFaceError)

	return nil
}

func (c *Client) makeRequest(ctx context.Context, req *http.Request) (*http.Response, []byte, error) {
	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)

	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, nil, ctx.Err()
		default:
			// Do nothing
		}

		return nil, nil, err
	}

	b, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return nil, nil, readErr
	}

	defer func() {
		// Drain up to 512 bytes and close the body to let the Transport reuse the connection
		io.CopyN(ioutil.Discard, resp.Body, 512)
		resp.Body.Close()
	}()

	return resp, b, err
}

// TokenAuthTransport is an http.RoundTripper that authenticates all requests
// using token-based HTTP Authentication with the provided token.
type TokenAuthTransport struct {
	Token string // FindFace authentication token

	// Transport is the underlying HTTP transport to use when making requests.
	// It will default to http.DefaultTransport if nil.
	Transport http.RoundTripper
}

// RoundTrip implements the RoundTripper interface.
func (t *TokenAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req = cloneRequest(req) // per RoundTrip contract
	if t.Token != "" {
		req.Header.Set(headerTokenAuth, fmt.Sprintf("Token %s", t.Token))
	}
	return t.transport().RoundTrip(req)
}

// Client returns an *http.Client that makes requests that are authenticated
// using token-based HTTP Authentication.
func (t *TokenAuthTransport) Client() *http.Client {
	return &http.Client{Transport: t}
}

func (t *TokenAuthTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}

// cloneRequest returns a clone of the provided *http.Request. The clone is a
// shallow copy of the struct and its Header map.
func cloneRequest(r *http.Request) *http.Request {
	// shallow copy of the struct
	r2 := new(http.Request)
	*r2 = *r
	// deep copy of the Header
	r2.Header = make(http.Header, len(r.Header))
	for k, s := range r.Header {
		r2.Header[k] = append([]string(nil), s...)
	}
	return r2
}

// FindFaceError represents the error response object that is returned when
// making a request to findface.
type FindFaceError struct {
	Code  string `json:"code"`
	Faces []struct {
		X1 int `json:"x1"`
		X2 int `json:"x2"`
		Y1 int `json:"y1"`
		Y2 int `json:"y2"`
	} `json:"faces"`
	Param  string `json:"param"`
	Reason string `json:"reason"`
}
