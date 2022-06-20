package responseModelStub

import commonModel "github.com/6reduk/smsaeroclient/smsaero/dispatcher/common/model"

func GetPaginator() *commonModel.Paginator {
	links := commonModel.NewPaginator()
	links.Self = "/test/self?page=1"
	links.Prev = ""
	links.Next = "/test/next?page=2"
	links.First = "/test/first?page=1"
	links.Last = "/test/last?page=5"
	return links
}

func GetPaginatorOnLastPage() *commonModel.Paginator {
	links := commonModel.NewPaginator()
	links.Self = "/test/self?page=1"
	links.Prev = ""
	links.Next = ""
	links.First = "/test/first?page=1"
	links.Last = "/test/last?page=1"
	return links
}
