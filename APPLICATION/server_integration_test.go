package poker

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {

	//store := NewInMemoryPlayerStore()
	database, cleanDatabase := createTempFile(t, "[]")
	defer cleanDatabase()

	store, err := NewFileSystemPlayerStore(database)
	assertNoError(t, err)

	srv := NewPlayerServer(store)
	player := "Nicholas"

	srv.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	srv.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	srv.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		srv.ServeHTTP(response, newGetScoreRequest(player))
		assertStatus(t, response.Code, http.StatusOK)

		assertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		srv.ServeHTTP(response, newLeagueRequest())
		assertStatus(t, response.Code, http.StatusOK)

		got := getLeagueFromResponse(t, response.Body)
		want := League{
			{"Nicholas", 3},
		}
		assertLeague(t, got, want)
	})

}

func Benchmark(b *testing.B) {
	//store := NewInMemoryPlayerStore()
	database, cleanDatabase := createTempFile(b, "[]")
	defer cleanDatabase()

	store, err := NewFileSystemPlayerStore(database)
	assertNoError(b, err)

	srv := NewPlayerServer(store)
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
