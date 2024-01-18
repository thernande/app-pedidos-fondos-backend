package loginhandler

import (
	"github.com/thernande/app-pedidos-fondos-backend/internal/config/database"
	"github.com/thernande/app-pedidos-fondos-backend/internal/model/tables"
	v1 "github.com/thernande/app-pedidos-fondos-backend/proto/login/v1"
)

type register struct {
	db   *database.Db
	User tables.Usuario
}

func New(db *database.Db) *register {
	return &register{
		db:   db,
		User: tables.Usuario{},
	}
}

func (r *register) SetUserWithProtobuf(user *v1.User) {
	r.User.Documento = user.Document
	r.User.Clave = user.Password
	r.User.Nombres = user.Name
	r.User.Apellidos = user.Lastname
	r.User.Telefono = user.Phone
	r.User.Correo = user.Email
	r.User.Empresa = user.Company
}

func (r *register) RegisterUser() error {
	r.db.InitTransaction()

	if createErr := r.db.Instance.Create(&r.User).Error; createErr != nil {
		r.db.Instance.Rollback()
		r.db.Log.ErrorLogPrint(createErr.Error())
		return createErr
	}
	r.db.CommitTransaction()
	return nil
}
