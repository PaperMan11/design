## 1 设计模式

> 原文链接 [https://www.yuque.com/aceld/lfhu8y/rg6nsf](https://www.yuque.com/aceld/lfhu8y/rg6nsf)

### 1.1 软件设计模式的种类

GoF提出的设计模式有23个，包括：
1. 创建型(Creational)模式：如何创建对象；
1. 结构型(Structural )模式：如何实现类或对象的组合；
1. 行为型(Behavioral)模式：类或对象怎样交互以及怎样分配职责。

设计模式目前种类： GoF的23种   + “简单工厂模式” = 24种。


### 1.2 设计模式总览表
<table class="ne-table" style="width: 685px;"><colgroup><col
                    width="124"><col width="187"><col width="187"><col
                    width="186"></colgroup><tbody class="ne-table-inner"><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(71, 119, 29);" data-col="0"><div
                            class="ne-td-content"><ne-p id="ub04345ee"
                                data-lake-id="ub04345ee" ne-alignment="center"><ne-text
                                    id="ud02b6f98" ne-bold="true" style="color:
                                    rgb(255, 255, 255);">模式名称</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(71, 119,
                        29);" data-col="1"><div class="ne-td-content"><ne-p
                                id="uf85781e7" data-lake-id="uf85781e7"
                                ne-alignment="center"><ne-text id="ud1d6e3bd"
                                    ne-bold="true" style="color: rgb(255, 255,
                                    255);">模式名称</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(71, 119,
                        29);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="u6a7add66"
                                data-lake-id="u6a7add66" ne-alignment="center"><ne-text
                                    id="u731e36f5" ne-bold="true" style="color:
                                    rgb(255, 255, 255);">作用</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(108, 108, 108);" rowspan="6" data-col="0"><div
                            class="ne-td-content"><ne-p id="ua901b4e4"
                                data-lake-id="ua901b4e4" ne-alignment="center"><ne-text
                                    id="udb6986cb" ne-bold="true" style="color:
                                    rgb(255, 255, 255);">创建型模式</ne-text><ne-text
                                    id="u85dc547d" ne-bold="true" style="color:
                                    rgb(255, 255, 255);"> </ne-text><ne-text
                                    id="ufd2e75c1" ne-bold="true" style="color:
                                    rgb(255, 255, 255);">Creational Pattern</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="u93d243b1" data-lake-id="u93d243b1"
                                ne-alignment="center"><ne-text id="u99322eec"
                                    ne-bold="true" style="color: rgb(255, 255,
                                    255);">（</ne-text><ne-text id="uad9610fd"
                                    ne-bold="true" style="color: rgb(255, 255,
                                    255);">6</ne-text><ne-text id="u8f6d76d3"
                                    ne-bold="true" style="color: rgb(255, 255,
                                    255);">）</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(254, 220,
                        79);" data-col="1"><div class="ne-td-content"><ne-p
                                id="u726f56f7" data-lake-id="u726f56f7"
                                ne-alignment="center"><ne-text id="u20570e51"
                                    style="color: rgb(0, 0, 0);">单例模式</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="u57c661bf" data-lake-id="u57c661bf"
                                ne-alignment="center"><ne-text id="u9fcdb887"
                                    style="color: rgb(251, 2, 7);">★★★★☆</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(255, 255,
                        255);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="u30d94006"
                                data-lake-id="u30d94006" ne-alignment="center"><ne-text
                                    id="u23e8b655" style="color: rgb(0, 0, 0);">是保证一个类仅有一个实例，并提供一个访问它的全局访问点。</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(254, 220, 79);" data-col="1"><div
                            class="ne-td-content"><ne-p id="u78041b9b"
                                data-lake-id="u78041b9b" ne-alignment="center"><ne-text
                                    id="ubc806d2c" style="color: rgb(0, 0, 0);">简单工厂模式</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="u93fe324e" data-lake-id="u93fe324e"
                                ne-alignment="center"><ne-text id="u73c2c954"
                                    style="color: rgb(251, 2, 7);">★★★☆☆</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(235, 235,
                        235);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="ud58f12b3"
                                data-lake-id="ud58f12b3" ne-alignment="center"><ne-text
                                    id="ubd97f423" style="color: rgb(0, 0, 0);">通过专门定义一个类来负责创建其他类的实例，被创建的实例通常都具有共同的父类。</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(254, 220, 79);" data-col="1"><div
                            class="ne-td-content"><ne-p id="ube8d881d"
                                data-lake-id="ube8d881d" ne-alignment="center"><ne-text
                                    id="u29e3babb" style="color: rgb(0, 0, 0);">工厂方法模式</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="u1da9ff38" data-lake-id="u1da9ff38"
                                ne-alignment="center"><ne-text id="uec4e3ebd"
                                    style="color: rgb(251, 2, 7);">★★★★★</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(255, 255,
                        255);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="u5d121852"
                                data-lake-id="u5d121852" ne-alignment="center"><ne-text
                                    id="u6c809f11" style="color: rgb(0, 0, 0);">定义一个创建产品对象的工厂接口，将实际创建工作推</ne-text><ne-text
                                    id="ue37734f9" style="color: rgb(0, 0, 0);">迟到子类中。</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(254, 220, 79);" data-col="1"><div
                            class="ne-td-content"><ne-p id="u066e1f46"
                                data-lake-id="u066e1f46" ne-alignment="center"><ne-text
                                    id="u1aacbd90" style="color: rgb(0, 0, 0);">抽象工厂模式</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="ufef0776e" data-lake-id="ufef0776e"
                                ne-alignment="center"><ne-text id="u34e3b12c"
                                    style="color: rgb(251, 2, 7);">★★★★★</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(235, 235,
                        235);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="u0e77695a"
                                data-lake-id="u0e77695a" ne-alignment="center"><ne-text
                                    id="u6155ee30" style="color: rgb(0, 0, 0);">提供一个创建一系列相关或者相互依赖的接口，而无需指定它们具体的类。</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(255, 255, 255);" data-col="1"><div
                            class="ne-td-content"><ne-p id="uf12c314c"
                                data-lake-id="uf12c314c" ne-alignment="center"><ne-text
                                    id="ue8eeb1c8" style="color: rgb(0, 0, 0);">原型模式</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="u9337f48f" data-lake-id="u9337f48f"
                                ne-alignment="center"><ne-text id="uc48b3205"
                                    style="color: rgb(251, 2, 7);">★★★☆☆</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(255, 255,
                        255);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="ua3ddc555"
                                data-lake-id="ua3ddc555" ne-alignment="center"><ne-text
                                    id="ua5a74252" style="color: rgb(0, 0, 0);">用原型实例指定创建对象的种类，并且通过拷贝这些原型创建新的对象。</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(235, 235, 235);" data-col="1"><div
                            class="ne-td-content"><ne-p id="u38e12954"
                                data-lake-id="u38e12954" ne-alignment="center"><ne-text
                                    id="ufc085c92" style="color: rgb(0, 0, 0);">建造者模式</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="u2b90dcee" data-lake-id="u2b90dcee"
                                ne-alignment="center"><ne-text id="ub5759b1c"
                                    style="color: rgb(251, 2, 7);">★★☆☆☆</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(235, 235,
                        235);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="u1c03b286"
                                data-lake-id="u1c03b286" ne-alignment="center"><ne-text
                                    id="u7faf7623" style="color: rgb(0, 0, 0);">将一个复</ne-text><ne-text
                                    id="u5db4016e" style="color: rgb(0, 0, 0);">杂的构建与其表示相分离，使得同样的构建过程可以创建不同的表示。</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(108, 108, 108);" rowspan="7" data-col="0"><div
                            class="ne-td-content"><ne-p id="uec5e7331"
                                data-lake-id="uec5e7331" ne-alignment="center"><ne-text
                                    id="ub614ee20" ne-bold="true" style="color:
                                    rgb(255, 255, 255);">结构型模式</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="u7e6fd7e3" data-lake-id="u7e6fd7e3"
                                ne-alignment="center"><ne-text id="ubd3b3e7b"
                                    ne-bold="true" style="color: rgb(255, 255,
                                    255);">Structural Pattern</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="uc3592aaf" data-lake-id="uc3592aaf"
                                ne-alignment="center"><ne-text id="ua640fd07"
                                    ne-bold="true" style="color: rgb(255, 255,
                                    255);">（</ne-text><ne-text id="u36e63f5b"
                                    ne-bold="true" style="color: rgb(255, 255,
                                    255);">7</ne-text><ne-text id="uc8d3e003"
                                    ne-bold="true" style="color: rgb(255, 255,
                                    255);">）</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(254, 220,
                        79);" data-col="1"><div class="ne-td-content"><ne-p
                                id="u357ab596" data-lake-id="u357ab596"
                                ne-alignment="center"><ne-text id="uc6d1a4f7"
                                    style="color: rgb(0, 0, 0);">适配器模式</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="u5c12ba3d" data-lake-id="u5c12ba3d"
                                ne-alignment="center"><ne-text id="u9fe66dc5"
                                    style="color: rgb(251, 2, 7);">★★★★☆</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(255, 255,
                        255);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="u8d4c9ea8"
                                data-lake-id="u8d4c9ea8" ne-alignment="center"><ne-text
                                    id="uc284d983" style="color: rgb(0, 0, 0);">将一个类的接口转换成客户希望的另外一个接口。使得原本由于接口不兼容而不能一起工作的那些类可以一起工作。</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" data-col="1"><div
                            class="ne-td-content"><ne-p id="ucb5aff9f"
                                data-lake-id="ucb5aff9f" ne-alignment="center"><ne-text
                                    id="u851a605d" style="color: rgb(0, 0, 0);">桥接模式</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="ucd919d71" data-lake-id="ucd919d71"
                                ne-alignment="center"><ne-text id="ud1fbc406"
                                    style="color: rgb(251, 2, 7);">★★★☆☆</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(235, 235,
                        235);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="ueac3751e"
                                data-lake-id="ueac3751e" ne-alignment="center"><ne-text
                                    id="u8c73a130" style="color: rgb(0, 0, 0);">将抽象部分与实际部分分离，使它们都可以独立的变化。</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(255, 255, 255);" data-col="1"><div
                            class="ne-td-content"><ne-p id="u4cce8660"
                                data-lake-id="u4cce8660" ne-alignment="center"><ne-text
                                    id="u5df43745" style="color: rgb(0, 0, 0);">组合模式</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="uac9efa36" data-lake-id="uac9efa36"
                                ne-alignment="center"><ne-text id="ub817da50"
                                    style="color: rgb(251, 2, 7);">★★☆☆☆</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(255, 255,
                        255);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="ua46cae4d"
                                data-lake-id="ua46cae4d" ne-alignment="center"><ne-text
                                    id="uad430a18" style="color: rgb(0, 0, 0);">将对象组合成树形结构以表示</ne-text><ne-text
                                    id="ue1d6fd8d" style="color: rgb(0, 0, 0);">“</ne-text><ne-text
                                    id="u5cbbe1a7" style="color: rgb(0, 0, 0);">部分</ne-text><ne-text
                                    id="u9e20eea5" style="color: rgb(0, 0, 0);">--</ne-text><ne-text
                                    id="uac9400be" style="color: rgb(0, 0, 0);">整体</ne-text><ne-text
                                    id="u6ce81db1" style="color: rgb(0, 0, 0);">”</ne-text><ne-text
                                    id="uf779320a" style="color: rgb(0, 0, 0);">的层次结构。使得用户对单个对象和组合对象的使用具有一致性。</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(254, 220, 79);" data-col="1"><div
                            class="ne-td-content"><ne-p id="u7eb03d9d"
                                data-lake-id="u7eb03d9d" ne-alignment="center"><ne-text
                                    id="u7526e1b1" style="color: rgb(0, 0, 0);">装饰模式</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="u7100fa0e" data-lake-id="u7100fa0e"
                                ne-alignment="center"><ne-text id="ub5386037"
                                    style="color: rgb(251, 2, 7);">★★★☆☆</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(235, 235,
                        235);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="u98987601"
                                data-lake-id="u98987601" ne-alignment="center"><ne-text
                                    id="ub7659be4" style="color: rgb(0, 0, 0);">动态的给一个对象添加一些额外的职责。就增加功能来说，此模式比生成子类更为灵活。</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(254, 220, 79);" data-col="1"><div
                            class="ne-td-content"><ne-p id="ufbd3b2fa"
                                data-lake-id="ufbd3b2fa" ne-alignment="center"><ne-text
                                    id="ub2ba8209" style="color: rgb(0, 0, 0);">外观模式</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="u21def322" data-lake-id="u21def322"
                                ne-alignment="center"><ne-text id="ubf36bf61"
                                    style="color: rgb(251, 2, 7);">★★★★★</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(255, 255,
                        255);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="u17cf8e28"
                                data-lake-id="u17cf8e28" ne-alignment="center"><ne-text
                                    id="u34c96725" style="color: rgb(0, 0, 0);">为子系统中的一组接口提供一个一致的界面，此模式定义了一个高层接口，这个接口使得这一子系统更加容易使用。</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(235, 235, 235);" data-col="1"><div
                            class="ne-td-content"><ne-p id="u9cb53dc6"
                                data-lake-id="u9cb53dc6" ne-alignment="center"><ne-text
                                    id="u9a0958e0" style="color: rgb(0, 0, 0);">享元模式</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="u7142bac8" data-lake-id="u7142bac8"
                                ne-alignment="center"><ne-text id="uce5d3530"
                                    style="color: rgb(251, 2, 7);">★☆☆☆☆</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(235, 235,
                        235);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="u85efdd02"
                                data-lake-id="u85efdd02" ne-alignment="center"><ne-text
                                    id="uaac933b7" style="color: rgb(0, 0, 0);">以共享的方式高效的支持大量的细粒度的对象。</ne-text><ne-text
                                    id="ud44b647e" style="color: rgb(0, 0, 0);">
                                </ne-text><ne-text id="u2f250266" style="color:
                                    rgb(0, 0, 0);"> </ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(254, 220, 79);" data-col="1"><div
                            class="ne-td-content"><ne-p id="u7940cb36"
                                data-lake-id="u7940cb36" ne-alignment="center"><ne-text
                                    id="u04c513d7" style="color: rgb(0, 0, 0);">代理模式</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="u368b3272" data-lake-id="u368b3272"
                                ne-alignment="center"><ne-text id="u187f4d0c"
                                    style="color: rgb(251, 2, 7);">★★★★☆</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(255, 255,
                        255);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="u36e2520b"
                                data-lake-id="u36e2520b" ne-alignment="center"><ne-text
                                    id="u553b5325" style="color: rgb(0, 0, 0);">为其他对象提供一种代理以控制对这个对象的访问。</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(108, 108, 108);" rowspan="11" data-col="0"><div
                            class="ne-td-content"><ne-p id="ue3bec636"
                                data-lake-id="ue3bec636" ne-alignment="center"><ne-text
                                    id="uffdd619f" ne-bold="true" style="color:
                                    rgb(255, 255, 255);">行为型模式</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="uc2216cef" data-lake-id="uc2216cef"
                                ne-alignment="center"><ne-text id="u7bdbda03"
                                    ne-bold="true" style="color: rgb(255, 255,
                                    255);">Behavioral Pattern</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="u82f2e917" data-lake-id="u82f2e917"
                                ne-alignment="center"><ne-text id="u6ecb9a55"
                                    ne-bold="true" style="color: rgb(255, 255,
                                    255);">（</ne-text><ne-text id="ue1351457"
                                    ne-bold="true" style="color: rgb(255, 255,
                                    255);">11</ne-text><ne-text id="u85bd7f48"
                                    ne-bold="true" style="color: rgb(255, 255,
                                    255);">）</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(235, 235,
                        235);" data-col="1"><div class="ne-td-content"><ne-p
                                id="u66d1cfb8" data-lake-id="u66d1cfb8"
                                ne-alignment="center"><ne-text id="u3e3b3a55"
                                    style="color: rgb(0, 0, 0);">职责链模式</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="u4084302b" data-lake-id="u4084302b"
                                ne-alignment="center"><ne-text id="u81fba3a0"
                                    style="color: rgb(251, 2, 7);">★★☆☆☆</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(235, 235,
                        235);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="u869dda0f"
                                data-lake-id="u869dda0f" ne-alignment="center"><ne-text
                                    id="u742a39b7" style="color: rgb(0, 0, 0);">在该模式里，很多对象由每一个对象对其下家的引用而连接起来形成一条链。请求在这个链上传递，直到链上的某一个对象决定处理此请求，这使得系统可以在不影响客户端的情况下动态地重新组织链和分配责任。</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(254, 220, 79);" data-col="1"><div
                            class="ne-td-content"><ne-p id="uc2bfbc7d"
                                data-lake-id="uc2bfbc7d" ne-alignment="center"><ne-text
                                    id="ue8394efa" style="color: rgb(0, 0, 0);">命令模式</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="u77619198" data-lake-id="u77619198"
                                ne-alignment="center"><ne-text id="u311a321e"
                                    style="color: rgb(251, 2, 7);">★★★★☆</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(255, 255,
                        255);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="uc531220d"
                                data-lake-id="uc531220d" ne-alignment="center"><ne-text
                                    id="u3956814d" style="color: rgb(0, 0, 0);">将一个请求封装为一个对象，从而使你可用不同的请求对客户端进行参数化；对请求排队或记录请求日志，以及支持可撤销的操作。</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(235, 235, 235);" data-col="1"><div
                            class="ne-td-content"><ne-p id="u0f494dd8"
                                data-lake-id="u0f494dd8" ne-alignment="center"><ne-text
                                    id="uc5312433" style="color: rgb(0, 0, 0);">解释器模式</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="u93c77b91" data-lake-id="u93c77b91"
                                ne-alignment="center"><ne-text id="ua672a4d0"
                                    style="color: rgb(251, 2, 7);">★☆☆☆☆</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(235, 235,
                        235);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="u34b70851"
                                data-lake-id="u34b70851" ne-alignment="center"><ne-text
                                    id="u5dcf0bb7" style="color: rgb(0, 0, 0);">如何为简单的语言定义一个语法，如何在该语言中表示一个句子，以及如何解释这些句子。</ne-text><ne-text
                                    id="u1a77fe6b" style="color: rgb(0, 0, 0);">
                                </ne-text><span class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(255, 255, 255);" data-col="1"><div
                            class="ne-td-content"><ne-p id="u716e5dfb"
                                data-lake-id="u716e5dfb" ne-alignment="center"><ne-text
                                    id="ud17ea554" style="color: rgb(0, 0, 0);">迭代器模式</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="u3ebec942" data-lake-id="u3ebec942"
                                ne-alignment="center"><ne-text id="u422d7c18"
                                    style="color: rgb(251, 2, 7);">★☆☆☆☆</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(255, 255,
                        255);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="ufa3af7a8"
                                data-lake-id="ufa3af7a8" ne-alignment="center"><ne-text
                                    id="u3eadc898" style="color: rgb(0, 0, 0);">提供了一种方法顺序来访问一个聚合对象中的各个元素，而又不需要暴露该对象的内部表示。</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(235, 235, 235);" data-col="1"><div
                            class="ne-td-content"><ne-p id="u8076df3c"
                                data-lake-id="u8076df3c" ne-alignment="center"><ne-text
                                    id="ucef4e03a" style="color: rgb(0, 0, 0);">中介者模式</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="u240513c6" data-lake-id="u240513c6"
                                ne-alignment="center"><ne-text id="u1ab4ba7c"
                                    style="color: rgb(251, 2, 7);">★★☆☆☆</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(235, 235,
                        235);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="u1968ade4"
                                data-lake-id="u1968ade4" ne-alignment="center"><ne-text
                                    id="ue6c6b005" style="color: rgb(0, 0, 0);">定义一个中介对象来封装系列对象之间的交互。终结者使各个对象不需要显示的相互调用</ne-text><ne-text
                                    id="u00ce4bc1" style="color: rgb(0, 0, 0);">
                                </ne-text><ne-text id="u19fee7c6" style="color:
                                    rgb(0, 0, 0);">，从而使其耦合性松散，而且可以独立的改变他们之间的交互。</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(255, 255, 255);" data-col="1"><div
                            class="ne-td-content"><ne-p id="ufe6bcdd4"
                                data-lake-id="ufe6bcdd4" ne-alignment="center"><ne-text
                                    id="u92409155" style="color: rgb(0, 0, 0);">备忘录模式</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="u5acbc56a" data-lake-id="u5acbc56a"
                                ne-alignment="center"><ne-text id="u98c74dc8"
                                    style="color: rgb(251, 2, 7);">★★☆☆☆</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(255, 255,
                        255);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="u3f95e797"
                                data-lake-id="u3f95e797" ne-alignment="center"><ne-text
                                    id="u50e1b4fb" style="color: rgb(0, 0, 0);">是在不破坏封装的前提下，捕获一个对象的内部状态，并在该对象之外保存这个状态。</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(254, 220, 79);" data-col="1"><div
                            class="ne-td-content"><ne-p id="u13540607"
                                data-lake-id="u13540607" ne-alignment="center"><ne-text
                                    id="ua72d7cf7" style="color: rgb(0, 0, 0);">观察者模式</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="ue9404ba1" data-lake-id="ue9404ba1"
                                ne-alignment="center"><ne-text id="u436b054c"
                                    style="color: rgb(251, 2, 7);">★★★★★</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(235, 235,
                        235);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="ub2e5954e"
                                data-lake-id="ub2e5954e" ne-alignment="center"><ne-text
                                    id="u53ece606" style="color: rgb(0, 0, 0);">定义对象间的一种一对多的依赖关系，当一个对象的状态发生改变时，所有依赖于它的对象都得到通知并被自动更新。</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(255, 255, 255);" data-col="1"><div
                            class="ne-td-content"><ne-p id="u17589dfd"
                                data-lake-id="u17589dfd" ne-alignment="center"><ne-text
                                    id="u4b5a9fbe" style="color: rgb(0, 0, 0);">状态模式</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="u8fd4db88" data-lake-id="u8fd4db88"
                                ne-alignment="center"><ne-text id="ua6c4763e"
                                    style="color: rgb(251, 2, 7);">★★☆☆☆</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(255, 255,
                        255);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="u656c9f96"
                                data-lake-id="u656c9f96" ne-alignment="center"><ne-text
                                    id="u3a99574a" style="color: rgb(0, 0, 0);">对象的行为，依赖于它所处的状态。</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="u0490f3eb" data-lake-id="u0490f3eb"><ne-text
                                    id="u785d1ee3">
                                </ne-text><span class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(254, 220, 79);" data-col="1"><div
                            class="ne-td-content"><ne-p id="u02312d9c"
                                data-lake-id="u02312d9c" ne-alignment="center"><ne-text
                                    id="u27ffb049" style="color: rgb(0, 0, 0);">策略模式</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="u469a4467" data-lake-id="u469a4467"
                                ne-alignment="center"><ne-text id="ue33c26c4"
                                    style="color: rgb(251, 2, 7);">★★★★☆</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(235, 235,
                        235);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="u374ea560"
                                data-lake-id="u374ea560" ne-alignment="center"><ne-text
                                    id="ua00f602a" style="color: rgb(0, 0, 0);">准备一组算法，并将每一个算法封装起来，使得它们可以互换。</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(254, 220, 79);" data-col="1"><div
                            class="ne-td-content"><ne-p id="uf0e935cb"
                                data-lake-id="uf0e935cb" ne-alignment="center"><ne-text
                                    id="u4c4acac9" style="color: rgb(0, 0, 0);">模板方法模式</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="u26a12be6" data-lake-id="u26a12be6"
                                ne-alignment="center"><ne-text id="u963b1819"
                                    style="color: rgb(251, 2, 7);">★★★☆☆</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(255, 255,
                        255);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="ud95cb679"
                                data-lake-id="ud95cb679" ne-alignment="center"><ne-text
                                    id="u7ff8915d" style="color: rgb(0, 0, 0);">得子类可以不改变一个算法的结构即可重定义该算法的某些特定步骤。</ne-text><ne-text
                                    id="ua5a286d4" style="color: rgb(0, 0, 0);">
                                </ne-text><span class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(235, 235, 235);" data-col="1"><div
                            class="ne-td-content"><ne-p id="uf9223153"
                                data-lake-id="uf9223153" ne-alignment="center"><ne-text
                                    id="ub4b4eaac" style="color: rgb(0, 0, 0);">访问者模式</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="u2fb4fc2b" data-lake-id="u2fb4fc2b"
                                ne-alignment="center"><ne-text id="u98edbde4"
                                    style="color: rgb(251, 2, 7);">★☆☆☆☆</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(235, 235,
                        235);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="ub8e306ff"
                                data-lake-id="ub8e306ff" ne-alignment="center"><ne-text
                                    id="ub5fbacc4" style="color: rgb(0, 0, 0);">表示一个作用于某对象结构中的各元素的操作，它使你可以在不改变各元素的类的前提下定义作用于这些元素的新操作。</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr></tbody></table>


## 2 面向对象设计原则

> 原则的目的：高内聚，低耦合

**面向对象设计原则表**：
<table class="ne-table" style="width: 750px;"><colgroup><col
                    width="375"><col width="374"></colgroup><tbody
                class="ne-table-inner"><tr class="ne-tr"><td class="ne-td"
                        style="background-color: rgb(71, 119, 29);"
                        data-col="0"><div class="ne-td-content"><ne-p
                                id="ua18c25a1" data-lake-id="ua18c25a1"
                                ne-alignment="center"><ne-text id="u9fbb434c"
                                    ne-bold="true" style="color: rgb(255, 255,
                                    255);">名称</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(71, 119,
                        29);" data-col="1"><div class="ne-td-content"><ne-p
                                id="u51baa67c" data-lake-id="u51baa67c"
                                ne-alignment="center"><ne-text id="uad0d4283"
                                    ne-bold="true" style="color: rgb(255, 255,
                                    255);">定义</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(108, 108, 108);" data-col="0"><div
                            class="ne-td-content"><ne-p id="u6284af3a"
                                data-lake-id="u6284af3a" ne-alignment="center"><ne-text
                                    id="u8720f7b8" ne-bold="true" style="color:
                                    rgb(255, 255, 255);">单一职责原则</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="ub52fd534" data-lake-id="ub52fd534"
                                ne-alignment="center"><ne-text id="u0393e5e3"
                                    ne-bold="true" style="color: rgb(255, 255,
                                    255);">(Single Responsibility Principle,
                                    SRP)</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="u9ad82f07" data-lake-id="u9ad82f07"
                                ne-alignment="center"><ne-text id="ube6a8743"
                                    ne-bold="true" style="color: rgb(251, 2,
                                    7);">★★★★☆</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(255, 255,
                        255);" data-col="1"><div class="ne-td-content"><ne-p
                                id="u7103f101" data-lake-id="u7103f101"><ne-text
                                    id="u8a1a788d" style="color: rgb(0, 0, 0);">类的职责单一，对外只提供一种功能，而引起类变化的原因都应该只有一个。</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(108, 108, 108);" data-col="0"><div
                            class="ne-td-content"><ne-p id="u2c4eabb1"
                                data-lake-id="u2c4eabb1" ne-alignment="center"><ne-text
                                    id="u831f6b66" ne-bold="true" style="color:
                                    rgb(255, 255, 255);">开闭原则</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="u62de9c27" data-lake-id="u62de9c27"
                                ne-alignment="center"><ne-text id="u02bb9be1"
                                    ne-bold="true" style="color: rgb(255, 255,
                                    255);">(Open-Closed Principle, OCP)</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="u7750ccc6" data-lake-id="u7750ccc6"
                                ne-alignment="center"><ne-text id="u92c3d9c7"
                                    ne-bold="true" style="color: rgb(251, 2,
                                    7);">★★★★★</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(235, 235,
                        235);" data-col="1"><div class="ne-td-content"><ne-p
                                id="uc77124d5" data-lake-id="uc77124d5"><ne-text
                                    id="u5f508092" style="color: rgb(251, 2,
                                    7);">类的改动是通过增加代码进行的，而不是修改源代码。</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(108, 108, 108);" data-col="0"><div
                            class="ne-td-content"><ne-p id="u361d0570"
                                data-lake-id="u361d0570" ne-alignment="center"><ne-text
                                    id="u37d8282a" ne-bold="true" style="color:
                                    rgb(255, 255, 255);">里氏代换原则</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="u51f7acf4" data-lake-id="u51f7acf4"
                                ne-alignment="center"><ne-text id="u46cb7985"
                                    ne-bold="true" style="color: rgb(255, 255,
                                    255);">(Liskov Substitution Principle, LSP</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="ufe54c3f6" data-lake-id="ufe54c3f6"
                                ne-alignment="center"><ne-text id="u2d4f4442"
                                    ne-bold="true" style="color: rgb(251, 2,
                                    7);">★★★★★</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(255, 255,
                        255);" data-col="1"><div class="ne-td-content"><ne-p
                                id="ufef02504" data-lake-id="ufef02504"><ne-text
                                    id="u34901593" style="color: rgb(0, 0, 0);">任何抽象类（interface接口）出现的地方都可以用他的实现类进行替换，实际就是虚拟机制，语言级别实现面向对象功能。</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(108, 108, 108);" data-col="0"><div
                            class="ne-td-content"><ne-p id="uf6c15886"
                                data-lake-id="uf6c15886" ne-alignment="center"><ne-text
                                    id="u6410be9b" ne-bold="true" style="color:
                                    rgb(255, 255, 255);">依赖倒转原则</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="u0c788a32" data-lake-id="u0c788a32"
                                ne-alignment="center"><ne-text id="ufb196368"
                                    ne-bold="true" style="color: rgb(255, 255,
                                    255);">(Dependence</ne-text><ne-text
                                    id="uc369bea0" ne-bold="true" style="color:
                                    rgb(255, 255, 255);"> </ne-text><ne-text
                                    id="uc4e985a5" ne-bold="true" style="color:
                                    rgb(255, 255, 255);">Inversion Principle,
                                    DIP)</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="ucc50c9ce" data-lake-id="ucc50c9ce"
                                ne-alignment="center"><ne-text id="u4e3135bc"
                                    ne-bold="true" style="color: rgb(251, 2,
                                    7);">★★★★★</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(235, 235,
                        235);" data-col="1"><div class="ne-td-content"><ne-p
                                id="u9ee8c3f4" data-lake-id="u9ee8c3f4"><ne-text
                                    id="ua790ad6f" style="color: rgb(251, 2,
                                    7);">依赖于抽象</ne-text><ne-text id="u7ba1615d"
                                    style="color: rgb(251, 2, 7);">(</ne-text><ne-text
                                    id="u1b4a17fe" style="color: rgb(251, 2,
                                    7);">接口</ne-text><ne-text id="ucaa18e0e"
                                    style="color: rgb(251, 2, 7);">)</ne-text><ne-text
                                    id="u0ef5f3f1" style="color: rgb(251, 2,
                                    7);">，不要依赖具体的实现</ne-text><ne-text
                                    id="uf91bea4b" style="color: rgb(251, 2,
                                    7);">(</ne-text><ne-text id="ued8e11fb"
                                    style="color: rgb(251, 2, 7);">类</ne-text><ne-text
                                    id="u7cca3839" style="color: rgb(251, 2,
                                    7);">)</ne-text><ne-text id="ub10b0afb"
                                    style="color: rgb(251, 2, 7);">，也就是针对接口编程。</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(108, 108, 108);" data-col="0"><div
                            class="ne-td-content"><ne-p id="uaa40c9ca"
                                data-lake-id="uaa40c9ca" ne-alignment="center"><ne-text
                                    id="u74acc498" ne-bold="true" style="color:
                                    rgb(255, 255, 255);">接口隔离原则</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="ud287cbf9" data-lake-id="ud287cbf9"
                                ne-alignment="center"><ne-text id="ucdf8b63b"
                                    ne-bold="true" style="color: rgb(255, 255,
                                    255);">(Interface Segregation Principle, ISP</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="u2b8578e9" data-lake-id="u2b8578e9"
                                ne-alignment="center"><ne-text id="ue72af631"
                                    ne-bold="true" style="color: rgb(251, 2,
                                    7);">★★☆☆☆</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(255, 255,
                        255);" data-col="1"><div class="ne-td-content"><ne-p
                                id="u0625d642" data-lake-id="u0625d642"><ne-text
                                    id="ud9a3c0aa" style="color: rgb(0, 0, 0);">不应该强迫用户的程序依赖他们不需要的接口方法。一个接口应该只提供一种对外功能，不应该把所有操作都封装到一个接口中去。</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(108, 108, 108);" data-col="0"><div
                            class="ne-td-content"><ne-p id="uf7fbda86"
                                data-lake-id="uf7fbda86" ne-alignment="center"><ne-text
                                    id="ufc532b48" ne-bold="true" style="color:
                                    rgb(255, 255, 255);">合成复用原则</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="u33f42a12" data-lake-id="u33f42a12"
                                ne-alignment="center"><ne-text id="u26cf2984"
                                    ne-bold="true" style="color: rgb(255, 255,
                                    255);">(Composite Reuse Principle, CRP)</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="ue21204e7" data-lake-id="ue21204e7"
                                ne-alignment="center"><ne-text id="uf112d509"
                                    ne-bold="true" style="color: rgb(251, 2,
                                    7);">★★★★☆</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(235, 235,
                        235);" data-col="1"><div class="ne-td-content"><ne-p
                                id="u2dac04ce" data-lake-id="u2dac04ce"><ne-text
                                    id="ue09a2751" style="color: rgb(0, 0, 0);">如果使用继承，会导致父类的任何变换都可能影响到子类的行为。如果使用对象组合，就降低了这种依赖关系。对于继承和组合，优先使用组合。</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(108, 108, 108);" data-col="0"><div
                            class="ne-td-content"><ne-p id="u5f85d5de"
                                data-lake-id="u5f85d5de" ne-alignment="center"><ne-text
                                    id="uc17d3cc7" ne-bold="true" style="color:
                                    rgb(255, 255, 255);">迪米特法则</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="u883ae6e9" data-lake-id="u883ae6e9"
                                ne-alignment="center"><ne-text id="ue6068f94"
                                    ne-bold="true" style="color: rgb(255, 255,
                                    255);">(Law of Demeter, LoD</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="u5aa571a1" data-lake-id="u5aa571a1"
                                ne-alignment="center"><ne-text id="u860d0b08"
                                    ne-bold="true" style="color: rgb(251, 2,
                                    7);">★★★☆☆</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(255, 255,
                        255);" data-col="1"><div class="ne-td-content"><ne-p
                                id="ufc1c8368" data-lake-id="ufc1c8368"><ne-text
                                    id="ufa0c74d1" style="color: rgb(251, 2,
                                    7);">一个对象应当对其他对象尽可能少的了解，从而降低各个对象之间的耦合，提高系统的可维护性。例如在一个程序中，各个模块之间相互调用时，通常会提供一个统一的接口来实现。这样其他模块不需要了解另外一个模块的内部实现细节，这样当一个模块内部的实现发生改变时，不会影响其他模块的使用。（黑盒原理）</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr></tbody></table>

- 单一职责原则：一个类对外只提供一种功能
- 开闭原则：增加功能时去增加代码而不是修改代码
- 依赖倒转原则：模块与模块依赖抽象而不是具体实现
- 合成复用原则：通过组合来实现父类方法
- 迪米特法则：依赖第三方来实现解耦

## 3 创建型模式
<table class="ne-table" style="width: 748px;"><colgroup><col
                    width="187"><col width="187"><col width="187"><col
                    width="186"></colgroup><tbody class="ne-table-inner"><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(71, 119, 29);" data-col="0"><div
                            class="ne-td-content"><ne-p id="udce4ae03"
                                data-lake-id="udce4ae03" ne-alignment="center"><ne-text
                                    id="u5ddf9a7f" ne-bold="true" style="color:
                                    rgb(255, 255, 255);">模式名称</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(71, 119,
                        29);" data-col="1"><div class="ne-td-content"><ne-p
                                id="ub495af5a" data-lake-id="ub495af5a"
                                ne-alignment="center"><ne-text id="u754533c8"
                                    ne-bold="true" style="color: rgb(255, 255,
                                    255);">模式名称</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(71, 119,
                        29);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="u51ec2ffa"
                                data-lake-id="u51ec2ffa" ne-alignment="center"><ne-text
                                    id="u9487cdc6" ne-bold="true" style="color:
                                    rgb(255, 255, 255);">作用</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(108, 108, 108);" rowspan="6" data-col="0"><div
                            class="ne-td-content"><ne-p id="udceb7b20"
                                data-lake-id="udceb7b20" ne-alignment="center"><ne-text
                                    id="ua3f4f4e0" ne-bold="true" style="color:
                                    rgb(255, 255, 255);">创建型模式 Creational
                                    Pattern</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="uef51016a" data-lake-id="uef51016a"
                                ne-alignment="center"><ne-text id="uc2fb32e3"
                                    ne-bold="true" style="color: rgb(255, 255,
                                    255);">（</ne-text><ne-text id="ubacfb65b"
                                    ne-bold="true" style="color: rgb(255, 255,
                                    255);">6</ne-text><ne-text id="ub219fb68"
                                    ne-bold="true" style="color: rgb(255, 255,
                                    255);">）</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(254, 220,
                        79);" data-col="1"><div class="ne-td-content"><ne-p
                                id="ud3d81e7d" data-lake-id="ud3d81e7d"
                                ne-alignment="center"><ne-text id="uca4ed84c"
                                    style="color: rgb(0, 0, 0);">单例模式</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="u43478cd4" data-lake-id="u43478cd4"
                                ne-alignment="center"><ne-text id="ue8cba071"
                                    style="color: rgb(251, 2, 7);">★★★★☆</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(255, 255,
                        255);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="u7573e376"
                                data-lake-id="u7573e376" ne-alignment="center"><ne-text
                                    id="ue925754d" style="color: rgb(0, 0, 0);">是保证一个类仅有一个实例，并提供一个访问它的全局访问点。</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(254, 220, 79);" data-col="1"><div
                            class="ne-td-content"><ne-p id="u4845d3c2"
                                data-lake-id="u4845d3c2" ne-alignment="center"><ne-text
                                    id="uc6f17be6" style="color: rgb(0, 0, 0);">简单工厂模式</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="u8d98fc2d" data-lake-id="u8d98fc2d"
                                ne-alignment="center"><ne-text id="u0ebf58b1"
                                    style="color: rgb(251, 2, 7);">★★★☆☆</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(235, 235,
                        235);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="u0bee8773"
                                data-lake-id="u0bee8773" ne-alignment="center"><ne-text
                                    id="u95799c9b" style="color: rgb(0, 0, 0);">通过专门定义一个类来负责创建其他类的实例，被创建的实例通常都具有共同的父类。</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(254, 220, 79);" data-col="1"><div
                            class="ne-td-content"><ne-p id="u84f17d1a"
                                data-lake-id="u84f17d1a" ne-alignment="center"><ne-text
                                    id="u1dfbb8d8" style="color: rgb(0, 0, 0);">工厂方法模式</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="ucce06e80" data-lake-id="ucce06e80"
                                ne-alignment="center"><ne-text id="u1d8d93d4"
                                    style="color: rgb(251, 2, 7);">★★★★★</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(255, 255,
                        255);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="u72e1dd45"
                                data-lake-id="u72e1dd45" ne-alignment="center"><ne-text
                                    id="u4f985603" style="color: rgb(0, 0, 0);">定义一个创建产品对象的工厂接口，将实际创建工作推迟到子类中。</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(254, 220, 79);" data-col="1"><div
                            class="ne-td-content"><ne-p id="ue8ce8657"
                                data-lake-id="ue8ce8657" ne-alignment="center"><ne-text
                                    id="uadaf237b" style="color: rgb(0, 0, 0);">抽象工厂模式</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="u8c97c4af" data-lake-id="u8c97c4af"
                                ne-alignment="center"><ne-text id="u99d8a078"
                                    style="color: rgb(251, 2, 7);">★★★★★</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(235, 235,
                        235);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="u4fade2d3"
                                data-lake-id="u4fade2d3" ne-alignment="center"><ne-text
                                    id="ub1142aec" style="color: rgb(0, 0, 0);">提供一个创建一系列相关或者相互依赖的接口，而无需指定它们具体的类。</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(255, 255, 255);" data-col="1"><div
                            class="ne-td-content"><ne-p id="ue393aed8"
                                data-lake-id="ue393aed8" ne-alignment="center"><ne-text
                                    id="u7a43ea11" style="color: rgb(0, 0, 0);">原型模式</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="u7dd666f9" data-lake-id="u7dd666f9"
                                ne-alignment="center"><ne-text id="u30e03258"
                                    style="color: rgb(251, 2, 7);">★★★☆☆</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(255, 255,
                        255);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="u082e44f5"
                                data-lake-id="u082e44f5" ne-alignment="center"><ne-text
                                    id="u69f9800c" style="color: rgb(0, 0, 0);">用原型实例指定创建对象的种类，并且通过拷贝这些原型创建新的对象。</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(235, 235, 235);" data-col="1"><div
                            class="ne-td-content"><ne-p id="u522c919b"
                                data-lake-id="u522c919b" ne-alignment="center"><ne-text
                                    id="u6a52cc1b" style="color: rgb(0, 0, 0);">建造者模式</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="uf792fd80" data-lake-id="uf792fd80"
                                ne-alignment="center"><ne-text id="u4e1bc657"
                                    style="color: rgb(251, 2, 7);">★★☆☆☆</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(235, 235,
                        235);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="u163b4435"
                                data-lake-id="u163b4435" ne-alignment="center"><ne-text
                                    id="u1f0b26cb" style="color: rgb(0, 0, 0);">将一个复杂的构建与其表示相分离，使得同样的构建过程可以创建不同的表示。</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr></tbody></table>

## 4 结构型模式
  <table class="ne-table" style="width: 685px;"><colgroup><col
                    width="124"><col width="187"><col width="187"><col
                    width="186"></colgroup><tbody class="ne-table-inner"><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(71, 119, 29);" data-col="0"><div
                            class="ne-td-content"><ne-p id="ub04345ee"
                                data-lake-id="ub04345ee" ne-alignment="center"><ne-text
                                    id="ud02b6f98" ne-bold="true" style="color:
                                    rgb(255, 255, 255);">模式名称</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(71, 119,
                        29);" data-col="1"><div class="ne-td-content"><ne-p
                                id="uf85781e7" data-lake-id="uf85781e7"
                                ne-alignment="center"><ne-text id="ud1d6e3bd"
                                    ne-bold="true" style="color: rgb(255, 255,
                                    255);">模式名称</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(71, 119,
                        29);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="u6a7add66"
                                data-lake-id="u6a7add66" ne-alignment="center"><ne-text
                                    id="u731e36f5" ne-bold="true" style="color:
                                    rgb(255, 255, 255);">作用</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(108, 108, 108);" rowspan="7" data-col="0"><div
                            class="ne-td-content"><ne-p id="uec5e7331"
                                data-lake-id="uec5e7331" ne-alignment="center"><ne-text
                                    id="ub614ee20" ne-bold="true" style="color:
                                    rgb(255, 255, 255);">结构型模式</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="u7e6fd7e3" data-lake-id="u7e6fd7e3"
                                ne-alignment="center"><ne-text id="ubd3b3e7b"
                                    ne-bold="true" style="color: rgb(255, 255,
                                    255);">Structural Pattern</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="uc3592aaf" data-lake-id="uc3592aaf"
                                ne-alignment="center"><ne-text id="ua640fd07"
                                    ne-bold="true" style="color: rgb(255, 255,
                                    255);">（</ne-text><ne-text id="u36e63f5b"
                                    ne-bold="true" style="color: rgb(255, 255,
                                    255);">7</ne-text><ne-text id="uc8d3e003"
                                    ne-bold="true" style="color: rgb(255, 255,
                                    255);">）</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(254, 220,
                        79);" data-col="1"><div class="ne-td-content"><ne-p
                                id="u357ab596" data-lake-id="u357ab596"
                                ne-alignment="center"><ne-text id="uc6d1a4f7"
                                    style="color: rgb(0, 0, 0);">适配器模式</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="u5c12ba3d" data-lake-id="u5c12ba3d"
                                ne-alignment="center"><ne-text id="u9fe66dc5"
                                    style="color: rgb(251, 2, 7);">★★★★☆</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(255, 255,
                        255);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="u8d4c9ea8"
                                data-lake-id="u8d4c9ea8" ne-alignment="center"><ne-text
                                    id="uc284d983" style="color: rgb(0, 0, 0);">将一个类的接口转换成客户希望的另外一个接口。使得原本由于接口不兼容而不能一起工作的那些类可以一起工作。</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" data-col="1"><div
                            class="ne-td-content"><ne-p id="ucb5aff9f"
                                data-lake-id="ucb5aff9f" ne-alignment="center"><ne-text
                                    id="u851a605d" style="color: rgb(0, 0, 0);">桥接模式</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="ucd919d71" data-lake-id="ucd919d71"
                                ne-alignment="center"><ne-text id="ud1fbc406"
                                    style="color: rgb(251, 2, 7);">★★★☆☆</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(235, 235,
                        235);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="ueac3751e"
                                data-lake-id="ueac3751e" ne-alignment="center"><ne-text
                                    id="u8c73a130" style="color: rgb(0, 0, 0);">将抽象部分与实际部分分离，使它们都可以独立的变化。</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(255, 255, 255);" data-col="1"><div
                            class="ne-td-content"><ne-p id="u4cce8660"
                                data-lake-id="u4cce8660" ne-alignment="center"><ne-text
                                    id="u5df43745" style="color: rgb(0, 0, 0);">组合模式</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="uac9efa36" data-lake-id="uac9efa36"
                                ne-alignment="center"><ne-text id="ub817da50"
                                    style="color: rgb(251, 2, 7);">★★☆☆☆</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(255, 255,
                        255);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="ua46cae4d"
                                data-lake-id="ua46cae4d" ne-alignment="center"><ne-text
                                    id="uad430a18" style="color: rgb(0, 0, 0);">将对象组合成树形结构以表示</ne-text><ne-text
                                    id="ue1d6fd8d" style="color: rgb(0, 0, 0);">“</ne-text><ne-text
                                    id="u5cbbe1a7" style="color: rgb(0, 0, 0);">部分</ne-text><ne-text
                                    id="u9e20eea5" style="color: rgb(0, 0, 0);">--</ne-text><ne-text
                                    id="uac9400be" style="color: rgb(0, 0, 0);">整体</ne-text><ne-text
                                    id="u6ce81db1" style="color: rgb(0, 0, 0);">”</ne-text><ne-text
                                    id="uf779320a" style="color: rgb(0, 0, 0);">的层次结构。使得用户对单个对象和组合对象的使用具有一致性。</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(254, 220, 79);" data-col="1"><div
                            class="ne-td-content"><ne-p id="u7eb03d9d"
                                data-lake-id="u7eb03d9d" ne-alignment="center"><ne-text
                                    id="u7526e1b1" style="color: rgb(0, 0, 0);">装饰模式</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="u7100fa0e" data-lake-id="u7100fa0e"
                                ne-alignment="center"><ne-text id="ub5386037"
                                    style="color: rgb(251, 2, 7);">★★★☆☆</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(235, 235,
                        235);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="u98987601"
                                data-lake-id="u98987601" ne-alignment="center"><ne-text
                                    id="ub7659be4" style="color: rgb(0, 0, 0);">动态的给一个对象添加一些额外的职责。就增加功能来说，此模式比生成子类更为灵活。</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(254, 220, 79);" data-col="1"><div
                            class="ne-td-content"><ne-p id="ufbd3b2fa"
                                data-lake-id="ufbd3b2fa" ne-alignment="center"><ne-text
                                    id="ub2ba8209" style="color: rgb(0, 0, 0);">外观模式</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="u21def322" data-lake-id="u21def322"
                                ne-alignment="center"><ne-text id="ubf36bf61"
                                    style="color: rgb(251, 2, 7);">★★★★★</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(255, 255,
                        255);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="u17cf8e28"
                                data-lake-id="u17cf8e28" ne-alignment="center"><ne-text
                                    id="u34c96725" style="color: rgb(0, 0, 0);">为子系统中的一组接口提供一个一致的界面，此模式定义了一个高层接口，这个接口使得这一子系统更加容易使用。</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(235, 235, 235);" data-col="1"><div
                            class="ne-td-content"><ne-p id="u9cb53dc6"
                                data-lake-id="u9cb53dc6" ne-alignment="center"><ne-text
                                    id="u9a0958e0" style="color: rgb(0, 0, 0);">享元模式</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="u7142bac8" data-lake-id="u7142bac8"
                                ne-alignment="center"><ne-text id="uce5d3530"
                                    style="color: rgb(251, 2, 7);">★☆☆☆☆</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(235, 235,
                        235);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="u85efdd02"
                                data-lake-id="u85efdd02" ne-alignment="center"><ne-text
                                    id="uaac933b7" style="color: rgb(0, 0, 0);">以共享的方式高效的支持大量的细粒度的对象。</ne-text><ne-text
                                    id="ud44b647e" style="color: rgb(0, 0, 0);">
                                </ne-text><ne-text id="u2f250266" style="color:
                                    rgb(0, 0, 0);"> </ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(254, 220, 79);" data-col="1"><div
                            class="ne-td-content"><ne-p id="u7940cb36"
                                data-lake-id="u7940cb36" ne-alignment="center"><ne-text
                                    id="u04c513d7" style="color: rgb(0, 0, 0);">代理模式</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p><ne-p
                                id="u368b3272" data-lake-id="u368b3272"
                                ne-alignment="center"><ne-text id="u187f4d0c"
                                    style="color: rgb(251, 2, 7);">★★★★☆</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(255, 255,
                        255);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="u36e2520b"
                                data-lake-id="u36e2520b" ne-alignment="center"><ne-text
                                    id="u553b5325" style="color: rgb(0, 0, 0);">为其他对象提供一种代理以控制对这个对象的访问。</ne-text><span
                                    class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr></tbody></table>

## 5 行为型模式
<table class="ne-table" style="width: 684px;"><colgroup><col
                    width="124"><col width="187"><col width="187"><col
                    width="185"></colgroup><tbody class="ne-table-inner"><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(71, 119, 29);" data-col="0"><div
                            class="ne-td-content"><ne-p id="u5b1a5694"
                                data-lake-id="u5b1a5694"><ne-text id="uf36d7e23"
                                    ne-bold="true" style="color: rgb(255, 255,
                                    255);" ne-fontsize="14">模式名称</ne-text><ne-text
                                    id="u54119bf2" style="color: rgb(38, 38,
                                    38);" ne-fontsize="14">
                                </ne-text><span class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(71, 119,
                        29);" data-col="1"><div class="ne-td-content"><ne-p
                                id="ub3759123" data-lake-id="ub3759123"><ne-text
                                    id="udc4a9a00" ne-bold="true" style="color:
                                    rgb(255, 255, 255);" ne-fontsize="14">模式名称</ne-text><ne-text
                                    id="u2cf4de8f" style="color: rgb(38, 38,
                                    38);" ne-fontsize="14">
                                </ne-text><span class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(71, 119,
                        29);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="ub76e5cb7"
                                data-lake-id="ub76e5cb7"><ne-text id="uce90351a"
                                    ne-bold="true" style="color: rgb(255, 255,
                                    255);" ne-fontsize="14">作用</ne-text><ne-text
                                    id="u990daf58" style="color: rgb(38, 38,
                                    38);" ne-fontsize="14">
                                </ne-text><span class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(108, 108, 108);" rowspan="11" data-col="0"><div
                            class="ne-td-content"><ne-p id="uc7a79d32"
                                data-lake-id="uc7a79d32"><ne-text id="ubc01d184"
                                    ne-bold="true" style="color: rgb(255, 255,
                                    255);" ne-fontsize="14">行为型模式</ne-text><ne-text
                                    id="u9d821664" style="color: rgb(38, 38,
                                    38);" ne-fontsize="14">
                                </ne-text><ne-text id="u9d94812e" ne-bold="true"
                                    style="color: rgb(255, 255, 255);"
                                    ne-fontsize="14">Behavioral Pattern</ne-text><ne-text
                                    id="udac5738e" style="color: rgb(38, 38,
                                    38);" ne-fontsize="14">
                                </ne-text><ne-text id="u1613c933" ne-bold="true"
                                    style="color: rgb(255, 255, 255);"
                                    ne-fontsize="14">（11）</ne-text><ne-text
                                    id="uebaf489a" style="color: rgb(38, 38,
                                    38);" ne-fontsize="14">
                                </ne-text><span class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(235, 235,
                        235);" data-col="1"><div class="ne-td-content"><ne-p
                                id="uf5d3ded8" data-lake-id="uf5d3ded8"><ne-text
                                    id="u417ee999" style="color: rgb(0, 0, 0);"
                                    ne-fontsize="14">职责链模式</ne-text><ne-text
                                    id="u496d3861" style="color: rgb(38, 38,
                                    38);" ne-fontsize="14">
                                </ne-text><ne-text id="u115776a7" style="color:
                                    rgb(251, 2, 7);" ne-fontsize="14">★★☆☆☆</ne-text><ne-text
                                    id="u2fcd1a6c" style="color: rgb(38, 38,
                                    38);" ne-fontsize="14">
                                </ne-text><span class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(235, 235,
                        235);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="u46999399"
                                data-lake-id="u46999399"><ne-text id="u52219efc"
                                    style="color: rgb(0, 0, 0);"
                                    ne-fontsize="14">在该模式里，很多对象由每一个对象对其下家的引用而连接起来形成一条链。请求在这个链上传递，直到链上的某一个对象决定处理此请求，这使得系统可以在不影响客户端的情况下动态地重新组织链和分配责任。</ne-text><ne-text
                                    id="u7ed300de" style="color: rgb(38, 38,
                                    38);" ne-fontsize="14">
                                </ne-text><span class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(254, 220, 79);" data-col="1"><div
                            class="ne-td-content"><ne-p id="ub92bdf1e"
                                data-lake-id="ub92bdf1e"><ne-text id="u11ae663d"
                                    style="color: rgb(0, 0, 0);"
                                    ne-fontsize="14">命令模式</ne-text><ne-text
                                    id="u6a55e4c5" style="color: rgb(38, 38,
                                    38);" ne-fontsize="14">
                                </ne-text><ne-text id="u239abe0d" style="color:
                                    rgb(251, 2, 7);" ne-fontsize="14">★★★★☆</ne-text><ne-text
                                    id="u02687b04" style="color: rgb(38, 38,
                                    38);" ne-fontsize="14">
                                </ne-text><span class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="u4ef7c8e0"
                                data-lake-id="u4ef7c8e0"><ne-text id="ub193d130"
                                    style="color: rgb(0, 0, 0);"
                                    ne-fontsize="14">将一个请求封装为一个对象，从而使你可用不同的请求对客户端进行参数化；对请求排队或记录请求日志，以及支持可撤销的操作。</ne-text><ne-text
                                    id="u99ae1049" style="color: rgb(38, 38,
                                    38);" ne-fontsize="14">
                                </ne-text><span class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(235, 235, 235);" data-col="1"><div
                            class="ne-td-content"><ne-p id="u5636f8c1"
                                data-lake-id="u5636f8c1"><ne-text id="u2860fd82"
                                    style="color: rgb(0, 0, 0);"
                                    ne-fontsize="14">解释器模式</ne-text><ne-text
                                    id="u6f0272a9" style="color: rgb(38, 38,
                                    38);" ne-fontsize="14">
                                </ne-text><ne-text id="u4a3fffbb" style="color:
                                    rgb(251, 2, 7);" ne-fontsize="14">★☆☆☆☆</ne-text><ne-text
                                    id="ueaba340d" style="color: rgb(38, 38,
                                    38);" ne-fontsize="14">
                                </ne-text><span class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(235, 235,
                        235);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="uee748b3d"
                                data-lake-id="uee748b3d"><ne-text id="uc8211678"
                                    style="color: rgb(0, 0, 0);"
                                    ne-fontsize="14">如何为简单的语言定义一个语法，如何在该语言中表示一个句子，以及如何解释这些句子。</ne-text><ne-text
                                    id="u6d30efee" style="color: rgb(38, 38,
                                    38);" ne-fontsize="14">
                                </ne-text><span class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" data-col="1"><div
                            class="ne-td-content"><ne-p id="uc2074a1f"
                                data-lake-id="uc2074a1f"><ne-text id="u385706e8"
                                    style="color: rgb(0, 0, 0);"
                                    ne-fontsize="14">迭代器模式</ne-text><ne-text
                                    id="ubcfbb53f" style="color: rgb(38, 38,
                                    38);" ne-fontsize="14">
                                </ne-text><ne-text id="u403eb13d" style="color:
                                    rgb(251, 2, 7);" ne-fontsize="14">★☆☆☆☆</ne-text><ne-text
                                    id="uf197fe42" style="color: rgb(38, 38,
                                    38);" ne-fontsize="14">
                                </ne-text><span class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="uf04e2584"
                                data-lake-id="uf04e2584"><ne-text id="u321fd884"
                                    style="color: rgb(0, 0, 0);"
                                    ne-fontsize="14">提供了一种方法顺序来访问一个聚合对象中的各个元素，而又不需要暴露该对象的内部表示。</ne-text><ne-text
                                    id="u3f8f88e8" style="color: rgb(38, 38,
                                    38);" ne-fontsize="14">
                                </ne-text><span class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(235, 235, 235);" data-col="1"><div
                            class="ne-td-content"><ne-p id="uc7b0e66d"
                                data-lake-id="uc7b0e66d"><ne-text id="u30020c63"
                                    style="color: rgb(0, 0, 0);"
                                    ne-fontsize="14">中介者模式</ne-text><ne-text
                                    id="u28e49e1c" style="color: rgb(38, 38,
                                    38);" ne-fontsize="14">
                                </ne-text><ne-text id="ubcec17b9" style="color:
                                    rgb(251, 2, 7);" ne-fontsize="14">★★☆☆☆</ne-text><ne-text
                                    id="uf42c5449" style="color: rgb(38, 38,
                                    38);" ne-fontsize="14">
                                </ne-text><span class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(235, 235,
                        235);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="u742a8e8f"
                                data-lake-id="u742a8e8f"><ne-text id="u1bbb82b0"
                                    style="color: rgb(0, 0, 0);"
                                    ne-fontsize="14">定义一个中介对象来封装系列对象之间的交互。终结者使各个对象不需要显示的相互调用</ne-text><ne-text
                                    id="u01246bbf" style="color: rgb(0, 0, 0);"
                                    ne-fontsize="14">，从而使其耦合性松散，而且可以独立的改变他们之间的交互。</ne-text><ne-text
                                    id="u38ec074f" style="color: rgb(38, 38,
                                    38);" ne-fontsize="14">
                                </ne-text><span class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" data-col="1"><div
                            class="ne-td-content"><ne-p id="ub4376703"
                                data-lake-id="ub4376703"><ne-text id="u66ff29db"
                                    style="color: rgb(0, 0, 0);"
                                    ne-fontsize="14">备忘录模式</ne-text><ne-text
                                    id="ud4819267" style="color: rgb(38, 38,
                                    38);" ne-fontsize="14">
                                </ne-text><ne-text id="ua2a3dbc9" style="color:
                                    rgb(251, 2, 7);" ne-fontsize="14">★★☆☆☆</ne-text><ne-text
                                    id="u8184a200" style="color: rgb(38, 38,
                                    38);" ne-fontsize="14">
                                </ne-text><span class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="ue2c4007f"
                                data-lake-id="ue2c4007f"><ne-text id="u8c758477"
                                    style="color: rgb(0, 0, 0);"
                                    ne-fontsize="14">是在不破坏封装的前提下，捕获一个对象的内部状态，并在该对象之外保存这个状态。</ne-text><ne-text
                                    id="u4fe1f2cc" style="color: rgb(38, 38,
                                    38);" ne-fontsize="14">
                                </ne-text><span class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(254, 220, 79);" data-col="1"><div
                            class="ne-td-content"><ne-p id="u3cdc975e"
                                data-lake-id="u3cdc975e"><ne-text id="u8666c8c4"
                                    style="color: rgb(0, 0, 0);"
                                    ne-fontsize="14">观察者模式</ne-text><ne-text
                                    id="u825bb440" style="color: rgb(38, 38,
                                    38);" ne-fontsize="14">
                                </ne-text><ne-text id="u71b7a92f" style="color:
                                    rgb(251, 2, 7);" ne-fontsize="14">★★★★★</ne-text><ne-text
                                    id="ub5fabd9b" style="color: rgb(38, 38,
                                    38);" ne-fontsize="14">
                                </ne-text><span class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(235, 235,
                        235);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="u78b66520"
                                data-lake-id="u78b66520"><ne-text id="udb76356e"
                                    style="color: rgb(0, 0, 0);"
                                    ne-fontsize="14">定义对象间的一种一对多的依赖关系，当一个对象的状态发生改变时，所有依赖于它的对象都得到通知并被自动更新。</ne-text><ne-text
                                    id="uf72a6c6b" style="color: rgb(38, 38,
                                    38);" ne-fontsize="14">
                                </ne-text><span class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" data-col="1"><div
                            class="ne-td-content"><ne-p id="u0365aeb5"
                                data-lake-id="u0365aeb5"><ne-text id="ucd85d5b0"
                                    style="color: rgb(0, 0, 0);"
                                    ne-fontsize="14">状态模式</ne-text><ne-text
                                    id="u25ab981e" style="color: rgb(38, 38,
                                    38);" ne-fontsize="14">
                                </ne-text><ne-text id="u95f71416" style="color:
                                    rgb(251, 2, 7);" ne-fontsize="14">★★☆☆☆</ne-text><ne-text
                                    id="ue61f7f14" style="color: rgb(38, 38,
                                    38);" ne-fontsize="14">
                                </ne-text><span class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="u34462503"
                                data-lake-id="u34462503"><ne-text id="ud66338ae"
                                    style="color: rgb(0, 0, 0);"
                                    ne-fontsize="14">对象的行为，依赖于它所处的状态。</ne-text><ne-text
                                    id="ue75e8003" style="color: rgb(38, 38,
                                    38);" ne-fontsize="14">
                                </ne-text><ne-text id="u442a31ac" style="color:
                                    rgb(38, 38, 38);" ne-fontsize="14">
                                </ne-text><span class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(254, 220, 79);" data-col="1"><div
                            class="ne-td-content"><ne-p id="u3fa2c020"
                                data-lake-id="u3fa2c020"><ne-text id="u5c8121d0"
                                    style="color: rgb(0, 0, 0);"
                                    ne-fontsize="14">策略模式</ne-text><ne-text
                                    id="ua7651430" style="color: rgb(38, 38,
                                    38);" ne-fontsize="14">
                                </ne-text><ne-text id="u3205cc44" style="color:
                                    rgb(251, 2, 7);" ne-fontsize="14">★★★★☆</ne-text><ne-text
                                    id="ud5ebdef5" style="color: rgb(38, 38,
                                    38);" ne-fontsize="14">
                                </ne-text><span class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(235, 235,
                        235);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="u787621a6"
                                data-lake-id="u787621a6"><ne-text id="u2c34745e"
                                    style="color: rgb(0, 0, 0);"
                                    ne-fontsize="14">准备一组算法，并将每一个算法封装起来，使得它们可以互换。</ne-text><ne-text
                                    id="u8e8d9008" style="color: rgb(38, 38,
                                    38);" ne-fontsize="14">
                                </ne-text><span class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(254, 220, 79);" data-col="1"><div
                            class="ne-td-content"><ne-p id="u4f8d4b57"
                                data-lake-id="u4f8d4b57"><ne-text id="u18cab91b"
                                    style="color: rgb(0, 0, 0);"
                                    ne-fontsize="14">模板方法模式</ne-text><ne-text
                                    id="u10ada684" style="color: rgb(38, 38,
                                    38);" ne-fontsize="14">
                                </ne-text><ne-text id="u1c6b7ea9" style="color:
                                    rgb(251, 2, 7);" ne-fontsize="14">★★★☆☆</ne-text><ne-text
                                    id="u0d72c11a" style="color: rgb(38, 38,
                                    38);" ne-fontsize="14">
                                </ne-text><span class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="u0d99892f"
                                data-lake-id="u0d99892f"><ne-text id="u6b6111c9"
                                    style="color: rgb(0, 0, 0);"
                                    ne-fontsize="14">得子类可以不改变一个算法的结构即可重定义该算法的某些特定步骤。</ne-text><ne-text
                                    id="u8a497490" style="color: rgb(38, 38,
                                    38);" ne-fontsize="14">
                                </ne-text><span class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr><tr
                    class="ne-tr"><td class="ne-td" style="background-color:
                        rgb(235, 235, 235);" data-col="1"><div
                            class="ne-td-content"><ne-p id="ubc3053ef"
                                data-lake-id="ubc3053ef"><ne-text id="uf2f3323f"
                                    style="color: rgb(0, 0, 0);"
                                    ne-fontsize="14">访问者模式</ne-text><ne-text
                                    id="ucd2caad9" style="color: rgb(38, 38,
                                    38);" ne-fontsize="14">
                                </ne-text><ne-text id="u5dc3b997" style="color:
                                    rgb(251, 2, 7);" ne-fontsize="14">★☆☆☆☆</ne-text><ne-text
                                    id="u2da2ad7b" style="color: rgb(38, 38,
                                    38);" ne-fontsize="14">
                                </ne-text><span class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td><td
                        class="ne-td" style="background-color: rgb(235, 235,
                        235);" colspan="2" data-col="2"><div
                            class="ne-td-content"><ne-p id="u4b7e4649"
                                data-lake-id="u4b7e4649"><ne-text id="u8861daca"
                                    style="color: rgb(0, 0, 0);"
                                    ne-fontsize="14">表示一个作用于某对象结构中的各元素的操作，它使你可以在不改变各元素的类的前提下定义作用于这些元素的新操作。</ne-text><ne-text
                                    id="u51a7e5c7" style="color: rgb(38, 38,
                                    38);" ne-fontsize="14">
                                </ne-text><span class="ne-viewer-b-filler"
                                    ne-filler="block"><br></span></ne-p></div><div
                            class="ne-td-break" contenteditable="false"></div></td></tr></tbody></table>
