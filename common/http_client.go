package common

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/teraflops-droid/rbh-interview-service/common/logger"
	"github.com/teraflops-droid/rbh-interview-service/exception"
	"io"
	"math/rand"
	"net"
	"net/http"
	"reflect"
	"time"
)

type HttpHeader struct {
	Key   string
	Value string
}

type ClientComponent[T any, E any] struct {
	HttpMethod     string
	UrlApi         string
	ConnectTimeout uint32
	ActiveTimeout  uint32
	Headers        []HttpHeader
	RequestBody    *T
	ResponseBody   *E
}

func (c *ClientComponent[T, E]) Execute(ctx context.Context) error {

	client := &http.Client{
		Timeout: time.Duration(rand.Int31n(int32(c.ActiveTimeout))) * time.Millisecond,
		Transport: &http.Transport{
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
			TLSHandshakeTimeout: 5 * time.Second,
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return net.DialTimeout(network, addr, time.Duration(rand.Int31n(int32(c.ConnectTimeout)))*time.Millisecond)
			},
		},
	}

	var request *http.Request
	var response *http.Response
	var err error = nil

	//set request body
	if reflect.ValueOf(c.RequestBody).IsZero() || c.RequestBody == nil {
		request, err = http.NewRequest(c.HttpMethod, c.UrlApi, nil)
	} else {
		requestBody, err := json.Marshal(c.RequestBody)
		exception.PanicLogging(err)

		//logging request body
		logger.InitLogger("local")

		requestBodyByte := bytes.NewBuffer(requestBody)

		request, err = http.NewRequestWithContext(ctx, c.HttpMethod, c.UrlApi, requestBodyByte)
		exception.PanicLogging(err)
	}

	//set header
	request.Header.Set("Content-Type", "application/json")
	for _, header := range c.Headers {
		request.Header.Set(header.Key, header.Value)
	}

	//logging before
	logger.Info(ctx, fmt.Sprintf("Request Url: %s", c.UrlApi))
	logger.Info(ctx, fmt.Sprintf("Request Method %s", c.HttpMethod))
	logger.Info(ctx, fmt.Sprintf("Request Header %s", request.Header))

	//time
	start := time.Now()

	response, err = client.Do(request)
	//error handling for http client
	if err != nil {
		return err
	}

	//time
	elapsed := time.Now().Sub(start)

	responseBody, err := io.ReadAll(response.Body)
	exception.PanicLogging(err)

	err = json.Unmarshal(responseBody, &c.ResponseBody)
	exception.PanicLogging(err)

	logger.Info(ctx, fmt.Sprintf("Received response for %s in %d ms", c.UrlApi, elapsed.Milliseconds()))
	logger.Info(ctx, fmt.Sprintf("Response Header: %v", response.Header))
	logger.Info(ctx, fmt.Sprintf("Response HTTP Status: %d", response.StatusCode))
	logger.Info(ctx, fmt.Sprintf("Response HTTP Version: %s", response.Proto))
	logger.Info(ctx, fmt.Sprintf("Response Body: %s", string(responseBody)))

	return nil
}
