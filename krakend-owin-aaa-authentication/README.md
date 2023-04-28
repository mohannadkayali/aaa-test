
Backend Middleware
- Configuration needs to come from krakend config
  - host
  -

{
  "url_pattern": "/api/v1/productlisting/products/?vendor_id={id}",
  "method": "GET",
  "host": [],
  "extra_config": {
    "qos/circuit-breaker": {
      "interval": 60,
      "log_status_change": true,
      "max_errors": 50,
      "name": "verticals-api",
      "timeout": 30
    }
  }
},

