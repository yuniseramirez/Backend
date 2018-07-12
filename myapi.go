package main

import ("encoding/json"

"log"

"net/http"

"github.com/gorilla/mux"
"github.com/gorilla/handlers")

type Articles struct {
        ID string
        Title string
        Image string
        Author string

}

 

var articles = []Articles{

        Articles{ID: "1", Image: "https://lh3.googleusercontent.com/ho3xK_RikLgKvXcgvt5IkE86UxI64atJrzEzPmbhr7oGtubgVlOHrV7Ukb615sKRD4-Wn-sJf_8ZPl7_fk8=pf-w200-h200", Title: "All 12 Boys And Their Coach Are Rescued From Thai Cave After 2 Weeks", Author: "NPR"},
        Articles{ID: "2", Image: "http://www2.pictures.livingly.com/mp/zvOAvZDJb0vl.jpg", Title: "You will love these rare andstunning photos of Princess Diana", Author: "Kimia Madani"},
        Articles{ID: "3", Image: "https://cbsnews1.cbsistatic.com/hub/i/r/2018/07/10/1350f97a-d083-455a-b994-845440650a84/resize/620x/f946cfb0cee567a2d52ced58120292da/thailand.jpg", Title: "People on the street cheered and clapped when ambulances ferrying the boys arrived at the hospital in Chiang Rai city.", Author: "CBS NEWS"},
        Articles{ID: "4", Image: "https://assets.bwbx.io/images/users/iqjWHBFdfxIU/iADgPrqC7yoI/v0/1000x-1.jpg", Title: "Meet the 40-Year-Old Erdogan Son-in-Law Running Turkey’s Economy", Author: "Onur Ant"},
        Articles{ID: "5", Image: "https://fm.cnbc.com/applications/cnbc.com/resources/img/editorial/2018/07/02/105306577-1530537945343rts1ub47.530x298.jpg?v=1530542057", Title: "Tesla ditches reservations, opens up Model 3 car sales to all customers in North America", Author: "Lora Kolodny"}}

 

func GetData(w http.ResponseWriter, r *http.Request) { //yunise- add the PostData function

                json.NewEncoder(w).Encode(articles)

}

 

func main() {

        router := mux.NewRouter()

        router.HandleFunc("/data", GetData).Methods("GET")

        log.Print("localhost:8000")
        corsObj := handlers.AllowedOrigins([]string{"*"})

        log.Fatal(http.ListenAndServe(":8000", handlers.CORS(corsObj)(router)))

}
