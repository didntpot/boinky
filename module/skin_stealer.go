package module

import (
	"bufio"
	"fmt"
	"github.com/cqdetdev/boinky/data"
	"github.com/cqdetdev/boinky/util"
	"github.com/sandertv/gophertunnel/minecraft"
	"github.com/sandertv/gophertunnel/minecraft/protocol/login"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
	"github.com/df-mc/atomic"
	"image"
	"image/png"
	"log"
	"os"
)

type SkinStealer struct {
	Module

	kill atomic.Bool

	conn *minecraft.Conn

	l *log.Logger
}

func (m *SkinStealer) Start(addr string) {
	s := bufio.NewScanner(os.Stdin)
	conn, err := minecraft.Dialer{
		ClientData:  login.ClientData{},
		TokenSource: data.TokenSrc,
	}.Dial("raknet", addr)

	m.conn = conn

	if err != nil {
		m.l.Println(err)
		return
	}

	err = conn.DoSpawn()

	if err != nil {
		m.l.Println(err)
		return
	}

	fmt.Println("Logged in as " + conn.IdentityData().DisplayName)

	go func() {
		for {
			pk, _ := conn.ReadPacket()

			if err != nil {
				m.l.Println(err)
				return
			}

			if pl, ok := pk.(*packet.PlayerList); ok {
				for _, player := range pl.Entries {
					skin := player.Skin
					name := util.StripColor(player.Username)
					img := image.NewRGBA(image.Rect(0, 0, int(skin.SkinImageWidth), int(skin.SkinImageHeight)))
					img.Pix = skin.SkinData
					_ = os.Mkdir("skins/", 0777)
					file, _ := os.Create(fmt.Sprintf("skins/%s_skin.png", name))
					err = png.Encode(file, img)
					if err != nil {
						m.l.Println(err)
						return
					}
					file.Close()
					m.l.Printf("Successfully stole %s's skin\n", name)
				}
				m.Quit()
			}
		}
	}()

	for s.Scan() {
		msg := s.Text()
		if msg == "!exit" {
			m.Quit()
			break
		}
	}

	if s.Err() != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func (m *SkinStealer) Quit() {
	m.kill.Swap(true)
}
