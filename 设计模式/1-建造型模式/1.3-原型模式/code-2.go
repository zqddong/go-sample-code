package main

import (
	"bytes"
	"fmt"
)

//是一个类似 DOM 树对象的例子，因为 DOM 对象往往层级会很深，
//那么要创建类似的DOM树的时候能让我们更好的理解原型模式的优势

// Node a document object model node
type Node interface {
	// String return nodes text representation
	String() string
	// Parent returns the node parent
	Parent() Node
	// SetParent sets the node parent
	SetParent(node Node)
	// Children returns the node children nodes
	Children() []Node
	// AddChild adds a child node
	AddChild(child Node)
	// Clone clones a node
	Clone() Node
}

type Element struct {
	text     string
	parent   Node
	children []Node
}

func NewElement(text string) *Element {
	return &Element{
		text:     text,
		parent:   nil,
		children: make([]Node, 0),
	}
}

func (e *Element) Parent() Node {
	return e.parent
}

func (e *Element) SetParent(node Node) {
	e.parent = node
}

func (e *Element) Children() []Node {
	return e.children
}

func (e *Element) AddChild(child Node) {
	cp := child.Clone()
	cp.SetParent(e)
	e.children = append(e.children, cp)
}

func (e *Element) Clone() Node {
	cp := &Element{
		text:     e.text,
		parent:   nil,
		children: make([]Node, 0),
	}

	for _, child := range e.children {
		cp.AddChild(child)
	}
	return cp
}

func (e *Element) String() string {
	buffer := bytes.NewBufferString(e.text)
	for _, c := range e.Children() {
		text := c.String()
		fmt.Fprintf(buffer, "\n %s", text)
	}

	return buffer.String()
}

func main() {
	// 职级节点 总监
	directorNode := NewElement("Director of Engineering")

	// 职级节点 研发经理
	engManagerNode := NewElement("Engineering Manager")
	engManagerNode.AddChild(NewElement("Lead Software Engineer"))

	// 研发经理是总监的下级
	directorNode.AddChild(engManagerNode)
	directorNode.AddChild(engManagerNode)

	// 办公室经理也是总监的下级
	officeManagerNode := NewElement("Office Manager")
	directorNode.AddChild(officeManagerNode)

	fmt.Println("")
	fmt.Println("# Company Hierarchy")
	fmt.Print(directorNode)
	fmt.Println("")

	// 从研发经理节点克隆出一颗新的树
	fmt.Println("# Team Hiearachy")
	fmt.Print(engManagerNode.Clone())
}
