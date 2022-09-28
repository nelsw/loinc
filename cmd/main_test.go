package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
	"testing"
)

var req = events.APIGatewayV2HTTPRequest{
	RequestContext: events.APIGatewayV2HTTPRequestContext{
		HTTP: events.APIGatewayV2HTTPRequestContextHTTPDescription{
			Method: http.MethodGet,
		},
	},
}

func Test_OK(t *testing.T) {

	r := req

	r.QueryStringParameters = map[string]string{"id": "83805-2"}
	expectOK(t, r)

	r.RequestContext.HTTP.Method = http.MethodHead
	expectOK(t, r)

	r.RequestContext.HTTP.Method = http.MethodConnect
	expectOK(t, r)

	r.RequestContext.HTTP.Method = http.MethodOptions
	expectOK(t, r)

	r.RequestContext.HTTP.Method = http.MethodTrace
	expectOK(t, r)
}

func Test_NotImplemented(t *testing.T) {
	req.RequestContext.HTTP.Method = http.MethodPost
	expectStatusNotImplemented(t, req)

	req.RequestContext.HTTP.Method = http.MethodPut
	expectStatusNotImplemented(t, req)

	req.RequestContext.HTTP.Method = http.MethodDelete
	expectStatusNotImplemented(t, req)

	req.RequestContext.HTTP.Method = http.MethodPatch
	expectStatusNotImplemented(t, req)
}

func Test_BadRequest(t *testing.T) {

	r := req

	r.RequestContext.HTTP.Method = http.MethodGet
	r.QueryStringParameters = map[string]string{}
	expectBadRequest(t, r)

	r.QueryStringParameters = map[string]string{"id": ""}
	expectBadRequest(t, r)
}

func expectOK(t *testing.T, r events.APIGatewayV2HTTPRequest) {
	execute(t, r, http.StatusOK)
}

func expectBadRequest(t *testing.T, r events.APIGatewayV2HTTPRequest) {
	execute(t, r, http.StatusBadRequest)
}

func expectStatusNotImplemented(t *testing.T, r events.APIGatewayV2HTTPRequest) {
	execute(t, r, http.StatusNotImplemented)
}

func execute(t *testing.T, req events.APIGatewayV2HTTPRequest, status int) {
	if res, _ := Handler(context.TODO(), req); res.StatusCode != status {
		t.Error(res.StatusCode, res.Body)
	}
}
