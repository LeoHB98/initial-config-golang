package utils

import (
	"fmt"
	"strings"

	"gitalpha.com/go/nome_projeto/config"
	_ "github.com/godror/godror"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
)

// ConnectDB realiza a conexão com o banco de dados
func ConnectDB(u *config.DBConfig) (*sqlx.DB, error) {

	db, err := sqlx.Open("godror", fmt.Sprintf("%s/%s@%s:%d/%s", u.DbUser, u.DbPass, u.DbHost, u.DbPort, u.DbSid))
	if err != nil {

		return nil, fmt.Errorf("erro ao abrir conexao:%v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("erro ao pingar no banco: %v", err)
	}

	db.Mapper = reflectx.NewMapperTagFunc("db",
		nil,
		func(s string) string {
			return strings.ToUpper(s)
		},
	)

	sqlx.BindDriver("godror", sqlx.NAMED)

	return db, nil
}
