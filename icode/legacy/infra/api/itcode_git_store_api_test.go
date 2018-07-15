package api

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/src-d/go-git.v4"
)

var gitUrl = "git@github.com:it-chain/tesseract.git"

func TestClone(t *testing.T) {

	os.RemoveAll("./.tmp")
	defer os.RemoveAll("./.tmp")
	api := NewGitApi(nil)

	sc, err := api.Clone(gitUrl)

	if err != nil {
		assert.NoError(t, err)
	}

	assert.Equal(t, "./.tmp/"+"tesseract", sc.Path)
	assert.Equal(t, gitUrl, sc.GitUrl)
	assert.Equal(t, "tesseract", sc.RepositoryName)
}

func TestChangeRemote(t *testing.T) {

	//given
	os.RemoveAll("./.tmp")
	defer os.RemoveAll("./.tmp")
	api := NewGitApi(nil)
	iCodeMeta, err := api.Clone(gitUrl)
	assert.NoError(t, err)

	//when
	err = api.ChangeRemote(iCodeMeta.Path, "https://github.com/steve-buzzni"+"/"+iCodeMeta.RepositoryName)
	assert.NoError(t, err)

	//then
	r, err := git.PlainOpen(iCodeMeta.Path)
	assert.NoError(t, err)
	remote, err := r.Remote(git.DefaultRemoteName)
	assert.NoError(t, err)
	assert.Equal(t, "https://github.com/steve-buzzni"+"/"+iCodeMeta.RepositoryName, remote.Config().URLs[0])
}

func TestGitApi_Push(t *testing.T) {

	//given
	b, err := NewBackupGithubStoreApi("", "")
	assert.NoError(t, err)
	api := NewGitApi(b)
	iCodeMeta, err := api.Clone(gitUrl)
	assert.NoError(t, err)
	defer func() {
		os.RemoveAll(tmp)
		ctx := context.Background()

		//then
		_, err := client.Repositories.Delete(ctx, "steve-buzzni", iCodeMeta.RepositoryName)
		assert.NoError(t, err)
	}()

	assert.NoError(t, err)

	//when
	err = api.Push(iCodeMeta)

	//then
	assert.NoError(t, err)
}

func TestGetNameFromGitUrl(t *testing.T) {
	name := getNameFromGitUrl(gitUrl)

	assert.Equal(t, "tesseract", name)
}

func TestDirExist(t *testing.T) {

	p := "tmp"
	wp := "tmp2"
	defer os.RemoveAll(p)

	err := os.MkdirAll(p, 0755)

	assert.NoError(t, err)
	assert.True(t, dirExists(p))
	assert.False(t, dirExists(wp))
}
