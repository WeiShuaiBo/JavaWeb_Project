package logic

import (
	"bluebell_backend/dao/mysql"
	"bluebell_backend/dao/redis"
	"bluebell_backend/models"
	"bluebell_backend/pkg/snowflake"
	"fmt"
	"go.uber.org/zap"
)

func CreatePost(post *models.Post) (err error) {
	// 生成帖子ID
	postID, err := snowflake.GetID()
	if err != nil {
		zap.L().Error("snowflake.GetID() failed", zap.Error(err))
		return
	}
	post.PostID = postID
	// 创建帖子
	if err := mysql.CreatePost(post); err != nil {
		zap.L().Error("mysql.CreatePost(&post) failed", zap.Error(err))
		return err
	}
	community, err := mysql.GetCommunityNameByID(fmt.Sprint(post.CommunityID))
	if err != nil {
		zap.L().Error("mysql.GetCommunityNameByID failed", zap.Error(err))
		return err
	}
	if err := redis.CreatePost(
		fmt.Sprint(post.PostID),
		fmt.Sprint(post.AuthorId),
		post.Title,
		TruncateByWords(post.Content, 120),
		community.CommunityName); err != nil {
		zap.L().Error("redis.CreatePost failed", zap.Error(err))
		return err
	}

	return
}

func GetPost(postID string) (post *models.ApiPostDetail, err error) {
	fmt.Println(postID + "asdfasdf")
	post, err = mysql.GetPostByID(postID)
	fmt.Println(post, "111222")
	if err != nil {
		zap.L().Error("mysql.GetPostByID(postID) failed", zap.String("post_id", postID), zap.Error(err))
		return nil, err
	}
	user, err := mysql.GetUserByID(fmt.Sprint(post.AuthorId))
	if err != nil {
		zap.L().Error("mysql.GetUserByID() failed", zap.String("author_id", fmt.Sprint(post.AuthorId)), zap.Error(err))
		return
	}
	post.AuthorName = user.UserName
	community, err := mysql.GetCommunityByID(fmt.Sprint(post.CommunityID))
	if err != nil {
		zap.L().Error("mysql.GetCommunityByID() failed", zap.String("community_id", fmt.Sprint(post.CommunityID)), zap.Error(err))
		return
	}
	post.CommunityName = community.CommunityName
	return post, nil
}

func GetPostList2() (data []*models.ApiPostDetail, err error) {
	postList, err := mysql.GetPostList()
	if err != nil {
		zap.L().Error("获取帖子列表出错，请你重新查看")
		return
	}
	data = make([]*models.ApiPostDetail, 0, len(postList))
	results, err := redis.Client.ZRevRangeWithScores("bluebell:post:score", 0, -1).Result()

	for i := range results {
		member := results[i].Member.(string)
		score := results[i].Score
		fmt.Println(member, score)
		post, _ := GetPost(member)
		fmt.Println(post, "111")
		//fmt.Println(post)
		//	user, err1 := mysql.GetUserByID(fmt.Sprint(post.AuthorId))
		//	if err1 != nil {
		//		zap.L().Error("mysql.GetUserByID() failed", zap.String("author_id", fmt.Sprint(post.AuthorId)), zap.Error(err))
		//		continue
		//	}
		//	post.AuthorName = user.UserName
		//	community, err1 := mysql.GetCommunityByID(fmt.Sprint(post.CommunityID))
		//	if err1 != nil {
		//		zap.L().Error("mysql.GetCommunityByID() failed", zap.String("community_id", fmt.Sprint(post.CommunityID)), zap.Error(err))
		//		continue
		//	}
		//	post.CommunityName = community.CommunityName
		//	data = append(data, post)
	}
	//fmt.Println(data)
	return
}
