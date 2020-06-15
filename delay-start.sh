#!/bin/bash
#set -x

: ${SLEEP_SECOND:=10}
wait_for() {
  echo Waiting for $1 to listen on $2...
  for i in `seq $SLEEP_SECOND` ; do
    nc -z $1 $2 > /dev/null 2>&1

    result=$?
    if [ $result -eq 0 ] ; then
      return
    fi
    sleep 1
  done
  echo "Operation timed out" >&2
  exit 1
}

usage() {
  exitcode="$1"
  cat << USAGE >&2
Usage:
  [-d host:port] [-t timeout] [-c command args] [-h]
  -d host:port                        Port detection
  -t timeout                          Timeout in seconds, zero for no timeout
  -c command args                     Execute command with args after the test finishes
  -h                                  Help
USAGE
  exit "$exitcode"
}

if [ $# -eq 0 ] ; then
    usage 1
fi

#declare
DEPENDS=''
CMD=''
while getopts "d:t:c:h" arg
do
    case $arg in
        d)
            DEPENDS=$OPTARG
            ;;
        c)
            CMD=$OPTARG
            ;;
        t)
            SLEEP_SECOND=$OPTARG
            ;;
        h)
            usage 2
            ;;
        ?)
            echo "unkonw argument"
            usage 1
            ;;
    esac
done

for var in ${DEPENDS//,/}
do
    host=${var%:*}
    port=${var#*:}
    wait_for $host $port
done

eval $CMD