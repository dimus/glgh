/*
Copyright © 2020 Dmitry Mozzherin <dmozzherin@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/dimus/glgh"
	"github.com/dimus/glgh/config"
	"github.com/gnames/gnlib/sys"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

const configText = `# WorkDir is a place to keep cached data on disk.
# WorkDir ~/.local/share/glgh

# GitLabToken is OpenAuth personal token.
# Go to https://gitlab.com/profile/personal_access_tokens.
# Create a new Access Token with "api" and "read_repository" scopes.
# GitLabToken: abcd

# GitLabOwner is your account of organization on GitLab.
# For example, gitlab.com/gogna/gnparser has "gogna" as the owner.
# GitLabOwner: owner

# GitLabRepo is the repository on GitLab.
# For example, gitlab.com/gogna/gnparser has "gnparser" as the repo.
# GitLabRepo: repo

# GitHubToken is OpenAuth personal token.
# Go to https://github.com/settings/tokens. Generate a new token with
# "repo" scope.
# GitHubToken: 1234

# GitHubOwner is your account of organization on GitHub.
# For example, github.com/gnames/gnparser has "gnames" as the owner.
# GitHubOwner: owner

# GitHubRepo is the repository on GitHub.
# For example, github.com/gnames/gnparser has "gnparser" as the repo.
# GitHubRepo: repo
`

var (
	opts []config.Option
)

// cfgData purpose is to achieve automatic import of data from the
// configuration file, if it exists.
type cfgData struct {
	WorkDir     string
	GitLabToken string
	GitLabOwner string
	GitLabRepo  string
	GitHubToken string
	GitHubOwner string
	GitHubRepo  string
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "glgh",
	Short: "glgh imports issues from a GitLab repo to a GitHub repo.",
	Run: func(cmd *cobra.Command, args []string) {
		if showVersionFlag(cmd) {
			os.Exit(0)
		}
		reimport, _ := cmd.Flags().GetBool("reimport")
		opts = append(opts, config.OptReimport(reimport))
		cfg := config.NewConfig(opts...)
		g := glgh.NewGlGh(cfg)
		err := g.Issues()
		if err != nil {
			log.Fatal(err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("version", "V", false, "version and build timestamp")
	rootCmd.Flags().BoolP("reimport", "r", false, "delete cache and reimport data")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	var home string
	var err error
	configFile := "glgh"

	// Find home directory.
	home, err = homedir.Dir()
	if err != nil {
		log.Fatalf("Cannot find home directory: %s.", err)
	}
	home = filepath.Join(home, ".config")

	viper.AddConfigPath(home)
	viper.SetConfigName(configFile)

	configPath := filepath.Join(home, fmt.Sprintf("%s.yaml", configFile))
	touchConfigFile(configPath, configFile)

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Printf("Using config file: %s.", viper.ConfigFileUsed())
	}
	getOpts()
}

// getOpts imports data from the configuration file. Some of the settings can
// be overriden by command line flags.
func getOpts() {
	cfg := &cfgData{}
	err := viper.Unmarshal(cfg)
	if err != nil {
		log.Fatalf("Cannot deserialize config data: %s.", err)
	}

	if cfg.WorkDir != "" {
		opts = append(opts, config.OptWorkDir(cfg.WorkDir))
	}
	if cfg.GitLabToken != "" {
		opts = append(opts, config.OptGitLabToken(cfg.GitLabToken))
	}
	if cfg.GitLabOwner != "" {
		opts = append(opts, config.OptGitLabOwner(cfg.GitLabOwner))
	}
	if cfg.GitLabRepo != "" {
		opts = append(opts, config.OptGitLabRepo(cfg.GitLabRepo))
	}
	if cfg.GitHubToken != "" {
		opts = append(opts, config.OptGitHubToken(cfg.GitHubToken))
	}
	if cfg.GitHubOwner != "" {
		opts = append(opts, config.OptGitHubOwner(cfg.GitHubOwner))
	}
	if cfg.GitHubRepo != "" {
		opts = append(opts, config.OptGitHubRepo(cfg.GitHubRepo))
	}
}

// touchConfigFile checks if config file exists, and if not, it gets created.
func touchConfigFile(configPath string, configFile string) {
	if sys.FileExists(configPath) {
		return
	}

	log.Printf("Creating config file: %s.", configPath)
	createConfig(configPath, configFile)
}

// createConfig creates config file.
func createConfig(path string, file string) {
	err := sys.MakeDir(filepath.Dir(path))
	if err != nil {
		log.Fatalf("Cannot create dir %s: %s.", path, err)
	}

	err = ioutil.WriteFile(path, []byte(configText), 0600)
	if err != nil {
		log.Fatalf("Cannot write to file %s: %s", path, err)
	}
}

// showVersionFlag provides version and the build timestamp. If it returns
// true, it means that version flag was given.
func showVersionFlag(cmd *cobra.Command) bool {
	hasVersionFlag, err := cmd.Flags().GetBool("version")
	if err != nil {
		log.Fatalf("Cannot get version flag: %s.", err)
	}

	if hasVersionFlag {
		fmt.Printf("\nversion: %s\nbuild: %s\n\n", glgh.Version, glgh.Build)
	}
	return hasVersionFlag
}
