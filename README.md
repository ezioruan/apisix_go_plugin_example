# Apache APISIX Golang plugin Example

## Prerequisites
- Go (>= 1.15)
- APISIX (>= 2.9.0)


## Development

# Run 







## create a route poi

`
curl --location --request PUT 'http://127.0.0.1:9080/apisix/admin/routes/1' \
--header 'X-API-KEY: edd1c9f034335f136f87ad84b625c8f1' \
--header 'Content-Type: application/json' \
--data-raw '{
    "uri": "/get",
    "plugins": {
        "ext-plugin-pre-req": {
            "conf": [
                {
                    "name": "go-plugin-basic-auth",
                    "value": "{\"username\":\"foo\",\"password\":\"bar\"}"
                }
            ]
        }
    },
    "upstream": {
        "type": "roundrobin",
        "nodes": {
            "httpbin.org": 1
        }
    }
}'


`

