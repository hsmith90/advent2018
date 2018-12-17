package day8

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

var nodeList []*Node

func TestPartA(t *testing.T) {
	input, _ := ioutil.ReadFile("input")
	var data []int
	strings := strings.Split(string(input), " ")

	for _, s := range strings {
		d, _ := strconv.Atoi(s)
		data = append(data, d)
	}

	CreateNode(data, &Node{})

	var totalMeta int
	for _, n := range nodeList {
		for _, m := range n.metadata {
			totalMeta += m
		}
	}

	fmt.Printf("Part A: %v\n", totalMeta)

	value := nodeList[0].NodeValue()

	fmt.Printf("Part B: %v\n", value)
}

type Node struct {
	childNodeCount int
	metadataCount  int
	metadata       []int
	childNodes     []*Node
	parentNode     *Node
}

func CreateNode(d []int, parent *Node) (tail []int) {
	header := d[:2]
	tail = d[2:]

	newNode := Node{}
	newNode.childNodeCount = header[0]
	newNode.metadataCount = header[1]
	nodeList = append(nodeList, &newNode)

	if parent.childNodeCount != 0 {
		newNode.parentNode = parent
		parent.childNodes = append(parent.childNodes, &newNode)
	}

	for c := 0; c < newNode.childNodeCount; c++ {
		tail = CreateNode(tail, &newNode)
	}

	for m := 0; m < newNode.metadataCount; m++ {
		newNode.metadata = append(newNode.metadata, tail[0])
		tail = tail[1:]
	}

	return tail

}

func (n *Node) NodeValue() (value int) {
	if n.childNodeCount == 0 {
		for _, m := range n.metadata {
			value += m
		}
	} else {
		for _, m := range n.metadata {
			if m <= n.childNodeCount {
				value += n.childNodes[m-1].NodeValue()
			}
		}
	}

	return value
}
