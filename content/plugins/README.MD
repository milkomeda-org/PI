go build -gcflags="all=-N -l" -buildmode=plugin

### 自定义插件
Instance() interface{} Instance方法返回插件实例

实例应实现Plugin接口和对应类型Plugin接口

> 例如，下例实现了Plugin和RenderingPlugin

```
package main

type SmartMD struct {
}

func (s SmartMD) Render(src string) (string, error) {
	return src, nil
}

func (s SmartMD) Name() string {
	return "start_md"
}

func (s SmartMD) Type() string {
	return "render"
}

func Instance() interface{} {
	return SmartMD{}
}
```