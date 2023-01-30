package utils

import (
	"gitalpha.com/go/nome_projeto/config"
	"github.com/jlaffaye/ftp"
	"github.com/jmoiron/sqlx"
)

type Utils struct {
	*config.Config
	*sqlx.DB
	ftp *ftp.ServerConn
}

func Start(cfg *config.Config) *Utils {

	utils := New(cfg)
	dataConnect, err := ConnectDB(utils.DbConfig)
	if err != nil {
		ReceiverErrors(err, cfg.Email)
		return &Utils{}
	}

	ftpConnect, err := ConnectFTP(cfg.FTP)
	if err != nil {
		ReceiverErrors(err, cfg.Email)
		return &Utils{}
	}

	utils.DB = dataConnect
	utils.ftp = ftpConnect

	return &utils

}

func New(cfg *config.Config) Utils {
	return Utils{Config: cfg}
}

// // Close finaliza a conexão com o servidor FTP
// func (f *FTPServer) Close() error {
// 	return f.ftp.Quit()
// }

// // FTP retorna a conexão com o ftp
// func (f *FTPServer) FTP() *ftp.ServerConn {
// 	return f.ftp
// }
