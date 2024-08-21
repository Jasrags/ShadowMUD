package screen

import (
	"github.com/Jasrags/ShadowMUD/common/user"
	"github.com/Jasrags/ShadowMUD/config"

	"github.com/sirupsen/logrus"
)

type Screens struct {
	log  *logrus.Entry
	user *user.User
	cfg  *config.Server
}

func New(u *user.User, cfg *config.Server) *Screens {
	s := &Screens{
		user: u,
		cfg:  cfg,
	}
	s.log = logrus.WithFields(logrus.Fields{"user": u.Username, "id": u.ID, "package": "screen"})
	return s
}
