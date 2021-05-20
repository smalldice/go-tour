package main

import "fmt"

type Vertex struct {
	x, y int
}

var (
	v1 = Vertex{1, 2}  // 创建一个Vertex 类型的结构体
	v2 = Vertex{x: 1}  // y: 0 被隐式地赋予
	v3 = Vertex{}      // x: 0 y:0 被隐士地赋予
	p  = &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体
)

func main() {
	fmt.Println(v1, v2, v3, p)

	c := p

	c.x = 3

	fmt.Println(c.x, c.y, p.x)
}
