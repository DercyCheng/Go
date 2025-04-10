package interview

import (
	"sort"
	"strings"
)

// FileNode 文件系统节点结构
type FileNode struct {
	isFile   bool                 // 是否为文件，true表示文件，false表示目录
	content  string               // 文件内容，仅在isFile为true时有效
	children map[string]*FileNode // 子节点映射表，用于存储目录下的文件和子目录
}

// FileSystem 文件系统结构
type FileSystem struct{ 
	root *FileNode // 根节点
}

// NewFileSystem 创建新的文件系统实例
func NewFileSystem() FileSystem {
	return FileSystem{&FileNode{children: make(map[string]*FileNode)}}
}

// split 辅助函数：将路径字符串分割成路径数组
func split(path string) []string {
	if path == "/" {
		return []string{} // 根路径返回空数组
	}
	parts := strings.Split(path, "/")[1:] // 分割路径，移除第一个空字符串
	if parts[len(parts)-1] == "" {
		parts = parts[:len(parts)-1] // 移除末尾的空字符串（处理末尾斜杠的情况）
	}
	return parts
}

// Ls 列出指定路径下的所有文件和目录
func (fs *FileSystem) Ls(path string) []string {
	node := fs.root
	parts := split(path)

	for _, p := range parts { // 遍历到目标节点
		node = node.children[p]
	}

	if node.isFile { // 文件路径返回文件名
		return []string{parts[len(parts)-1]}
	}

	res := make([]string, 0, len(node.children))
	for k := range node.children {
		res = append(res, k)
	}
	sort.Strings(res)
	return res
}

// Mkdir 创建目录（支持递归创建）
func (fs *FileSystem) Mkdir(path string) {
	node := fs.root
	for _, p := range split(path) {
		if _, ok := node.children[p]; !ok {
			node.children[p] = &FileNode{children: make(map[string]*FileNode)}
		}
		node = node.children[p]
	}
}

// AddContentToFile 向文件添加内容
// 如果文件不存在则创建新文件，如果文件存在则追加内容
func (fs *FileSystem) AddContentToFile(filePath, content string) {
	parts := split(filePath)
	node := fs.root
	fileName := parts[len(parts)-1]

	// 遍历到父目录
	for _, p := range parts[:len(parts)-1] {
		node = node.children[p]
	}

	// 创建/追加文件
	if node.children[fileName] == nil {
		node.children[fileName] = &FileNode{isFile: true}
	}
	node.children[fileName].content += content
}

// ReadContentFromFile 读取文件内容
// 返回指定文件的完整内容
func (fs *FileSystem) ReadContentFromFile(filePath string) string {
	node := fs.root
	for _, p := range split(filePath) {
		node = node.children[p]
	}
	return node.content
}
