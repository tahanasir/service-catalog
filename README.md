# service-catalog
Service-catalog is a backend Go application that allows organizations to see a list of their microservices and their versions. Implemented as a REST API backed by a MySQL database. 

## How to run
```
docker-compose up --build
```

All services: 
* Default: `http://localhost:8080/v1/services`
* Query params: `http://localhost:8080/v1/services?search=fx&page=1&limit=5`

Single service:
* `http://localhost:8080/v1/services/5b9a2437-0834-4f01-82b5-f92b7c79d43e`
## Design Considerations

### Database
I decided to go with MySQL as the data storage solution because there was a relationship between services and versions. I am most familiar with MySQL and the sorting and filtering requirements made it a good choice. 

#### Assumptions
* Services are unique
* Service names can be duplicated
* Service descriptions are optional
* Service can have 0 or more versions
* UUID over autoincrement ID for security
* sort by created date desc
* versions are unique based on version name and service id

#### Trade-offs
Although a relational database makes it easy to query the data, the list of versions for a service could be represented as a document. We could have either had a column as a JSON blob or gone with a NoSQL database. Filtering and sorting data would be potentially more difficult. 

### REST API
I decided to go with two GET APIs, one to retrieve a single service and its contents (including versions), the other to fetch all services with contents and the number of versions. Using 3 query params I was able to search, paginate and limit the size of the request.

#### Assumptions
* Service description can be omitted if empty
* VersionCount has to be visible, even if 0
* Versions can be omitted if none for single service request
* return pagination data for frontend

### Testing
Wrote a sample handler test to showcase how we would write unit tests. Mocked the storage interface using mockgen. Tested the happy path.

#### Assumptions
* Test for each API
* Test all paths (happy, 404, 500, etc.)
* Test database method using integration tests and sample data

### Local Development
Used docker-compose to run the application locally running two containers, one for the app, the other for mysql. Pre-loaded seed data to test responses and mapped ports appropriately.
