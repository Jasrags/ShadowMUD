package core

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/Jasrags/ShadowMUD/common"
	"github.com/Jasrags/ShadowMUD/config"
	"github.com/Jasrags/ShadowMUD/utils"
	"github.com/fatih/color"

	"github.com/gliderlabs/ssh"
	"github.com/sirupsen/logrus"
)

const (
	ConfigFilepath = "_data/config/server.yaml"
)

type World struct {
	config     config.Server
	characters common.Charcters
	zones      common.Zones
	rooms      common.Rooms
}

func NewWorld() *World {
	var serverConfig config.Server
	utils.LoadStructFromYAML(ConfigFilepath, &serverConfig)

	logrus.WithField("serverConfig", serverConfig).Info("Loaded server configuration")

	logrusLevel, err := logrus.ParseLevel(serverConfig.LogLevel)
	if err != nil {
		logrus.WithError(err).Warn("Could not parse log level, defaulting to INFO")
		logrusLevel = logrus.InfoLevel
	}
	logrus.SetLevel(logrusLevel)

	logrus.WithField("log_level", logrusLevel).Info("Logger level set")

	w := &World{
		config: serverConfig,
		zones:  make(common.Zones),
		rooms:  make(common.Rooms),
	}

	return w
}

func (w *World) Start() {
	logrus.WithFields(logrus.Fields{"host": w.config.Host, "port": w.config.Port}).Info("Starting server")
	ssh.Handle(w.Handler)

	server := &ssh.Server{
		BannerHandler:            w.BannerHandler,
		Handler:                  w.Handler,
		ConnectionFailedCallback: w.ConnectionFailedCallback,
		Addr:                     net.JoinHostPort(w.config.Host, w.config.Port),
		IdleTimeout:              w.config.Timeouts.Idle,
	}

	// Load data
	w.LoadData()

	if err := server.ListenAndServe(); err != nil {
		logrus.WithError(err).Error("Could not start server")
	}
}

func (w *World) ConnectionFailedCallback(conn net.Conn, err error) {
	defer conn.Close()
	logrus.WithError(err).Error("Connection failed")
}

func (w *World) BannerHandler(ctx ssh.Context) string {
	logrus.Debug("Sending banner")
	return color.New(color.FgHiGreen).Sprintln("Welcome to my SSH server, friend!")
	// return "Welcome to my SSH server, friend!\r\n"
}

func (w *World) HandleTimeout(s ssh.Session) {
	logrus.Debug("Timeout handler started")
	i := 0
	for {
		i += 1
		select {
		case <-time.After(time.Second):
			continue
		case <-s.Context().Done():
			logrus.Info("Connection closed: timeout")
			return
		}
	}
}

func (w *World) Handler(s ssh.Session) {
	logrus.WithFields(logrus.Fields{"user": s.User(), "remote_addr": s.RemoteAddr()}).Info("New connection")

	go w.HandleTimeout(s)

	c := common.NewCharacter(s)
	if ok := c.Authenticate(); !ok {
		color.New(color.FgHiRed).Fprintln(s, "Username or Password is incorrect")
		return
	}

	c.Load()

	w.characters[c.ID] = c
	defer func() {
		delete(w.characters, c.ID)
	}()

	if err := c.GameLoop(); err != nil {
		logrus.WithError(err).Error("Error in game loop")
	}
}

func (w *World) LoadData() {
	w.LoadZones()
	w.LoadRooms()
	w.BuildRooms()
}

func (w *World) LoadZones() {
	logrus.Info("Loading zones")
	rootpath := w.config.Data.BaseDir + w.config.Data.ZonesDir

	if err := filepath.Walk(rootpath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) == ".yaml" {
			var spec common.ZoneSpec
			if err := utils.LoadStructFromYAML(path, &spec); err != nil {
				logrus.WithError(err).WithFields(logrus.Fields{"path": path}).Error("Failed to load zone")
				return err
			}

			// Load the zone
			v := common.NewZone(&spec)
			w.zones[v.ID] = v
			logrus.WithFields(logrus.Fields{"path": path, "id": v.ID}).Debug("Found zone")
		}

		return nil
	}); err != nil {
		logrus.WithError(err).Error("file walk error")
	}

	logrus.Infof("Finished loading %d zones", len(w.zones))
}

func (w *World) LoadRooms() {
	logrus.Info("Loading rooms")

	rootpath := w.config.Data.BaseDir + w.config.Data.RoomsDir

	if err := filepath.Walk(rootpath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) == ".yaml" {
			var spec common.RoomSpec
			if err := utils.LoadStructFromYAML(path, &spec); err != nil {
				logrus.WithError(err).WithFields(logrus.Fields{"path": path}).Error("Failed to load room")
				return err
			}

			// Load our rooms
			v := common.NewRoom(&spec)
			logrus.WithFields(logrus.Fields{"zone_id": v.Spec.ZoneID, "path": path, "id": v.ID}).Debug("Found room")
			w.rooms[fmt.Sprintf("%s:%s", v.Spec.ZoneID, v.ID)] = v

		}

		return nil
	}); err != nil {
		logrus.WithError(err).Error("file walk error")
	}

	logrus.Infof("Finished loading %d rooms", len(w.rooms))
}

func (w *World) BuildRooms() {
	logrus.Info("Building rooms")
	for _, r := range w.rooms {
		logrus.WithFields(logrus.Fields{"zone_id": r.Spec.ZoneID, "room_id": r.ID}).Debug("Building room")
		// Add this room to the zone's room list
		w.zones[r.Spec.ZoneID].Rooms[r.ID] = r
		// Set the room's zone
		r.Zone = w.zones[r.Spec.ZoneID]
	}
}
