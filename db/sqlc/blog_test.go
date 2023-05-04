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

func createRandomBlog(t *testing.T) Blog {
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
	}

	blog, err := testQueries.InsertArticle(context.Background(), arg)
	require.NoError(t, err)

	require.Equal(t, user.ID, blog.OwnerID)
	require.Equal(t, title, blog.Title)
	require.Equal(t, content, blog.Content)
	require.Equal(t, image, blog.Image)

	return blog
}

func TestInseertBlog(t *testing.T) {
	createRandomBlog(t)
}

func TestIncrViews(t *testing.T) {
	blog := createRandomBlog(t)

	err := testQueries.IncrViews(context.Background(), blog.ID)
	require.NoError(t, err)
}

func TestGetBlog(t *testing.T) {

	blog1 := createRandomBlog(t)

	blog2, err := testQueries.GetArticle(context.Background(), blog1.ID)
	require.NoError(t, err)

	require.Equal(t, blog1.ID, blog2.ID)
	require.Equal(t, blog1.OwnerID, blog2.OwnerID)
	require.Equal(t, blog1.Title, blog2.Title)
	require.Equal(t, blog1.Content, blog2.Content)
	require.Equal(t, blog1.Image, blog2.Image)
	require.Equal(t, blog1.Views, blog2.Views)
	require.WithinDuration(t, blog1.CreatedAt, blog2.CreatedAt, time.Second)
	require.WithinDuration(t, blog1.UpdatedAt, blog2.UpdatedAt, time.Second)
}

func TestListBlogs(t *testing.T) {

	for i := 0; i < 10; i++ {
		createRandomBlog(t)
	}

	blogs, err := testQueries.ListArticles(context.Background(), &ListArticlesParams{
		Limit:  5,
		Offset: 5,
	})
	require.NoError(t, err)
	require.Len(t, blogs, 5)

	for _, blog := range blogs {
		require.NotEmpty(t, blog)
	}
}

func TestDeleteBlog(t *testing.T) {

	blog := createRandomBlog(t)

	err := testQueries.DeleteArticle(ctx, blog.ID)
	require.NoError(t, err)
}

func TestUpdateBlogOnlyTitle(t *testing.T) {

	blog := createRandomBlog(t)

	title := utils.RandomString(6)

	newBlog, err := testQueries.UpdateArticle(ctx, &UpdateArticleParams{
		ID: blog.ID,
		Title: sql.NullString{
			String: title,
			Valid:  true,
		},
		UpdatedAt: time.Now(),
	})
	require.NoError(t, err)

	require.NotEqual(t, blog.Title, newBlog.Title)
	require.Equal(t, title, newBlog.Title)
}

func TestUpdateBlogOnlyContent(t *testing.T) {

	blog := createRandomBlog(t)

	content := utils.RandomString(50)

	newBlog, err := testQueries.UpdateArticle(ctx, &UpdateArticleParams{
		ID: blog.ID,
		Content: sql.NullString{
			String: content,
			Valid:  true,
		},
		UpdatedAt: time.Now(),
	})
	require.NoError(t, err)

	require.NotEqual(t, blog.Content, newBlog.Content)
	require.Equal(t, content, newBlog.Content)
}

func TestUpdateBlogOnlyImage(t *testing.T) {

	blog := createRandomBlog(t)

	image := fmt.Sprintf("%s.jpg", utils.RandomString(32))

	newBlog, err := testQueries.UpdateArticle(ctx, &UpdateArticleParams{
		ID: blog.ID,
		Image: sql.NullString{
			String: image,
			Valid:  true,
		},
		UpdatedAt: time.Now(),
	})
	require.NoError(t, err)

	require.NotEqual(t, blog.Image, newBlog.Image)
	require.Equal(t, image, newBlog.Image)
}

func TestUpdateBlogAll(t *testing.T) {
	oldBlog := createRandomBlog(t)

	title := utils.RandomString(6)
	image := fmt.Sprintf("%s.jpg", utils.RandomString(32))
	content := utils.RandomString(50)

	newBlog, err := testQueries.UpdateArticle(ctx, &UpdateArticleParams{
		ID: oldBlog.ID,
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

	require.Equal(t, title, newBlog.Title)
	require.Equal(t, content, newBlog.Content)
	require.Equal(t, image, newBlog.Image)
}

func TestSearchBlog(t *testing.T) {

	str := fmt.Sprintf("%%%s%%", "s")

	var limit int32 = 5

	arg := &SearchArticleParams{
		Title:  str,
		Limit:  limit,
		Offset: (limit - 1) * limit,
	}

	_, err := testQueries.SearchArticle(ctx, arg)
	require.NoError(t, err)
}
