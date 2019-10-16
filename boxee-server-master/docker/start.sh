#!/bin/sh

if [[ -z "${EXTERNAL_IP}" ]]; then
	echo "Please set a value for EXTERNAL_IP, usually it's the IP of the host running the container"
	exit 1
fi

envsubst < /etc/dnsmasq.conf.template > /etc/dnsmasq.conf
envsubst < /etc/hosts.dnsmasq.template > /etc/hosts.dnsmasq

/boxee-server proxy &
exec webproc --config /etc/dnsmasq.conf,/etc/hosts.dnsmasq -- dnsmasq --no-daemon
