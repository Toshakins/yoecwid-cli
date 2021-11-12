package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"strings"
)

type storeParams struct {
	templateId string
	password   string
	apikey     string
}

type Subcommand struct {
	flagSet *flag.FlagSet
}

func readArguments() (storeParams, error) {
	const STORE = "store"
	subcommandMap := map[string]Subcommand{
		"store": {flagSet: flag.NewFlagSet(STORE, flag.ExitOnError)},
	}

	templateIdPtr := subcommandMap[STORE].flagSet.String(
		"template-id", "", "Template ID, like 'bookstore' (Required)")
	passwordPtr := subcommandMap[STORE].flagSet.String(
		"password", "", "Account password (Required)")
	apiKeyPtr := subcommandMap[STORE].flagSet.String(
		"apikey", "", "API key for YOLA_B2C_PROFESSIONAL_TIERED plan (Required)")

	if len(os.Args) < 2 {
		fmt.Println("Please enter a subcommand")
		fmt.Println()
		printDefaults(subcommandMap)
		return storeParams{}, errors.New("no subcommand provided")
	}

	switch os.Args[1] {
	case STORE:
		parseError := subcommandMap[STORE].flagSet.Parse(os.Args[2:])
		if parseError != nil {
			fmt.Printf("Parsing error!")
			return storeParams{}, errors.New("cannot parse store command")
		}
	default:
		printDefaults(subcommandMap)
		return storeParams{}, errors.New("unknown command")
	}

	if subcommandMap[STORE].flagSet.Parsed() {
		if *templateIdPtr == "" || *passwordPtr == "" || *apiKeyPtr == "" {
			fmt.Println("Missing parameter")
			subcommandMap[STORE].flagSet.PrintDefaults()
			return storeParams{}, errors.New("store missing required parameters")
		}
		return storeParams{*templateIdPtr, *passwordPtr, *apiKeyPtr}, nil
	}

	fmt.Println("Cannot parse input")
	printDefaults(subcommandMap)
	return storeParams{}, errors.New("cannot parse commands")
}

func printDefaults(subcommands map[string]Subcommand) {
	for command, data := range subcommands {
		fmt.Println("Usage: yoecwid-cli subcommand options")
		fmt.Println()
		fmt.Println("Available subcommands below")
		fmt.Println()
		fmt.Printf("%s\n", command)
		data.flagSet.PrintDefaults()
		fmt.Println()
	}
}

func main() {
	params, err := readArguments()
	if err != nil {
		return
	}
	remoteCall(params)
}

func remoteCall(params storeParams) {
	var client *http.Client
	var remoteURL string

	// Commented code below was useful for debugging 🙃
	//proxyURL, _ := url.Parse("http://127.0.0.1:8080")
	//proxy := http.ProxyURL(proxyURL)
	//transport := &http.Transport{Proxy: proxy}
	//transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	//client = &http.Client{Transport: transport}

	client = &http.Client{}
	remoteURL = "https://my.ecwid.com/resellerapi/v1/register?register=y"

	//prepare the reader instances to encode
	email := fmt.Sprintf("%s@ecwidstore.yola.net", params.templateId)
	storeTemplate := fmt.Sprintf(storeBaseTemplate, email, email, email)
	values := map[string]io.Reader{
		"email":           strings.NewReader(email),
		"password":        strings.NewReader(params.password),
		"is_test":         strings.NewReader("true"),
		"name":            strings.NewReader(email),
		"defaultlanguage": strings.NewReader("en"),
		"plan":            strings.NewReader("YOLA_B2C_PROFESSIONAL_TIERED"),
		"register":        strings.NewReader("y"),
		"ip":              strings.NewReader("6.0.0.1"),
		"key":             strings.NewReader(params.apikey),
		"template":        strings.NewReader(storeTemplate),
	}
	err := Upload(client, remoteURL, values)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n\nEmail: %s\n", email)
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

func Upload(client *http.Client, url string, values map[string]io.Reader) (err error) {
	// Prepare a form that you will submit to that URL.
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	for key, r := range values {
		var fw io.Writer
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		if key == "template" {
			h := make(textproto.MIMEHeader)
			h.Set("Content-Disposition",
				fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
					escapeQuotes(key), escapeQuotes("default_store.xml")))
			if fw, err = w.CreatePart(h); err != nil {
				return
			}

		} else {
			// Add other fields
			if fw, err = w.CreateFormField(key); err != nil {
				return
			}
		}
		if _, err = io.Copy(fw, r); err != nil {
			return err
		}

	}
	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	w.Close()

	// Now that you have a form, you can submit it to your handler.
	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		return
	}
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())

	// Submit the request
	res, err := client.Do(req)
	if err != nil {
		return
	}

	// Check the response
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", res.Status)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	fmt.Printf("%s", body)
	return
}
