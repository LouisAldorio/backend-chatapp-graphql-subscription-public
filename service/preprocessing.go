package service

import (
	"fmt"
	"myapp/auth"
	"myapp/config"
	"myapp/graph/model"

	"gorm.io/gorm"
	"github.com/vektah/gqlparser/gqlerror"
)

func Login(email string, password string) (*string,error) {

	db, sql := config.ConnectGorm()
	defer sql.Close()

	var subscriber model.Subscriber
	err := db.Table("subscribers").Where("email = ?",email).First(&subscriber).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &gqlerror.Error{
				Message: "User is not Registered!",
			}
		}
	}

	if !auth.CheckPassword(subscriber.HashedPassword, password) {
		return nil, &gqlerror.Error{
			Message: "Password is Invalid, Please try again!",
		}
	}

	token, err := auth.CreateToken(subscriber.ID)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func Register(input model.NewSubscriber) (*string,error) {

	var createdSubscriber *model.Subscriber

	db, sql := config.ConnectGorm()
	defer sql.Close()

	var subscriber model.Subscriber
	err := db.Table("subscribers").Where("email = ?",input.Email).First(&subscriber).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			createdSubscriber = SubcriberCreate(input)
		}else {
			return nil, fmt.Errorf(fmt.Sprintf("%s", err))
		}		
	}else if err == nil{
		return nil, &gqlerror.Error{
			Message: "Email has Already been Registered!",
		}
	}

	token, err := auth.CreateToken(createdSubscriber.ID)
	if err != nil {
		return nil, err
	}

	return &token, nil
}