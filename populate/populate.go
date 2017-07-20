package populate

import(
	"bytes"
	"io/ioutil"
	"log"
	"fmt"
	"os/exec"
	"html/template"
)



var TemplateFile, _ =  template.ParseFiles("templates/form.html")


type Providers struct {
	Providers_name string
	Plan_number string
}

type FormData struct {
	Member string
	CurrentAddr string
	UKAddr string
	ProvidersData []Providers
	DateOfBirth string
	NIN string
}



func FillTempl(Member string, CurrentAddr string, UKAddr string, Providers_data[]Providers, DateOfBirth string, NIN string){
	fmt.Printf("FillTempl")
	buff := bytes.NewBufferString("")
	pdf_path := "templates/letterofauthority.pdf"
	// Compile and allocate in buffer
	err := TemplateFile.Execute(buff, FormData{
		Member: Member,
		CurrentAddr: CurrentAddr,
		UKAddr: UKAddr,
		ProvidersData: Providers_data,
		DateOfBirth: DateOfBirth,
		NIN: NIN,
	})
	if err != nil {
		log.Fatalln(err)
	}
	err = ioutil.WriteFile("form_compiled.html", buff.Bytes(), 0666)
	if err != nil {
		log.Fatalln(err)
	}
	//convert compiled file to pdf
	err = exec.Command("wkhtmltopdf", "form_compiled.html", pdf_path).Run()
	if err == nil {
		fmt.Printf("[+ TEMPLATE] Save successful")
	} else {
		fmt.Printf("[- TEMPLATE] Error generating PDF %s", err)
	}
}
