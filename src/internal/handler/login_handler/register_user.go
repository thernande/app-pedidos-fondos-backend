package loginhandler

import (
	"fmt"

	"github.com/thernande/app-pedidos-fondos-backend/internal/config/database"
	"github.com/thernande/app-pedidos-fondos-backend/internal/model/tables"
	v1 "github.com/thernande/app-pedidos-fondos-backend/proto/login/v1"
)

type registerUser struct {
	db   *database.Db
	User tables.Usuario
}

func NewRegisterUser(db *database.Db) *registerUser {
	return &registerUser{
		db:   db,
		User: tables.Usuario{},
	}
}

func (r *registerUser) SetUserWithProtobuf(user *v1.User) {
	r.User.Documento = user.Document
	r.User.Clave = user.Password
	r.User.Nombres = user.Name
	r.User.Apellidos = user.Lastname
	r.User.Telefono = user.Phone
	r.User.Correo = user.Email
	r.User.Empresa = user.Company
}

func (r *registerUser) RegisterUser() error {
	r.db.InitTransaction()

	if createErr := r.db.Instance.Create(&r.User).Error; createErr != nil {
		r.db.Instance.Rollback()
		r.db.Log.ErrorLogPrint(createErr.Error())
		return createErr
	}
	r.db.CommitTransaction()
	return nil
}

func (r *registerUser) DeleteUser(document string) error {
	r.db.InitTransaction()
	query := fmt.Sprintf("DELETE FROM usuarios WHERE documento = '%s'", document)
	if deleteErr := r.db.ExecByQuery(query); deleteErr != nil {
		r.db.Instance.Rollback()
		r.db.Log.ErrorLogPrint(deleteErr.Error())
		return deleteErr
	}
	r.db.CommitTransaction()
	return nil
}
