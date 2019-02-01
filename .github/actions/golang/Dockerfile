FROM golang:1.11

LABEL "name"="golang tools action"
LABEL "maintainer"="Marko Mikulicic <mkm@bitnami.com>"
LABEL "version"="1.0.0"

LABEL "com.github.actions.icon"="code"
LABEL "com.github.actions.color"="purple"
LABEL "com.github.actions.name"="Go Tools"
LABEL "com.github.actions.description"="This is an Action to run go commands."

ENV GO111MODULE=on

COPY entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]

