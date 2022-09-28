package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	log "github.com/sirupsen/logrus"
	"loinc-lambda/pkg/store"
	"net/http"
	"os"
	"time"
)

const missingParam = "missing param [id]"

func init() {

	log.SetOutput(os.Stdout)

	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.JSONFormatter{
		DisableTimestamp:  true,
		DisableHTMLEscape: true,
	})

	if os.Getenv("WHEREAMI") == "localhost" {
		log.SetLevel(log.DebugLevel)
		log.SetFormatter(&log.TextFormatter{
			ForceColors:     true,
			FullTimestamp:   true,
			TimestampFormat: time.Stamp,
		})
	}
}

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

	log.WithContext(ctx).
		WithField("method", req.RequestContext.HTTP.Method).
		WithField("params", req.QueryStringParameters).
		Info()

	switch req.RequestContext.HTTP.Method {

	case http.MethodGet:
		if ID, ok := req.QueryStringParameters["id"]; !ok {
			log.WithContext(ctx).
				WithField("code", http.StatusBadRequest).
				WithField("body", missingParam).
				Warn()
			return events.APIGatewayV2HTTPResponse{StatusCode: http.StatusBadRequest, Body: missingParam}, nil
		} else if val, err := store.Get(ID); err != nil {
			log.WithContext(ctx).
				WithField("code", http.StatusBadRequest).
				WithField("body", err.Error()).
				Warn()
			return events.APIGatewayV2HTTPResponse{StatusCode: http.StatusBadRequest, Body: err.Error()}, nil
		} else {
			bytes, _ := json.Marshal(&val)
			log.WithContext(ctx).
				WithField("code", http.StatusOK).
				WithField("body", val).
				Info()
			return events.APIGatewayV2HTTPResponse{StatusCode: http.StatusOK, Body: string(bytes)}, nil
		}

	case http.MethodHead:
		fallthrough
	case http.MethodConnect:
		fallthrough
	case http.MethodOptions:
		fallthrough
	case http.MethodTrace:
		log.WithContext(ctx).
			WithField("code", http.StatusOK).
			Info()
		return events.APIGatewayV2HTTPResponse{StatusCode: http.StatusOK}, nil

	case http.MethodPost:
		fallthrough
	case http.MethodPut:
		fallthrough
	case http.MethodPatch:
		fallthrough
	case http.MethodDelete:
		fallthrough
	default:
		log.WithContext(ctx).
			WithField("code", http.StatusNotImplemented).
			Warn()
		return events.APIGatewayV2HTTPResponse{StatusCode: http.StatusNotImplemented}, nil
	}
}

func main() {
	lambda.Start(Handler)
}
