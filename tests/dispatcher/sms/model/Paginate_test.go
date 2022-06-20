package model

import (
	"fmt"
	commonModel "github.com/6reduk/smsaeroclient/internal/dispatcher/common/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type PaginateShould struct {
	suite.Suite
}

func TestPaginateShould(t *testing.T) {
	suite.Run(t, &PaginateShould{})
}

func (s *PaginateShould) TestPaginate_MissingPaginateLinks_ReturnError() {
	links := commonModel.NewPaginator()

	currentPage, err := links.CurrentPage()
	require.NotNil(s.T(), err)
	assert.Equal(s.T(), 0, currentPage)
	assert.Equal(s.T(), commonModel.ErrPaginateIsMissing, err)

	nextPage, err := links.NextPage()
	require.NotNil(s.T(), err)
	assert.Equal(s.T(), 0, nextPage)
	assert.Equal(s.T(), commonModel.ErrPaginateIsMissing, err)

	prevPage, err := links.PrevPage()
	require.NotNil(s.T(), err)
	assert.Equal(s.T(), 0, prevPage)
	assert.Equal(s.T(), commonModel.ErrPaginateIsMissing, err)

	firstPage, err := links.FirstPage()
	require.NotNil(s.T(), err)
	assert.Equal(s.T(), 0, firstPage)
	assert.Equal(s.T(), commonModel.ErrPaginateIsMissing, err)

	lastPage, err := links.LastPage()
	require.NotNil(s.T(), err)
	assert.Equal(s.T(), 0, lastPage)
	assert.Equal(s.T(), commonModel.ErrPaginateIsMissing, err)
}

func (s *PaginateShould) TestPaginate_ValidPaginateLinks_ReturnError() {
	links := commonModel.NewPaginator()
	urlTemplate := "/v2/api?page=%d"
	expectedCurrentPage := 2
	expectedNextPage := 3
	expectedPrevPage := 1
	expectedFirstPage := 1
	expectedLastPage := 3

	links.Self = fmt.Sprintf(urlTemplate, expectedCurrentPage)
	links.Next = fmt.Sprintf(urlTemplate, expectedNextPage)
	links.Prev = fmt.Sprintf(urlTemplate, expectedPrevPage)
	links.First = fmt.Sprintf(urlTemplate, expectedFirstPage)
	links.Last = fmt.Sprintf(urlTemplate, expectedLastPage)

	currentPage, err := links.CurrentPage()
	require.NoError(s.T(), err)
	assert.Equal(s.T(), expectedCurrentPage, currentPage)

	nextPage, err := links.NextPage()
	require.NoError(s.T(), err)
	assert.Equal(s.T(), expectedNextPage, nextPage)

	prevPage, err := links.PrevPage()
	require.NoError(s.T(), err)
	assert.Equal(s.T(), expectedPrevPage, prevPage)

	firstPage, err := links.FirstPage()
	require.NoError(s.T(), err)
	assert.Equal(s.T(), expectedFirstPage, firstPage)

	lastPage, err := links.LastPage()
	require.NoError(s.T(), err)
	assert.Equal(s.T(), expectedLastPage, lastPage)
}
