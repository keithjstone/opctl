package core

import (
	"errors"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opspec-io/opctl/util/updater"
)

var _ = Describe("selfUpdate", func() {

	Context("Execute", func() {
		It("should call exiter with expected args when provided invalid channel", func() {
			/* arrange */
			fakeExiter := new(fakeExiter)
			providedReleaseChannel := "invalidChannel"

			objectUnderTest := _core{
				exiter: fakeExiter,
			}

			/* act */
			objectUnderTest.SelfUpdate(providedReleaseChannel)

			/* assert */
			Expect(fakeExiter.ExitArgsForCall(0)).
				Should(Equal(ExitReq{Message: fmt.Sprintf(
					"%v is not an available release channel. "+
						"Available release channels are 'beta' 'stable'. \n", providedReleaseChannel), Code: 1}))
		})
		It("should call exiter with expected args when updater.GetUpdateIfExists returns error", func() {
			/* arrange */
			fakeExiter := new(fakeExiter)
			returnedError := errors.New("dummyError")

			fakeUpdater := new(updater.FakeUpdater)
			fakeUpdater.GetUpdateIfExistsReturns(&updater.Update{}, returnedError)

			objectUnderTest := _core{
				updater: fakeUpdater,
				exiter:  fakeExiter,
			}

			/* act */
			objectUnderTest.SelfUpdate("beta")

			/* assert */
			Expect(fakeExiter.ExitArgsForCall(0)).
				Should(Equal(ExitReq{Message: returnedError.Error(), Code: 1}))
		})
		It("should call updater.GetUpdateIfExists with expected args", func() {
			/* arrange */
			fakeUpdater := new(updater.FakeUpdater)

			objectUnderTest := _core{
				updater: fakeUpdater,
				exiter:  new(fakeExiter),
			}

			providedChannel := "beta"

			/* act */
			objectUnderTest.SelfUpdate(providedChannel)

			/* assert */
			Expect(fakeUpdater.GetUpdateIfExistsArgsForCall(0)).Should(Equal(providedChannel))
		})
		Describe("when updater.GetUpdateIfExists returns no update", func() {
			It("should call exiter with expected args", func() {
				/* arrange */
				fakeExiter := new(fakeExiter)

				fakeUpdater := new(updater.FakeUpdater)
				fakeUpdater.GetUpdateIfExistsReturns(nil, nil)

				objectUnderTest := _core{
					updater: fakeUpdater,
					exiter:  fakeExiter,
				}

				/* act */
				objectUnderTest.SelfUpdate("beta")

				/* assert */
				Expect(fakeExiter.ExitArgsForCall(0)).
					Should(Equal(ExitReq{Message: "No update available, already at the latest version!", Code: 0}))
			})
		})
		Describe("when update.GetUpdateIfExists returns an update", func() {
			Describe("and updater.ApplyUpdate is successful", func() {
				It("should call exiter with expected args", func() {
					/* arrange */
					fakeExiter := new(fakeExiter)

					fakeUpdater := new(updater.FakeUpdater)
					returnedUpdate := &updater.Update{Version: "dummyVersion"}

					fakeUpdater.GetUpdateIfExistsReturns(returnedUpdate, nil)

					objectUnderTest := _core{
						updater: fakeUpdater,
						exiter:  fakeExiter,
					}

					/* act */
					objectUnderTest.SelfUpdate("beta")

					/* assert */
					Expect(fakeExiter.ExitArgsForCall(0)).
						Should(Equal(ExitReq{
							Message: fmt.Sprintf("Updated to new version: %s!\n", returnedUpdate.Version),
							Code:    0,
						}))
				})
			})
			Describe("and updater.ApplyUpdate returns error", func() {
				It("should call exiter with expected args", func() {
					/* arrange */
					fakeExiter := new(fakeExiter)
					returnedError := errors.New("dummyError")

					fakeUpdater := new(updater.FakeUpdater)

					fakeUpdater.GetUpdateIfExistsReturns(&updater.Update{Version: "dummyVersion"}, nil)

					fakeUpdater.ApplyUpdateReturns(returnedError)

					objectUnderTest := _core{
						updater: fakeUpdater,
						exiter:  fakeExiter,
					}

					/* act */
					objectUnderTest.SelfUpdate("beta")

					/* assert */
					Expect(fakeExiter.ExitArgsForCall(0)).
						Should(Equal(ExitReq{Message: returnedError.Error(), Code: 1}))
				})
			})
		})
	})
})