# translate-bot 
Real time translation on Slack by using Real Time Message API and Cloud Translation API.   
Bot will detect the given sentense and detect which language were used, then translate Japanese into English, or other languages to Japanese.

# How to use

1. Set up Golang environment. https://golang.org/doc/install
2. Set up Cloud SDK environment. https://cloud.google.com/sdk/
3. Follow from step 1 through 3 in the getting started guide. https://cloud.google.com/translate/docs/getting-started
4. `gcloud auth application-default login` and login with your google account that can access the created project.
4. Get a Real Time Messaging API key from slack by adding Bots integration.
5. Add the API key to your environment variable as "
5. Clone the repository
6. `go get -u ../...`
7. `SLACKAPIKEY="[your-api-key-from-slack]" go run main.go`
