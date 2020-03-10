/**
借助stack实现浏览器前进、后退功能
*/
package _3_stack

type Browser struct {
	backStack    *linkedListStack
	forwardStack *linkedListStack
	local        interface{}
}

func NewBrowser() *Browser {
	return &Browser{
		backStack:    NewLinkListStack(),
		forwardStack: NewLinkListStack(),
	}
}

// 是否可以后退
func (browser *Browser) CanBack() bool {
	return !browser.backStack.IsEmpty()
}

// 是否可以前进
func (browser *Browser) CanForward() bool {
	return !browser.forwardStack.IsEmpty()
}

// 打开新网页
func (browser *Browser) Open(addr interface{}) {
	if browser.local != nil {
		browser.backStack.Push(browser.local)
	}
	browser.local = addr
	if !browser.forwardStack.IsEmpty() {
		browser.forwardStack.Flush()
	}
}

// 后退
func (browser *Browser) Back() {
	if !browser.CanBack() {
		return
	}

	browser.forwardStack.Push(browser.local)
	browser.backStack.Pop()
	browser.local = browser.backStack.Top()
}

// 前进
func (browser *Browser) Forward() {
	if !browser.CanForward() {
		return
	}
	browser.backStack.Push(browser.local)
	browser.forwardStack.Pop()
	browser.local = browser.forwardStack.Top()
}
