# Apache APISIX Golang plugin Example

## Prerequisites
- Go (>= 1.15)
- APISIX (>= 2.9.0)


## Development

### Run APISIX with docker
`
docker-compose up
`
Once the process is complete, execute the curl command on the host running Docker to access the Admin API, and determine if Apache APISIX was successfully started based on the returned data.
`
# Note: Please execute the curl command on the host where you are running Docker.
curl "http://127.0.0.1:9080/apisix/admin/services/" -H 'X-API-KEY: edd1c9f034335f136f87ad84b625c8f1'
`

The following data is returned to indicate that Apache APISIX was successfully started:
`
{
  "count":1,
  "action":"get",
  "node":{
    "key":"/apisix/services",
    "nodes":{},
    "dir":true
  }
}
`

### Run the plugin
` APISIX_LISTEN_ADDRESS=unix:/tmp/runner.sock go run main.go `



### create a test routes with and enable the plugin

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

### Go to `http://127.0.0.1:9080/get` and you'll need to provide username and password. otherwise it will give you an 401 error

