package controller

import (
	"net/http"
	"os"
  "path"
)

func Server(w http.ResponseWriter, r *http.Request) {
  
  fp := path.Join("web", r.URL.Path)

  // Return a 404 if the template doesn't exist
  _, err := os.Stat(fp)
  if err != nil {
    if os.IsNotExist(err) {
      http.NotFound(w, r)
      return
    }
  }

  // // Return a 404 if the request is for a directory
  // if info.IsDir() {
  //   http.NotFound(w, r)
  //   return
  // }

  // tmpl, err := template.ParseFiles(lp, fp)
  // if err != nil {
  //   // Log the detailed error
  //   log.Println(err.Error())
  //   // Return a generic "Internal Server Error" message
  //   http.Error(w, http.StatusText(500), 500)
  //   return
  // }

  // if err := tmpl.ExecuteTemplate(w, "layout", nil); err != nil {
  //   log.Println(err.Error())
  //   http.Error(w, http.StatusText(500), 500)
  // }
}
