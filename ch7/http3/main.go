package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var shoplist = template.Must(template.New("shop").Parse(`
<table>
<tr style='text-align: left'>
  <th>Item</th>
  <th>Price</th>
</tr>
{{ range $k, $v := .}}
<tr>
  <td>{{ $k }}</td>
  <td>{{ $v }}</td>
</tr>
{{end}}
</table>
`))

var mu sync.Mutex
var db = database{"shoes": 50, "socks": 5}

func main() {
	// mux := http.NewServeMux()

	// mux.Handle("/list", http.HandlerFunc(db.list))
	// mux.Handle("/price", http.HandlerFunc(db.price))

	//same as before
	// mux.HandleFunc("/list", db.list)
	// mux.HandleFunc("/price", db.price)

	//Using the default
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/read", db.read)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)

	// log.Fatal(http.ListenAndServe("localhost:8000", mux))
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	// for item, price := range db {
	// 	fmt.Fprintf(w, "%s: %s\n", item, price)
	// }
	log.Print(shoplist.Execute(w, db))

	mu.Unlock()
}

func (db database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")

	mu.Lock()
	price, ok := db[item]
	mu.Unlock()
	if !ok {
		msg := fmt.Sprintf("no such item: %q\n", item)
		http.Error(w, msg, http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) read(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")

	mu.Lock()
	price, ok := db[item]
	mu.Unlock()
	if !ok {
		msg := fmt.Sprintf("no such item: %q\n", item)
		http.Error(w, msg, http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "%s: %s\n", item, price)
}

func (db database) update(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	newprice, err := strconv.ParseFloat(r.URL.Query().Get("price"), 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "could not parse newprice: %v\n", err)
		return
	}

	mu.Lock()

	defer mu.Unlock()

	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	db[item] = dollars(newprice)
	fmt.Fprintf(w, "updated %s: %s\n", item, dollars(newprice))
}

func (db database) create(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, err := strconv.ParseFloat(r.URL.Query().Get("price"), 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "could not parse price: %v\n", err)
		return
	}

	mu.Lock()

	defer mu.Unlock()

	if p, ok := db[item]; ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "item already in database %s: %s\n", item, p)
		return
	}

	db[item] = dollars(price)
	fmt.Fprintf(w, "added to database %q: %s\n", item, dollars(price))
}

func (db database) delete(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")

	mu.Lock()
	defer mu.Unlock()

	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "item not in database: %s\n", item)
		return
	}

	delete(db, item)
	fmt.Fprintf(w, "deleted %q from database\n", item)
}
