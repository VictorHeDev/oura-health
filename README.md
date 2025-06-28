# Oura

This project is to track Victor's personal health data using an Oura ring and [Oura's API](https://cloud.ouraring.com/v2/docs#section/Getting-Started/What-is-an-API).

## Setup

### Authentication

Since I will be accessing my own personal data, I can use a Personal Access Token. Oura provides 2 methods for authentication:

1. OAuth2
2. Personal Access Tokens
    ```
    GET /v2/usercollection/personal_info HTTP/1.1
    Host: api.ouraring.com
    Authorization: Bearer <token>
    ```

Set your own personal access token into the `.env.example` file and simply copy it over to an `.env` file. Be sure not to commit your actual personal access token.

```bash
cp .env.example .env
```

### Webhooks
- Set up webhooks to consume Oura data whenever there are new notications available
- Make a single request for historical data when a user first connects, then use webhooks for ongoing updates

### Essential concepts

#### Types of data

- daily summaries
- time series data
- events

#### Timeline

- data collection
- syncing
- processing
- API

### How to get historical data

1. Use reasonable date ranges
2. Implement pagination
3. Consider running historical data collection as a bg job
4. Store data in your own DB to avoid repeated API calls

### Best practices for data access
- initial load: pull all historical data
- ongoing updates: use webhooks for all subsequent data updates
- webhook integration - minimize the number of API calls and ensure that you always have the latest data
- error handling: deal with occassional gaps of data when user does not sync their ring



