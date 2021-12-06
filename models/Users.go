package models

type User struct {
	//UserID      int         `json:"userId" bson:"userId"`
	ID       interface{} `json:"id,omitempty" bson:"_id,omitempty"`
	Login    string      `json:"login" bson:"login"`
	Password string      `json:"password" bson:"password"`
}

/* type TodoUpdate struct {
	ModifiedCount int64 `json:"modifiedCount"`
	Result        Todo
} */

/* type TodoDelete struct {
	DeletedCount int64 `json:"deletedCount" bson:"deletedCount"`
} */
