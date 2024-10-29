package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {

	store := NewInMemoryPlayerStore()
	srv := PlayerServer{Store: store}
	player := "Nicholas"

	srv.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	srv.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	srv.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	srv.ServeHTTP(response, newGetScoreRequest(player))
	assertStatus(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), "3")
}

func Benchmark(b *testing.B) {
	store := NewInMemoryPlayerStore()
	srv := PlayerServer{Store: store}
	player := "Nicholas"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		srv.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	}

	response := httptest.NewRecorder()
	srv.ServeHTTP(response, newGetScoreRequest(player))
	assertStatus(b, response.Code, http.StatusOK)

	assertResponseBody(b, response.Body.String(), fmt.Sprintf("%d", b.N))
}
