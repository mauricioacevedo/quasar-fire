# quasar-fire
Repositorio para prueba Fuego de Quasar

## Features
- GetLocation: retrieve the posible location of the Imperial ship
- GetMessage: using the Resistance satellites the service returns the message that was broadcast by the Imperial ship
- POST - /topsecret: service used to obtain the complete information of the Imperial ship 


## How to use
1. First you to have in your system git and golang@1.4.0
2. clone the project: 
```
git clonehttps://github.com/mauricioacevedo/quasar-fire.git 
```
into golang native folder: `$GOHOME/github.com/`

3. For some environments its necessary to have the env var GO111MODULE : `export GO111MODULE=on`

4. Execute with: `go run main.go`
