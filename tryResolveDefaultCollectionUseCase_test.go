package opspec

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "github.com/opspec-io/sdk-golang/models"
  "errors"
  "os"
  "time"
  "path"
)

// dummy fileInfo
type fileInfo struct {
  name    string      // base name of the file
  size    int64       // length in bytes for regular files; system-dependent for others
  mode    os.FileMode // file mode bits
  modTime time.Time   // modification time
  isDir   bool        // abbreviation for Mode().IsDir()
  sys     interface{} // underlying data source (can return nil)
}

func (this fileInfo) Name() string {
  return this.name
}

func (this fileInfo) Size() int64 {
  return this.size
}
func (this fileInfo) Mode() os.FileMode {
  return this.mode
}
func (this fileInfo) ModTime() time.Time {
  return this.modTime
}
func (this fileInfo) IsDir() bool {
  return this.isDir
}
func (this fileInfo) Sys() interface{} {
  return &this.sys
}

var _ = Describe("_tryResolveDefaultCollectionUseCase", func() {

  Context("Execute", func() {

    Context("when Filesystem.ListChildFileInfosOfDir returns an error", func() {

      It("should be returned", func() {

        /* arrange */
        expectedError := errors.New("ListChildFileInfosOfDirError")

        fakeFilesystem := new(FakeFilesystem)
        fakeFilesystem.ListChildFileInfosOfDirReturns(nil, expectedError)

        objectUnderTest := newTryResolveDefaultCollectionUseCase(
          fakeFilesystem,
        )

        /* act */
        _, actualError := objectUnderTest.Execute(
          models.TryResolveDefaultCollectionReq{},
        )

        /* assert */
        Expect(actualError).To(Equal(expectedError))

      })

    })

    Context("when default op collection doesn't exist", func() {

      It("should return an empty pathToDefaultCollection", func() {

        /* arrange */
        fakeFilesystem := new(FakeFilesystem)
        fakeFilesystem.ListChildFileInfosOfDirReturns(nil, nil)

        objectUnderTest := newTryResolveDefaultCollectionUseCase(
          fakeFilesystem,
        )

        /* act */
        pathToDefaultCollection, _ := objectUnderTest.Execute(
          models.TryResolveDefaultCollectionReq{},
        )

        /* assert */
        Expect(pathToDefaultCollection).To(BeEmpty())

      })

      It("should return a nil err", func() {

        /* arrange */
        fakeFilesystem := new(FakeFilesystem)
        fakeFilesystem.ListChildFileInfosOfDirReturns(nil, nil)

        objectUnderTest := newTryResolveDefaultCollectionUseCase(
          fakeFilesystem,
        )

        /* act */
        _, err := objectUnderTest.Execute(
          models.TryResolveDefaultCollectionReq{},
        )

        /* assert */
        Expect(err).To(BeNil())

      })

    })
    Context("when default op collection is in provided PathToDir", func() {

      It("should return expected pathToDefaultCollection", func() {

        /* arrange */
        providedPathToDir := "/dummy/path"

        fakeFilesystem := new(FakeFilesystem)
        fakeFilesystem.ListChildFileInfosOfDirReturns(
          []os.FileInfo{
            fileInfo{
              name:NameOfDefaultOpCollection,
              isDir:true,
            },
          },
          nil,
        )

        expectedPathToDefaultCollection := path.Join(providedPathToDir, NameOfDefaultOpCollection)

        objectUnderTest := newTryResolveDefaultCollectionUseCase(
          fakeFilesystem,
        )

        /* act */
        actualPathToDefaultCollection, _ := objectUnderTest.Execute(
          models.TryResolveDefaultCollectionReq{
            PathToDir:providedPathToDir,
          },
        )

        /* assert */
        Expect(actualPathToDefaultCollection).To(Equal(expectedPathToDefaultCollection))

      })

      It("should return a nil err", func() {

        /* arrange */
        fakeFilesystem := new(FakeFilesystem)
        fakeFilesystem.ListChildFileInfosOfDirReturns(
          []os.FileInfo{
            fileInfo{
              name:NameOfDefaultOpCollection,
              isDir:true,
            },
          },
          nil,
        )

        objectUnderTest := newTryResolveDefaultCollectionUseCase(
          fakeFilesystem,
        )

        /* act */
        _, err := objectUnderTest.Execute(
          models.TryResolveDefaultCollectionReq{},
        )

        /* assert */
        Expect(err).To(BeNil())

      })

    })

    It("should call Filesystem.ListChildFileInfosOfDir with expected args", func() {

      /* arrange */
      expectedListChildFileInfosOfDirArg := "/dummy/path"

      fakeFilesystem := new(FakeFilesystem)
      fakeFilesystem.ListChildFileInfosOfDirReturns(nil, nil)

      objectUnderTest := newTryResolveDefaultCollectionUseCase(
        fakeFilesystem,
      )

      /* act */
      objectUnderTest.Execute(
        models.TryResolveDefaultCollectionReq{
          PathToDir:expectedListChildFileInfosOfDirArg,
        },
      )

      /* assert */
      actualListChildFileInfosOfDirArg := fakeFilesystem.ListChildFileInfosOfDirArgsForCall(0)
      Expect(actualListChildFileInfosOfDirArg).To(Equal(expectedListChildFileInfosOfDirArg))

    })

  })

})
