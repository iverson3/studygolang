package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"log"
	"os"
	"sort"
)

// defelate 压缩算法

const (
	filePath = ""
	outPutPath = ""
	depressPath = ""
	contentBufferSize = 5000000
)

type TreeNode struct {
	Val int     // 每一个字节
	Times int   // 字节出现的次数
	Left *TreeNode
	Right *TreeNode
}

type treeHeap []*TreeNode

func (p treeHeap) Less(i, j int) bool {
	return p[i].Times <= p[j].Times
}

func (p treeHeap) Len() int {
	return len(p)
}

func (p treeHeap) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *treeHeap) Push(node interface{}) {
	*p = append(*p, node.(*TreeNode))
}

func (p *treeHeap) Pop() interface{} {
	n := len(*p)
	node := (*p)[n-1]
	*p = (*p)[:n-1]
	return node
}

func main() {
	HuffmanEncoding(filePath, outPutPath)
}

func HuffmanEncoding(filePath, outPath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	imap := getFrequencyMap(file)
	if len(imap) == 0 {
		log.Fatal("file content is empty")
		return
	}

	plist := make(treeHeap, 0)
	for k, v := range imap {
		plist = append(plist, &TreeNode{Val: k, Times: v})
	}
	// 按照频率进行排序
	sort.Sort(plist)

	// 初始化哈夫曼树
	hTree := initHuffmanTree(plist)

	encodeMap := make(map[int]string)
	// 遍历哈夫曼树,生成哈夫曼编码表(正表，用于编码),key(ASSCII),value(路径痕迹)
	createEncodingTable(hTree, encodeMap)

	// 将输入文件的字符通过码表编码，输出到另一个文件, 压缩模块完成
	encoding(filePath, outPath, encodeMap)
}

// 读取文件内容，为每个字节统计其出现的次数
func getFrequencyMap(file *os.File) map[int]int {
	imap := make(map[int]int)
	reader := bufio.NewReader(file)
	buf := make([]byte, contentBufferSize)
	count, _ := reader.Read(buf)
	for i := 0; i < count; i++ {
		imap[int(buf[i])]++
	}
	return imap
}

// 构建哈夫曼树
// 使用优先队列构建最小路径权值哈夫曼树
func initHuffmanTree(plist treeHeap) *TreeNode {
	// 使用递归也可以生成哈夫曼树，但是不能保证树的结构是最优的，我们需要保证构建的哈夫曼树具有最小权值路径和，这样才能使得压缩的效率最大化
	// 从代码可以清晰的看到，实现了堆接口后的哈夫曼树构造非常简单。因为heap是优先队列的缘故，每次插入都会按Times（频次）升序排序保证下次合成的两两节点权值之和是所有节点中最小的
	heap.Init(&plist)
	for plist.Len() > 1 {
		node1 := heap.Pop(&plist).(*TreeNode)
		node2 := heap.Pop(&plist).(*TreeNode)
		root := &TreeNode{Times: node1.Times + node2.Times}
		if node1.Times > node2.Times {
			root.Right, root.Left = node1, node2
		} else {
			root.Right, root.Left = node2, node1
		}
		heap.Push(&plist, root)
	}
	return plist[0]
}

// 生成码表，考虑到性能必须要生成map, key(int对应ASSCII????，string对应bit编码，后续转成bit)
// encodeMap中key为字符的ASCII码值，用int表示，value为哈夫曼树从根节点到叶子节点的左右路径编码值，用string表示，其中tmp记录递归过程中所走过的路径，往左走用'0'表示，往右走用'1'表示，最后再转换成string就得到了对应ASCII码的编码值。
func createEncodingTable(node *TreeNode, encodeMap map[int]string) {
	// 回溯遍历二叉树，每往左加一个0 每往右加一个1 到达叶子节点则将经过的路径(01串)转成string存入码表
	treePath := make([]byte, 0)
	var depth func(treeNode *TreeNode)
	depth = func(root *TreeNode) {
		if root == nil {
			return
		}

		// 当前节点是叶子节点
		if root.Left == nil && root.Right == nil {
			encodeMap[root.Val] = string(treePath)
			return
		}

		// 当前节点是普通节点，则进行左右回溯遍历 (先左后右)
		// 向左则加0
		treePath = append(treePath, '0')
		depth(root.Left)

		// 向右则加1 (将最后一个字节修改为1)
		treePath[len(treePath) - 1] = '1'
		depth(root.Right)

		// 移除最后一个0/1
		treePath = treePath[:len(treePath) - 1]
	}

	depth(node)
}

func encoding(inPath, outPath string, encodeMap map[int]string) {
	inFile, err := os.Open(inPath)
	if err != nil {
		return
	}
	defer inFile.Close()

	reader := bufio.NewReader(inFile)
	fileContent := make([]byte, contentBufferSize)
	count, _ := reader.Read(fileContent)

	var buff bytes.Buffer
	for i := 0; i < count; i++ {
		v := fileContent[i]
		// code是01序列，即树的路径
		if code, ok := encodeMap[int(v)]; len(code) != 0 && ok {
			buff.WriteString(code)
		}
	}

	var buf byte
	res := make([]byte, 0)
	for idx, bit := range buff.Bytes() {
		// 每八个位用一个byte表示，结果放入res切片中
		pos := idx % 8
		if pos == 0 && idx > 0 {
			res = append(res, buf)
			buf = 0
		}

		if bit == '1' {
			buf |= 1 << pos
		}
	}

	//left := buff.Len() % 8
	res = append(res, buf)

	outFile, err := os.OpenFile(outPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer outFile.Close()

	writeCount, err := outFile.Write(res)
	if err != nil {
		fmt.Println(writeCount)
		return
	}
}