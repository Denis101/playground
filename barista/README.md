Idea for a log forwarder in Node.js

Would be a fun excuse to play with TypeScript.

Add API with the following endpoints:
* Version of barista used
* Inputs enabled, any listening endpoints (HTTP etc)
* Outputs enabled
* Filters enabled

API output:
```
{
  "version": "0.0.1",
  "inputs": {
    "http": {
      "port": 80
    },
    "tail": {
      "path": "/var/log/nginx/error.log"
    }
  },
  "outputs": {
    "elasticsearch": {
      "host": "localhost"
      "port": 1234
    },
    "s3": {
      "bucket": "s3-bucket-name",
      "region": "eu-west-1"
    }
  },
  "filters": [
    {
      "id": "someFilterName",
      "resolutions": [
        "elasticsearch",
        "s3"
      ]
    }
  ],
}
```
