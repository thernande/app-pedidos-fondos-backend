package loginhandler

import (
	"github.com/pkg/errors"
	"github.com/thernande/app-pedidos-fondos-backend/internal/config/database"
	"github.com/thernande/app-pedidos-fondos-backend/internal/model/tables"
	"github.com/thernande/app-pedidos-fondos-backend/pkg/encrypt"
	"github.com/thernande/app-pedidos-fondos-backend/pkg/errores"
	"golang.org/x/crypto/bcrypt"
)

type loginUser struct {
	Db    *database.Db
	User  tables.Usuario
	Token string
}

func NewLoginUser(db *database.Db) *loginUser {
	return &loginUser{
		Db:   db,
		User: tables.Usuario{},
	}
}

func (l *loginUser) LoginUser(password string) error {
	if err := l.Db.Instance.Where("documento = ?", l.User.Documento).First(&l.User).Error; err != nil {
		l.Db.Log.ErrorLogPrint(err.Error())
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(l.User.Clave), []byte(password)); err != nil {
		l.Db.Log.ErrorLogPrint(err.Error())
		return err
	}

	var err = make(chan errores.ChannelErrors)
	go encrypt.GenerateJWTLoginUsuario(l.User.ID, l.User.Documento, &l.Token, err)

	if jwtError := <-err; jwtError.Condition {
		l.Db.Log.ErrorLogPrint(jwtError.Error)
		return errors.New(jwtError.Error)
	}
	return nil
}
