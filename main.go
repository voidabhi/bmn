package main

import(
	"io/ioutil"
	"encoding/json"
	"net/http"
	"os"
	"io"
	"text/template"
	"fmt"
	"strconv"
	"flag"
)

// Account Model
type AccountEntry struct {
	Username string
	Password string
	SuccessRate string `json:"success_rate"`
	Upvotes string
	Posted string
}

// Account Collection
type AccountBook struct {
	Accounts []*AccountEntry
}


func main(){

	flag.Usage = usage
	flag.Parse()
	
	args:=flag.Args()
	
	if len(args) != 1 {
		usage()
	}
	
	query := args[0]
	
	accountBook,_ := bmn(query)
	if len(accountBook.Accounts) > 0 {
		printMessage(os.Stdout,strconv.Itoa(len(accountBook.Accounts))+" accounts for "+query)
	} else {
		printMessage(os.Stdout,"No accounts for this domain.")
	}
	
	for _,account := range accountBook.Accounts {
		printAccount(os.Stdout,account)
	}
}


// Usage

const usageTmpl = `BMN is a command-line utility to find logins/passwords for websites that force you to register.
Usage:
  bmn [website]
`

func usage() {
	printUsage(os.Stderr)
	os.Exit(2)
}

func printUsage(w io.Writer) {
	tmpl(w, usageTmpl, nil)
}

// Result

const accountTmpl = `
	Username: {{.Username}}
	Password: {{.Password}}
	SuccessRate: {{.SuccessRate}}
	Upvotes: {{.Upvotes}}
	Posted: {{.Posted}}
	
`

func printAccount(w io.Writer,a *AccountEntry){
	tmpl(w,accountTmpl,a)
}

// Messages

const messageTmpl = `
	{{.Message}}
`

func printMessage(w io.Writer,msg string) {
	tmpl(w,messageTmpl,struct {Message string}{msg})
}

func printErr(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(2)
}


// Helpers

// Fetching data from bmn webservice
func bmn(website string)(*AccountBook,error) {

	res, err := http.Get("http://bugmenotapi.herokuapp.com/"+website)
	if err != nil {
		return nil,err
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil,err
	}
	var accountBook AccountBook
	json.Unmarshal(body,&accountBook)
	
	return &accountBook,nil
}

// Output template
func tmpl(w io.Writer, text string, data interface{}) {
	t := template.New("top")
	template.Must(t.Parse(text))
	if err := t.Execute(w, data); err != nil {
		panic(err)
	}
}
