package findface

import "net/http"

type FindFaceResponse struct {
	RawResponseBody []byte
	Response        *http.Response
	Error           *FindFaceError
}

type responser interface {
	setResponse(*http.Response) responser
	setRawResponseBody([]byte) responser
	setError(*FindFaceError) responser
}

func (r *FindFaceResponse) setResponse(response *http.Response) responser {
	r.Response = response
	return r
}

func (r *FindFaceResponse) setRawResponseBody(rawResponseBody []byte) responser {
	r.RawResponseBody = rawResponseBody
	return r
}

func (r *FindFaceResponse) setError(err *FindFaceError) responser {
	r.Error = err
	return r
}
