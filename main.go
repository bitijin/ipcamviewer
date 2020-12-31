package main

import (
	"bufio"
	"fmt"
	vlc "github.com/adrg/libvlc-go/v3"
	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"log"
	"os"
	"strings"
)

const appID = "com.bitijin.IPCamViewer"

var loadCamera = 0

var cameraTable = [5][5]string{
	{"example","0.0.0.0.0","video","username","password"},
	{"example","0.0.0.0.0","video","username","password"},
	{"example","0.0.0.0.0","video","username","password"},
	{"example","0.0.0.0.0","video","username","password"},
	{"example","0.0.0.0.0","video","username","password"},
}

func setPlayerWindow(player *vlc.Player, window *gdk.Window) error {
	return player.SetXWindow(window.GetXID())
}

func builderGetObject(builder *gtk.Builder, name string) glib.IObject {
	obj, _ := builder.GetObject(name)
	return obj
}

func assertErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func assertConv(ok bool) {
	if !ok {
		log.Panic("invalid widget conversion")
	}
}

func playerReleaseMedia(player *vlc.Player) {
	player.Stop()
	if media, _ := player.Media(); media != nil {
		media.Release()
	}
}

func createCameras(){
	homedir, _ := os.UserHomeDir()
	if _, err := os.Stat(homedir+"/.local/share/ipcamviewer/cameras.list"); os.IsNotExist(err) {
		fmt.Println("Creating Example Camera List")
		err := os.Mkdir(homedir+"/.local/share/ipcamviewer/", 0777)
		if err != nil {
			log.Fatal(err)
		}
		file, err := os.Create(homedir+"/.local/share/ipcamviewer/cameras.list")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		for i := 0; i < 5; i++ {
			file.WriteString("Example:0.0.0.0:video:user:pass\n")
		}
		file.Close()
	}
}

func readCameras(){
	homedir, _ := os.UserHomeDir()
	file, err := os.Open(homedir+"/.local/share/ipcamviewer/cameras.list")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	x := 0
	for scanner.Scan(){
		if strings.HasPrefix(scanner.Text(),"#") == false{
			data := strings.Split(scanner.Text(),":")
			for i := 0; i < len(data); i++ {
				cameraTable[x][i] = data[i]
			}
			x++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func writeCameras(){
	fmt.Println(cameraTable)
	homedir, _ := os.UserHomeDir()
	file, err := os.Create(homedir+"/.local/share/ipcamviewer/cameras.list")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for i := 0; i < 5; i++ {
		camera := ""
		for x:= 0; x <5; x++ {
			camera += cameraTable[i][x]
			if x != 4 {
				camera += ":"
			}
		}
		file.WriteString(camera+"\n")
		fmt.Println(camera)
	}
	file.Close()
}

func main() {
	// Initialize libVLC module.
	err := vlc.Init("--quiet", "--no-xlib")
	assertErr(err)

	// Create a new player.
	player, err := vlc.NewPlayer()
	assertErr(err)

	// Create new GTK application.
	app, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	assertErr(err)

	app.Connect("activate", func() {
		// Load application layout.
		//builder, err := gtk.BuilderNewFromFile("layout.glade")
		builder, err := gtk.BuilderNewFromString(layout)
		assertErr(err)

		// Get application window.
		appWin, ok := builderGetObject(builder, "appWindow").(*gtk.ApplicationWindow)
		assertConv(ok)

		cameraWin, ok := builderGetObject(builder, "cameraListWindow").(*gtk.ApplicationWindow)
		assertConv(ok)

		addCameraWin, ok := builderGetObject(builder, "addCamera").(*gtk.ApplicationWindow)
		assertConv(ok)

		name0, _ := builderGetObject(builder, "name0").(*gtk.Label)
		name1, _ := builderGetObject(builder, "name1").(*gtk.Label)
		name2, _ := builderGetObject(builder, "name2").(*gtk.Label)
		name3, _ := builderGetObject(builder, "name3").(*gtk.Label)
		name4, _ := builderGetObject(builder, "name4").(*gtk.Label)

		address0, _ := builderGetObject(builder, "address0").(*gtk.Label)
		address1, _ := builderGetObject(builder, "address1").(*gtk.Label)
		address2, _ := builderGetObject(builder, "address2").(*gtk.Label)
		address3, _ := builderGetObject(builder, "address3").(*gtk.Label)
		address4, _ := builderGetObject(builder, "address4").(*gtk.Label)

		stream0, _ := builderGetObject(builder, "stream0").(*gtk.Label)
		stream1, _ := builderGetObject(builder, "stream1").(*gtk.Label)
		stream2, _ := builderGetObject(builder, "stream2").(*gtk.Label)
		stream3, _ := builderGetObject(builder, "stream3").(*gtk.Label)
		stream4, _ := builderGetObject(builder, "stream4").(*gtk.Label)

		newName, _ := builderGetObject(builder, "newName").(*gtk.Entry)
		newAddress, _ := builderGetObject(builder, "newAddress").(*gtk.Entry)
		newStream, _ := builderGetObject(builder, "newStream").(*gtk.Entry)
		newUsername, _ := builderGetObject(builder, "newUsername").(*gtk.Entry)
		newPassword, _ := builderGetObject(builder, "newPassword").(*gtk.Entry)
		cameraPosition, _ := builderGetObject(builder, "cameraPosition").(*gtk.SpinButton)

		// Get play button.
		//playButton, ok := builderGetObject(builder, "playButton").(*gtk.Button)
		//assertConv(ok)

		// Get connect button.

		// Get urlInput
		urlInput, ok := builderGetObject(builder, "urlInput").(*gtk.Entry)
		assertConv(ok)

		// Add builder signal handlers.
		signals := map[string]interface{}{
			"onRealizePlayerArea": func(playerArea *gtk.DrawingArea) {
				// Set window for the player.
				playerWindow, err := playerArea.GetWindow()
				assertErr(err)
				err = setPlayerWindow(player, playerWindow)
				assertErr(err)
			},
			"onDrawPlayerArea": func(playerArea *gtk.DrawingArea, cr *cairo.Context) {
				cr.SetSourceRGB(0, 0, 0)
				cr.Paint()
			},
			/*"onClickConnectButton": func() {
				connectButton.SetLabel("Connecting...")
				url, _ := urlInput.GetText()
				fmt.Println(url)
				if _, err := player.LoadMediaFromURL(url); err != nil {
					log.Printf("Cannot load selected media: %s\n", err)
					return
				}
				player.Play()
				playButton.SetLabel("gtk-media-stop")
				connectButton.SetLabel("Connect")
			},*/
			"onActivateQuit": func() {
				app.Quit()
			},
			"onClickPlayButton": func(playButton *gtk.Button) {
				if media, _ := player.Media(); media == nil {
					return
				}

				if player.IsPlaying() {
					player.Stop()
					playButton.SetLabel("gtk-media-play")
				} else {
					player.Play()
					playButton.SetLabel("gtk-media-stop")
				}
			},
			"onClickCameraList": func() {
				cameraWin.SetVisible(true)
				name0.SetText(cameraTable[0][0])
				name1.SetText(cameraTable[1][0])
				name2.SetText(cameraTable[2][0])
				name3.SetText(cameraTable[3][0])
				name4.SetText(cameraTable[4][0])

				address0.SetText(cameraTable[0][1])
				address1.SetText(cameraTable[1][1])
				address2.SetText(cameraTable[2][1])
				address3.SetText(cameraTable[3][1])
				address4.SetText(cameraTable[4][1])

				stream0.SetText(cameraTable[0][2])
				stream1.SetText(cameraTable[1][2])
				stream2.SetText(cameraTable[2][2])
				stream3.SetText(cameraTable[3][2])
				stream4.SetText(cameraTable[4][2])
			},
			"onClickCloseCameraList": func() {
				cameraWin.SetVisible(false)
			},
			"onClickAddCamera": func() {
				addCameraWin.SetVisible(true)
			},
			"onClickCloseAddCamera": func() {
				addCameraWin.SetVisible(false)
			},
			"onClickConfirmAddCamera": func() {
				pos := int(cameraPosition.GetValue())-1
				cameraTable[pos][0], _ = newName.GetText()
				cameraTable[pos][1], _ = newAddress.GetText()
				cameraTable[pos][2], _ = newStream.GetText()
				cameraTable[pos][3], _ = newUsername.GetText()
				cameraTable[pos][4], _ = newPassword.GetText()

				writeCameras()

				name0.SetText(cameraTable[0][0])
				name1.SetText(cameraTable[1][0])
				name2.SetText(cameraTable[2][0])
				name3.SetText(cameraTable[3][0])
				name4.SetText(cameraTable[4][0])

				address0.SetText(cameraTable[0][1])
				address1.SetText(cameraTable[1][1])
				address2.SetText(cameraTable[2][1])
				address3.SetText(cameraTable[3][1])
				address4.SetText(cameraTable[4][1])

				stream0.SetText(cameraTable[0][2])
				stream1.SetText(cameraTable[1][2])
				stream2.SetText(cameraTable[2][2])
				stream3.SetText(cameraTable[3][2])
				stream4.SetText(cameraTable[4][2])

				addCameraWin.SetVisible(false)
			},
			"onClickView0": func() {
				loadCamera = 0
				urlInput.SetText(cameraTable[0][0])
				cameraWin.SetVisible(false)
				url := "rtsp://"+cameraTable[loadCamera][3]+":"+cameraTable[loadCamera][4]+"@"+cameraTable[loadCamera][1]+":554/"+cameraTable[loadCamera][2]
				fmt.Println(url)
				player.LoadMediaFromURL(url)
			},
			"onClickView1": func() {
				loadCamera = 1
				urlInput.SetText(cameraTable[1][0])
				cameraWin.SetVisible(false)
				url := "rtsp://"+cameraTable[loadCamera][3]+":"+cameraTable[loadCamera][4]+"@"+cameraTable[loadCamera][1]+":554/"+cameraTable[loadCamera][2]
				fmt.Println(url)
				player.LoadMediaFromURL(url)
			},
			"onClickView2": func() {
				loadCamera = 2
				urlInput.SetText(cameraTable[2][0])
				cameraWin.SetVisible(false)
				url := "rtsp://"+cameraTable[loadCamera][3]+":"+cameraTable[loadCamera][4]+"@"+cameraTable[loadCamera][1]+":554/"+cameraTable[loadCamera][2]
				fmt.Println(url)
				player.LoadMediaFromURL(url)
			},
			"onClickView3": func() {
				loadCamera = 3
				urlInput.SetText(cameraTable[3][0])
				cameraWin.SetVisible(false)
				url := "rtsp://"+cameraTable[loadCamera][3]+":"+cameraTable[loadCamera][4]+"@"+cameraTable[loadCamera][1]+":554/"+cameraTable[loadCamera][2]
				fmt.Println(url)
				player.LoadMediaFromURL(url)
			},
			"onClickView4": func() {
				loadCamera = 4
				urlInput.SetText(cameraTable[4][0])
				cameraWin.SetVisible(false)
				url := "rtsp://"+cameraTable[loadCamera][3]+":"+cameraTable[loadCamera][4]+"@"+cameraTable[loadCamera][1]+":554/"+cameraTable[loadCamera][2]
				fmt.Println(url)
				player.LoadMediaFromURL(url)
			},
		}
		builder.ConnectSignals(signals)

		createCameras()
		readCameras()
		fmt.Println(cameraTable)

		appWin.ShowAll()
		app.AddWindow(appWin)
		urlInput.SetText(cameraTable[0][0])
		url := "rtsp://"+cameraTable[loadCamera][3]+":"+cameraTable[loadCamera][4]+"@"+cameraTable[loadCamera][1]+":554/"+cameraTable[loadCamera][2]
		fmt.Println(url)
		player.LoadMediaFromURL(url)

	})

	// Cleanup on exit.
	app.Connect("shutdown", func() {
		playerReleaseMedia(player)
		player.Release()
		vlc.Release()
	})

	// Launch the application.
	os.Exit(app.Run(os.Args))
}
