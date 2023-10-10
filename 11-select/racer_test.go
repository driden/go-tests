package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("Compares speeds of the servers, returning the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowUrl, fastURL)

		if err != nil {
			t.Fatalf("Did not expect to get an error but got %q", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Returns an error if no server responds within 10 seconds", func(t *testing.T) {
		serverA := makeDelayedServer(25 * time.Millisecond)
		defer serverA.Close()

		_, err := ConfigurableRacer(serverA.URL, serverA.URL, 10*time.Millisecond)

		if err == nil {
			t.Error("Expected an error and didn't get one.")
		}
	})

}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				time.Sleep(delay)
				w.WriteHeader(http.StatusOK)
			}))
}
