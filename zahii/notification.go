package zahii

import (
	"fmt"
)

type NotificationService service

type Notification struct {
	Base
	Title  string `json:"title"`
	Body   string `json:"body"`
	Type   string `json:"type"`
	IsRead bool   `json:"is_read"`
	Data   string `json:"data"`
}

type ListNotificationResponse struct {
	BaseResponse
	Body []*Notification `json:"body"`
}

func (s *NotificationService) SetLocationID(id string) *NotificationService {
	s.locationID = id
	return s
}

func (s *NotificationService) GetCurrentNotifications() (*ListNotificationResponse, error) {
	var result ListNotificationResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get("/customer/notification/current")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *NotificationService) MarkAllRead() (*BaseResponse, error) {
	var result BaseResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Post("/customer/notification/base/mark-all-read")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *NotificationService) Read(id uint) (*BaseResponse, error) {
	var result BaseResponse
	_, err := s.client.newRequest(s.locationID).
		SetPathParam("id", fmt.Sprintf("%d", id)).
		SetResult(&result).
		Post("/customer/notification/base/read/{id}")
	if err != nil {
		return nil, err
	}
	return &result, nil
}
