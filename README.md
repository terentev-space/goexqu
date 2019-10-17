# GoExQu
This project represents my attempt to combine business with pleasure: learn golang and write microservice for sending webhooks to our clients.

# Quickstart

### Build project
...

### Copy config
##### example.config.json -> config.json
```json
{
  "queue_service": {
    "type": "rabbitmq",
    "connection": "amqp://guest:guest@localhost:5672/",
    "name": "goexqu"
  }
}
```
##### Params
* type - service type name:
  * rabbitmq - RabbitMQ Service
* connection - service connection string:
  * amqp://guest:guest@localhost:5672/ - RabbitMQ Service connection string
    * guest - login
    * guest - password
    * localhost - host
    * 5672 - port
* name - service queue name:
  * goexqu - RabbitMQ Service queue name

### Send message to queue
##### Format (string)
```json
{
  "urls": ["http://test.url/first", "http://test.url/second"],
  "data": "{\"hello\": \"world\"}",
  "options": {}
}
```
##### Params
* urls - URL list
* data - The data that will be sent to the URL (POST)
* options - ???
