package module

import (
	"github.com/JustTalDevelops/nbs"
	"github.com/sandertv/gophertunnel/minecraft"
	"go.uber.org/atomic"
	"log"
)

var songs = []*nbs.Song{}

type MusicPlayer struct {
	Module

	kill atomic.Bool

	conn *minecraft.Conn

	song *nbs.Song

	l *log.Logger
}

func (m *MusicPlayer) Start(addr string) {}

func (m *MusicPlayer) Quit() {}
