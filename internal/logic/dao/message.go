package dao

import (
	"go-IM/internal/logic/db"
	"go-IM/internal/logic/model"
)

type messageDao struct{}

var MessageDao = new(messageDao)

// Add 插入一条消息
func (*messageDao) Add(message model.Message) error {
	_, err := db.Cli.Exec(`
		INSERT INTO message(app_id,object_type,object_id,request_id,sender_type,sender_id,sender_device_id,receiver_type,receiver_id,to_user_ids,type,content,seq,send_time) 
		VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?)
	`, message.AppId, message.ObjectType, message.ObjectId, message.RequestId, message.SenderType, message.SenderId, message.SenderDeviceId,
		message.ReceiverType, message.ReceiverId, message.ToUserIds, message.Type, message.Content, message.Seq, message.SendTime)
	return err
}

// 根据类型和id查询大于序号大于seq的消息
func (*messageDao) ListBySeq(appId, objectType, objectId, seq int64) (*[]model.Message, error) {
	rows, e := db.Cli.Query(`
			SELECT app_id,object_type,object_id,request_id,sender_type,sender_id,sender_device_id,receiver_type,receiver_id,to_user_ids,type,content,seq,send_time 
			FROM message 
			WHERE app_id = ? 
			AND object_type = ?
			AND object_id = ?
			AND seq > ?`, appId, objectType, objectId, seq)
	if e != nil {
		return nil, e
	}
	messages := make([]model.Message, 0, 5)
	for rows.Next() {
		message := new(model.Message)
		e := rows.Scan(&message.AppId, &message.ObjectType, &message.ObjectId, &message.RequestId, &message.SenderType, &message.SenderId, &message.SenderDeviceId, &message.ReceiverType,
			&message.ReceiverId, &message.ToUserIds, &message.Type, &message.Content, &message.Seq, &message.SendTime)
		if e != nil {
			return nil, e
		}
		messages = append(messages, *message)
	}
	return &messages, nil
}
