package service

import (
	"myapp/config"
	"myapp/graph/model"

	"github.com/vektah/gqlparser/gqlerror"
	"gorm.io/gorm"
)

func Invite(receiverId int, roomId int, inviterId int) (*model.PendingInvitation, error) {
	db, sql := config.ConnectGorm()
	defer sql.Close()

	if err := db.Table("subscribers_have_rooms").
		Where("room_id = ? AND subscriber_id = ?", roomId, receiverId).
		First(&model.SubscribersHaveRooms{}).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			newInvitation := model.PendingInvitation{
				InviterID:  inviterId,
				ReceiverID: receiverId,
				RoomID:     roomId,
			}

			err := db.Table("pending_invitation").Create(&newInvitation).Error
			if err != nil {
				return &model.PendingInvitation{}, err
			}

			return &newInvitation, nil
		}
	}

	return &model.PendingInvitation{}, &gqlerror.Error{
		Message: "User Already Exist In Room!",
	}
}

func SendInvitationRequestGet(subscriberId int) ([]*model.PendingInvitation, error) {
	db, sql := config.ConnectGorm()
	defer sql.Close()

	var invitations []*model.PendingInvitation
	err := db.Table("pending_invitation").Where("inviter_id = ?", subscriberId).Find(&invitations).Error

	return invitations, err
}

func ReceivedInvitationRequestGet(subscriberId int) ([]*model.PendingInvitation, error) {
	db, sql := config.ConnectGorm()
	defer sql.Close()

	var invitations []*model.PendingInvitation
	err := db.Table("pending_invitation").Where("receiver_id = ?", subscriberId).Find(&invitations).Error

	return invitations, err
}

func AcceptInvitation(roomId int, invitationId int, subscriberId int) (bool, error) {

	_, err := AddNewSubscriberToRoom(roomId, subscriberId)
	if err != nil {
		return false, err
	}

	err = DeleteByInvitationId(invitationId)
	if err != nil {
		return false, err
	}

	return true, nil
}

func DeleteByInvitationId(invitationId int) error {
	db, sql := config.ConnectGorm()
	defer sql.Close()

	err := db.Table("pending_invitation").Where("id = ?", invitationId).Delete(&model.PendingInvitation{}).Error

	if err != nil {
		return err
	}

	return nil
}
