package main

import "fmt"

// Node merepresentasikan Rak atau Buku
type Node struct {
	ID    int    // ID unik buku/rak
	Label string // Nama Kategori atau Judul Buku
	Left  *Node  // Sub-kategori Kiri (Arah Fiksi)
	Right *Node  // Sub-kategori Kanan (Arah Non-Fiksi)
}

// 1. ALGORITMA REKURSIF
func PreorderRecursive(root *Node) {
	if root == nil {
		return
	}
	// Operasi Dasar: Catat Kategori/Buku saat ini
	fmt.Printf("[%s] ", root.Label)

	// Masuk ke lorong Kiri (Fiksi)
	PreorderRecursive(root.Left)
	// Masuk ke lorong Kanan (Non-Fiksi)
	PreorderRecursive(root.Right)
}

// 2. ALGORITMA ITERATIF
func PreorderIterative(root *Node) {
	if root == nil {
		return
	}
	stack := []*Node{root}
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Operasi Dasar: Catat
		fmt.Printf("[%s] ", curr.Label)

		// Masukkan Kanan dulu (Non-Fiksi) ke tumpukan
		// Agar Kiri (Fiksi) diambil duluan di putaran berikutnya
		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
	}
}

// 3. GENERATOR POHON BUKU
func BuildTree(n int) *Node {
	if n == 0 {
		return nil
	}

	// Root selalu menjadi Lobi Utama
	root := &Node{ID: 1, Label: "LOBI UTAMA"}

	if n == 1 {
		return root
	}

	for i := 2; i <= n; i++ {
		// Tentukan nama label berdasarkan ID agar sesuai skenario
		var namaLabel string
		
		if i == 2 {
			namaLabel = "RAK FIKSI"
		} else if i == 3 {
			namaLabel = "RAK NON-FIKSI"
		} else {
			// Untuk ID > 3, kita anggap buku biasa
			// Kita beri label generic "Buku-ID"
			namaLabel = fmt.Sprintf("Buku-%d", i)
		}

		newNode := &Node{ID: i, Label: namaLabel}
		insertRoot(root, newNode)
	}

	return root
}

// Fungsi pembantu untuk memasukkan node (Level Order Insertion)
func insertRoot(root *Node, newNode *Node) {
	antrian := []*Node{root}

	for len(antrian) > 0 {
		temp := antrian[0]
		antrian = antrian[1:]

		// Cek Kiri (Jalur Fiksi)
		if temp.Left == nil {
			temp.Left = newNode
			return
		} else {
			antrian = append(antrian, temp.Left)
		}

		// Cek Kanan (Jalur Non-Fiksi)
		if temp.Right == nil {
			temp.Right = newNode
			return
		} else {
			antrian = append(antrian, temp.Right)
		}
	}
}

func main() {
	fmt.Println("=== SISTEM PERPUSTAKAAN: FIKSI (Kiri) vs NON-FIKSI (Kanan) ===")
	fmt.Println("Metode: Preorder Traversal (Cek Induk -> Cek Anggota)")
	fmt.Println("============================================================")

	n := 7
	root := BuildTree(n)

	fmt.Printf("Simulasi Audit dengan %d item:\n", n)
	
	fmt.Print("\n[REKURSIF] Pustakawan A Mencatat:\n")
	PreorderRecursive(root)
	fmt.Println()

	fmt.Print("\n[ITERATIF] Pustakawan B Mencatat:\n")
	PreorderIterative(root)
	fmt.Println()
	
	fmt.Println("\n------------------------------------------------------------")
}