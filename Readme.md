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
    </li>
    <li>
      <a href="#api-endpoints">API Endpoints</a>
      <ul>
        <li><a href="#register-drone">Register Drone</a></li>
        <li><a href="#load-drone-with-medications">Load drone</a></li>
      </ul>
    </li>
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



## API Endpoints

This section documents how to consume this service APIs, listed below all requests with request/response payloads examples.

### Register Drone

Used to register a drone 

#### Request URL

`http://localhost:6504/drones/register`

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

##### request URL

`http://localhost:6504/drones/load`

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


