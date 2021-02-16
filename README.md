# Artifact Store

stores artifact

## Installation


## Go Test
```bash
cd go

go test ./pkg/*

```
## BUILD docker 
Artifact Store.

```bash
docker build -t artifactstore  ./go  --build-arg service_name  
```



## Usage

```go
package file

import (
	"net/http"

	
	"path"
	"os"
	"io"
	"sync"
	log "github.com/sirupsen/logrus"
)

type FileCache struct {
	Dir        string 
	Items      map[string]os.FileInfo
	In         chan string
	Mutex      sync.Mutex
	Shutdown   chan interface{}
}


func NewFileCache(dir string) *FileCache {
    fileItems := make(map[string]os.FileInfo)
	return  &FileCache{
		Items: fileItems,
		Dir: dir,
	}
}

func (f *FileCache) GetFile(address string) ( os.FileInfo, error) {
    var fileInfo os.FileInfo
	if value , found := f.Items[address] ; found == false{
		file , err	:= downloadFile( f.Dir, address )
		if err != nil{
			return  fileInfo , err
		} 
		return file, nil
		f.Items[address] = file
	} else{
		log.Info(value)
		fileInfo = value
	}

	return fileInfo , nil
}

func downloadFile(filepath string, url string) (os.FileInfo, error) {
	filename := path.Base(url)
	
	var fileInfo os.FileInfo
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return fileInfo, err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath +"/" + filename )
	if err != nil {
		return fileInfo, err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	
	fileInfo , err  = os.Lstat(filepath)
	return fileInfo,  err
}

func FileGC(tick int, expiry float64 , fileCache *FileCache) {

	// TODO: ticker time to cleanup 
}
```



## HELM
```
helm install manifest 

```