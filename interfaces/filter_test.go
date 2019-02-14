package interfaces

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	applicationMocks "flamingo.me/csrf/application/mocks"
	"flamingo.me/flamingo/v3/framework/web"
	"github.com/stretchr/testify/suite"
)

type (
	CsrfFilterTestSuite struct {
		suite.Suite

		filter  *CsrfFilter
		service *applicationMocks.Service
		chain   *web.FilterChain

		context        context.Context
		webRequest     *web.Request
		responseWriter http.ResponseWriter
	}
)

func TestCsrfFilterTestSuite(t *testing.T) {
	suite.Run(t, &CsrfFilterTestSuite{})
}

func (t *CsrfFilterTestSuite) SetupSuite() {
	t.context = context.Background()
	t.responseWriter = httptest.NewRecorder()
	t.webRequest = web.CreateRequest(nil, nil)
}

func (t *CsrfFilterTestSuite) SetupTest() {
	t.service = &applicationMocks.Service{}

	t.filter = &CsrfFilter{}
	t.filter.Inject(&web.Responder{}, t.service)

	//t.chainFilter = &routerMocks.Filter{}
	//t.chain = &web.FilterChain{
	//	Filters: []web.Filter{
	//		t.chainFilter,
	//	},
	//}
}

func (t *CsrfFilterTestSuite) TearDown() {
	t.service.AssertExpectations(t.T())
	//t.chainFilter.AssertExpectations(t.T())
	//t.chainFilter = nil
	t.chain = nil
	t.service = nil
}

func (t *CsrfFilterTestSuite) TestFilter_WrongToken() {
	t.service.On("IsValid", t.webRequest).Return(false).Once()

	response := t.filter.Filter(t.context, t.webRequest, t.responseWriter, t.chain)
	forbidden, ok := response.(*web.ServerErrorResponse)
	t.True(ok)
	t.Equal(uint(http.StatusForbidden), forbidden.Status)
}

func (t *CsrfFilterTestSuite) TestFilter_Success() {
	//t.chainFilter.On("Filter", t.context, t.webRequest, t.responseWriter, t.chain).Return(&web.BasicResponse{}).Once()
	t.service.On("IsValid", t.webRequest).Return(true).Once()

	//response := t.filter.Filter(t.context, t.webRequest, t.responseWriter, t.chain)
	//t.IsType(&web.BasicResponse{}, response)
}
