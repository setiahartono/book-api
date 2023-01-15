# Book API

## Overview

An API to facilitate book pickup schedule to prevent overcapacity during pandemic time. To prevent widespreak outbreak it is safer to propose pickup quota in a timeslot in a crowd control attempt. A user should be able to make books pickup on proposed schedule. Should the user be able to book a schedule successfuly, the timeslot quota will be reduced by one, thus preventing another user to book the mentioned timeslot whenever the quota hits zero.


This API was made by using Go 1.18 without any persistent storage (only making use of Array as a means of storage). The data state is highly simulated with dummy record to illustrate simple API creation by using Go.

This repository included both the API and the unit test.

## How to run

### Running Go directly

- Make sure that Go is properly installed `go version`

- Install dependencies `go mod tidy` and `go mod download`

- Run the server by accessing the root directory and type `go run .`

#### Running with Docker

- Make sure that you have Docker installed `docker -v`

- Build the Docker Image, type `docker build . && docker-compose build --no-cache`

- Run the container, type `docker-compose up`

## The Unit Test

- Just go to the root directory

- Run `go test ./api/tests/`

- The test cases can be viewed at `api/tests` folder

## The API

The followings are the short documentation of the API.

### Books

- Show Book List: `GET /books/`
    - Description:
        - Show a list available book and the information of the book
    - Response (200)  
        ```json
        {
            "OL1709032M": {
                "Title": "Reliable computer systems",
                "Authors": [
                "Daniel P. Siewiorek"
                ],
                "Edition": "OL1709032M"
            },
            "OL19609345M": {
                "Title": "Digital design",
                "Authors": [
                "M. Morris Mano"
                ],
                "Edition": "OL19609345M"
            },

            .
            .
            .

            "OL4731539M": {
                "Title": "Gödel, Escher, Bach",
                "Authors": [
                "Douglas R. Hofstadter"
                ],
                "Edition": "OL4731539M"
            }
        }
        ```
    - cURL Import
        ```cURL
        curl -X 'GET' 'http://<URL>/books/' -H 'accept: application/json'
        ```

- Show Schedule List: `GET /schedules/`
    - Description:
        - Show a list of schedule timeslot and quota, schedule ID is a reference to book the timeslot
    - Response (200)  
        ```json
        {
            "0001": {
                "Timeslot": "2023-12-31 09:00:00 to 2023-12-31 09:45:00",
                "ScheduleID": "0001",
                "Quota": 20
            },
            "0002": {
                "Timeslot": "2023-12-31 09:45:00 to 2023-12-31 10:30:00",
                "ScheduleID": "0002",
                "Quota": 0
            },
            "0003": {
                "Timeslot": "2023-12-31 10:30:00 to 2023-12-31 11:00:00",
                "ScheduleID": "0003",
                "Quota": 10
            }
        }
        ```
    - cURL Import
        ```cURL
        curl -X 'GET' 'http://<URL>/schedules/' -H 'accept: application/json'
        ```

- Book a Schedule: `POST /schedules/`
    - Description
        - Book a schedule
    - Request Data
        - schedule_id - ID of the schedule (string)
        - books - list of object containing edition_number
        ```json
        {
            "schedule_id": "0001",
            "books": [
                {"edition_number": "OL1709032M"},
                {"edition_number": "OL19609345M"}
            ]
        }
        ```
    - Response (200)  
        ```json
        {
           "message": "Timeslot booked at 2023-12-31 09:00:00 to 2023-12-31 09:45:00"
        }
        ```
    - cURL Import
        ```cURL
        curl -X POST \
        'http://<URL>/schedules/' \
        --header 'Accept: application/json' \
        --data-raw '{
            "schedule_id": "0001",
            "books": [
                {"edition_number": "OL1709032M"},
                {"edition_number": "OL19609345M"}
            ]
        }'
        ```

- Show My Schedule: `GET /myschedules/`
    - Description:
        - Show the timeslot that you have booked along with the book information
    - Response (200)  
        ```json
        {
            "0001": {
                "ScheduleID": "0001",
                "Timeslot": "2023-12-31 09:00:00 to 2023-12-31 09:45:00",
                "Books": [
                {
                    "Title": "Reliable computer systems",
                    "Authors": [
                    "Daniel P. Siewiorek"
                    ],
                    "Edition": "OL1709032M"
                },
                {
                    "Title": "Digital design",
                    "Authors": [
                    "M. Morris Mano"
                    ],
                    "Edition": "OL19609345M"
                }
                ]
            }
        }
        ```
    - cURL Import
        ```cURL
        curl -X 'GET' 'http://<URL>/myschedules/' -H 'accept: application/json'
        ```

## Author

### Bayu Sasrabau Setiahartono

[LinkedIn](https://linkedin.com/in/setiahartono)