package app

import "github.com/Rindrics/require-label-prefix-on-closed/domain"

type Command interface {
	Execute() error
}

type PostCommentParams struct {
	RepoInfo domain.RepoInfo
	Number   int
	Body     string
}

type Commenter interface {
	PostComment(PostCommentParams) error
}

type PostCommentCommand struct {
	Params    PostCommentParams
	Commenter Commenter
	onSuccess Action
}

type Labeler interface {
	AddLabels(AddLabelsParams) error
}

type AddLabelsParams struct {
	RepoInfo domain.RepoInfo
	Number   int
	Labels   domain.Labels
}

type AddLabelsCommand struct {
	Params    AddLabelsParams
	Labeler   Labeler
	onSuccess PostCommentCommand
}
