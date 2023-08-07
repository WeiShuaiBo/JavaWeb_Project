package mysql

import (
	"bluebell_backend/models"
	"fmt"
	"gorm.io/gorm"

	"go.uber.org/zap"
)

// CreatePost 创建帖子
func CreatePost(post *models.Post) (err error) {
	result := DB.Create(post)
	if result.Error != nil {
		zap.L().Error("insert post failed", zap.Error(result.Error))
		err = ErrorInsertFailed
		return
	}
	return
}

func GetPostByID(idStr string) (*models.ApiPostDetail, error) {
	var post *models.ApiPostDetail
	fmt.Println(idStr)
	err := DB.Table("post").Debug().Select("post_id, title, content, author_id, community_id, update_time").Where("post_id = ?", idStr).First(&post).Error
	if err == gorm.ErrRecordNotFound {
		zap.L().Error("数据没有找到")
		err = ErrorInvalidID
		return nil, err
	}
	if err != nil {
		zap.L().Error("query post failed", zap.String("sql", err.Error()), zap.Error(err))
		err = ErrorQueryFailed
		return nil, err
	}
	return post, nil
}

func GetPostListByIDs(ids []string) (postList []*models.Post, err error) {
	err = DB.Table("post").Select("post_id, title, content, author_id, community_id, create_time").
		Where("post_id IN (?)", ids).
		Find(&postList).Error
	return
}

func GetPostList() ([]*models.ApiPostDetail, error) {
	var posts []*models.Post
	err := DB.Table("post").
		Select("post_id, title, content, author_id, community_id, create_time").
		Limit(8).
		Find(&posts).Error
	if err != nil {
		return nil, err
	}

	apiPosts := make([]*models.ApiPostDetail, len(posts))
	for i, post := range posts {
		var user models.User
		var community models.Community
		DB.Where("user_id=?", post.AuthorId).First(&user)
		DB.Where("community_id", post.CommunityID).First(&community)
		apiPosts[i] = &models.ApiPostDetail{
			Post:          post,
			AuthorName:    user.UserName,
			CommunityName: community.CommunityName,
		}
	}
	return apiPosts, nil
}
