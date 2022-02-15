# Url Shortener

###### This web application was build with Gin framework, you can install dependencies using:
`go 1.17`
`go get github.com/url-shortener`

###### You need to install redis,i used docker for this:
`docker run -d --name=redis -p 6379:6379 redis`

###### Also add the following Envrinments:
`URL_LENGTH=8` `REDIS_URL=localhost:6379`