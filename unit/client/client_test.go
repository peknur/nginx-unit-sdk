package client

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/peknur/nginx-unit-sdk/unit"
	"github.com/peknur/nginx-unit-sdk/unit/config"
	"github.com/peknur/nginx-unit-sdk/unit/config/route"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientGet(t *testing.T) {
	expected, err := ioutil.ReadFile("testdata/client_get_config.json")
	require.NoError(t, err)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, string(expected))
	}))
	client, err := New(srv.URL, nil)
	require.NoError(t, err)
	cfg := config.Config{}
	err = client.Get(context.Background(), unit.ConfigPath, &cfg)
	assert.NoError(t, err)
	assert.Equal(t, "/var/log/access.log", cfg.AccessLog)
	assert.Equal(t, "https", cfg.Routes["main"][0].Match.Scheme)
}

func TestClientPut(t *testing.T) {
	expected, err := ioutil.ReadFile("testdata/client_put_config.json")
	require.NoError(t, err)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		assert.NoError(t, err)
		assert.JSONEq(t, string(expected), string(body))
		assert.Equal(t, http.MethodPut, r.Method)
		fmt.Fprint(w, string(expected))
	}))
	client, err := New(srv.URL, nil)
	require.NoError(t, err)
	cfg := config.Routes{
		"main": {
			{
				Match: &route.Match{
					URI: "/return/*",
				},
				Action: &route.Action{
					Return:   301,
					Location: "https://www.example.com",
				},
			},
		},
	}
	err = client.Put(context.Background(), unit.RoutesPath, &cfg)
	assert.NoError(t, err)
}
func TestClientPost(t *testing.T) {
	expected, err := ioutil.ReadFile("testdata/client_post_config.json")
	require.NoError(t, err)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		assert.NoError(t, err)
		assert.JSONEq(t, string(expected), string(body))
		assert.Equal(t, http.MethodPost, r.Method)
		fmt.Fprint(w, string(expected))
	}))
	client, err := New(srv.URL, nil)
	require.NoError(t, err)
	cfg := config.Routes{
		"main": {
			{
				Match: &route.Match{
					URI: "/return/*",
				},
				Action: &route.Action{
					Return:   301,
					Location: "https://www.example.com",
				},
			},
		},
	}
	err = client.Post(context.Background(), unit.RoutesPath, &cfg)
	assert.NoError(t, err)
}

func TestClienDelete(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
	}))
	client, err := New(srv.URL, nil)
	require.NoError(t, err)
	err = client.Delete(context.Background(), unit.ConfigPath)
	assert.NoError(t, err)
}
