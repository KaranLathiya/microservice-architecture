
# Microservice Architecture

## Summary

Two different services named Expert and Portfolio. 

Expert service for list of expert details and Portfolio service for list of portfolio  details created by expert.

Expert service running on port 8081 and portfolio service running on 8082.

## Run Locally

First clone the project

    git clone https://github.com/KaranLathiya/microservice-architecture.git

Then download the dependency on both service seperately using 

    go mod tidy

Then start the cockroachdb and create the databse named expert and portfolio then migrate both database seperately from respective migration.sql file.

Set the .env file from .env.example in bot service (Set the same JWTKEY)

Then start the both service seperately using 

    go run main.go

## Internal Structure

To get the list of expert details there is paginated api --GET method

    http://localhost:8081/expert?includePortfolios=true&includeNumberOfPortfolios=4&page=0

Set the page parameter to get the list of expert in number of 10.(page default set to 0)

Set the includePortfolios parameter to get the list of portfolios of expert and for get the exact number of portfolios set the includeNumberOfPortfolios.(includePortfolios default set to true and includeNumberOfPortfolios set to default 3)

If includePortfolios value is set to false then only expert details fetched otherwise if set value set to true expert details with portfolis details fetched.

For fetching portfolis details internally JWT token created with default expiration time of 5 minutes on expert service and then set it to the authorization field on header of request and in body set the array of expertIds and includeNumberOfPortfolios for internal call on portfolio service of api --POST method 

    http://localhost:8082/user/portfolio

On the side of portfolio service first verify the token through middleware and after that get the details from request body and then provide the response in format of map[expert]{ totalPortfolios: int portfolios: array of portfolios}

After getting response from internal call on portfolio service on the side of expert service map the portfolios details with respective expert.





