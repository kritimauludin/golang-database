package repository

import (
	"context"
	"fmt"
	golangdatabase "golang-database"
	"golang-database/model"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(golangdatabase.GetConnection())

	ctx := context.Background()

	comment := model.Comment{
		Email:   "repository@test.com",
		Comment: "Test repository lg",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(golangdatabase.GetConnection())

	result, err := commentRepository.FindById(context.Background(), 22)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)

}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(golangdatabase.GetConnection())

	comments, err := commentRepository.FindAll(context.Background())

	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}
