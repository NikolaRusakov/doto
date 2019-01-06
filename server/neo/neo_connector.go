package neo

import (
	"fmt"
	"github.com/jmcvetta/neoism"
	"os"
)

var Db, err = neoism.Connect(Login)

var Login = fmt.Sprintf("http://%s:%s@localhost:7474",
	os.Getenv("NEO_NAME"), os.Getenv("NEO_PASS"))
