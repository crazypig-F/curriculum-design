package main

import (
	"HuffmanCoding/data"
	"HuffmanCoding/tree"
	"fmt"
	"github.com/fzipp/canvas"
	"image/color"
	"io/ioutil"
	"log"
	"math"
	"strconv"
)

const (
	radio = 30
	k = 6
	b = 0
)


func DrawLine(ctx *canvas.Context, mx, xy, lx, ly float64) {
	ctx.BeginPath()
	ctx.MoveTo(mx, xy)
	ctx.LineTo(lx, ly)
	ctx.Stroke()
}

func DrawNode(ctx *canvas.Context, x, y float64) {
	ctx.BeginPath()
	ctx.Arc(x, y, radio, 0, 2*math.Pi, false)
	ctx.Stroke()
}

func PutText(ctx *canvas.Context, text string, x, y float64) {
	ctx.FillText(text, x, y+5)
}

func DrawTree(ctx *canvas.Context, root *tree.HuffmanTreeNode, nx, ny float64, deep float64) {
	if root == nil {
		return
	}
	DrawNode(ctx, nx, ny)
	PutText(ctx, root.Val+"_"+strconv.Itoa(root.Number), nx, ny)
	if root.Left != nil {
		DrawLine(ctx, nx, ny+radio, nx-k*deep-b, ny+50)
	}
	if root.Right != nil {
		DrawLine(ctx, nx, ny+radio, nx+k*deep+b, ny+50)
	}
	DrawTree(ctx, root.Left, nx-k*deep-b, ny+radio+50, deep/2)
	DrawTree(ctx, root.Right, nx+k*deep+b, ny+radio+50, deep/2)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func GetDeep(root *tree.HuffmanTreeNode) int {
	if root == nil {
		return 0
	}
	return max(GetDeep(root.Left)+1, GetDeep(root.Right)+1)
}

func PrintTree(ctx *canvas.Context) {
	file1, err := ioutil.ReadFile("data/inputFile.txt")
	if err != nil {
		fmt.Print(err)
	}
	queue := tree.ChangeNode(data.CountCharacter(file1, "data/hfmTree.txt"))
	root := tree.CreateTree(queue)
	ctx.SetFillStyle(color.RGBA{R: 200, A: 255})
	ctx.SetStrokeStyle(color.RGBA{G: 200, A: 255})
	ctx.SetFont("13px Arial")
	ctx.SetTextAlign(canvas.AlignCenter)
	DrawTree(ctx, root, 800, 50, math.Pow(2, float64(GetDeep(root))))
	ctx.Flush()
}

func main() {
	err := canvas.ListenAndServe(":8080", PrintTree,
		canvas.Size(2500, 2500),
		canvas.Title("Drawing"),
	)
	if err != nil {
		log.Fatal(err)
	}
}
