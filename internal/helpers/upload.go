package helpers

import (
	"io"
	"net/http"
	"os"
)

func FileUpload(r *http.Request, files []string) (map[string]string, error){
	// this function returns the filename(to save in database) of the saved file or an error if it occurs
	if err := r.ParseMultipartForm(32 << 20); err != nil{
		return nil, err
	}

	var filenames = make(map[string]string)

	for _, i := range files{
		file, handler, err := r.FormFile(i)
		if err != nil {
			return nil, err
		}

		defer file.Close() //close the file when we finish

		//this is path which  we want to store the file
		f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			return nil, err
		}

		defer f.Close()
		_, _ = io.Copy(f, file)

		filenames[i] = f.Name()
	}


	//here we save our file to our path
	return filenames, nil
}
