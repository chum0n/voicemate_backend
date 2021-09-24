# D-4_2

Repository for backend of project "voicemate".

## voicemate API

An awesome API for project "voicemate".

### Requirements

See modules list in `go.mod`.

### Setup

1. Start application

    ```sh
    docker-compose build
    docker-compose up --detach
    ```

    Server starts on localhost.
    See "http://127.0.0.1:8000/".


## voicemate Database

### Structure

- voicemate_database/data

  Raw data to preserve in a container.

- voicemate_database/initialization

  Initialization DDLs.  
  They should be always synced with production database schema.

## Usage

- Initialize (Reinitialize)

    ```sh
    rm -rfi voicemate_database/data/*
    docker-compose up --detach --build
    ```

- Start

    ```sh
    docker-compose up --detach
    ```

- Stop

    ```sh
    docker-compose down
    ```

- Destroy

    ```sh
    rm -rfi voicemate_database/data/*
    docker-compose down --rmi all --volumes
    ```

