package primitive

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// This file contains supporting function for transformation

//Mode Primitive mode
type Mode int

//Modes supported by prmitive package
const (
	ModeCombo Mode = iota
	ModeTriangle
	ModeRect
	ModeEllipse
	ModeCircle
	ModeRotatedrect
	ModeBeziers
	ModeRotatedellipse
	ModePolygon
)

//WithMode is option for transform function for defining mode
func WithMode(mode Mode) func() []string {
	return func() []string {
		return []string{"-m", fmt.Sprintf("%d", mode)}
	}
}

//Transform will take image and aplly a primitive transformation
func Transform(image io.Reader, ext string, numShapes int, opts ...func() []string) (io.Reader, error) {
	var args []string
	for _, opt := range opts {
		args = append(args, opt()...)
	}
	in, err := tempfile("in_", ext)

	if err == nil {

		defer os.Remove(in.Name())

		out, err := tempfile("outtmp_", ext)
		if err == nil {

			defer os.Remove(out.Name())

			//Read image into in file
			_, err = io.Copy(in, image)
			if err == nil {

				//Run primitives with  -i in.name() -o out.name
				stdCombo, err := primitve(in.Name(), out.Name(), numShapes, args...)
				if err == nil {

					fmt.Println(stdCombo)
					//read out in reader and return reader and delete the outfile
					b := bytes.NewBuffer(nil)
					_, err = io.Copy(b, out)
					return b, err
				}
			}
		}

	}
	return nil, errors.New("primitive : failed to create temp input file")
}

//Calls primitive command line function
func primitve(inputFile, outFile string, numShapes int, args ...string) (string, error) {
	argStr := fmt.Sprintf("-i %s -o %s -n %d ", inputFile, outFile, numShapes)
	args = append(strings.Fields(argStr), args...)
	cmd := exec.Command("primitive", args...)
	b, err := cmd.CombinedOutput()
	return string(b), err
}

//Creates temporary file for primitive
func tempfile(prefix, ext string) (*os.File, error) {
	in, err := ioutil.TempFile("", prefix)
	if err != nil {
		return nil, errors.New("primitive: failed to create temporary file")
	}
	defer os.Remove(in.Name())
	return os.Create(fmt.Sprintf("%s.%s", in.Name(), ext))
}
