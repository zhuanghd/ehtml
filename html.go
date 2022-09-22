package ehtml

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

// Node html 节点
type Node struct {
	selection *goquery.Selection
}

// SingleNode 根据html文本查询单个子级元素
func SingleNode(content string, sel string) *Node {
	all := Nodes(content, sel)
	if all == nil || len(all) == 0 {
		return nil
	} else {
		return all[0]
	}
}

// Nodes 根据html文本查询子级元素
func Nodes(content string, selector string) []*Node {
	els := make([]*Node, 0)
	if content == "" {
		return els
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil || doc == nil {
		return els
	}
	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		els = append(els, &Node{
			selection: s,
		})
	})
	return els
}

// SingleNode 查询某节点的单个子级元素
func (e *Node) SingleNode(selector string) *Node {
	all := e.Nodes(selector)
	if all == nil || len(all) == 0 {
		return nil
	} else {
		return all[0]
	}
}

// Nodes 查询某节点的子级元素列表
func (e *Node) Nodes(selector string) []*Node {
	els := make([]*Node, 0)
	e.selection.Find(selector).Each(func(i int, s *goquery.Selection) {
		els = append(els, &Node{
			selection: s,
		})
	})
	return els
}

// Parent 获取某节点的父级元素
func (e *Node) Parent() *Node {
	return &Node{
		selection: e.selection.Parent(),
	}
}

// Child 获取某节点的单个直接子级元素
func (e *Node) Child() *Node {
	all := e.Children()
	if all == nil || len(all) == 0 {
		return nil
	} else {
		return all[0]
	}
}

// Children 获取某节点的直接子级元素列表
func (e *Node) Children() []*Node {
	els := make([]*Node, 0)
	e.selection.Children().Each(func(i int, s *goquery.Selection) {
		els = append(els, &Node{
			selection: s,
		})
	})
	return els
}

// Attr 获取某节点的属性值
func (e *Node) Attr(name string) string {
	val, _ := e.selection.Attr(name)
	return strings.TrimSpace(val)
}

// AttrOr 获取某节点的属性值，如果属性为空则返回默认值
func (e *Node) AttrOr(name string, defaultValue string) string {
	val := e.Attr(name)
	if val == "" {
		return defaultValue
	} else {
		return val
	}
}

// Text 获取某节点的内部文本
func (e *Node) Text() string {
	return strings.TrimSpace(e.selection.Text())
}

// TextOr 获取某节点的内部文本，如果文本为空则返回默认值
func (e *Node) TextOr(defaultValue string) string {
	ret := e.Text()
	if ret == "" {
		return defaultValue
	} else {
		return ret
	}
}
