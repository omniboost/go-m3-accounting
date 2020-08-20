package m3

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-m3-accounting/utils"
)

func (c *Client) NewPostActivitiesRequest() PostActivitiesRequest {
	r := PostActivitiesRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewPostActivitiesQueryParams()
	r.pathParams = r.NewPostActivitiesPathParams()
	r.requestBody = r.NewPostActivitiesRequestBody()
	return r
}

type PostActivitiesRequest struct {
	client      *Client
	queryParams *PostActivitiesQueryParams
	pathParams  *PostActivitiesPathParams
	method      string
	headers     http.Header
	requestBody PostActivitiesRequestBody
}

func (r PostActivitiesRequest) NewPostActivitiesQueryParams() *PostActivitiesQueryParams {
	return &PostActivitiesQueryParams{}
}

type PostActivitiesQueryParams struct {
}

func (p PostActivitiesQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *PostActivitiesRequest) QueryParams() *PostActivitiesQueryParams {
	return r.queryParams
}

func (r PostActivitiesRequest) NewPostActivitiesPathParams() *PostActivitiesPathParams {
	return &PostActivitiesPathParams{}
}

type PostActivitiesPathParams struct {
}

func (p *PostActivitiesPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PostActivitiesRequest) PathParams() *PostActivitiesPathParams {
	return r.pathParams
}

func (r *PostActivitiesRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostActivitiesRequest) Method() string {
	return r.method
}

func (r PostActivitiesRequest) NewPostActivitiesRequestBody() PostActivitiesRequestBody {
	return PostActivitiesRequestBody{}
}

type PostActivitiesRequestBody Activity

func (r *PostActivitiesRequest) RequestBody() *PostActivitiesRequestBody {
	return &r.requestBody
}

func (r *PostActivitiesRequest) SetRequestBody(body PostActivitiesRequestBody) {
	r.requestBody = body
}

func (r *PostActivitiesRequest) NewResponseBody() *PostActivitiesResponseBody {
	return &PostActivitiesResponseBody{}
}

type PostActivitiesResponseBody struct{}

func (r *PostActivitiesRequest) URL() url.URL {
	return r.client.GetEndpointURL("daily-report/activities", r.PathParams())
}

func (r *PostActivitiesRequest) Do() (PostActivitiesResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r.Method(), r.URL(), r.RequestBody())
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
