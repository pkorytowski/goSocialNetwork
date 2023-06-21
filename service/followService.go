package service

import "socialNetwork/model"

func FollowUser(follow model.Follow) error {
	if err := model.DB.Create(&follow).Error; err != nil {
		return err
	}

	return nil
}

func UnfollowUser(follow model.Follow) error {
	if err := model.DB.Delete(&follow).Error; err != nil {
		return err
	}

	return nil
}

func GetFollowById(id int) model.Follow {
	var follow model.Follow
	model.DB.Where("id = ?", id).First(&follow)
	return follow
}

func GetFollowedByUserId(userId int) []model.Follow {
	var follows []model.Follow
	model.DB.Where("follows_id = ?", userId).Find(&follows)
	return follows
}
