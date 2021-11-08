package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

func main() {

	var client *http.Client
	var remoteURL string

	templateIdPtr := flag.String("template-id", "", "Template ID, like 'bookstore' (Required)")
	passwordPtr := flag.String("password", "", "Account password (Required)")
	apiKeyPtr := flag.String("apikey", "", "API key for YOLA_B2C_PROFESSIONAL_TIERED plan (Required)")
	flag.Parse()

	if *templateIdPtr == "" || *passwordPtr == "" || *apiKeyPtr == "" {
		fmt.Printf("\nMissing parameter\n\n\n")
		flag.PrintDefaults()
		os.Exit(1)
	}


	client = &http.Client{}
	remoteURL = "https://my.ecwid.com/resellerapi/v1/register?register=y"

	//prepare the reader instances to encode
	email := fmt.Sprintf("%s@ecwidstore.yola.net", *templateIdPtr)
	values := map[string]io.Reader{
		"email": strings.NewReader(email),
		"password": strings.NewReader(*passwordPtr),
		"is_test": strings.NewReader("true"),
		"name": strings.NewReader("product+demostore@yola.com"),
		"defaultlanguage": strings.NewReader("en"),
		"plan": strings.NewReader("YOLA_B2C_PROFESSIONAL_TIERED"),
		"register": strings.NewReader("y"),
		"ip": strings.NewReader("6.0.0.1"),
		"key": strings.NewReader(*apiKeyPtr),
	}
	err := Upload(client, remoteURL, values)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n\nEmail: %s\n", email)
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
		// Add an image file
		if x, ok := r.(*os.File); ok {
			if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
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
