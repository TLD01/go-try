package geofence

import (
	"fmt"

	"aerowatch.com/api/common"
	"aerowatch.com/api/geolocation"
)

type NotificationType string

const (
	NotificationTypeEmail NotificationType = "Email"
	NotificationTypeSMS   NotificationType = "SMS"
	NotificationTypePush  NotificationType = "Push"
	NotificationTypeInApp NotificationType = "InApp"
)

var validNotificationTypes = map[NotificationType]struct{}{
    NotificationTypeEmail: {},
    NotificationTypeSMS:   {},
    NotificationTypePush:  {},
    NotificationTypeInApp: {},
}

func (nt NotificationType) IsValid() bool {
	_, exists := validNotificationTypes[nt]
	return exists
}


func NewNotificationType(value string) (NotificationType, error) {
	var nt NotificationType = NotificationType(value)
	if !nt.IsValid() {		
		return "", fmt.Errorf("invalid notification type: %s", value)
	}
	return nt, nil
}

type NotificationSettings struct {
	Type       NotificationType `json:"type" bson:"type"`
	Recipients []string         `json:"recipients" bson:"recipients"`
}

type Geofence struct {
	common.Persisted
	Name                 string                 `json:"name"`
	Enabled              bool                   `json:"enabled"`
	Polygon              geolocation.Polygon    `json:"polygon"`
	NotificationSettings []NotificationSettings `json:"notificationSettings"`
}
