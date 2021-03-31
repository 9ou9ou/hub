// Code generated by goa v3.3.1, DO NOT EDIT.
//
// user client HTTP transport
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the user service endpoint HTTP clients.
type Client struct {
	// RefreshAccessToken Doer is the HTTP client used to make requests to the
	// RefreshAccessToken endpoint.
	RefreshAccessTokenDoer goahttp.Doer

	// NewRefreshToken Doer is the HTTP client used to make requests to the
	// NewRefreshToken endpoint.
	NewRefreshTokenDoer goahttp.Doer

	// CORS Doer is the HTTP client used to make requests to the  endpoint.
	CORSDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the user service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		RefreshAccessTokenDoer: doer,
		NewRefreshTokenDoer:    doer,
		CORSDoer:               doer,
		RestoreResponseBody:    restoreBody,
		scheme:                 scheme,
		host:                   host,
		decoder:                dec,
		encoder:                enc,
	}
}

// RefreshAccessToken returns an endpoint that makes HTTP requests to the user
// service RefreshAccessToken server.
func (c *Client) RefreshAccessToken() goa.Endpoint {
	var (
		encodeRequest  = EncodeRefreshAccessTokenRequest(c.encoder)
		decodeResponse = DecodeRefreshAccessTokenResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildRefreshAccessTokenRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.RefreshAccessTokenDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("user", "RefreshAccessToken", err)
		}
		return decodeResponse(resp)
	}
}

// NewRefreshToken returns an endpoint that makes HTTP requests to the user
// service NewRefreshToken server.
func (c *Client) NewRefreshToken() goa.Endpoint {
	var (
		encodeRequest  = EncodeNewRefreshTokenRequest(c.encoder)
		decodeResponse = DecodeNewRefreshTokenResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildNewRefreshTokenRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.NewRefreshTokenDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("user", "NewRefreshToken", err)
		}
		return decodeResponse(resp)
	}
}
