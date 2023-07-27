package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.35

import (
	"context"
	"fmt"
	"github.com/google/uuid"

	"github.com/dave-andrew/gqlgen-todos/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, newUser model.NewUser) (*model.User, error) {
	users := &model.User{
		ID:       uuid.NewString(),
		Name:     newUser.Name,
		Password: newUser.Password,
	}

	if err := r.Db.Create(users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, newUser model.NewUser) (*model.User, error) {
	var users model.User
	if err := r.Db.Where("Name = ?", newUser.Name).First(&users).Error; err != nil {
		return nil, err
	}

	if newUser.Password != "" {
		users.Password = newUser.Password
	}

	if err := r.Db.Where("Name = ?", newUser.Name).Updates(&users).Error; err != nil {
		return nil, err
	}

	return &users, nil
}

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, post model.NewPost) (*model.Post, error) {
	var user model.User
	if err := r.Db.Where("ID = ?", post.UserID).First(&user).Error; err != nil {
		return nil, err
	}

	posts := &model.Post{
		ID:          uuid.NewString(),
		Title:       post.Title,
		Description: post.Description,
		UserID:      post.UserID,
	}

	if err := r.Db.Create(posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, userID string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: GetUser - getUser"))
}

// GetAllUser is the resolver for the getAllUser field.
func (r *queryResolver) GetAllUser(ctx context.Context) ([]*model.User, error) {
	var allUser []*model.User
	if err := r.Db.Find(&allUser).Error; err != nil {
		return nil, err
	}

	return allUser, nil
}

// GetAllPost is the resolver for the getAllPost field.
func (r *queryResolver) GetAllPost(ctx context.Context) ([]*model.Post, error) {
	var allPost []*model.Post
	if err := r.Db.Find(&allPost).Preload("User").Find(&allPost).Error; err != nil {
		return nil, err
	}

	return allPost, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }