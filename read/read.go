package read

import (
	"io/ioutil"
	"sync"
)

type Rtype struct {
	sync.Mutex
	file string
	e    error
}

func (r *Rtype) Read() ([]byte, error) {
	r.Lock()
	defer r.Unlock()

	content, err := ioutil.ReadFile(r.file)
	if err != nil {
		r.e = err
	}

	return content, err

}
