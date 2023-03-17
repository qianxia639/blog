package db

import (
	"Blog/utils"
	"context"
	"fmt"
	"testing"

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

	require.Equal(t, user.ID, blog.ID)
	require.Equal(t, ty.ID, blog.ID)
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
