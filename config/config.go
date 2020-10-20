package config

import (
	"io/ioutil"
	"log"
	us "nospin/user"
	"os"
	"os/user"
	"strings"
)

func New(u us.User) {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	home := usr.HomeDir
	os.Mkdir(home+"/nospin", 0777)
	f, _ := os.OpenFile(home+"/nospin/config", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	f.WriteString("name=" + u.ID)
}
func Set(name string, value string) {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	home := usr.HomeDir
	f, _ := os.OpenFile(home+"/nospin/config", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	f.WriteString("\n" + name + "=" + value)
}
func Get(name string) string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	home := usr.HomeDir
	f, fErr := ioutil.ReadFile(home + "/nospin/config")
	if fErr != nil {
		log.Fatalln(fErr)
	}
	file := strings.Split(string(f), "\n")
	for _, l := range file {
		if len(l) != 0 {
			nv := strings.Split(l, "=")
			if nv[0] == name {
				return nv[1]
			}
		}
	}
	return ""
}