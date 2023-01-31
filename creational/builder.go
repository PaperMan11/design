package creational

// 创建者模式: 将复杂对象的构建分成多个简单对象的构建

// Builder 是生成器接口
type Builder interface {
	Part1()
	Part2()
	Part3()
}

type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

// 建造产品
func (d *Director) Construct() {
	d.builder.Part1()
	d.builder.Part2()
	d.builder.Part3()
}

// --------------------------------------------------------------

type Builder1 struct {
	res string
}

func (b *Builder1) Part1() {
	b.res += "1"
}

func (b *Builder1) Part2() {
	b.res += "2"
}

func (b *Builder1) Part3() {
	b.res += "3"
}

func (b *Builder1) GetRes() string {
	return b.res
}

// --------------------------------------------------------------

type Builder2 struct {
	res int
}

func (b *Builder2) Part1() {
	b.res += 1
}

func (b *Builder2) Part2() {
	b.res += 2
}

func (b *Builder2) Part3() {
	b.res += 3
}

func (b *Builder2) GetRes() int {
	return b.res
}
