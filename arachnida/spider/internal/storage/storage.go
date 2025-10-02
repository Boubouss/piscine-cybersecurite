package storage

import (
	// "fmt"
	// "io"
	// "net/http"
	// "os"
	// "path/filepath"
	"spider/internal/types"
	// "github.com/google/uuid"
)

// func createFolder(path string) (string, error) {
//   _, err := os.Stat(path)
//   if err != nil {
//     err = os.MkdirAll(path, 0755)
//     if err != nil {
//       _, err = os.Stat("./data/")
//       if err != nil {
//         err = os.MkdirAll("./data/", 0755)
//         if err != nil {
//           return "", err
//         }
//       }
//     }
//   }
//   return path, nil
// }

func SaveImage(image *types.Image) (error) {
  // resp, err := http.Get(image.URL)
  //
  // if err != nil {
  //   return fmt.Errorf("Failed to fetch")
  // }
  //
  // defer resp.Body.Close()
  //
  // if resp.StatusCode != http.StatusOK {
  //   return fmt.Errorf("Unauthorized fetch")
  // }
  //
  // filename := uuid.New().String() + "_" + filepath.Base(image.URL)
  //
  // dir, err := createFolder(image.OPT.Path)
  //
  // if err != nil {
  //   return fmt.Errorf("Creation folder")
  // }
  //
  // path := filepath.Join(dir, filename)
  //
  // file, err := os.Create(path)
  //
  // if err != nil {
  //   return fmt.Errorf("Creation file")
  // }
  //
  // defer file.Close()
  //
  // _, err = io.Copy(file, resp.Body)
  //
  // if err != nil {
  //   return fmt.Errorf("Copy failed")
  // }
  //
  // return nil

  return nil
}
