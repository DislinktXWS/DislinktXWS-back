package api

import (
	pb "github.com/dislinktxws-back/common/proto/message_service"
	"github.com/dislinktxws-back/message_service/domain"
	"time"
)

func mapNewConversation(userPb *pb.Participants) *domain.Participants {
	conversation := &domain.Participants{
		Sender:   userPb.Sender,
		Receiver: userPb.Receiver,
	}
	return conversation
}

func mapNewMessage(messagePb *pb.Message) *domain.Message {
	message := &domain.Message{
		Sender:   messagePb.Sender,
		Receiver: messagePb.Receiver,
		Content:  messagePb.Content,
		Date:     time.Now(),
		IsRead:   messagePb.IsRead,
	}
	return message
}

func mapConversation(conversation *domain.Conversation) *pb.GetConversationResponse {
	id := conversation.Id.Hex()

	conversationPb := &pb.GetConversationResponse{
		Id:                 id,
		FirstParticipator:  conversation.FirstParticipator,
		SecondParticipator: conversation.SecondParticipator,
	}
	for _, message := range conversation.Messages {
		conversationPb.Messages = append(conversationPb.Messages, &pb.Message{
			Sender:   message.Sender,
			Receiver: message.Receiver,
			Content:  message.Content,
			Date:     message.Date.String(),
		})
	}

	return conversationPb
}
