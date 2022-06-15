package betterttv_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mrmarble/telegram-emote-bot/pkg/betterttv"
	"github.com/stretchr/testify/require"
)

func TestSearch(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(`[{"id":"1","code":"KEKW","imageType":"png"}]`))
	}))
	defer server.Close()

	client := &betterttv.Client{
		HttpClient: server.Client(),
		URL:        server.URL,
	}

	results, err := client.Search("KEKW")

	require.Nil(t, err)
	require.Equal(t, results, []betterttv.SearchResult{{ID: "1", Code: "KEKW", Type: "png"}})
}

func TestUrl(t *testing.T) {
	result := betterttv.SearchResult{ID: "1", Code: "KEKW", Type: "png"}
	require.Equal(t, result.ID.Url(), "https://cdn.betterttv.net/emote/1/3x")
}
