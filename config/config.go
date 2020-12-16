package config

import (
	"fmt"
	"path/filepath"

	"github.com/gnames/gnlib/sys"
)

type Config struct {
	WorkDir     string
	GitLabToken string
	GitLabOwner string
	GitLabRepo  string
	GitHubToken string
	GitHubOwner string
	GitHubRepo  string
	Reimport    bool
}

func NewConfig(opts ...Option) Config {
	cfg := Config{
		WorkDir: sys.ConvertTilda("~/.local/share/glgh"),
	}
	for _, opt := range opts {
		opt(&cfg)
	}
	return cfg
}

func (c Config) GitLabRepoPath() string {
	return fmt.Sprintf("%s/%s", c.GitLabOwner, c.GitHubRepo)
}

func (c Config) GitLabFilePath() string {
	return filepath.Join(c.WorkDir, "gitlab_issues_dump.gob")
}

type Option func(*Config)

func OptWorkDir(s string) Option {
	return func(c *Config) {
		s = sys.ConvertTilda(s)
		c.GitLabToken = s
	}
}

func OptGitLabToken(s string) Option {
	return func(c *Config) {
		c.GitLabToken = s
	}
}

func OptGitLabOwner(s string) Option {
	return func(c *Config) {
		c.GitLabOwner = s
	}
}

func OptGitLabRepo(s string) Option {
	return func(c *Config) {
		c.GitLabRepo = s
	}
}

func OptGitHubToken(s string) Option {
	return func(c *Config) {
		c.GitHubToken = s
	}
}

func OptGitHubOwner(s string) Option {
	return func(c *Config) {
		c.GitHubOwner = s
	}
}

func OptGitHubRepo(s string) Option {
	return func(c *Config) {
		c.GitHubRepo = s
	}
}

func OptReimport(b bool) Option {
	return func(c *Config) {
		c.Reimport = b
	}
}
