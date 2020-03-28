package player

import (
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/gioapp/goj/pkg/gel"
	"github.com/gioapp/goj/pkg/gelook"
	"golang.org/x/exp/shiny/materialdesign/icons"
	"image"
)

type item struct {
	i int
}

func (it *item) doSlide(n int) {
	it.i = it.i + n
}

type ScrollerGauge struct {
	Label      string
	Value      int
	Control    func(i int)
	scrollUnit float32

	ColorBg      string
	BorderRadius [4]float32
	OperateValue interface{}
	//Height       float32
	body *ScrollBarBody
	up   *ScrollerGaugeButton
	down *ScrollerGaugeButton
}

type ScrollBarBody struct {
	pressed      bool
	Do           func(interface{})
	ColorBg      string
	Position     float32
	Cursor       float32
	OperateValue interface{}
	Width        int
	CursorHeight float32
}

type ScrollerGaugeButton struct {
	icon         *gelook.DuoUIicon
	button       gelook.IconButton
	widgetButton *gel.Button
	Height       float32
	iconColor    string
	iconBgColor  string
	insetTop     float32
	insetRight   float32
	insetBottom  float32
	insetLeft    float32
	iconSize     float32
	iconPadding  float32
}

func (g *GoJoy) scrollerGauge(control func(i int)) *ScrollerGauge {
	iconUp, _ := gelook.NewDuoUIicon(icons.AVVolumeUp)
	iconDown, _ := gelook.NewDuoUIicon(icons.AVVolumeDown)
	itemValue := item{
		i: 0,
	}
	up := &ScrollerGaugeButton{
		icon:         iconUp,
		button:       g.Theme.IconButton(iconUp),
		widgetButton: new(gel.Button),
		Height:       16,
		iconColor:    "ff445588",
		iconBgColor:  "ff882266",
		insetTop:     0,
		insetRight:   0,
		insetBottom:  0,
		insetLeft:    0,
		iconSize:     32,
		iconPadding:  0,
	}
	down := &ScrollerGaugeButton{
		icon:         iconUp,
		button:       g.Theme.IconButton(iconDown),
		widgetButton: new(gel.Button),
		Height:       16,
		iconSize:     16,
		iconColor:    "ff445588",
		iconBgColor:  "ff882266",
	}
	body := &ScrollBarBody{
		pressed:  false,
		ColorBg:  "",
		Position: 0,
		Cursor:   0,
		Do: func(n interface{}) {
			itemValue.doSlide(n.(int))
		},
		OperateValue: 1,
		CursorHeight: 30,
	}
	return &ScrollerGauge{
		ColorBg:      "ff885566",
		BorderRadius: [4]float32{},
		OperateValue: 1,
		Control:      control,
		//Value:value,
		//ListPosition: 0,
		//Height: 16,
		body: body,
		up:   up,
		down: down,
	}
}
func (s *ScrollerGaugeButton) ScrollerGaugeButton() *gelook.IconButton {
	button := s.button
	button.Inset.Top = unit.Dp(0)
	button.Inset.Bottom = unit.Dp(0)
	button.Inset.Right = unit.Dp(0)
	button.Inset.Left = unit.Dp(0)
	button.Size = unit.Dp(32)
	button.Padding = unit.Dp(0)
	return &button
}
func (s *ScrollerGauge) Layout(gtx *layout.Context) {
	//fmt.Println(s.Value)

	layout.Flex{
		Axis: layout.Horizontal,
	}.Layout(gtx,
		layout.Rigid(func() {
			for s.up.widgetButton.Clicked(gtx) {
				//p.panelContent.Position.Offset = p.panelContent.Position.Offset - int(s.body.CursorHeight)
			}
			s.up.ScrollerGaugeButton().Layout(gtx, s.up.widgetButton)
		}),
		layout.Flexed(1, func() {
			s.bodyLayout(gtx)
		}),
		layout.Rigid(func() {
			for s.down.widgetButton.Clicked(gtx) {
				//p.panelContent.Position.Offset = p.panelContent.Position.Offset + int(s.body.CursorHeight)
			}
			s.down.ScrollerGaugeButton().Layout(gtx, s.down.widgetButton)
		}),
	)
}

func (s *ScrollerGauge) bodyLayout(gtx *layout.Context) {
	for _, e := range gtx.Events(s.body) {
		if e, ok := e.(pointer.Event); ok {
			//s.body.Position = e.Position.X - (s.body.CursorHeight / 2)
			s.body.Position = e.Position.X
			switch e.Type {
			case pointer.Press:
				s.body.pressed = true
				s.body.Do(s.body.OperateValue)
				//list.Position.First = int(s.Position)
			case pointer.Release:
				s.body.pressed = false
			}
		}
	}
	cs := gtx.Constraints
	//s.body.Width = cs.Width.Max
	s.scrollUnit = float32(cs.Width.Max) / float32(s.Value)
	sliderBg := HexARGB("ff558899")
	colorBg := HexARGB("ff30cfcf")
	colorBorder := HexARGB("ffcf3030")
	border := unit.Dp(0)
	if s.body.pressed {
		if s.body.Position >= 0 && s.body.Position <= (float32(cs.Width.Max)) {
			s.body.Cursor = s.body.Position
			//s.Value = int(s.body.Position)
			s.Control(int(s.body.Position / s.scrollUnit))
			//fmt.Println("aaaaa", int(s.body.Position/s.scrollUnit))
			//fmt.Println("value", s.Value)
			//fmt.Println("scrollUnit", s.scrollUnit)
			//fmt.Println("scrollUnit", s.scrollUnit)
			//fmt.Println("Position", s.body.Position)
			//s.Position.Offset = int(s.body.Cursor / s.scrollUnit)
		}
		colorBg = HexARGB("ffcf30cf")
		colorBorder = HexARGB("ff303030")
		border = unit.Dp(0)
	}
	pointer.Rect(
		image.Rectangle{Max: image.Point{X: cs.Width.Max, Y: cs.Height.Max}},
	).Add(gtx.Ops)
	pointer.InputOp{Key: s.body}.Add(gtx.Ops)
	DrawRectangle(gtx, cs.Width.Max, cs.Height.Max, colorBorder, [4]float32{0, 0, 0, 0}, unit.Dp(0))
	layout.UniformInset(border).Layout(gtx, func() {
		cs := gtx.Constraints
		DrawRectangle(gtx, cs.Width.Max, cs.Height.Max, colorBg, [4]float32{0, 0, 0, 0}, unit.Dp(0))
		//cs := gtx.Constraints
		layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func() {
				layout.W.Layout(gtx, func() {
					DrawRectangle(gtx, int(s.body.Cursor), int(s.body.CursorHeight), sliderBg, [4]float32{5, 5, 5, 5}, unit.Dp(0))
				})
			}),
		)
	})
	//fmt.Println("Cursor", s.body.Cursor)
	//fmt.Println("Position", s.body.Position)
}
