#!/bin/bash
wrk -t8 -c100 -d20s --script=wrk.lua --latency http://localhost:8080
