<!-- PROJECT LOGO -->
<br />
<div align="center">

<h3 align="center" id="readme-top">Wallet</h3>

  <p align="center">
    Julo Programming Test
    <br />
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

E - Wallet Service

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

this application using Golang with Echo Framework and Sqlite Database
for sqlite you don't have to install locally, because the database driver already has sqlite embedded.

### Prerequisites

copy .env.example and rename to .env and change APP_LOCATION to project directory
```editorconfig
APP_LOCATION=<project directory>
```


first install the dependencies
* golang
  ```sh
  go get -v 
  ```

### Installation

1. Run
   ```sh
   go run cmd/main.go server
   ```
2. after that Migration and database file already up and created

<p align="right">(<a href="#readme-top">back to top</a>)</p>


### migration
this migration using go-migrate as dependency, please refer to the [Documentation](https://github.com/golang-migrate/migrate) for installation 

1. Create migration
```shell
migrate create -ext sql -dir ./migrations -seq add_mood_to_users
```

2. Running all migration
```shell
go run cmd/main.go migration up
```
_For more examples, please refer to the [Documentation](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)_


<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- CONTACT -->
## Contact

Muhammad Fatihul Ikhsan - [Instagram](-https://instagram.com/mfikhsan) - mfatihul8902@gmail.com

<p align="right">(<a href="#readme-top">back to top</a>)</p>


