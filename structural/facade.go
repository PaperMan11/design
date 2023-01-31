package structural

import "fmt"

/*	外观模式
 *	优点：
 *		1. 它对客户端屏蔽了子系统组件，减少了客户端所需处理的对象数目，并使得子系统使用起来更加容易。
 *		   通过引入外观模式，客户端代码将变得很简单，与之关联的对象也很少。
 * 		2. 它实现了子系统与客户端之间的松耦合关系，这使得子系统的变化不会影响到调用它的客户端，
 *		   只需要调整外观类即可。
 *		3. 一个子系统的修改对其他子系统没有任何影响。
 *
 *	缺点：
 *		1. 不能很好地限制客户端直接使用子系统类，如果对客户端访问子系统类做太多的限制则减少了可变性和灵活性。
 *		2. 如果设计不当，增加新的子系统可能需要修改外观类的源代码，违背了开闭原则。
 *
 *	适用场景：
 * 		1. 复杂系统需要简单入口使用。
 *		2. 客户端程序与多个子系统之间存在很大的依赖性。
 *		3. 在层次化结构中，可以使用外观模式定义系统中每一层的入口，层与层之间不直接产生联系，
 *		   而通过外观类建立联系，降低层之间的耦合度。
 */

// -------------------子系统--------------------------

// 电视机
type TV struct{}

func (t *TV) On() {
	fmt.Println("打开 电视机")
}

func (t *TV) Off() {
	fmt.Println("关闭 电视机")
}

// 电视机
type VoiceBox struct{}

func (v *VoiceBox) On() {
	fmt.Println("打开 音箱")
}

func (v *VoiceBox) Off() {
	fmt.Println("关闭 音箱")
}

// 灯光
type Light struct{}

func (l *Light) On() {
	fmt.Println("打开 灯光")
}

func (l *Light) Off() {
	fmt.Println("关闭 灯光")
}

// 游戏机
type Xbox struct{}

func (x *Xbox) On() {
	fmt.Println("打开 游戏机")
}

func (x *Xbox) Off() {
	fmt.Println("关闭 游戏机")
}

// 麦克风
type MicroPhone struct{}

func (m *MicroPhone) On() {
	fmt.Println("打开 麦克风")
}

func (m *MicroPhone) Off() {
	fmt.Println("关闭 麦克风")
}

// 投影仪
type Projector struct{}

func (p *Projector) On() {
	fmt.Println("打开 投影仪")
}

func (p *Projector) Off() {
	fmt.Println("关闭 投影仪")
}

// -------------------外观角色--------------------------

// 家庭影院(外观)
type HomePlayerFacade struct {
	tv    TV
	vb    VoiceBox
	light Light
	xbox  Xbox
	mp    MicroPhone
	pro   Projector
}

// KTV模式
func (hp *HomePlayerFacade) DoKTV() {
	fmt.Println("家庭影院进入KTV模式")
	hp.tv.On()
	hp.pro.On()
	hp.mp.On()
	hp.light.Off()
	hp.vb.On()
}

// 游戏模式
func (hp *HomePlayerFacade) DoGame() {
	fmt.Println("家庭影院进入Game模式")
	hp.tv.On()
	hp.light.On()
	hp.xbox.On()
}
