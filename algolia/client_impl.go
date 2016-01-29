package algolia

import (
	"errors"
)

type client struct {
	service *Service
}

func (c *client) AddObject(s string, i interface{}) (ObjectTask, error) {
	panic(errors.New("Not implemented"))
}

func (c *client) ClearIndex(s string) (Task, error) {
	panic(errors.New("Not implemented"))
}

func (c *client) DeleteIndex(s string) (Task, error) {
	panic(errors.New("Not implemented"))
}

func (c *client) DeleteObject(s string) (Task, error) {
	panic(errors.New("Not implemented"))
}

func (c *client) GetObject(s string, s1 []string) Value {
	panic(errors.New("Not implemented"))
}

func (c *client) UpsertObject(s string, s1 string, i interface{}) (ObjectTask, error) {
	panic(errors.New("Not implemented"))
}
