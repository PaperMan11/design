package structural

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	var huawei = new(HuaWei)
	huawei.Show() //调用原构件方法

	fmt.Println("---------")
	//用贴膜装饰器装饰，得到新功能构件
	var moHuawei = NewMoDecorator(huawei) //通过HueWei ---> MoHuaWei
	moHuawei.Show()                       //调用装饰后新构件的方法

	fmt.Println("---------")
	var keHuawei = NewKeDecorator(huawei) //通过HueWei ---> KeHuaWei
	keHuawei.Show()

	fmt.Println("---------")
	var keMoHuaWei = NewMoDecorator(keHuawei) //通过KeHuaWei ---> KeMoHuaWei
	keMoHuaWei.Show()
}
