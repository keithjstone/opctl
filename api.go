package sdk

//go:generate counterfeiter -o ./fakeApi.go --fake-name FakeApi ./ Api

import "github.com/opspec-io/sdk-golang/models"

type Api interface {
  AddOp(
  req models.AddOpReq,
  ) (err error)
  
  SetDescriptionOfOp(
  req models.SetDescriptionOfOpReq,
  ) (err error)
}

func New(
filesys Filesystem,
) (api Api) {

  var compositionRoot compositionRoot
  compositionRoot = newCompositionRoot(
    filesys,
  )

  api = &_api{
    compositionRoot:compositionRoot,
  }

  return
}

type _api struct {
  compositionRoot compositionRoot
}

func (this _api) AddOp(
req models.AddOpReq,
) (err error) {
  return this.
  compositionRoot.
  AddOpUseCase().
  Execute(req)
}

func (this _api) SetDescriptionOfOp(
req models.SetDescriptionOfOpReq,
) (err error) {
  return this.
  compositionRoot.
  SetDescriptionOfOpUseCase().
  Execute(req)
}
