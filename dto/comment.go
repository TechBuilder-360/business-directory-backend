package dto

type Comment struct {
	UserId     int    `db:"user_id"`
	BusinessId int    `db:"business_id"`
	Text       string `db:"text"`
	Parent     int    `db:"parent"` // Id of comment been referred to, if null then it is a root comment
	Kids       []int  `db:"kids"`   // Id of sub-comments; meant to allow easy counting of child comments
	Deleted    bool   `db:"deleted"`
	CreationOn string `db:"created_on"`
	Updated    bool   `db:"updated"`
}
