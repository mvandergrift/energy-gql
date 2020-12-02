# GraphQL Data API
### Providing data access to the Fortitude project

Provides a basic GraphQL endpoint to query and mutate the data used by the Fortitude project. 

#### Measures
- Meals
  - Food
  - Food measures (macro nutrients, calories)
  - Food measure units
  - Food eaten within a meal
  - Meal types
- Activities
  - User activities
  - Activity types
- Users
- Query predicates (supporting dynamic querying within GraphQL)
  - Predicate structure
  - Logic 
  - Groups (nested querying)
  
#### Authentication
- Auth0 based oauth2
- jdk supplied by caller contains embeded permissions
- jdk decoded (rsa) using well-known public key
- _Note: This is handled by a authorizer module jwtRsaCustomAuthorizer when requests are routed through AWS APIGateway_

#### Data storage & access
- 3NF relational model
- GORM data access library serves as ORM
- Postgres 12+ backend
- DB credentials stored in either .env file (autoload) or Serverless enviornment (i.e. AWS Lambda config)

#### Infrastructure / deployment
- AWS Lambda [handlers/main.go] use aws-lambda-go module to register GorillaMuxAdapter that proxies requests from there
- Standalone uses same GorillaMuxAdapter as AWS Lambda for simplicity, registers handlers
- Serverless framework used to handle registration of AWS Lambda, API Gateway, and API CloudWatch
 
 ---
[Fortitude Project Status](https://strategic-dev.com/fortitude#sdk)
