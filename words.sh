#!/bin/sh

(
echo "package main"
echo "var words = []Word{"
curl -s https://raw.githubusercontent.com/whilp/words/master/wordlist.csv | sed -e 's/^\(.*\),\(.*\)$/	{"\1", \2},/'
echo "}"
) | gofmt > words.go

