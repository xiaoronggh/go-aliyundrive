package aliyundrive

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/mdp/qrterminal"
	qrcode2 "github.com/skip2/go-qrcode"
	"os"
)

type LoginQrCode interface {
	Show(content string)
	Close()
	EventNew(res *queryQrCodeResp)
	EventScanned(res *queryQrCodeResp)
	EventExpired(res *queryQrCodeResp)
	EventCanceled(res *queryQrCodeResp)
	EventConfirmed(res *queryQrCodeResp)
}

type LoginQrTerminal struct {
	scaned bool
}

func (*LoginQrTerminal) Show(content string) {
	qrterminal.Generate(content, qrterminal.L, os.Stdout)
}
func (*LoginQrTerminal) EventNew(res *queryQrCodeResp) {
}
func (t *LoginQrTerminal) EventScanned(res *queryQrCodeResp) {
	if !t.scaned {
		fmt.Println("扫描成功, 请在手机上根据提示确认登录")
	}
	t.scaned = true
}
func (*LoginQrTerminal) EventExpired(res *queryQrCodeResp) {
}
func (*LoginQrTerminal) EventCanceled(res *queryQrCodeResp) {
}
func (*LoginQrTerminal) EventConfirmed(res *queryQrCodeResp) {
}
func (*LoginQrTerminal) Close() {
}

type LoginSmallQrCode struct {
	scaned bool
}

func (s *LoginSmallQrCode) Show(content string) {
	obj, _ := qrcode2.New(content, qrcode2.Low)
	fmt.Print(obj.ToSmallString(false))
}
func (s *LoginSmallQrCode) EventNew(res *queryQrCodeResp) {
}
func (s *LoginSmallQrCode) EventScanned(res *queryQrCodeResp) {
}
func (s *LoginSmallQrCode) EventExpired(res *queryQrCodeResp) {
}
func (s *LoginSmallQrCode) EventCanceled(res *queryQrCodeResp) {
}
func (s *LoginSmallQrCode) EventConfirmed(res *queryQrCodeResp) {
}
func (s *LoginSmallQrCode) Close() {
}

type LoginUIQrCode struct {
	w    fyne.Window
	l    *widget.Label
	show bool
}

func (u *LoginUIQrCode) Show(content string) {
	app := app.New()
	u.w = app.NewWindow("登录")
	u.l = widget.NewLabel("Please use the mobile client to scan the code to log in")
	png, err := qrcode2.Encode(content, qrcode2.Medium, 256)
	if err != nil {
		panic(err)
	}
	res := &fyne.StaticResource{
		StaticName:    "qrcode",
		StaticContent: png,
	}
	image := &canvas.Image{
		Resource: res,
	}
	image.FillMode = canvas.ImageFillOriginal

	u.w.SetContent(container.NewVBox(
		image,
		u.l,
	))
	u.show = true
	u.w.ShowAndRun()
}
func (*LoginUIQrCode) EventNew(res *queryQrCodeResp) {
}
func (u *LoginUIQrCode) EventScanned(res *queryQrCodeResp) {
	u.l.SetText("The scan is successful. Please confirm your login according to the prompt on your mobile phone")
}
func (u *LoginUIQrCode) EventExpired(res *queryQrCodeResp) {
	u.l.SetText("qr code expired")
}
func (u *LoginUIQrCode) EventCanceled(res *queryQrCodeResp) {
	u.l.SetText("cancel login")
}
func (u *LoginUIQrCode) EventConfirmed(res *queryQrCodeResp) {
}
func (u *LoginUIQrCode) Close() {
	if u.w != nil && u.show {
		u.show = false
		u.w.Close()
	}
}
