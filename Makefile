.PHONY: all gen clean

all: clean gen

gen:
	oapi-codegen --config config.yaml openapi.yaml > twitch.go

clean:
	rm -f twitch.go	