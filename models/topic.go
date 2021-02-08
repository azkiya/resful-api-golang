package models

type Topic struct {
	ID   interface{} `json:"id,omitempty" bson:"_id,omitempty"`
	Name string      `json:"name,omitempty" bson:"name,omitempty"`
}

type TopicUpdate struct {
	ModifiedCount int64 `json:"modifiedCount"`
	Result        Topic
}

type TopicDelete struct {
	DeletedCount int64 `json:"deletedCount"`
}
