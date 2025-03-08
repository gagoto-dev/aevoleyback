# aevoleyback

## Para desplegar en desarrollo aceptando cambios
```
$ npm i -g nodemon

# Ejecutamos nodemon, para escuchar la extensión "go", cuando escuche cambios
# en ficheros .go, se reiniciará el servicio.
$ nodemon -e go --signal SIGTERM --exec 'go' run ./cmd
```
