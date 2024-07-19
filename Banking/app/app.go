package app

import (
	"Banking/domain"
	"Banking/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type TimeResponse map[string]string

func Start() {

	router := mux.NewRouter()
	// router.HandleFunc("/greet", greet)
	// router.HandleFunc("/getAllCustomers", getAllCustomers).Methods(http.MethodGet)
	// router.HandleFunc("/getCustomer/{id:[0-9]+}", getCustomer).Methods(http.MethodGet)
	// router.HandleFunc("/createCustomer", createCustomer).Methods(http.MethodPost)

	//router.HandleFunc("/api/time", getTime).Methods(http.MethodGet)
	//log.Fatal(http.ListenAndServe("localhost:8000", router))

	// http.HandleFunc("/greet", greet)
	// http.HandleFunc("/getAllCustomers", getAllCustomers)
	//log.Fatal(http.ListenAndServe("localhost:8000", nil))

	//wiring -banking
	//	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	//router.HandleFunc("/getAllCustomers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/getCustomers/{customer_id:[0-9]+}", ch.GetCustomerById).Methods(http.MethodGet)
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

/*func getTime(w http.ResponseWriter, r *http.Request) {
	//extract tz
	response := TimeResponse{}
	timezones := r.URL.Query().Get("tz")

	//if not  present //return utc else convert
	if timezones == "" {
		response["UTC"] = time.Now().UTC().Format(time.RFC3339)
	} else {
		tzarr := strings.Split(timezones, ",")
		for _, timezone := range tzarr {
			loc, err := time.LoadLocation(timezone)
			log.Println("err", err)
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				response["error"] = "Not a valid timezone"
				json.NewEncoder(w).Encode(response)
				return
			}
			response[timezone] = time.Now().In(loc).Format(time.RFC3339)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}*/

/*func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //returns the route variables
	fmt.Fprint(w, vars["id"])
	//panic("unimplemented")
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "post req rexd")
	//panic("unimplemented")
}*/
