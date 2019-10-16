run-dns: build stop
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
	boxee-server

build:
	docker build -t hazward/boxee-server .

release: build
	docker login
	docker push hazward/boxee-server:latest

stop:
	docker stop dnsmasq || true
	docker rm dnsmasq || true
