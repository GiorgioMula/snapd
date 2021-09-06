package osutil

import (
	"os"
)

const randChunkSize = 1024
const numShredCycles = 10
const urandomDevice = "/dev/urandom"

func Shred(filename string) error {
	// Open file to scrumble and get len information
	fi, err := os.Stat(filename)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(filename, os.O_RDWR, 0755)
	if err != nil {
		return err
	}

	// Shred file content with random data
	for shredCycle := 0; shredCycle < numShredCycles; shredCycle++ {
		var i int64
		var numBytes int
		for i = 0; i < fi.Size(); {
			p_random_chunk, _ := randomDataSlice()
			write_slice := *p_random_chunk
			if (fi.Size() - i) < int64(numBytes) {
				write_slice = (*p_random_chunk)[0 : fi.Size()-i]
			}
			numBytes, err = f.WriteAt(write_slice, i)
			if err != nil {
				f.Close()
				return err
			}
			i += int64(numBytes)
		}
	}
	err = f.Close()
	if err != nil {
		return err
	}

	// Get rid of file
	return os.Remove(filename)
}

// randomDataSlice generate a slice of random data,
// returns pointer to data
func randomDataSlice() (*[]byte, error) {
	f, err := os.Open(urandomDevice)
	if err != nil {
		return nil, err
	}
	var randomData [randChunkSize]byte
	randomSlice := randomData[:]
	_, err = f.Read(randomSlice)

	return &randomSlice, err
}
