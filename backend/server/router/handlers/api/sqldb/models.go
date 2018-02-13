package sqldb

import (
	"time"
)

func getTime() (r string) {
	//p := fmt.Println
	t := time.Now()
	//_, z := t.Zone()

	//p(z)
	//p(t.Format("2006-01-02 15:04:05"))
	r = t.Format("2006-01-02 15:04:05")
	return r
}
