# Listeners : Golang HTTP service

This is a Golang HTTP service listening for Callbacks and Timeouts from the following API providers.

# Usage

Creates two binaries, which can be used to test the HTTP listners.

```bash
$ go build ./cmd/listener && go build ./cmd/listener-cli
```

Run both binaries on separate terminals using `./listener` and `./listener-cli --help`

`./listener` is a Golang HTTP server that listens on `localhost:3000`. It runs and provides an output similar to the following.

```bash
(base) wondenge@hpelitebook:~/projects/sdks/listeners$ ./listener
ts=2020-05-20T14:37:59.275981517Z caller=http.go:155 info="HTTP \"../../gen/http/openapi.json\" mounted on GET /swagger/swagger.json"
ts=2020-05-20T14:37:59.276063985Z caller=http.go:158 info="HTTP \"Show\" mounted on GET /health/"
ts=2020-05-20T14:37:59.276091612Z caller=http.go:161 info="HTTP \"AccountBalanceTimeout\" mounted on POST /mpesa/accountbalance/v1/timeout"
ts=2020-05-20T14:37:59.276118617Z caller=http.go:161 info="HTTP \"AccountBalanceResultEndpoint\" mounted on POST /mpesa/accountbalance/v1/result"
ts=2020-05-20T14:37:59.276143064Z caller=http.go:161 info="HTTP \"TransactionStatusTimeout\" mounted on POST /mpesa/transactionstatus/v1/timeout"
ts=2020-05-20T14:37:59.276163721Z caller=http.go:161 info="HTTP \"TransactionStatusResultEndpoint\" mounted on POST /mpesa/transactionstatus/v1/result"
ts=2020-05-20T14:37:59.27617754Z caller=http.go:161 info="HTTP \"ReversalTimeout\" mounted on POST /mpesa/reversal/v1/timeout"
ts=2020-05-20T14:37:59.276189432Z caller=http.go:161 info="HTTP \"ReversalResultEndpoint\" mounted on POST /mpesa/reversal/v1/result"
ts=2020-05-20T14:37:59.276201031Z caller=http.go:161 info="HTTP \"B2CTimeout\" mounted on POST /mpesa/b2c/v1/timeout"
ts=2020-05-20T14:37:59.276212529Z caller=http.go:161 info="HTTP \"B2CResult\" mounted on POST /mpesa/b2c/v1/result"
ts=2020-05-20T14:37:59.276223977Z caller=http.go:161 info="HTTP \"C2BValidation\" mounted on POST /mpesa/c2b/v1/validation"
ts=2020-05-20T14:37:59.276239161Z caller=http.go:161 info="HTTP \"C2BConfirmation\" mounted on POST /mpesa/c2b/v1/confirmation"
ts=2020-05-20T14:37:59.276264395Z caller=http.go:170 info="HTTP server listening on \"localhost:3000\""
```

The server has the ability to shutdown nicely to ensure that all requests have been completed. This is usually referred to as graceful shutdown.

`./listener-cli --help` is a CLI client for the the listener API.

```bash
(base) wondenge@hpelitebook:~/projects/sdks/listeners$ ./listener-cli --help
./listener-cli is a command line client for the listener API.

Usage:
    ./listener-cli [-host HOST][-url URL][-timeout SECONDS][-verbose|-v] SERVICE ENDPOINT [flags]

    -host HOST:  server host (local). valid values: local
    -url URL:    specify service URL overriding host URL (http://localhost:8080)
    -timeout:    maximum number of seconds to wait for response (30)
    -verbose|-v: print request and response details (false)

Commands:
    health show
    mpesa (account-balance-timeout|account-balance-result|transaction-status-timeout|transaction-status-result|reversal-timeout|reversal-result|b2-c-timeout|b2-c-result|c2-b-validation|c2-b-confirmation)

Additional help:
    ./listener-cli SERVICE [ENDPOINT] --help

Example:
    ./listener-cli health show
    ./listener-cli mpesa account-balance-timeout --body '{
          "ConversationId": "236543-276372-2",
          "MpesaResultCode": 0,
          "MpesaResultDesc": "Initiator information is invalid",
          "MpesaResultParameters": "Et omnis doloremque et molestiae esse ut.",
          "MpesaResultType": 0,
          "Occasion": "Occasion",
          "OriginatorConversationId": "AG_2376487236_126732989KJHJKH",
          "QueueTimeoutURL": "https://internalsandbox.safaricom.co.ke/mpesa/abresults/v1/submit",
          "TransactionID": "LHG31AA5TX"
       }'
```

Test the Account Balance Callback by posting the following on the terminal.

```bash
./listener-cli mpesa account-balance-timeout --body '{
        "ConversationId": "236543-276372-2",
        "MpesaResultCode": 0,
        "MpesaResultDesc": "Initiator information is invalid",
        "MpesaResultParameters": "Et omnis doloremque et molestiae esse ut.",
        "MpesaResultType": 0,
        "Occasion": "Occasion",
        "OriginatorConversationId": "AG_2376487236_126732989KJHJKH",
        "QueueTimeoutURL": "https://internalsandbox.safaricom.co.ke/mpesa/abresults/v1/submit",
        "TransactionID": "LHG31AA5TX"
    }'

```

It returns a Status Code of 201.

```bash
ts=2020-05-20T14:45:16.007591911Z caller=log.go:30 id=Z7kn6F7z req="POST /mpesa/accountbalance/v1/timeout" from=127.0.0.1
ts=2020-05-20T14:45:16.007855149Z caller=mpesa.go:24 info=mpesa.AccountBalanceTimeout
ts=2020-05-20T14:45:16.007891715Z caller=log.go:37 id=Z7kn6F7z status=201 bytes=3 time=302.972µs
```
