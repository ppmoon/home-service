package git

import (
	"github.com/go-git/go-git/v5"
	"os"
)

type Client struct {
}

func NewGitClient() *Client {
	return &Client{}
}

// git plain clone
func (c *Client) PlainClone(path, url string) (*git.Repository, error) {
	return git.PlainClone(path, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})
}

// git pull
func (c *Client) Pull(path string) error {
	r, err := git.PlainOpen(path)
	if err != nil {
		return err
	}
	w, err := r.Worktree()
	if err != nil {
		return err
	}
	return w.Pull(&git.PullOptions{})
}
