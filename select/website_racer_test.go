package selectpkg // select is a keyword

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("returns the url to the fastest server", func(t *testing.T) {
		slowServer := makeSleepyServer(20 * time.Millisecond)
		defer slowServer.Close()

		fastServer := makeSleepyServer(0 * time.Millisecond)
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Errorf("got an unexpected error: %q", err)
		}

		if got != want {
			t.Errorf("Racer(%q, %q) = %q, want %q", slowURL, fastURL, got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		server := makeSleepyServer(10 * time.Microsecond)
		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 1*time.Microsecond)

		if err == nil {
			t.Error("expected an error but did not get one")
		}
	})
}

func makeSleepyServer(d time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(d)
		w.WriteHeader(http.StatusOK)
	}))
}
