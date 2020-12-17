package glapi

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/dimus/glgh/config"
	"github.com/dimus/glgh/entity/gitlab"
	"github.com/gnames/gnlib/encode"
	"github.com/gnames/gnlib/sys"
)

type glapi struct {
}

type gitlabAPI struct {
	cfg config.Config
}

func NewGitLabAPI(cfg config.Config) gitlab.GitLab {
	sys.MakeDir(cfg.WorkDir)
	return gitlabAPI{cfg: cfg}
}

func (g gitlabAPI) ReadIssues() (gitlab.Data, error) {
	var res gitlab.Data
	if sys.FileExists(g.cfg.GitLabFilePath()) && !g.cfg.Reimport {
		return g.readDump()
	}

	res, err := g.runGraphQL()
	if err != nil {
		log.Fatal(err)
	}
	err = g.writeDump(res)
	return res, err
}

func (g gitlabAPI) readDump() (gitlab.Data, error) {
	enc := encode.GNjson{}
	var res gitlab.Data

	f, err := os.Open(g.cfg.GitLabFilePath())
	if err != nil {
		return res, err
	}
	defer f.Close()

	encoded, err := ioutil.ReadAll(f)
	if err != nil {
		return res, err
	}

	err = enc.Decode(encoded, &res)
	return res, err
}

func (g gitlabAPI) writeDump(res gitlab.Data) error {
	enc := encode.GNjson{Pretty: true}
	encoded, err := enc.Encode(res)
	if err != nil {
		return err
	}
	f, err := os.Create(g.cfg.GitLabFilePath())
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(encoded)
	return err
}
