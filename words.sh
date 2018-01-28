#!/bin/sh

(
echo "package main"
echo "var Words = map[string]Pos {"
curl -s https://raw.githubusercontent.com/whilp/words/master/wordlist.csv | sed -e 's/^\(.*\),\(.*\)$/	"\1": \2,/'
echo "}"
) > words.go

