package algolia

type Task interface {
	TaskID() int64
}
type ObjectTask interface {
	Task
	ObjectID() string
}

type objectTask struct {
	ID    int64  `json:"taskID"`
	ObjID string `json:"objectID"`
}

func (t *objectTask) TaskID() int64 {
	return t.ID
}

func (t *objectTask) ObjectID() string {
	return t.ObjID
}
