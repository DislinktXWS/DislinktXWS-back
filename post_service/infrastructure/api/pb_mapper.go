package api

import (
	pb "module/common/proto/post_service"
	"module/post_service/domain"
)

func mapPost(post *domain.Post) *pb.Post {
	postPb := &pb.Post{
		Id:      post.Id.Hex(),
		Content: post.Content,
		Date:    post.Date,
	}
	return postPb
}

func mapNewPost(postPb *pb.Post) *domain.Post {
	post := &domain.Post{
		Content: postPb.Content,
		Date:    postPb.Date,
	}
	return post
}
