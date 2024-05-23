package system

type NotificationPayload struct {
}

func NewNotificationMessage(tag string) Message[NotificationPayload] {
	return Message[NotificationPayload]{
		Type:    MessageNotification,
		Tag:     tag,
		Payload: NotificationPayload{},
	}
}
