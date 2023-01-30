package utils

import (
	"fmt"
	"time"

	"gitalpha.com/go/nome_projeto/config"
	"github.com/jlaffaye/ftp"
)

// ConnectFTP realiza a conexão com o servidor FTP
func ConnectFTP(u *config.FTP) (*ftp.ServerConn, error) {
	hostPort := fmt.Sprintf("%v:%v", u.FtpHost, u.FtpPort)

	c, err := ftp.Dial(hostPort, ftp.DialWithTimeout(60*time.Second))
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar no ftp: %v", err)
	}

	err = c.Login(u.FtpUser, u.FtpPass)
	if err != nil {
		return nil, fmt.Errorf("erro ao efetuar login no ftp: %v", err)
	}

	return c, nil
}
