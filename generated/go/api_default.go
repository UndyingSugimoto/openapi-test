/*
 * test-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// A DefaultApiController binds http requests to an api service and writes the service results to the http response
type DefaultApiController struct {
	service DefaultApiServicer
}

// NewDefaultApiController creates a default api controller
func NewDefaultApiController(s DefaultApiServicer) Router {
	return &DefaultApiController{ service: s }
}

// Routes returns all of the api route for the DefaultApiController
func (c *DefaultApiController) Routes() Routes {
	return Routes{ 
		{
			"GetNumber1",
			strings.ToUpper("Get"),
			"/number1",
			c.GetNumber1,
		},
		{
			"GetTest",
			strings.ToUpper("Get"),
			"/test",
			c.GetTest,
		},
		{
			"GetTest2",
			strings.ToUpper("Get"),
			"/test2",
			c.GetTest2,
		},
		{
			"PostHoge",
			strings.ToUpper("Post"),
			"/hoge",
			c.PostHoge,
		},
	}
}

// GetNumber1 - test
func (c *DefaultApiController) GetNumber1(w http.ResponseWriter, r *http.Request) { 
	number1Req := &Number1Req{}
	if err := json.NewDecoder(r.Body).Decode(&number1Req); err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.GetNumber1(*number1Req)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}

// GetTest - Your GET endpoint
func (c *DefaultApiController) GetTest(w http.ResponseWriter, r *http.Request) { 
	body := &interface{}{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.GetTest(*body)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}

// GetTest2 - Your GET endpoint
func (c *DefaultApiController) GetTest2(w http.ResponseWriter, r *http.Request) { 
	result, err := c.service.GetTest2()
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}

// PostHoge - hogetest
func (c *DefaultApiController) PostHoge(w http.ResponseWriter, r *http.Request) { 
	testmodel := &Testmodel{}
	if err := json.NewDecoder(r.Body).Decode(&testmodel); err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.PostHoge(*testmodel)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}
