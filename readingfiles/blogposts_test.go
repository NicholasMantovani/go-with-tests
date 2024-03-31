package blogposts_test // the _test is needed if we don't want to test the "internal" functions but only the exported ones

import (
	"errors"
	blogposts "hello/readingfiles"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

type StubFailingFs struct {
}

func (s StubFailingFs) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, i always fail")
}

func TestNewBlogPosts(t *testing.T) {

	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
		seconBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
S
I
U
M`
	)
	t.Run("ok", func(t *testing.T) {
		mapFS := fstest.MapFS{
			"hello world.md":  {Data: []byte(firstBody)},
			"hello-world2.md": {Data: []byte(seconBody)},
		}
		posts, err := blogposts.NewPostsFromFs(mapFS)

		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(mapFS) {
			t.Errorf("got %d posts, wanted %d posts", len(posts), len(mapFS))
		}

		assertPost(t, posts[0], blogposts.Post{
			Title:       "Post 1",
			Description: "Description 1",
			Tags:        []string{"tdd", "go"},
			Body: `Hello
World`,
		})
	})

	t.Run("ko", func(t *testing.T) {
		_, err := blogposts.NewPostsFromFs(StubFailingFs{})

		if err == nil {
			t.Error("expected error but not found")
		}
	})

}

func assertPost(t testing.TB, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v want %+v", got, want)
	}
}
