package urls_processor

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"net/http"
	"testing"
)

func TestVisitUrls(t *testing.T) {
	defer gock.Off()
	gock.New("https://google.com").
		Get("/").
		Reply(http.StatusOK).
		BodyString(`12345`)

	defer gock.Off()
	gock.New("https://github.com/calebikhuohon").
		Get("/").
		Reply(http.StatusOK).
		BodyString(`123456789`)

	result := VisitUrls(context.Background(), []string{"https://google.com", "https://github.com/calebikhuohon"})
	assert.Equal(t, []Pair{
		{
			Url:      "https://github.com/calebikhuohon",
			BodySize: 9,
		},
		{
			Url:      "https://google.com",
			BodySize: 5,
		},
	}, result)
}
