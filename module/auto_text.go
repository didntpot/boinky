package module

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/cqdetdev/boinky/data"
	"github.com/cqdetdev/boinky/util"
	"github.com/df-mc/atomic"
	"github.com/mgutz/ansi"
	"github.com/sandertv/gophertunnel/minecraft"
	"github.com/sandertv/gophertunnel/minecraft/protocol/login"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

type AutoText struct {
	Module

	kill atomic.Bool

	texts []string

	conn *minecraft.Conn

	l *log.Logger
}

func (m *AutoText) Start(addr string) {
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

			if text, ok := pk.(*packet.Text); ok && !text.NeedsTranslation {
				fmt.Println(util.MinecraftToAscii(text.Message) + ansi.Reset)
			}

			for _, str := range m.texts {
				conn.WritePacket(&packet.Text{TextType: packet.TextTypeChat, Message: str})
			}
		}
	}()

	for s.Scan() {
		msg := s.Text()
		if msg == "!exit" {
			m.Quit()
			break
		}
		conn.WritePacket(&packet.Text{TextType: packet.TextTypeChat, Message: msg})
	}

	if s.Err() != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func (m *AutoText) Quit() {
	m.kill.Store(true)
	m.conn.Close()
}
