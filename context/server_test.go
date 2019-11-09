package context

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	t        *testing.T
	response string
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				s.t.Log("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}

		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

type SpyResponseWriter bool

func (s *SpyResponseWriter) Header() http.Header {
	*s = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	*s = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statuscode int) {
	*s = true
}

func TestHandler(t *testing.T) {
	const data = "hello, world"

	t.Run("store returns the data", func(t *testing.T) {
		store := &SpyStore{t: t, response: data}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got %q, want %q", response.Body.String(), data)
		}
	})

	t.Run("store fetch is cancelled after timeout", func(t *testing.T) {
		store := &SpyStore{t: t, response: data}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := new(SpyResponseWriter)

		svr.ServeHTTP(response, request)

		if *response {
			t.Error("response was not cancelled")
		}
	})
}
