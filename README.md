Given
-----

You are given an external API endpoint which allows to query recipe information. Each recipe can be accessed by an `integer` id.
The recipe id enumeration starts from `1`.

Example HTTP calls

```
curl -X GET https://s3-eu-west-1.amazonaws.com/test-golang-recipes/1
curl -X GET https://s3-eu-west-1.amazonaws.com/test-golang-recipes/2
curl -X GET https://s3-eu-west-1.amazonaws.com/test-golang-recipes/5
```

Task
----

Design an application which would act as a reverse proxy and expose the _aggregated recipes_ from the external API over HTTP.

#### Requirements

- The recipes in the aggregated list **must** contain the same data as the original recipes. Data modifications are **not allowed**.
- The endpoint response **must** be `JSON` encoded.
- The endpoint response time **must** be lower than `1s`.
- The application should be stateless, i.e. it is **not allowed** to cache the recipe response on the application side.
- The endpoint **should not** render all the recipes in a single response. It is **allowed** to make [use of pagination](http://docs.oasis-open.org/odata/odata/v4.01/cs01/part2-url-conventions/odata-v4.01-cs01-part2-url-conventions.html#_Toc505773300).

##### Use Case #1 - all recipes

A user should be able to retrieve an aggregated list of **all the recipes** from the source API.

_Specific requirements_
- The endpoint **must** provide access to ALL available recipes.
- The order for the rendered recipes **is irrelevant**
- The solution should operate under the assumption that the source API contains an unlimited number of recipes.

> `All available recipes` are the recipes with the `id` lower than the `id` with the first `404 Not Found` HTTP response status code.
>
> For example, if
>  `curl -X GET https://s3-eu-west-1.amazonaws.com/test-golang-recipes/99999` returns HTTP status code `200 OK`
> and
> `curl -X GET https://s3-eu-west-1.amazonaws.com/test-golang-recipes/100000` returns HTTP status code `404 Not Found`
> then
> `all available recipes` are the ones with the `ids` from 1 to 99999



Example endoint: `GET http://myservice.io/recipes`

```json
[
    {
        "id": "5",
        // ...
    },
    {
        "id": "1",
        // ...
    },
    {
        "id": "2",
        // ...
    }
]
```

##### Use Case #2 - recipes by `id`

A user should be able to retrieve a list of **aggregated recipes** from the source API by a given `id`.

_Specific requirements_

- The endpoint **must** provide access to the recipes by the provided `id`.
- The recipes should be ordered by `prepTime` from lowest to highest.

Example endpoint and response: `GET http://myservice.io/recipes?ids=1,2,5`

```json
[
    {
        "id": "1",
        "prepTime": "PT30",
        // ...
    },
    {
        "id": "5",
        "prepTime": "PT30",
        // ...
    },
    {
        "id": "2",
        "prepTime": "PT35",
        // ...
    }
]
```
Implementation
==============
Used zero external libs

## Code organisation

Directories structure is following https://github.com/golang-standards/project-layout

## Prerequisite

* Go >= 1.13

## Run
Default listening URL: localhost:8080

### Local
```
make run
```
### Docker
To be fixed...

## Build

### Local
```
make build
```

### Docker
To be fixed...

## Test
```
make test
```

### Default configs
* TOP=100
* SKIP=0
* REQUEST_TIMEOUT_MILLISEC=1000 ... to be extracted from hardcode
* CONCURRENCY_LIMIT=10

### TODO
* Docker: to finish a battle on Mac
* Refactoring: decouple domain helpers
* Testing: integration test of the recipesHandler and consider using the assert lib
* Logger!


