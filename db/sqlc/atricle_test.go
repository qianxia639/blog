package db

import (
	"Blog/utils"
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomArticle(t *testing.T) Article {
	user := createRandomUser(t)

	title := utils.RandomString(6)
	content := fmt.Sprintf("%s-content", title)
	image := fmt.Sprintf("%s.jpg", title)

	arg := &InsertArticleParams{
		OwnerID:   user.ID,
		Title:     title,
		Content:   content,
		Image:     image,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	article, err := testQueries.InsertArticle(context.Background(), arg)
	require.NoError(t, err)

	require.Equal(t, user.ID, article.OwnerID)
	require.Equal(t, title, article.Title)
	require.Equal(t, content, article.Content)
	require.Equal(t, image, article.Image)

	return article
}

func TestInseertArticle(t *testing.T) {
	createRandomArticle(t)
}

func TestIncrViews(t *testing.T) {
	article := createRandomArticle(t)

	err := testQueries.IncrViews(context.Background(), article.ID)
	require.NoError(t, err)
}

func TestGetArticle(t *testing.T) {

	article1 := createRandomArticle(t)

	article2, err := testQueries.GetArticle(context.Background(), article1.ID)
	require.NoError(t, err)

	require.Equal(t, article1.ID, article2.ID)
	require.Equal(t, article1.OwnerID, article2.OwnerID)
	require.Equal(t, article1.Title, article2.Title)
	require.Equal(t, article1.Content, article2.Content)
	require.Equal(t, article1.Image, article2.Image)
	require.Equal(t, article1.Views, article2.Views)
	require.WithinDuration(t, article1.CreatedAt, article2.CreatedAt, time.Second)
	require.WithinDuration(t, article1.UpdatedAt, article2.UpdatedAt, time.Second)
}

func TestListArticles(t *testing.T) {

	for i := 0; i < 10; i++ {
		createRandomArticle(t)
	}

	articles, err := testQueries.ListArticles(context.Background(), &ListArticlesParams{
		Title:  "",
		Limit:  5,
		Offset: 0,
	})
	require.NoError(t, err)

	for _, article := range articles {
		require.NotEmpty(t, article)
	}
}

func TestDeleteArticle(t *testing.T) {

	article := createRandomArticle(t)

	err := testQueries.DeleteArticle(ctx, article.ID)
	require.NoError(t, err)
}

func TestUpdateArticleOnlyTitle(t *testing.T) {

	article := createRandomArticle(t)

	title := utils.RandomString(6)

	newArticle, err := testQueries.UpdateArticle(ctx, &UpdateArticleParams{
		ID: article.ID,
		Title: sql.NullString{
			String: title,
			Valid:  true,
		},
		UpdatedAt: time.Now(),
	})
	require.NoError(t, err)

	require.NotEqual(t, article.Title, newArticle.Title)
	require.Equal(t, title, newArticle.Title)
}

func TestUpdateArticleOnlyContent(t *testing.T) {

	article := createRandomArticle(t)

	content := utils.RandomString(50)

	newArticle, err := testQueries.UpdateArticle(ctx, &UpdateArticleParams{
		ID: article.ID,
		Content: sql.NullString{
			String: content,
			Valid:  true,
		},
		UpdatedAt: time.Now(),
	})
	require.NoError(t, err)

	require.NotEqual(t, article.Content, newArticle.Content)
	require.Equal(t, content, newArticle.Content)
}

func TestUpdateArticleOnlyImage(t *testing.T) {

	article := createRandomArticle(t)

	image := fmt.Sprintf("%s.jpg", utils.RandomString(32))

	newArticle, err := testQueries.UpdateArticle(ctx, &UpdateArticleParams{
		ID: article.ID,
		Image: sql.NullString{
			String: image,
			Valid:  true,
		},
		UpdatedAt: time.Now(),
	})
	require.NoError(t, err)

	require.NotEqual(t, article.Image, newArticle.Image)
	require.Equal(t, image, newArticle.Image)
}

func TestUpdateArticleAll(t *testing.T) {
	oldArticle := createRandomArticle(t)

	title := utils.RandomString(6)
	image := fmt.Sprintf("%s.jpg", utils.RandomString(32))
	content := utils.RandomString(50)

	newArticle, err := testQueries.UpdateArticle(ctx, &UpdateArticleParams{
		ID: oldArticle.ID,
		Title: sql.NullString{
			String: title,
			Valid:  true,
		},
		Image: sql.NullString{
			String: image,
			Valid:  true,
		},
		Content: sql.NullString{
			String: content,
			Valid:  true,
		},
		UpdatedAt: time.Now(),
	})
	require.NoError(t, err)

	require.Equal(t, title, newArticle.Title)
	require.Equal(t, content, newArticle.Content)
	require.Equal(t, image, newArticle.Image)
}
