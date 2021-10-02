package server

import (
	"context"
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestServer_GetKey(t *testing.T) {
	s := NewMockStore()
	s.GetFunc.SetDefaultReturn([]byte("ABCD"), nil)

	c := NewMockCoder()
	c.DecodeStringFunc.SetDefaultReturn([]byte("https://kar.moe"), nil)

	tests := []struct {
		s    *Server
		name string
		key  string
	}{
		{
			name: "basic",
			s: &Server{
				store: s,
				coder: c,
			},
			key: "ABCD",
		},
	}

	r := httptest.NewRecorder()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/"+tt.key, nil)

			ctx := chi.NewRouteContext()
			ctx.URLParams.Add("key", tt.key)

			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

			tt.s.redirectFromURL(r, req)
			if r.Code != http.StatusMovedPermanently {
				t.Errorf("Status code: %d, got %s", r.Code, r.Body.String())
			}
		})
	}
}

func TestServer_SetKey(t *testing.T) {
	g := NewMockGenerator()
	g.GenerateFunc.SetDefaultReturn([]byte("EFGH"), nil)

	c := NewMockCoder()
	c.EncodeToStringFunc.SetDefaultReturn("ABCD")

	tests := []struct {
		s    *Server
		name string
		want string
		url  string
	}{
		{
			name: "basic",
			s: &Server{
				store: NewMockStore(),
				gen:   g,
				coder: c,
			},
			want: "ABCD",
			url:  "https://kar.moe",
		},
	}

	r := httptest.NewRecorder()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(
				"PUT",
				"/key?url="+base64.RawURLEncoding.EncodeToString([]byte(tt.url)),
				nil,
			)
			tt.s.newShortURL(r, req)
			if r.Body.String() != tt.want {
				t.Errorf("test failed, got %s wanted %s", r.Body.String(), tt.want)
			}
		})
	}
}
