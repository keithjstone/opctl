package core

import (
	"context"
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opctl/opctl/sdks/go/model"
	nodeFakes "github.com/opctl/opctl/sdks/go/node/fakes"
)

var _ = Context("Eventser", func() {
	Context("Events", func() {
		It("should call nodeHandle.APIClient().GetEventStream w/ expected args", func() {
			/* arrange */
			providedCtx := context.Background()

			fakeCore := new(nodeFakes.FakeOpNode)
			eventChannel := make(chan model.Event)
			close(eventChannel)
			fakeCore.GetEventStreamReturns(eventChannel, nil)

			objectUnderTest := _eventser{
				core: fakeCore,
			}

			/* act */
			objectUnderTest.Events(
				providedCtx,
			)

			/* assert */
			actualCtx,
				actualGetEventStreamReq := fakeCore.GetEventStreamArgsForCall(0)

			Expect(actualCtx).To(Equal(providedCtx))
			Expect(*actualGetEventStreamReq).To(Equal(model.GetEventStreamReq{}))

		})
		Context("client.GetEventStream errors", func() {
			It("should return expected error", func() {
				/* arrange */
				returnedError := errors.New("dummyError")

				fakeCore := new(nodeFakes.FakeOpNode)
				fakeCore.GetEventStreamReturns(nil, returnedError)

				objectUnderTest := _eventser{
					core: fakeCore,
				}

				/* act */
				err := objectUnderTest.Events(context.Background())

				/* assert */
				Expect(err).To(MatchError(returnedError))
			})
		})
		Context("client.GetEventStream doesn't error", func() {
			Context("channel closes unexpectedly", func() {
				It("should return expected error", func() {
					/* arrange */
					fakeCore := new(nodeFakes.FakeOpNode)
					eventChannel := make(chan model.Event)
					close(eventChannel)
					fakeCore.GetEventStreamReturns(eventChannel, nil)

					objectUnderTest := _eventser{
						core: fakeCore,
					}

					/* act */
					err := objectUnderTest.Events(context.Background())

					/* assert */
					Expect(err).To(MatchError(err))
				})
			})
		})
	})
})
