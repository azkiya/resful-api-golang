package models

type News struct {
	ID     interface{} `json:"id,omitempty" bson:"_id,omitempty"`
	Title  string      `json:"title,omitempty" bson:"title,omitempty"`
	Tags   []string    `json:"tags,omitempty" bson:"tags,omitempty"`
	Status string      `json:"status,omitempty" bson:"status,omitempty"`
}

type NewsUpdate struct {
	ModifiedCount int64 `json:"modifiedCount"`
	Result        News
}

type NewsDelete struct {
	DeletedCount int64 `json:"deletedCount"`
}
