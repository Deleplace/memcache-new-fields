package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "POST only", 400)
			return
		}
		useMemcache(r.Context(), w)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, page)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	addr := os.Getenv("ADDR") + ":" + port
	log.Printf("Listening on %s\n", addr)
	err := http.ListenAndServe(addr, nil)
	log.Fatal(err)
}

const page = `
<style>
	#msg, #errors {
		margin: 1em;
		white-space: pre;
		font-family: monospace;
	}
	#errors {
		color: #800;
		font-weight: bold;
	}
</style>
<button id="do_it">Do</button>

<div id="msg"></div>
<div id="errors"></div>

<script>
do_it.addEventListener("click", (e) => {
	msg.innerText = "Executing...";
	errors.innerText = "";
	fetch( "/test", {method: "POST"})
		.then( (r) => {
			if (r.ok) {
			  return r.text();
			}
			throw new Error(r.status + ": " + r.statusText );
		}).then( (t) => {
			msg.innerText = t;
		}).catch( (err) => {
			errors.innerText = err;
		});
});
</script>
`
