# quasar-fire
Repositorio para prueba Fuego de Quasar

## Features
- GetLocation: retrieve the posible location of the Imperial ship
- GetMessage: using the Resistance satellites the service returns the message that was broadcast by the Imperial ship
- POST - /topsecret: service used to obtain the complete information of the Imperial ship 
- POST - /topsecret_split: service used to setup information for one satellite 
- GET - /topsecret_split: retrieve the posible location of the Imperial ship  with the split information


## How to use
1. First you to have in your system git and golang@1.4.0
2. clone the project: 
```
git clonehttps://github.com/mauricioacevedo/quasar-fire.git 
```
into golang native folder: `$GOHOME/github.com/`

3. For some environments its necessary to have the env var GO111MODULE : `export GO111MODULE=on`

4. Execute with: `go run main.go`


## Development Work flow

The repository development workflow will be based in the gitflow workflow. For more information please refer to: `https://www.atlassian.com/git/tutorials/comparing-workflows/gitflow-workflow`
 
1. Every change need to be posted in a separated branch 
2. A pull request neeeds to be generated to the develop branch
3. With the approve of 2 repository collaborators the change can be merged to de develop branch
4. Develop a Master pull request need to be reviewed and approved by 3 collaborators before it could be merged
5. After de merge in master the proyect is deployed.

## CI/CD

 Github actions are used to ensure that the project is ok:
    - Lint: all project is passed by a linter to validate code style.
    - Build and Tests: the project is build to verify that the repo is in a good shape
    - Dockerhub push: a docker image is build and published to Docker hub
    - ECS Deploy: the base project is deployed to 

## Whats next?
- Support for multiple Resistance satellites.
- Improve algorithms for message interpretation.
- load tests.
- Sync satellites info with goroutines.

## Authors

#### Mauricio Acevedo - mauricio.acevedo@gmail.com
