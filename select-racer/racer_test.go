package racer

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRacer(t *testing.T) {

	t.Run("When they both complete in time", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		e := fastURL
		a, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Errorf("received unexpected error: %v", err)
		}

		if a != e {
			t.Errorf("expected: %q actual %q", e, a)
		}
	})

	t.Run("returns an error if a server does not respond within 10 seconds", func(t *testing.T) {
		timeout := time.Duration(3 * time.Millisecond)
		serverA := makeDelayedServer(4 * time.Millisecond)
		serverB := makeDelayedServer(5 * time.Millisecond)

		defer serverA.Close()
		defer serverB.Close()
		fmt.Println("A")
		_, err := ConfigurableRacer(serverA.URL, serverB.URL, timeout)
		fmt.Println("B")
		if err == nil {
			fmt.Println("C")
			t.Error("expected an error but didn't get one")
		}
	})

}
