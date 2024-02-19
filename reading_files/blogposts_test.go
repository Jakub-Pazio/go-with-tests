package blogpost_test

import (
	blogposts "go-with-tests/reading_files"
	"testing"
	"testing/fstest"
)

func TestNewBlogPost(t *testing.T) {
	fs := fstest.MapFS{
		"hello-world.md":  {Data: []byte("Title: htmx")},
		"hello-world2.md": {Data: []byte("Title: Go")},
	}

	titles := []string{
		"htmx",
		"Go",
	}

	posts, err := blogposts.NewPostsFromFS(fs)

	for i, post := range posts {
		assertPostTitle(t, post, titles[i])
	}

	if err != nil {
		t.Errorf("exected to not fail but got: %q", err.Error())
	}

	if len(posts) != len(fs) {
		t.Errorf("exprected %d posts but got %d posts", len(fs), len(posts))
	}
}

func assertPostTitle(t testing.TB, post blogposts.Post, want string) {
	t.Helper()
	if post.Title != want {
		t.Errorf("exprected %s but got %s", want, post.Title)
	}
}
