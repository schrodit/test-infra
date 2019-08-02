#! /bin/bash

cat /dev/stdin | testrunner viz | dot -Tpng -o /tmp/test.png && open /tmp/test.png