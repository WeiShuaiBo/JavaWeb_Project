package mysql

import (
	"bluebell_backend/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"go.uber.org/zap"
)

func GetCommunityList() ([]*models.Community, error) {
	var communityList []*models.Community

	err := DB.Find(&communityList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return communityList, nil
}
func GetCommunityNameByID(idStr string) (*models.Community, error) {
	var community models.Community

	err := DB.First(&community, "community_id = ?", idStr).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = ErrorInvalidID
		} else {
			zap.L().Error("query community failed", zap.String("sql", err.Error()), zap.Error(err))
			err = ErrorQueryFailed
		}
		return nil, err
	}

	return &community, nil
}

func GetCommunityByID(idStr string) (*models.CommunityDetail, error) {
	var community models.CommunityDetail

	err := DB.First(&community, "community_id = ?", idStr).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = ErrorInvalidID
		} else {
			zap.L().Error("query community failed", zap.String("sql", err.Error()), zap.Error(err))
			err = ErrorQueryFailed
		}
		return nil, err
	}

	return &community, nil
}
