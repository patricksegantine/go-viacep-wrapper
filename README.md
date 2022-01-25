# Viacep Wrapper
Viacep Wrapper is an API wrapper built with Golang used to find address by zipcode (Brazil only). This project was developed for study purposes.

<br />

## Architecture

Let's imagine the following challenges to face:
- we have a 3rd party API that provides us with the postal code service
- we need to isolate our domain from changes in that service
- we need to provide all our internal applications with a centralized service that delivers CEP (anti-corruption layer?)

Based on the problems listed above, I created the container diagram below (C4 Model) to exemplify the adopted architecture:


![browser response](/docs/architecture/container-diagram.png)

<br />

## Project structure

Well, for now, let's just focus on the ViaCep Wrapper API. Although the project itself is very simple, I tried to apply the concepts of Hexagonal Architecture (Ports and Adapters) and the principles of SOLID, such as inversion of control and dependency injection.

NOTE: To develop the api layer and expose the application routes I chose the Fiber framework [https://gofiber.io].



```
src
 |__adapters
 |  |__api
 |
 |__internal
    |__application
    |__domain
```


<br />

## How to test it?

NOTE: before executing the commands below, make sure that Go has been installed correctly [https://go.dev/doc/install].

After clone this repo and inside the 'src' directory run the command:

`
go run main.go
`

<br />


## Sending a command-line request using cURL:

Request

`
curl http://127.0.0.1:8080/api/find-address-by-zipcode/04543900 \
-H "Accept: application/json"
`

<br />

Response:

```
{
    "message":"",
    "data": {
        "AddressLine1":"Avenida Presidente Juscelino Kubitschek",
        "AddressLine2":"Vila Nova Conceição",
        "AddressLine3":"1830",
        "City":"São Paulo",
        "State":"SP",
        "Country":"Brazil",
        "ZipCode":"04543-900"
    }
}
```


If you don't have familiarity with cURL you can use browser to see the response from API:

![browser response](/docs/architecture/viacep-wrapper.png)
