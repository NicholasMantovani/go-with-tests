package blogposts

import (
	"io/fs"
)

func NewPostsFromFs(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")

	if err != nil {
		return nil, err
	}

	var posts []Post
	for _, f := range dir {
		post, err := getPost(fileSystem, f.Name())
		if err != nil {
			return nil, err // todo: needs clarification
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func getPost(filesystem fs.FS, filename string) (Post, error) {
	postFile, err := filesystem.Open(filename)
	if err != nil {
		return Post{}, err
	}

	defer postFile.Close()

	return newPost(postFile)
}
