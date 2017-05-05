package consumenodeapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang-interfaces/vhttp"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opspec-io/sdk-golang/model"
	"net/http"
)

var _ = Describe("KillOp", func() {

	It("should call httpClient.Do() with expected args", func() {

		/* arrange */
		providedKillOpReq := model.KillOpReq{
			OpId: "dummyRootOpId",
		}

		expectedBytes, _ := json.Marshal(providedKillOpReq)

		expectedHttpReq, _ := http.NewRequest(
			"POST",
			fmt.Sprintf("http://%v/ops/kills", "localhost:42224"),
			bytes.NewBuffer(expectedBytes),
		)

		fakeHttpClient := new(vhttp.Fake)

		objectUnderTest := consumeNodeApi{
			httpClient: fakeHttpClient,
		}

		/* act */
		objectUnderTest.KillOp(providedKillOpReq)

		/* assert */
		Expect(expectedHttpReq).To(Equal(fakeHttpClient.DoArgsForCall(0)))

	})
})
