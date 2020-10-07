package main

import "fmt"

type bstnode struct {
	name  string
	left  *bstnode
	right *bstnode
}

type bst struct {
	root *bstnode
}

func (b *bst) reset() {
	b.root = nil
}

func (b *bst) insert(name string) {
	b.insertRec(b.root, name)
}

func (b *bst) insertRec(node *bstnode, name string) *bstnode {
	if b.root == nil {
		b.root = &bstnode{
			name: name,
		}
		return b.root
	}
	if node == nil {
		return &bstnode{name: name}
	}
	if name <= node.name {
		node.left = b.insertRec(node.left, name)
	}
	if name > node.name {
		node.right = b.insertRec(node.right, name)
	}
	return node
}

func (b *bst) find(name string) error {
	node := b.findRec(b.root, name)
	if node == nil {
		return fmt.Errorf("Name: %s not found in tree", name)
	}
	fmt.Printf("%s is found in tree\n", name)
	return nil
}

func (b *bst) findRec(node *bstnode, name string) *bstnode {
	if node == nil {
		return nil
	}
	if node.name == name {
		return b.root
	}
	if name < node.name {
		return b.findRec(node.left, name)
	}
	return b.findRec(node.right, name)
}

func (b *bst) inorder() {
	b.inorderRec(b.root)
}

func (b *bst) inorderRec(node *bstnode) {
	if node != nil {
		b.inorderRec(node.left)
		fmt.Println(node.name)
		b.inorderRec(node.right)
	}
}

func (b *bst) preorder() {
	b.preorderRec(b.root)
}

func (b *bst) preorderRec(node *bstnode) {
	if node != nil {
		fmt.Println(node.name)
		b.inorderRec(node.left)
		b.inorderRec(node.right)
	}
}

// Min returns the Item with min value stored in the tree
func (b *bst) Min() *bstnode {
	n := b.root
	if n == nil {
		return nil
	}
	for {
		if n.left == nil {
			return b.root
		}
		n = n.left
	}
}

func main() {
	bst := &bst{}
	eg := []string{"John", "Wick", "Mary", "Jane", "Michael", "Angelo", "Tom", "Paul", "Jaina"}
	for _, name := range eg {
		bst.insert(name)
	}

	fmt.Printf("Printing Inorder:\n")
	bst.inorder()
	fmt.Println()

	fmt.Printf("\nPrinting Preorder:\n")
	bst.preorder()
	fmt.Println()

	fmt.Printf("\nFinding Values:\n")
	ex := []string{"Mages", "Wick", "Freud", "Jane", "Michael", "Splinter", "Tom", "Jerry", "Jaina"}
	for _, name := range ex {
		bst.find(name)
	}

	fmt.Printf("\nFinding Min:\n")
	bst.Min()

}
