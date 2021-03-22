package select_concurrency

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("test mana yang lebih cepat kembali", func(t *testing.T) {
		slowServer := makeDelayServer(20 * time.Millisecond)
		fastServer := makeDelayServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, _ := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("GOT %q, WANT %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond with 10s", func(t *testing.T) {
		serverA := makeDelayServer(25 * time.Millisecond)
		defer serverA.Close()

		_, err := ConfigurableRacer(serverA.URL, serverA.URL, 20*time.Millisecond)
		if err == nil {
			t.Errorf("expected error, bud didn't get one")
		}
	})
}

func makeDelayServer(duration time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(duration)
		w.WriteHeader(http.StatusOK)
	}))
}
