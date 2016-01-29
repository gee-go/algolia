package algolia

import (
	"net/http"

	"github.com/drinkin/di/env"
)

// Client interacts with the Algolia REST API.
type Client interface {
	DeleteIndex(indexName string) (Task, error)
	ClearIndex(indexName string) (Task, error)

	// AddObject adds one object in the index with automatic assignation of objectID
	AddObject(indexName string, obj interface{}) (ObjectTask, error)
	UpsertObject(indexName string, objectID string, obj interface{}) (ObjectTask, error)
	GetObject(id string, attrs ...string) Value
	DeleteObject(id string) (Task, error)
}

func New(appId, apiKey string) Client {
	return &client{
		service: &Service{
			appId:      appId,
			apiKey:     apiKey,
			httpClient: http.DefaultClient,
		},
	}
}

func FromEnv() Client {
	return New(env.MustGet("ALGOLIA_APP_ID"), env.MustGet("ALGOLIA_API_KEY"))
}
