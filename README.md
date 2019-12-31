<img src="./doc/logo.png" width="300">

# ChartMuseumUI

ChartMuseumUI is a simple web app that (currently) provides GUI for your charts so you and your team can view and share the technologies your are using to any one at any time (in near future more capabilities will be added).
ChartMuseumUI was written in Go (Golang) with the help of Beego Framework.

<img src="./doc/combine-gif.gif" width="2000">

## Getting Started

These instructions will get you started with your very own private chart repository and UI.

### Usage

ChartMuseumUI uses [ChartMuseum](https://github.com/helm/chartmuseum) as a backend.
To get started quickly, you can build and run the app using docker-compose.

Clone this repo and run the following:

```
docker-compose up
```

This will start ChartMuseumUI at [http://localhost:3000](http://localhost:3000)
and ChartMuseum at [http://localhost:8080](http://localhost:8080).
Check out the source of [docker-compose.yaml](./docker-compose.yaml) and modify for your purposes.

Here is an example docker-compose file defining ChartMuseum with Amazon S3 as a storage and exposing ChartMuseumUI on port 80:
```
version: '2.0'
services:
   ui:
     image: idobry/chartmuseumui:latest
     ports:
      - 80:8080
   chartmuseum:
     image: chartmuseum/chartmuseum:latest
     volumes:
       - ~/.aws:/root/.aws:ro
     restart: always
     environment:
      PORT: 8080
      DEBUG: 1
      STORAGE: "amazon"
      STORAGE_AMAZON_BUCKET: "chartmuseum-bucket"
      STORAGE_AMAZON_PREFIX: ""
      STORAGE_AMAZON_REGION: "eu-west-1"
     ports:
      - 8080:8080
```

Copy this file and run

```
docker-compose up
```
Easy, right? now, we can add our private repository to our Helm client:
```
# choose any name you like
$ helm repo add chartmuseum <chartmuseum-url>
$ helm repo update
# to view our repos list
$ helm repo list
NAME        URL
stable      https://kubernetes-charts.storage.googleapis.com
incubator   http://storage.googleapis.com/kubernetes-charts-incubator
chartmuseum http://localhost:8080
```

Let's upload a chart into our private repository:

```
$ cd /chart/path
# create a chart package - this will create a .tgz file
$ helm package .
# copy packge name and run
$ curl -L --data-binary "@<packge-name>" <chartmuseum-url>/api/charts
```

In the browser, navigate to localhost and view your charts

## config

#### **`/conf/config.yaml`**
```yaml
chartmuseum:
  host: http://ChartMuseum:8080 # you cahrtmuseum host
  hostapi: # set if you are running in a Multitenancy environment
  username: # username for your backend
  password: # password for your backend
ui:
  username: # username for frontend, not yet implemented
  password: # password for the frontend, not yet implemented
```

#### **`/conf/app.conf`**
beego config file for the project

## Built With

* [beego](https://beego.me/) - The web framework used
* [go](https://golang.org/) - Programing language
* [docker](https://www.docker.com/) - Packaged with docker

## Contributing

Code contributions are very welcome. If you are interested in helping make chartmuseumui great then feel free!

## todo
* frontend login
* unit tests

## Authors

* **Ido Braunstain** - *Initial work*
* **Denis Dabischa** - *add backend auth via config file, move to go modules*

See also the list of [contributors](https://github.com/idobry/contributors) who participated in this project.
