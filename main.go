package main

import (
	"os"
	"path"

	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"github.com/xfrr/goffmpeg/transcoder"
)

func uniq(strings []string) (ret []string) {
	return
}

func main() {
	gtk.Init(nil)
    window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
    window.SetPosition(gtk.WIN_POS_CENTER)
    window.SetTitle("GTK Go!")
    window.SetIconName("./img/conv.svg")
    window.Connect("destroy", func(ctx *glib.CallbackContext) {
        println("got destroy!", ctx.Data().(string))
        gtk.MainQuit()
    }, "App Closed!")


	button := gtk.NewButtonWithLabel("Button with label")
	button.SetImage("./img/penguin.jpg")
    button.Clicked(func() {
		messagedialog := gtk.NewMessageDialog(
			menuitem.GetTopLevelAsWindow(),
			gtk.DIALOG_MODAL,
			gtk.MESSAGE_INFO,
			gtk.BUTTONS_OK,
			entry.GetText())
		messagedialog.Response(func() {
			println("Dialog OK!")
			filechooserdialog := gtk.NewFileChooserDialog(
				"Choose File...",
				menuitem.GetTopLevelAsWindow(),
				gtk.FILE_CHOOSER_ACTION_OPEN,
				gtk.STOCK_OK,
				gtk.RESPONSE_ACCEPT)
			filter := gtk.NewFileFilter()
			filter.AddPattern("*.jpg")
			filechooserdialog.AddFilter(filter)
			filechooserdialog.Response(func() {
				inputPath := filechooserdialog.GetFilename()
				filechooserdialog.Destroy()
				outputPath := "/home/abanoub/pictures/conved.png"
				trans := new(transcoder.Transcoder)
				err := trans.Initialize(inputPath, outputPath)
				if err != nil {
					println("Error Occurred")
				}
				done := trans.Run(false)
				err = <-done
				println("conversion done! file is saved at: " + outputPath)

			})
			filechooserdialog.Run()
			messagedialog.Destroy()
		})
		messagedialog.Run()
	}
	window.Add(button)
}
