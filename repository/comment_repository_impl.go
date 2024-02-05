package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-database/model"
	"strconv"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment model.Comment) (model.Comment, error) {
	script := "INSERT INTO comments(email, comment) VALUES(?,?)"

	result, err := repository.DB.ExecContext(ctx, script, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}

	comment.Id = int32(id)

	return comment, nil
}

func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int32) (model.Comment, error) {
	script := "SELECT id, email, comment FROM comments WHERE id=? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	comment := model.Comment{}
	if err != nil {
		return comment, err
	}
	defer rows.Close()

	if rows.Next() {
		//ada
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		//tidak ada
		return comment, errors.New("Comment id" + strconv.Itoa(int(id)) + "not found")
	}
}

func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]model.Comment, error) {
	script := "SELECT id, email, comment FROM comments"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []model.Comment
	for rows.Next() {
		//ada
		comment := model.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)

		comments = append(comments, comment)
	}
	return comments, nil
}
