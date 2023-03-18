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
	ty := createRandomType(t)

	title := utils.RandomString(6)
	content := fmt.Sprintf("%s-content", title)
	image := fmt.Sprintf("%s.jpg", title)

	arg := InsertBlogParams{
		OwnerID: user.ID,
		TypeID:  ty.ID,
		Title:   title,
		Content: content,
		Image:   image,
	}

	blog, err := testQueries.InsertBlog(context.Background(), arg)
	require.NoError(t, err)

	require.Equal(t, user.ID, blog.OwnerID)
	require.Equal(t, ty.ID, blog.TypeID)
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

	blog2, err := testQueries.GetBlog(context.Background(), blog1.ID)
	require.NoError(t, err)

	require.Equal(t, blog1.ID, blog2.ID)
	require.Equal(t, blog1.OwnerID, blog2.OwnerID)
	require.Equal(t, blog1.TypeID, blog2.TypeID)
	require.Equal(t, blog1.Title, blog2.Title)
	require.Equal(t, blog1.Content, blog2.Content)
	require.Equal(t, blog1.Image, blog2.Image)
	require.Equal(t, blog1.Views, blog2.Views)
	require.WithinDuration(t, blog1.CreatedAt, blog2.CreatedAt, time.Second)
	require.WithinDuration(t, blog1.UpdatedAt, blog2.UpdatedAt, time.Second)
}

func TestListBlogs(t *testing.T) {

	_, err := testQueries.ListBlogs(context.Background())
	require.NoError(t, err)
}

func TestDeleteBlog(t *testing.T) {

	blog := createRandomBlog(t)

	err := testQueries.DeleteBlog(ctx, blog.ID)
	require.NoError(t, err)
}

func TestUpdateBlogOnlyTypeID(t *testing.T) {

	blog := createRandomBlog(t)
	ty := createRandomType(t)

	newBlog, err := testQueries.UpdateBlog(ctx, UpdateBlogParams{
		ID: blog.ID,
		TypeID: sql.NullInt64{
			Int64: ty.ID,
			Valid: true,
		},
	})
	require.NoError(t, err)

	require.NotEqual(t, blog.TypeID, newBlog.TypeID)
	require.Equal(t, ty.ID, newBlog.TypeID)
}

func TestUpdateBlogOnlyTitle(t *testing.T) {

	blog := createRandomBlog(t)

	title := utils.RandomString(6)

	newBlog, err := testQueries.UpdateBlog(ctx, UpdateBlogParams{
		ID: blog.ID,
		Title: sql.NullString{
			String: title,
			Valid:  true,
		},
	})
	require.NoError(t, err)

	require.NotEqual(t, blog.Title, newBlog.Title)
	require.Equal(t, title, newBlog.Title)
}

func TestUpdateBlogOnlyContent(t *testing.T) {

	blog := createRandomBlog(t)

	content := utils.RandomString(50)

	newBlog, err := testQueries.UpdateBlog(ctx, UpdateBlogParams{
		ID: blog.ID,
		Content: sql.NullString{
			String: content,
			Valid:  true,
		},
	})
	require.NoError(t, err)

	require.NotEqual(t, blog.Content, newBlog.Content)
	require.Equal(t, content, newBlog.Content)
}

func TestUpdateBlogOnlyImage(t *testing.T) {

	blog := createRandomBlog(t)

	image := fmt.Sprintf("%s.jpg", utils.RandomString(32))

	newBlog, err := testQueries.UpdateBlog(ctx, UpdateBlogParams{
		ID: blog.ID,
		Image: sql.NullString{
			String: image,
			Valid:  true,
		},
	})
	require.NoError(t, err)

	require.NotEqual(t, blog.Image, newBlog.Image)
	require.Equal(t, image, newBlog.Image)
}

func TestUpdateBlogAll(t *testing.T) {
	oldBlog := createRandomBlog(t)

	ty := createRandomType(t)

	title := utils.RandomString(6)
	image := fmt.Sprintf("%s.jpg", utils.RandomString(32))
	content := utils.RandomString(50)

	newBlog, err := testQueries.UpdateBlog(ctx, UpdateBlogParams{
		ID: oldBlog.ID,
		TypeID: sql.NullInt64{
			Int64: ty.ID,
			Valid: true,
		},
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
	})
	require.NoError(t, err)

	require.Equal(t, ty.ID, newBlog.TypeID)
	require.Equal(t, title, newBlog.Title)
	require.Equal(t, content, newBlog.Content)
	require.Equal(t, image, newBlog.Image)
}