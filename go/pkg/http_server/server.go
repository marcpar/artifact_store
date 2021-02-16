package http_server

import(
	"net/http"
	"fmt"
	"strings"
	"github.com/marcpar/filecache/go/pkg/file"
)

type Server struct{
	fileCache file.FileCache 
}

func NewServer (fileCache file.FileCache) *Server{
	return &Server{ fileCache : fileCache}
}

func (s *Server)  HttpServer(w http.ResponseWriter, r *http.Request) {
	artifact := strings.TrimPrefix(r.URL.Path, "/artifact-store/3rdparty/")
	s.fileCache.GetFile(artifact)

	fmt.Fprint(w, "20")
}


