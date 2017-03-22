package core

import (
	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opctl/opctl/util/pubsub"
	"github.com/opctl/opctl/util/uniquestring"
	"github.com/opspec-io/sdk-golang/model"
)

var _ = Context("parallelCaller", func() {
	Context("newParallelCaller", func() {
		It("should return parallelCaller", func() {
			/* arrange/act/assert */
			Expect(newParallelCaller(
				new(fakeCaller),
				new(pubsub.Fake),
				new(uniquestring.Fake),
			)).Should(Not(BeNil()))
		})
	})
	Context("Call", func() {
		It("should call caller for every parallelCall w/ expected args", func() {
			/* arrange */
			providedCallId := "dummyCallId"
			providedInboundScope := map[string]*model.Data{}
			providedRootOpId := "dummyRootOpId"
			providedPkgRef := "dummyPkgRef"
			providedScgParallelCalls := []*model.Scg{
				{
					Container: &model.ScgContainerCall{},
				},
				{
					Op: &model.ScgOpCall{},
				},
				{
					Parallel: []*model.Scg{},
				},
				{
					Serial: []*model.Scg{},
				},
			}

			fakeCaller := new(fakeCaller)

			returnedUniqueString := "dummyUniqueString"
			fakeUniqueStringFactory := new(uniquestring.Fake)
			fakeUniqueStringFactory.ConstructReturns(returnedUniqueString)

			objectUnderTest := newParallelCaller(fakeCaller, new(pubsub.Fake), fakeUniqueStringFactory)

			/* act */
			objectUnderTest.Call(
				providedCallId,
				providedInboundScope,
				providedRootOpId,
				providedPkgRef,
				providedScgParallelCalls,
			)

			/* assert */
			actualScgParallelCalls := []*model.Scg{}
			for callIndex := range providedScgParallelCalls {
				actualNodeId,
					actualChildOutboundScope,
					actualScg,
					actualPkgRef,
					actualRootOpId := fakeCaller.CallArgsForCall(callIndex)
				Expect(actualNodeId).To(Equal(returnedUniqueString))
				Expect(actualChildOutboundScope).To(Equal(providedInboundScope))
				Expect(actualPkgRef).To(Equal(providedPkgRef))
				Expect(actualRootOpId).To(Equal(providedRootOpId))
				actualScgParallelCalls = append(actualScgParallelCalls, actualScg)
			}
			Expect(actualScgParallelCalls).To(ConsistOf(providedScgParallelCalls))
		})
		Context("caller errors", func() {
			It("shouldn't exit until all childCalls complete & return expected error", func() {
				/* arrange */
				providedCallId := "dummyCallId"
				providedInboundScope := map[string]*model.Data{}
				providedRootOpId := "dummyRootOpId"
				providedPkgRef := "dummyPkgRef"
				providedScgParallelCalls := []*model.Scg{
					{
						Container: &model.ScgContainerCall{},
					},
					{
						Op: &model.ScgOpCall{},
					},
					{
						Parallel: []*model.Scg{},
					},
					{
						Serial: []*model.Scg{},
					},
				}

				expectedError := errors.New("Error(s) encountered during parallel call")
				fakeCaller := new(fakeCaller)
				fakeCaller.CallReturns(errors.New("dummyError"))

				returnedUniqueString := "dummyUniqueString"
				fakeUniqueStringFactory := new(uniquestring.Fake)
				fakeUniqueStringFactory.ConstructReturns(returnedUniqueString)

				objectUnderTest := newParallelCaller(fakeCaller, new(pubsub.Fake), fakeUniqueStringFactory)

				/* act */
				actualError := objectUnderTest.Call(
					providedCallId,
					providedInboundScope,
					providedRootOpId,
					providedPkgRef,
					providedScgParallelCalls,
				)

				/* assert */
				actualScgParallelCalls := []*model.Scg{}
				for callIndex := range providedScgParallelCalls {
					actualNodeId,
						actualChildOutboundScope,
						actualScg,
						actualPkgRef,
						actualRootOpId := fakeCaller.CallArgsForCall(callIndex)
					Expect(actualNodeId).To(Equal(returnedUniqueString))
					Expect(actualChildOutboundScope).To(Equal(providedInboundScope))
					Expect(actualPkgRef).To(Equal(providedPkgRef))
					Expect(actualRootOpId).To(Equal(providedRootOpId))
					actualScgParallelCalls = append(actualScgParallelCalls, actualScg)
				}
				Expect(actualScgParallelCalls).To(ConsistOf(providedScgParallelCalls))
				Expect(actualError).To(Equal(expectedError))
			})
		})
		Context("caller doesn't error", func() {
			It("shouldn't exit until all childCalls complete & not error", func() {
				/* arrange */
				providedCallId := "dummyCallId"
				providedInboundScope := map[string]*model.Data{}
				providedRootOpId := "dummyRootOpId"
				providedPkgRef := "dummyPkgRef"
				providedScgParallelCalls := []*model.Scg{
					{
						Container: &model.ScgContainerCall{},
					},
					{
						Op: &model.ScgOpCall{},
					},
					{
						Parallel: []*model.Scg{},
					},
					{
						Serial: []*model.Scg{},
					},
				}

				fakeCaller := new(fakeCaller)

				returnedUniqueString := "dummyUniqueString"
				fakeUniqueStringFactory := new(uniquestring.Fake)
				fakeUniqueStringFactory.ConstructReturns(returnedUniqueString)

				objectUnderTest := newParallelCaller(fakeCaller, new(pubsub.Fake), fakeUniqueStringFactory)

				/* act */
				actualError := objectUnderTest.Call(
					providedCallId,
					providedInboundScope,
					providedRootOpId,
					providedPkgRef,
					providedScgParallelCalls,
				)

				/* assert */
				actualScgParallelCalls := []*model.Scg{}
				for callIndex := range providedScgParallelCalls {
					actualNodeId,
						actualChildOutboundScope,
						actualScg,
						actualPkgRef,
						actualRootOpId := fakeCaller.CallArgsForCall(callIndex)
					Expect(actualNodeId).To(Equal(returnedUniqueString))
					Expect(actualChildOutboundScope).To(Equal(providedInboundScope))
					Expect(actualPkgRef).To(Equal(providedPkgRef))
					Expect(actualRootOpId).To(Equal(providedRootOpId))
					actualScgParallelCalls = append(actualScgParallelCalls, actualScg)
				}
				Expect(actualScgParallelCalls).To(ConsistOf(providedScgParallelCalls))
				Expect(actualError).To(BeNil())
			})
		})
	})
})
