package file

import (
	"testing"
	"fmt"

)


func Helper() *FileCache{
	return  NewFileCache( "../../../.tmp")
}


func TestGetFile(t  *testing.T){
	// fileUrl := "https://golangcode.com/logo.svg"

	fileCache :=  Helper()
	t.Run("get artifact", func(t *testing.T) {
		
        got , err := fileCache.GetFile("https://golangcode.com/logo.svg")
		if err != nil{
			t.Errorf( "failed request")
		}

		fmt.Println(got)
        want := "../../../.tmp/logo.svg"
        if got.Name() != want {
            t.Errorf("got %q, want %q", got, want)
        }
    })


	t.Log("finish ....")
}