package pkg

import (
	"context"
	"github.com/golang-interfaces/iioutil"
	"github.com/golang-interfaces/ios"
	"github.com/opspec-io/sdk-golang/model"
	"path/filepath"
)

func newFSHandle(
	path string,
) Handle {
	return fsHandle{
		ioUtil: iioutil.New(),
		os:     ios.New(),
		path:   path,
	}
}

// fsHandle allows interacting w/ a package sourced from the filesystem
type fsHandle struct {
	ioUtil iioutil.IIOUtil
	os     ios.IOS
	path   string
}

func (lh fsHandle) GetContent(
	ctx context.Context,
	contentPath string,
) (
	model.ReadSeekCloser,
	error,
) {
	return lh.os.Open(filepath.Join(lh.path, contentPath))
}

func (lh fsHandle) ListContents(
	ctx context.Context,
) (
	[]*model.PkgContent,
	error,
) {
	return lh.rListContents(lh.path)
}

// rListContents recursively lists pkg contents at path
func (lh fsHandle) rListContents(
	path string,
) (
	[]*model.PkgContent,
	error,
) {
	childFileInfos, err := lh.ioUtil.ReadDir(path)
	if nil != err {
		return nil, err
	}

	var contents []*model.PkgContent
	for _, contentFileInfo := range childFileInfos {

		contentPath := filepath.Join(path, contentFileInfo.Name())

		if contentFileInfo.IsDir() {
			// recurse into child dirs
			childContents, err := lh.rListContents(contentPath)
			if nil != err {
				return nil, err
			}
			contents = append(contents, childContents...)
		} else {
			contents = append(
				contents,
				&model.PkgContent{
					Path: contentPath,
					Size: contentFileInfo.Size(),
				},
			)
		}

	}

	return contents, err
}

func (lh fsHandle) Ref() string {
	return lh.path
}