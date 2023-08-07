package models

import (
	"encoding/json"
	"errors"
	"time"
)

type Post struct {
	PostID      uint64    `json:"post_id" gorm:"column:post_id"`
	Title       string    `json:"title" gorm:"column:title"`
	Content     string    `json:"content" gorm:"column:content"`
	AuthorId    uint64    `json:"author_id" gorm:"column:author_id"`
	CommunityID int64     `json:"community_id" gorm:"column:community_id"`
	Status      int32     `json:"status" gorm:"column:status"`
	UpdateTime  time.Time `json:"" gorm:"column:update_time"`
}

func (p *Post) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		Title       string `json:"title" `
		Content     string `json:"content" `
		CommunityID int64  `json:"community_id"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.Title) == 0 {
		err = errors.New("帖子标题不能为空")
	} else if len(required.Content) == 0 {
		err = errors.New("帖子内容不能为空")
	} else if required.CommunityID == 0 {
		err = errors.New("未指定版块")
	} else {
		//id, err1 := strconv.ParseInt(required.CommunityID, 10, 64)
		//if err1 != nil {
		//	return err1
		//}
		p.Title = required.Title
		p.Content = required.Content
		p.CommunityID = required.CommunityID
	}
	return
}

type ApiPostDetail struct {
	*Post
	AuthorName    string `json:"author_name"`
	CommunityName string `json:"community_name"`
}

func (p Post) TableName() string {
	return "post"
}
