#!/bin/sh

case "$1" in
	'start') echo "machine will start";;
	'stop')  echo "machine will stop";;
	*)	 echo "I have no idea";;
esac
