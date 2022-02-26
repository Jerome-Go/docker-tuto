package main
//
// https://docs.docker.com/language/golang/build-images/
//
// ------------1 DEBUT image docker
// 1 $ git init
//
// 2 $ git clone https://github.com/olliefr/docker-gs-ping -> clone de github
//
// 3  creer le fichier  Dockerfile -> instructions pour assembler une image Docker
//
// 4 $ docker build --tag docker-gs-ping .   
//                     -> construire l'image docker
//                     ->'docker build' commande crée des
//                        images Docker à partir du Dockerfile
//                     -> --tag indicateur est utilisé pour étiqueter l'image avec
//                       une valeur de chaîne. Si vous ne transmettez pas a --tag,
//                       Docker l'utilisera latest comme valeur par défaut.
//                     -> --tag sans indicateur, Docker l'utilisera latest 
//
//
// 5 $ docker image ls   -> voir la liste des images sur notre machine locale
//|-                     -|-          -|-            -|-              -|-      -|
// REPOSITORY               TAG       IMAGE ID       CREATED              SIZE
// docker-gs-ping           latest    2a869e438ab2   About a minute ago   533MB
// docker/getting-started   latest    adfdb308d623   2 weeks ago          27.4MB
//|-                     -|-          -|-            -|-              -|-      -|
//
//
// 6 $  docker image tag docker-gs-ping:latest docker-gs-ping:v1.0
//                     -> créer une nouvelle balise pour notre image
//                     -> 1ere argument = l'image "source"
//                     -> 2eme argument = la nouvelle balise à créer
//|-                     -|-       -|-            -|-              -|-      -|
// REPOSITORY               TAG       IMAGE ID       CREATED         SIZE
// docker-gs-ping           latest    2a869e438ab2   9 minutes ago   533MB
// docker-gs-ping           v1.0      2a869e438ab2   9 minutes ago   533MB
// docker/getting-started   latest    adfdb308d623   2 weeks ago     27.4MB
//|-                     -|-       -|-            -|-              -|-      -|
//
//
// 7 docker image rm docker-gs-ping:v1.0     -> Supprimons la balise 
//
// 8 creer un fichier dockerfile.multistage beaucoup plus petit
//
// 9 $ docker build -t docker-gs-ping:multistage -f Dockerfile.multistage .
//                   -> Puisque nous avons maintenant deux dockerfiles, 
//                      nous devons dire à Docker que nous voulons construire 
//                      en utilisant notre nouveau Dockerfile.multistage
//
// 
// 10 $ docker image ls          -> taille du nouveau docker-gs-ping = 23.7MB
//|-                     -|-          -|-            -|-              -|-     -|
// REPOSITORY               TAG          IMAGE ID       CREATED          SIZE
// docker-gs-ping           multistage   7813350d755d   9 seconds ago    23.7MB
// docker-gs-ping           latest       2a869e438ab2   26 minutes ago   533MB
// docker/getting-started   latest       adfdb308d623   2 weeks ago      27.4MB
// |-                    -|-          -|-            -|-              -|-
// ------------1 FIN image docker
//
//
//
// ------------2 DEBUT executer/stope une image dans un conteneur
// 11 $ docker run docker-gs-ping     -> lancer l'image dans un conteneur
//                                    -> curl localhost:8080/ "ne fonctionne pas"
// 
// 12 $ docker run -p 8080:8080 docker-gs-ping -> lancer l'image 
//                                    -> curl localhost:8080/  "fonctionne"
//
// ---- executer une image mode détaché
// 13 $ docker run -d -p 8080:8080 docker-gs-ping   
//                                    -> docker à démarer  en arrièer plan et 
//                                       imprimer l'id du conteneur
//|-                                                                          -|
// jerome@MacBook-Pro docker-gs-ping % docker run -d -p 8080:8080 docker-gs-ping
// 4411edd56333ee30ab46f24b1a447618f6da111f84a2f51448a376cce38479ed
// jerome@MacBook-Pro docker-gs-ping % curl localhost:8080/
// Hello, Docker! <3% 
//|-                                                                          -|
//
// 14 $ docker ps                      -> lister les conteneur
//|-                     -|-                     -|-           -|-             -|
// ID   IMAGE                  COMMAND             PORT           NAMES
// 44   docker-gs-ping     "/docker-gs-ping"       8080->8080     great_feynman
//|-                     -|-                     -|-           -|-             -|
//
// 15 $ docker stop great_feyman       -> fermer conteneur + non du conteneur
//                                     -> Lorsque nous arrêtons un conteneur, il 
//                                  n'est pas supprimé mais le statut est arrêté 
//
// 16 $ docker ps -all                 -> afficher tout les conteneurs meme ceux 
//                                        qui sont à arreter
//|-                     -|-                     -|-           -|-             -|
// ID   IMAGE                  COMMAND             PORT           NAMES
// 44   docker-gs-ping     "/docker-gs-ping"                      great_feynman
// e8   docker/getting     "/docker-entrypoint.…"                 sleepy_shockley
//|-                     -|-                     -|-           -|-             -|
//
// 17 $ docker restart great_feyman    -> redemarrer un conteneur
//
// 18 $ docker rm great_feynman        -> supprimer le conteneur
//                                    -> contrairement a la commande $ docker stop
//                                      le conteneur n'existe plus
//
// 19 point 13 pour creer un redemarrer un conteneur
//
// 20 docker run -d -p 8080:8080 --name jeromeGo docker-gs-ping 
//                                     -> point 13 + name  pour nomer le conteneur
//|-                     -|-                     -|-           -|-             -|
// ID   IMAGE                  COMMAND             PORT           NAMES
// 44   docker-gs-ping     "/docker-gs-ping"      8080->8080      jeromeGO
//|-                     -|-                     -|-           -|-             -|

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello, Docker! <3")
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
