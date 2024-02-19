package blogpost

import (
	"io"
	"io/fs"
)

type Post struct {
	Title string
}

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	var files []Post
	fs, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}

	for _, f := range fs {
		post, err := getPost(fileSystem, f)
		if err != nil {
			return nil, err
		}
		files = append(files, post)
	}

	return files, nil
}

func getPost(fs fs.FS, file fs.DirEntry) (Post, error) {
	postFile, err := fs.Open(file.Name())
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	postContent, err := io.ReadAll(postFile)
	if err != nil {
		return Post{}, err
	}
	return Post{Title: string(postContent[7:])}, nil
}
