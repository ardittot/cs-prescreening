# cs-prescreening

Build prescreening stage for credit scoring microservices app with Go

Install required packages
```
go get -u github.com/gin-gonic/gin
go get -u gopkg.in/resty.v1
```

### Compile & run
```
go build
./cs-prescreening
```

### API Specs

SICD
```
curl -X POST -H "Accept: application/json" -H "Content-Type: application/json" -d @./json/ex-sicd.json http://localhost:8000/sicd
```

DH
```
curl -X POST -H "Accept: application/json" -H "Content-Type: application/json" -d @./json/ex-dhn.json http://localhost:8000/dhn
```

Kependudukan
```
curl -X POST -H "Accept: application/json" -H "Content-Type: application/json" -d @./json/ex-kemendagri.json http://localhost:8000/kemendagri
```
