package handler

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/sachinnay/Gophercises/assignment18/transform/primitive"
)

// This file contains handler function for transformation

//GetMux method will return all the handlers
func GetMux() http.Handler {
	h, _ := homedir.Dir()
	imgPath := filepath.Join(h, "img")
	log.Println("=======>" + imgPath)
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir(imgPath))
	mux.Handle("/img/", http.StripPrefix("/img", fs))
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/upload", Upload)
	mux.HandleFunc("/modify/", Modify)
	return mux
}

//Home will handle first request
func Home(w http.ResponseWriter, r *http.Request) {
	log.Println("Home service called")
	html := `<html><body> 
	<form action="/upload" method="post" enctype="multipart/form-data">
	<input type="file" name="image"/>
	<button type="submit" >Upload Image </button>
	</form>
</body></html>`
	fmt.Fprint(w, html)
}

//Upload will be called when user will upload the image for transformation
func Upload(w http.ResponseWriter, r *http.Request) {
	log.Println("Upload service called")
	file, h, err := r.FormFile("image")
	if err == nil {
		defer file.Close()
		e := filepath.Ext(h.Filename)[1:]
		finalFile, err := tempfile("", e)
		if err == nil {
			defer finalFile.Close()
			io.Copy(finalFile, file)
			http.Redirect(w, r, "/modify/"+filepath.Base(finalFile.Name()), http.StatusFound)
			return
		}
	}
	log.Println("Error Occured  :: ", err.Error())
	http.Error(w, "error occured in upload", http.StatusInternalServerError)
	return
}

//Modify will modify the uploaded image
func Modify(w http.ResponseWriter, r *http.Request) {
	log.Println("Modify service called")
	h, _ := homedir.Dir()
	imgPath := filepath.Join(h, "img")
	f, err := os.Open(imgPath + "/" + filepath.Base(r.URL.Path))
	//f, err := os.Open("./img/" + filepath.Base(r.URL.Path))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer f.Close()
	ext := filepath.Ext(f.Name())

	modeStr := r.FormValue("mode")
	if modeStr == "" {
		//render mode choices
		renderModeChoices(w, r, f, ext)
		return
	}

	mode, err := strconv.Atoi(modeStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_ = mode
	//Render noshapes
	nStr := r.FormValue("n")
	if nStr == "" {
		//render mode choices
		renderNoShapesChoices(w, r, f, ext, primitive.Mode(mode))
		return
	}

	noShapes, err := strconv.Atoi(nStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_ = noShapes
	http.Redirect(w, r, "/img/"+filepath.Base(f.Name()), http.StatusFound)
	io.Copy(w, f)
}

//Used to generate the images by taking variable No of shapes for transformation
func renderNoShapesChoices(w http.ResponseWriter, r *http.Request, rs io.ReadSeeker, ext string,
	mode primitive.Mode) {
	log.Println("RenderNoShapesChoices called")
	opts := []genOpts{
		{N: 10, M: mode},
		{N: 20, M: mode},
		{N: 30, M: mode},
		{N: 40, M: mode},
	}
	imges, err := genImages(rs, ext, opts...)
	if err == nil {

		html := `<html><body>
		{{range .}}
			<a href="/modify/{{.Name}}?mode={{.Mode}}&n={{.NumShapes}}">
				<img style="width: 20%;" src="/img/{{.Name}}">
			</a>
		{{end}}
		</body></html>`
		tpl := template.Must(template.New("").Parse(html))
		type dataStruct struct {
			Name      string
			Mode      primitive.Mode
			NumShapes int
		}
		var data []dataStruct
		for i, img := range imges {
			data = append(data, dataStruct{
				Name:      filepath.Base(img),
				Mode:      opts[i].M,
				NumShapes: opts[i].N,
			})
		}
		err = tpl.Execute(w, data)

		return
	}
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

//Used for applying multiple modes for transformation
func renderModeChoices(w http.ResponseWriter, r *http.Request, rs io.ReadSeeker, ext string) {
	log.Println("RenderModeChoices called")
	opts := []genOpts{
		{N: 10, M: primitive.ModeBeziers},
		{N: 10, M: primitive.ModeCircle},
		{N: 10, M: primitive.ModeCombo},
		{N: 10, M: primitive.ModePolygon},
	}
	imges, err := genImages(rs, ext, opts...)
	if err == nil {
		html := `<html><body>
		{{range .}}
			<a href="/modify/{{.Name}}?mode={{.Mode}}">
				<img style="width: 20%;" src="/img/{{.Name}}">
			</a>
		{{end}}
		</body></html>`
		tpl := template.Must(template.New("").Parse(html))
		type dataStruct struct {
			Name string
			Mode primitive.Mode
		}
		var data []dataStruct
		for i, img := range imges {
			data = append(data, dataStruct{
				Name: filepath.Base(img),
				Mode: opts[i].M,
			})
		}
		err = tpl.Execute(w, data)
		return
	}
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

type genOpts struct {
	N int
	M primitive.Mode
}

//Genrates group of images
func genImages(rs io.ReadSeeker, ext string, opts ...genOpts) ([]string, error) {
	var ret []string
	for _, opt := range opts {
		rs.Seek(0, 0)
		f, err := genImage(rs, ext, opt.N, opt.M)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		ret = append(ret, f)
	}
	return ret, nil

}

// Genrates single image
func genImage(file io.Reader, ext string, numShapes int, mode primitive.Mode) (string, error) {
	var outFile *os.File
	var err error
	var out io.Reader
	out, err = primitive.Transform(file, ext, numShapes, primitive.WithMode(mode))
	if err == nil {
		outFile, err = tempfile("", ext)
		if err == nil {
			defer outFile.Close()
			io.Copy(outFile, out)
			return outFile.Name(), err
		}
	}
	return "", err
}

//Used to create temp file for handler
func tempfile(prefix, ext string) (*os.File, error) {
	h, _ := homedir.Dir()
	imgPath := filepath.Join(h, "img")
	in, err := ioutil.TempFile(imgPath+"/", prefix)
	if err != nil {

		return nil, errors.New("handler: failed to create temporary file")
	}
	defer os.Remove(in.Name())
	return os.Create(fmt.Sprintf("%s.%s", in.Name(), ext))
}
