package service

import (
	"context"
	"errors"

	"AI-Study-Community/internal/dao"
	"AI-Study-Community/internal/model/do"
	"AI-Study-Community/internal/model/entity"
)

// Custom errors for Comment service
var (
	ErrCommentNotFound = errors.New("comment not found")
	ErrCommentNotOwned = errors.New("comment not owned by user")
)

// Comment is the service layer for comments
type Comment struct {
}

// CommentService is the singleton instance of Comment service
var CommentService = &Comment{}

// Create creates a new comment.
// It validates the input and calls the DAO layer to insert the comment.
// Returns the created comment ID.
func (s *Comment) Create(ctx context.Context, req *do.CommentCreateReq) (*do.CommentCreateResp, error) {
	// Call DAO layer to insert the comment
	id, err := dao.CommentDao.Insert(ctx, req)
	if err != nil {
		return nil, err
	}

	return &do.CommentCreateResp{
		Id: id,
	}, nil
}

// Reply replies to an existing comment.
// It validates the parent comment exists and inserts the reply.
// Returns the created reply comment ID.
func (s *Comment) Reply(ctx context.Context, req *do.CommentReplyReq) (*do.CommentReplyResp, error) {
	// Check if parent comment exists
	parent, err := dao.CommentDao.GetOne(ctx, req.ParentId)
	if err != nil {
		return nil, err
	}
	if parent == nil {
		return nil, ErrCommentNotFound
	}

	// Build the create request from reply request
	createReq := &do.CommentCreateReq{
		PostId:        req.PostId,
		ParentId:      &req.ParentId,
		Content:       req.Content,
		AuthorId:      req.AuthorId,
		ReplyToUserId: req.ReplyToUserId,
	}

	// Insert the reply comment
	id, err := dao.CommentDao.Insert(ctx, createReq)
	if err != nil {
		return nil, err
	}

	return &do.CommentReplyResp{
		Id: id,
	}, nil
}

// Delete soft-deletes a comment.
// It validates the comment exists and belongs to the current user.
// Returns success flag.
func (s *Comment) Delete(ctx context.Context, req *do.CommentDeleteReq) (*do.CommentDeleteResp, error) {
	// Check if comment exists
	comment, err := dao.CommentDao.GetOne(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if comment == nil {
		return nil, ErrCommentNotFound
	}

	// Check if comment belongs to the current user
	if comment.AuthorId != req.AuthorId {
		return nil, ErrCommentNotOwned
	}

	// Delete the comment
	_, err = dao.CommentDao.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &do.CommentDeleteResp{
		Success: true,
	}, nil
}

// GetDetail retrieves a single comment by ID.
// It calls the DAO layer to get the comment and converts to response structure.
func (s *Comment) GetDetail(ctx context.Context, req *do.CommentGetOneReq) (*do.CommentResp, error) {
	// Get comment from DAO
	comment, err := dao.CommentDao.GetOne(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if comment == nil {
		return nil, ErrCommentNotFound
	}

	// Convert entity to response
	resp := s.entityToResp(comment, nil)

	return resp, nil
}

// GetTreeByPostId retrieves all comments for a post in tree structure.
// It fetches all top-level comments and their children, then builds a tree.
func (s *Comment) GetTreeByPostId(ctx context.Context, req *do.CommentGetTreeReq) (*do.CommentTreeResp, error) {
	// Get top-level comments (paginated)
	page := req.Page
	if page < 1 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize < 1 {
		pageSize = 20
	}

	// Get top-level comments
	topLevelReq := &do.CommentGetByPostIdReq{
		PostId:   req.PostId,
		Page:     page,
		PageSize: pageSize,
	}
	topLevelComments, total, err := dao.CommentDao.GetByPostId(ctx, topLevelReq)
	if err != nil {
		return nil, err
	}

	// Get ALL comments for this post (for building tree - use large page size)
	allReq := &do.CommentGetByPostIdReq{
		PostId:   req.PostId,
		Page:     1,
		PageSize: 10000, // Large enough to get all comments
	}
	allComments, _, err := dao.CommentDao.GetByPostId(ctx, allReq)
	if err != nil {
		return nil, err
	}

	// Get all child comments for this post
	allChildren, err := s.getAllChildrenByPostId(ctx, req.PostId)
	if err != nil {
		return nil, err
	}

	// Combine all comments into a map for quick lookup
	allCommentsMap := make(map[uint64]*entity.Comment)
	commentMap := make(map[uint64]*do.CommentResp)

	// Add all top-level comments
	for _, c := range allComments {
		allCommentsMap[c.Id] = c
	}

	// Add all child comments
	for _, c := range allChildren {
		allCommentsMap[c.Id] = c
	}

	// Convert all entities to responses
	for _, c := range allCommentsMap {
		commentMap[c.Id] = s.entityToResp(c, nil)
	}

	// Build tree structure
	// First pass: build children map
	childrenMap := make(map[uint64][]*do.CommentResp)
	for _, c := range allCommentsMap {
		if c.ParentId != nil {
			childrenMap[*c.ParentId] = append(childrenMap[*c.ParentId], commentMap[c.Id])
		}
	}

	// Second pass: assign children to parents
	for _, resp := range commentMap {
		if children, ok := childrenMap[resp.Id]; ok {
			resp.Children = children
		} else {
			resp.Children = []*do.CommentResp{}
		}
	}

	// Collect root-level comments (tree structure) - these are top-level comments
	tree := make([]*do.CommentResp, 0)
	for _, c := range topLevelComments {
		tree = append(tree, commentMap[c.Id])
	}

	// Flatten list (all comments)
	flatList := make([]*do.CommentResp, 0, len(topLevelComments))
	for _, c := range topLevelComments {
		flatList = append(flatList, commentMap[c.Id])
	}

	return &do.CommentTreeResp{
		Total: total,
		List:  flatList,
		Tree:  tree,
	}, nil
}

// getAllChildrenByPostId retrieves all child comments for a post (non-paginated).
func (s *Comment) getAllChildrenByPostId(ctx context.Context, postId uint64) ([]*entity.Comment, error) {
	// Get all top-level comment IDs first
	topLevelReq := &do.CommentGetByPostIdReq{
		PostId:   postId,
		Page:     1,
		PageSize: 10000,
	}
	topLevelComments, _, err := dao.CommentDao.GetByPostId(ctx, topLevelReq)
	if err != nil {
		return nil, err
	}

	// Collect all child comments recursively
	var allChildren []*entity.Comment
	for _, parent := range topLevelComments {
		children, err := s.getChildrenRecursively(ctx, parent.Id)
		if err != nil {
			return nil, err
		}
		allChildren = append(allChildren, children...)
	}

	return allChildren, nil
}

// getChildrenRecursively retrieves all children of a comment recursively.
func (s *Comment) getChildrenRecursively(ctx context.Context, parentId uint64) ([]*entity.Comment, error) {
	var allChildren []*entity.Comment

	// Get direct children
	childrenReq := &do.CommentGetChildrenReq{
		ParentId: parentId,
		Page:     1,
		PageSize: 1000,
	}
	children, _, err := dao.CommentDao.GetChildren(ctx, childrenReq)
	if err != nil {
		return nil, err
	}

	for _, child := range children {
		allChildren = append(allChildren, child)
		// Recursively get children of this child
		grandchildren, err := s.getChildrenRecursively(ctx, child.Id)
		if err != nil {
			return nil, err
		}
		allChildren = append(allChildren, grandchildren...)
	}

	return allChildren, nil
}

// entityToResp converts a Comment entity to a CommentResp.
func (s *Comment) entityToResp(e *entity.Comment, children []*do.CommentResp) *do.CommentResp {
	return &do.CommentResp{
		Id:            e.Id,
		PostId:        e.PostId,
		UserId:        e.AuthorId,
		Content:       e.Content,
		ParentId:      e.ParentId,
		AuthorId:      e.AuthorId,
		AuthorName:    "", // TODO: Populate from user service
		ReplyToUserId: e.ReplyToUserId,
		ReplyToAuthor: "", // TODO: Populate from user service
		Depth:         e.Depth,
		LikeCount:     e.LikeCount,
		Status:        e.Status,
		CreatedAt:     e.CreatedAt,
		UpdatedAt:     e.UpdatedAt,
		Children:     children,
	}
}
