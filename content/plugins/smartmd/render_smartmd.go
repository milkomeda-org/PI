package main

type SmartMD struct {
}

func (s SmartMD) Render(src string) (string, error) {
	src += "\n 已渲染"
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
