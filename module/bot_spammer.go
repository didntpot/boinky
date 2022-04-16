package module

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/df-mc/atomic"
	"github.com/sandertv/gophertunnel/minecraft"
	"github.com/sandertv/gophertunnel/minecraft/protocol/login"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

type BotSpammer struct {
	Module

	name  string
	msg string
	limit int

	spam atomic.Bool
	kill atomic.Bool

	conns []*minecraft.Conn

	l *log.Logger
}

func (m *BotSpammer) Start(addr string) {
	s := bufio.NewScanner(os.Stdin)
	for i := 1; i != m.limit; i++ {
		if m.kill.Load() == true {
			break
		}
		go func(i int) {
			name := m.name + strconv.Itoa(i)
			conn, err := minecraft.Dialer{
				IdentityData: login.IdentityData{
					DisplayName: name,
				},
				ErrorLog: m.l,
			}.Dial("raknet", addr)

			
			if err != nil {
				m.l.Println(err)
				return
			}

			m.conns = append(m.conns, conn)

			err = conn.DoSpawn()

			if err != nil {
				m.l.Println(err)
				return
			}

			m.l.Printf("Bot '%s' has joined %s", name, addr)

			for {
				_, err := conn.ReadPacket()
				if err != nil {
					m.l.Println(err)
					return
				}

				if m.spam.Load() == true {
					conn.WritePacket(&packet.Text{
						TextType: packet.TextTypeChat,
						SourceName: m.name,
						Message: m.msg,
					})
				}				

			}
		}(i)
		for s.Scan() {
			msg := s.Text()
			if msg == "!exit" {
				m.Quit()
				break
			}
		}
		time.Sleep(1 * time.Second / 4)
	}
}

func (m *BotSpammer) Spam(spam bool) {
	m.spam.Swap(spam)
}

func (m *BotSpammer) Quit() {
	m.kill.Swap(true)
	for _, conn := range m.conns {
		conn.Close()
	}
}