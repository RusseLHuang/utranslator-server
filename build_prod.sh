cp -R ./config/config.development.toml ./config.toml
go build main.go
pmgo start $GOPATH/src/github.com/utranslator-server