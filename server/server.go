package server


import(
	"net/http"
	"fmt"
	"encoding/json"
	"log"
	"contactform/populate"
	"contactform/config"
)


func ReqHandler(rw http.ResponseWriter, request *http.Request){
	decoder := json.NewDecoder(request.Body)
	var json_res populate.FormData
	err := decoder.Decode(&json_res)
	fmt.Println("Logging req: \n")
	fmt.Println(json_res.Member + json_res.CurrentAddr + json_res.UKAddr + json_res.DateOfBirth + json_res.NIN )
	fmt.Println(json_res.ProvidersData)
	if err != nil {
		fmt.Println("Request Handler error\n")
		log.Fatalln(err)
	}
	populate.FillTempl(json_res.Member, json_res.CurrentAddr, json_res.UKAddr, json_res.ProvidersData, json_res.DateOfBirth, json_res.NIN)
}

func defaultHandler(rw http.ResponseWriter, request *http.Request){
	fmt.Printf("default handler")
}


func Serve(){
	tlscfg := config.GetTlsConfig("config/")
	log.Print("\nCreating server with cert file: " + tlscfg.GetCert() + "and key file: " + tlscfg.GetKey())
	http.HandleFunc("/services/loa_pdfgen", ReqHandler)
	http.HandleFunc("/", defaultHandler)
		log.Print("listening on port: " + tlscfg.GetPort())
	err := http.ListenAndServeTLS(tlscfg.GetPort(), tlscfg.GetCert(), tlscfg.GetKey(), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}



