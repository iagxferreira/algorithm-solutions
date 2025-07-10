package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func NewNode(data int) *Node {
	return &Node{Data: data}
}

func (node *Node) Insert(value int) {
	if value < node.Data {
		if node.Left == nil {
			node.Left = NewNode(value)
		} else {
			node.Left.Insert(value)
		}
	} else {
		if node.Right == nil {
			node.Right = NewNode(value)
		} else {
			node.Right.Insert(value)
		}
	}
}

func LevelOrder(root *Node, writer *bufio.Writer) {
	if root == nil {
		return
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		fmt.Fprintf(writer, "%d ", current.Data)

		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	writer := bufio.NewWriterSize(os.Stdout, 16*1024*1024)

	// Read number of nodes
	nStr := readLine(reader)
	n, err := strconv.Atoi(strings.TrimSpace(nStr))
	checkError(err)

	// Read node values line
	valuesLine := readLine(reader)
	values := strings.Fields(valuesLine)
	if len(values) == 0 {
		// No nodes to insert
		writer.Flush()
		return
	}

	// Parse first value and create root
	firstVal, err := strconv.Atoi(values[0])
	checkError(err)
	root := NewNode(firstVal)

	// Insert remaining values
	for i := 1; i < n; i++ {
		val, err := strconv.Atoi(values[i])
		checkError(err)
		root.Insert(val)
	}

	// Print level order traversal
	// TODO: change this call, and array should be returned and then printed to writer.
	LevelOrder(root, writer)
	writer.WriteString("\n")
	writer.Flush()
}
