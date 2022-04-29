package api

import (
	pb "module/common/proto/post_service"
	"module/post_service/domain"
)

func mapPost(post *domain.Post) *pb.Post {
	postPb := &pb.Post{
		Id:       post.Id.Hex(),
		Content:  post.Content,
		Links:    post.Links,
		Date:     post.Date,
		User:     post.User,
		Likes:    post.Likes,
		Dislikes: post.Dislikes,
	}
	return postPb
}

func mapNewPost(postPb *pb.Post) *domain.Post {
	post := &domain.Post{
		Content:  postPb.Content,
		Links:    postPb.Links,
		Date:     postPb.Date,
		User:     postPb.User,
		Likes:    postPb.Likes,
		Dislikes: postPb.Dislikes,
	}
	return post
}
