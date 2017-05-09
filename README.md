# Slow Endpoint

A small HTTP server that delays a response to a request by a given amount.

* `GET /{seconds}` - returns a response in `seconds` seconds

Ready for serverless AppEngine deployment - an `app.yaml` has been created.

## Run locally

```
go run main.go
```

## Deploy to Google Cloud

Make sure you've installed the [gcloud tool](https://cloud.google.com/sdk/downloads) and configured a project (including having billing configured on your google cloud project), then run:

```
gcloud app deploy
```

List versions: `gcloud app versions list`

Start: `gcloud app versions start VERSIONID`

Stop: `gcloud app versions stop VERSIONID`

Logs: `gcloud app logs tail`