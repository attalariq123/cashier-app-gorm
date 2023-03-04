package repository

import (
	"a21hc3NpZ25tZW50/model"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type SessionsRepository struct {
	db *gorm.DB
}

func NewSessionsRepository(db *gorm.DB) SessionsRepository {
	return SessionsRepository{db}
}

func (u *SessionsRepository) AddSessions(session model.Session) error {

	res := u.db.Create(&session)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (u *SessionsRepository) DeleteSessions(tokenTarget string) error {

	res := u.db.Model(model.Session{}).Where("token = ?", tokenTarget).Delete(&model.Session{})
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (u *SessionsRepository) UpdateSessions(session model.Session) error {

	res := u.db.Model(&model.Session{}).Where("username = ?", session.Username).Updates(session)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (u *SessionsRepository) TokenValidity(token string) (model.Session, error) {

	availSession, err := u.SessionAvailToken(token)
	if err != nil {
		return model.Session{}, err
	}

	if u.TokenExpired(availSession) {
		err := u.DeleteSessions(token)
		if err != nil {
			return model.Session{}, err
		}
		return model.Session{}, fmt.Errorf("Token is Expired!")
	}

	return availSession, nil
}

func (u *SessionsRepository) SessionAvailName(name string) (model.Session, error) {

	var session model.Session
	res := u.db.First(&session, "username = ?", name)
	if res.Error != nil {
		return model.Session{}, res.Error
	}

	return session, nil
}

func (u *SessionsRepository) SessionAvailToken(token string) (model.Session, error) {

	var session model.Session
	res := u.db.First(&session, "token = ?", token)
	if res.Error != nil {
		return model.Session{}, res.Error
	}

	return session, nil
}

func (u *SessionsRepository) TokenExpired(s model.Session) bool {
	return s.Expiry.Before(time.Now())
}
