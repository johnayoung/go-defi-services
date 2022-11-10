
#!/bin/bash

USERNAME=admin
PASSWORD=micro

# start the server
if ! $(ps aux | grep -q "micro server"); then
	echo "starting server"
	micro server &> /tmp/server.log
fi

# set to local env
micro env set local

# login
micro login --username=$USERNAME --password=$PASSWORD

# run services
for api in coingecko; do
	if ! $(micro auth list rules | grep -q ^${api}); then
		## generate an auth rule to open the api
		micro auth create rule --resource=service:$api:* --access=granted --priority=1 $api
	fi

	## run the service
	echo "starting $api"
	micro run github.com/johnayoung/go-defi-services/pkg/services/$api
done