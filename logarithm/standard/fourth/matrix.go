package fourth

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

func (p *point) smaller(other *point) bool {
	return p.x < other.x && p.y < other.y
}

func printEdge1(edge [][]int) {
	if len(edge) < 0 {
		fmt.Println("empty edge")
		return
	}

	lastLen := len(edge[0])
	for i := 1; i < len(edge); i++ {
		if lastLen != len(edge[i]) {
			fmt.Println("invalid edge")
			return
		}
	}

	leftTop := &point{
		0, 0,
	}
	rightBottom := &point{
		len(edge) - 1, len(edge[0]) - 1,
	}
	for leftTop.smaller(rightBottom) {
		for i := leftTop.y; i < rightBottom.y; i++ {
			fmt.Printf("%d ", edge[leftTop.x][i])
		}
		for i := leftTop.x; i < rightBottom.x; i++ {
			fmt.Printf("%d ", edge[i][rightBottom.y])
		}
		for i := rightBottom.y; i > leftTop.y; i-- {
			fmt.Printf("%d ", edge[rightBottom.x][i])
		}
		for i := rightBottom.x; i > leftTop.x; i-- {
			fmt.Printf("%d ", edge[i][leftTop.y])
		}
		leftTop.addBottom(1).addRight(1)
		rightBottom.addLeft(1).addTop(1)
	}
	fmt.Println()
}
