package forth

import "fmt"

/**
同矩阵相关的算法题，需要注意找到宏观视角，从坐标变换，点位移动的角度入手解决问题
*/

/**
1.假设存在矩阵
	 1  2  3
	 8  9  4
	 7  6  5
按顺时针的方式将矩阵打印
取左上角的坐标a(0,0),右下角的坐标b(n-1,n-1)
（1）从a出发，固定行号0，列号从0到n-1打印
（2）固定列号n-1，行号从0到n-1打印
 (3) 固定行号n-1，列号从n-1到0打印
（4）固定列号n-1，行号从n-1到0打印
 a向右下角迁移1格(1,1)，b向右上角迁移1格(n-2,n-2)。直到a的行/列都大于b的行/列，循环结束
*/
type point struct {
	x int
	y int
}

func (p *point) addLeft(x int) *point {
	p.x -= x
	return p
}

func (p *point) addTop(y int) *point {
	p.y -= y
	return p
}

func (p *point) addRight(x int) *point {
	p.x += x
	return p
}

func (p *point) addBottom(y int) *point {
	p.y += y
	return p
}

func (p *point) 小于(other *point) bool {
	return p.x < other.x && p.y < other.y
}

func printEdge1(edge [][]int) {
	leftTop := &point{
		0, 0,
	}
	rightBottom := &point{
		len(edge) - 1, len(edge) - 1,
	}
	for leftTop.小于(rightBottom) {

		for i := leftTop.y; i < rightBottom.y; i++ {
			fmt.Print(edge[leftTop.x][i])
			fmt.Print(" ")
		}
		fmt.Print("\n")

		for i := leftTop.x; i < rightBottom.x; i++ {
			fmt.Print(edge[i][rightBottom.y])
			fmt.Print(" ")
		}
		fmt.Print("\n")

		for i := rightBottom.y; i > leftTop.y; i-- {
			fmt.Print(edge[rightBottom.x][i])
			fmt.Print(" ")
		}

		fmt.Print("\n")
		for i := rightBottom.x; i > leftTop.x; i-- {
			fmt.Print(edge[i][leftTop.y])
			fmt.Print(" ")
		}
		fmt.Print("\n")

		leftTop.addBottom(1).addRight(1)

		rightBottom.addLeft(1).addTop(1)
	}
}
