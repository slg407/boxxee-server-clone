# boxee-server

This is a replacement Boxee server, it was made to replace the now defunc boxee.tv servers.

## Dependencies

* make (if you plan on doing development)
* Docker

## Usage

_Note: Not tested on Windows_

To run the application, there is no need to clone this repository. First make sure you have Docker installed on the computer you want to use as a server and run this command:

```bash
docker run \
	-d \
	--name dnsmasq \
	-p 53:53/udp \
	-p 5380:8080 \
	-p 80:80 \
	--log-opt "max-size=100m" \
	-e "HTTP_USER=boxee" \
	-e "HTTP_PASS=box" \
	-e "EXTERNAL_IP=<set IP of running host>" \
	--restart always \
	hazward/boxee-server
```

Make sure to change the ``EXTERNAL_IP`` value to the IP of the machine that will be running the server. A web interface for the DNS server will be running at http://EXTERNAL_IP:5380 and the credentials will be `HTTP_USER` and `HTTP_PASS`.

To make use of the server you'll have to add the value of `EXTERNAL_IP` to the DNS servers your router should use.

## Working endpoints

* ``http://app.boxee.tv/api/login``
* ``http://app.boxee.tv/location/weather``, the provider might change to make it easier for development and maintenance
* ``http://[0-9].ping.boxee.tv/``
