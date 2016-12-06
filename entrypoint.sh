#!/bin/bash
echo "Nuance set key...."

/usr/local/bin/oplicmgr -c 123456789abc -N $NUANCE_KEY

echo "Starting...."
exec $@
