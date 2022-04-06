<div align="center">
<h3 align="center">Medications Delivery Using Drones - Golang Service </h3>
</div>

<details open>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about">About</a>
    </li>
    <li>
      <a href="#dependencies">Dependencies</a>
    </li>
    <li>
      <a href="#database">Database</a>
    </li>
    <li>
      <a href="#setup-and-run">Setup and Run</a>
      <ul>
        <li><a href="#ping">Ping</a></li>
      </ul>
    </li>
    <li>
      <a href="#api-endpoints">API Endpoints</a>
      <ul>
        <li><a href="#postman-collection">Postman Collection</a></li>
        <li><a href="#register-drone">Register Drone</a></li>
        <li><a href="#load-drone-with-medications">Load Drone With Medications</a></li>
        <li><a href="#get-drone-medications">Get Drone Medications</a></li>
        <li><a href="#get-drone-battery-capacity">Get Drone Battery Capacity</a></li>
        <li><a href="#update-drone-status">Update Drone Status</a></li>
      </ul>
    </li>
    <li><a href="#possible-enhancements">Possible Enhancements</a></li>
  </ol>
</details>

## About

This service is built using the latest golang version `1.18` as a hiring task for musala soft company.

## Dependencies

* [Go-Gonic](https://github.com/gin-gonic/gin) The foundational http router and data binding.
* [Gorm.io](https://gorm.io/gorm) The DB ORM.
* [Cron](https://github.com/robfig/cron)  App scheduler for processed cron jobs.
* [Validator](https://github.com/go-playground/validator) Used for applying validation rules on user inputs.

## Database

This service uses Postgres as it's datastore.

## Setup

To run the service, make sure you have Docker and Compose installed on testing machine.

## Setup and Run

> **Warning**: before running the below command please make sure that ports `5432` and `6504` are not allocated.

```shell
make up
```

This command builds the service and deploys the built binary to a container, It also spins up a postgres DB instance. 

The default exposed service host URL is: `https://localhost:6504/`

### Ping
To make sure that the service is up and running you can hit this ping URl `http://localhost:6504/ping` and you should receive response: 

```json
{
    "message": "OK"
}
```



## API Endpoints

This section documents how to consume this service APIs, listed below all requests with request/response payloads examples.

### Postman Collection

The URL below contains all endpoints requests collection on postman

`https://www.getpostman.com/collections/0f4b66ade48ff3ed3b8c`

### Register Drone

Used to register a drone 

#### Request URL

`http://localhost:6504/drones/register`

#### Request Method

`POST`

#### Request example

```json
{
    "serial_number": "123455",
    "model": "LIGHT-WEIGHT",
    "weight_limit": 100,
    "battery_capacity": 50
}
```

> **Assumption** Registering a new drone will set it's status to 'IDLE'


#### Response examples

##### 1- successful registration

```json
{
    "code": 200,
    "message": "drone registered successfully",
    "success": false
}
```

##### 2- existing serial number

```json
{
    "code": 422,
    "errors": {
        "SerialNumber": "already registered"
    },
    "message": "invalid data",
    "success": false
}
```

##### 3- not listed model type 

```json
{
    "code": 422,
    "errors": {
        "Model": "not allowed value, one of: 'LIGHT-WEIGHT' 'MIDDLE-WEIGHT' 'CRUISER-WEIGHT' 'HEAVY-WEIGHT'"
    },
    "message": "invalid data",
    "success": false
}
```

##### 4- battery capacity > 100 

```json
{
    "code": 422,
    "errors": {
        "BatteryCapacity": "exceeded the limit, max: 100"
    },
    "message": "invalid data",
    "success": false
}
```

##### 5- weight limit > 500 

```json
{
  "code": 422,
  "errors": {
    "WeightLimit": "exceeded the limit, max: 500"
  },
  "message": "invalid data",
  "success": false
}
```

### Load Drone with medications

This request used to load drone with list of medications

#### request URL

`http://localhost:6504/drones/load`

#### Request Method 

`GET`


#### Request example

```json
{
    "serial_number": "123455",
    "medications": [
        {
            "code": "185459",
            "name": "Ranny",
            "weight": 10
        }
    ]
}
```

#### Response examples

##### 1- success response

```json
{
    "serial_number": "123455",
    "medications": [
        {
            "code": "185459",
            "name": "Ranny",
            "weight": 10
        }
    ]
}
```

##### 2- not registered drone response

```json
{
    "serial_number": "123455d",
    "medications": [
        {
            "code": "1854599",
            "name": "Ranny",
            "weight": 10
        }
    ]
}
```

##### 3- already registered medication code

```json
{
    "code": 422,
    "errors": {
        "Code": "already registered"
    },
    "message": "invalid data",
    "success": false
}
```

##### 4- medication code contains invalid character

```json
{
    "code": 422,
    "errors": {
        "Code": "invalid value pattern"
    },
    "message": "invalid data",
    "success": false
}
```

##### 5- medications weights exceeds the drone mak weight

```json
{
  "code": 422,
  "message": "medications exceeds the max allowed weight limit.",
  "success": false
}
```

### Get Drone Medications

Used to get loaded medications for a specific drone

#### Request URL
`http://localhost:6504/drones/:serialNumber/medications`

#### Request Method
`GET`

#### Response examples

##### 1- success response

```json
{
    "code": 200,
    "medications": [
        {
            "created_at": "2022-04-06T01:36:34.173265+02:00",
            "drone_serial_number": "123455",
            "name": "Ranny",
            "weight": 5,
            "code": "856945",
            "image_url": "/storage/uploads/91947779410_1649201794__.png"
        }
    ],
    "message": "operation done successfully",
    "success": false
}
```

##### 2- not registered drone

```json
{
    "code": 404,
    "message": "not registered drone",
    "success": false
}
```

### Get Drone Battery Capacity

Used to fetch current drone battery capacity

#### Request URL

`http://localhost:6504/drones/:serialNumber/battery`

#### Request Method

`GET`

#### Response examples

##### 1- success response 

```json
{
    "battery_capacity": 50,
    "code": 200,
    "message": "operation done successfully",
    "success": false
}
```

##### 2- not registered drone 

```json
{
    "code": 404,
    "message": "not registered drone",
    "success": false
}
```

### Get Idle Drones

Used to fetch drones with status `IDLE`

#### Request URL

`http://localhost:6504/drones/idle`

#### Request Method

`GET`

#### Response examples

##### 1- success response

```json
{
    "code": 200,
    "drones": [
        {
            "created_at": "2022-04-06T01:39:46.626267+02:00",
            "serial_number": "123459",
            "model": "LIGHT-WEIGHT",
            "weight_limit": 400,
            "battery_capacity": 50,
            "status": "IDLE"
        }
    ],
    "message": "operation done successfully",
    "success": false
}
```
### Update Drone Status

Used to update drone status, one of: ('IDLE' 'LOADING' 'LOADED' 'DELIVERED' 'RETURNING')

#### Request URl

#### Request Method

`PUT`

#### Request Payload 

```json
{
    "serial_number": "123455",
    "status": "LOADED"
}
```

#### Response Examples 

##### 1- success response

```json
{
    "code": 422,
    "errors": {
        "SerialNumber": "not exists in db",
        "Status": "not allowed value, one of: 'IDLE' 'LOADING' 'LOADED' 'DELIVERED' 'RETURNING'"
    },
    "message": "invalid data",
    "success": false
}
```

##### 2- validation errors response

```json
{
    "code": 422,
    "errors": {
        "SerialNumber": "not exists in db",
        "Status": "not allowed value, one of: 'IDLE' 'LOADING' 'LOADED' 'DELIVERED' 'RETURNING'"
    },
    "message": "invalid data",
    "success": false
}
```

### Possible Enhancements

Below list states app possible enhancements I couldn't mount more time to implement.

- Return errors with keys json tags instead of struct keys.
- Enhance layers abstractions by implementing interfaces to support replacing app components easily.
- Add unit & integration tests
- Deploy an online demo.