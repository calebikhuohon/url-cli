# URL CLI Application

This CLI takes in a list of URLs as inputs, visits each URL, and prints the list of url-response body size pairs in descending order. It uses goroutines to ensure speedy completion of the I/o tasks.

## How to Run the Application
* Run `go install ./cmds/url-cli` to install the CLI
* Run `go test -v ./...` to run the existing unit test
* Run `url-cli list -u url1,url2,...` e.g `url-cli list -u https://google.com,https://facebook.com,https://github.com/calebikhuohon,https://abc.com` to run the CLI

## Project Structure
* The `cmds/url-cli` package contains the CLI-related logic and code organization.
* The `pkg/urls-processor` package contains the code logic related to processing each URL.