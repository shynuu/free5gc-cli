package boot

import "os"

func init() {
	os.Mkdir("./logs", 0777)
}

func Initialize() {

}
