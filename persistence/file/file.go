package file

import (
	"github.com/gypsydave5/kickoff"
	"os"
)

type File struct {
	path string
}

func NewFile(path string) *File {
	return &File{path}
}

func (f File) Save(kickoff *kickoff.Kickoff) error {
	fl, err := os.Create(f.path)
	if err != nil {
		return err
	}
	defer fl.Close()
	fl.WriteString("# ")
	fl.WriteString(kickoff.Title)
	fl.WriteString("\n\n\n")
	fl.WriteString(kickoff.Body)
	return nil
}
