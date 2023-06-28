package mysql

import (
	"bluebell_backend/models"

	"go.uber.org/zap"
)

func CreateComment(comment *models.Comment) error {
	err := DB.Create(&comment).Error
	if err != nil {
		zap.L().Error("insert comment failed", zap.Error(err))
		return ErrorInsertFailed
	}
	return nil
}

func GetCommentListByIDs(ids []string) (commentList []*models.Comment, err error) {
	err = DB.Where("comment_id IN (?)", ids).Find(&commentList).Error
	return
}
