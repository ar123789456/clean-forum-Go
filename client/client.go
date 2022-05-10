package client

import (
	"log"
	"net/http"
	"time"
)

func Run() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// w.Header().Set("Content-Type", "text/plain")
		// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		log.Println("index")
		// w.Header().Set("Content-Type", "text/html")

		http.ServeFile(w, r, "../../web/index.html")
	})

	fs := http.FileServer(http.Dir("../../web/static"))

	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	srv := &http.Server{
		Handler:      mux,
		Addr:         ":6969",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

// // See https://stackoverflow.com/questions/26141953/custom-404-with-gorilla-mux-and-std-http-fileserver
// func intercept404(handler, on404 http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		hookedWriter := &hookedResponseWriter{ResponseWriter: w}
// 		handler.ServeHTTP(hookedWriter, r)

// 		if hookedWriter.got404 {
// 			on404.ServeHTTP(w, r)
// 		}
// 	})
// }

// type hookedResponseWriter struct {
// 	http.ResponseWriter
// 	got404 bool
// }

// func (hrw *hookedResponseWriter) WriteHeader(status int) {
// 	if status == http.StatusNotFound {
// 		// Don't actually write the 404 header, just set a flag.
// 		hrw.got404 = true
// 	} else {
// 		hrw.ResponseWriter.WriteHeader(status)
// 	}
// }

// func (hrw *hookedResponseWriter) Write(p []byte) (int, error) {
// 	if hrw.got404 {
// 		// No-op, but pretend that we wrote len(p) bytes to the writer.
// 		return len(p), nil
// 	}

// 	return hrw.ResponseWriter.Write(p)
// }

// func serveFileContents(file string, files http.FileSystem) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		// Restrict only to instances where the browser is looking for an HTML file
// 		if !strings.Contains(r.Header.Get("Accept"), "text/plain") {
// 			w.WriteHeader(http.StatusNotFound)
// 			fmt.Fprint(w, "404 not found")

// 			return
// 		}

// 		// Open the file and return its contents using http.ServeContent
// 		index, err := files.Open(file)
// 		if err != nil {
// 			w.WriteHeader(http.StatusNotFound)
// 			fmt.Fprintf(w, "%s not found", file)

// 			return
// 		}

// 		fi, err := index.Stat()
// 		if err != nil {
// 			w.WriteHeader(http.StatusNotFound)
// 			fmt.Fprintf(w, "%s not found", file)

// 			return
// 		}

// 		w.Header().Set("Content-Type", "text/javascript; charset=utf-8")
// 		http.ServeContent(w, r, fi.Name(), fi.ModTime(), index)
// 	}
// }
