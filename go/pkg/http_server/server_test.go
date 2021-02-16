package http_server

// import(
//     "net/http"
//     "net/http/httptest"
// 	"testing"
//     "fmt"
//     "github.com/marcpar/filecache/go/pkg/file"
// )


// func initializeHelper() ( *file.FileCache){
//     return  file.NewFileCache( "../.tmp")
// }

// func TestGETArtifact(t *testing.T) {
//     fileCache := initializeHelper()
//     t.Run("get artifact", func(t *testing.T) {
//         request, _ := http.NewRequest(http.MethodGet, "/artifact-store/3rdparty/intel.com/mkl.zip", nil)
//         response := httptest.NewRecorder()
  
//         Server(response, request)
//         fmt.Println(request.RequestURI)
//         got := response.Body.String()
//         want := "20"

//         if got != want {
//             t.Errorf("got %q, want %q", got, want)
//         }
//     })

    
// }