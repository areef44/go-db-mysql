package repository

import (
	"context"
	"fmt"
	gomysql "go-mysql"
	"go-mysql/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(gomysql.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repository2@test.com",
		Comment: "Test Repository",
	}

	result, err := commentRepository.Insert(ctx, comment)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(gomysql.GetConnection())

	comment, err := commentRepository.FindById(context.Background(), 47)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(gomysql.GetConnection())

	comments, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}
