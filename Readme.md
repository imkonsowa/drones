<div align="center">
<h3 align="center">Medications Delivery Using Drones - Golang Service </h3>
</div>

## About

This service is built using the latest golang version `1.18` as a hiring task for musala soft company.

### Used Dependencies

* [Go-Gonic](https://github.com/gin-gonic/gin) The foundational http router and data binding.
* [Gorm.io](https://gorm.io/gorm) The DB ORM.
* [Cron](https://github.com/robfig/cron)  App scheduler for processed cron jobs.
* [Validator](https://github.com/go-playground/validator) Used for applying validation rules on user inputs.

### Used Database

This service used Postgres as the main datastore

## Setup

To run the service, make sure you have Docker and Compose installed on testing machine.

### Run environment

```shell
make up
```

This command builds the service and deploys the built binary to a container, It also spins up a postgres DB instance. 
