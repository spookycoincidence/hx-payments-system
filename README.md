# 💸 HX Payments System – Wallet Service

Este microservicio es parte del ecosistema **HX Payments System**, una solución modular pensada para la gestión financiera digital.  
En este módulo, desarrollamos el **Wallet Service**, encargado de la administración de billeteras virtuales para usuarios: creación, depósitos, retiros y consulta de saldo.

Este proyecto fue concebido como un **challenge técnico autoasistido**, diseñado para simular un entorno profesional, con foco en buenas prácticas, arquitectura limpia, pruebas y documentación.  
Aunque no se conecta a una base de datos real ni se ejecuta, está estructurado como si fuera a deployearse en un entorno real.

---

## 🧠 ¿Por qué arquitectura hexagonal?

Elegí este enfoque (también llamado **Ports & Adapters**) porque separa claramente las responsabilidades entre:

- **Lógica de negocio** (servicios)
- **Entrada y salida de datos** (transport, repositorios)
- **Dependencias externas** (HTTP, base de datos)

Esto permite:

- Escalabilidad y facilidad para testear
- Bajo acoplamiento y alta cohesión
- Reemplazar cualquier dependencia (ej. base de datos real en lugar de mock) sin tocar la lógica central


## 📦 Otros servicios del ecosistema :

user-service

transaction-service

notification-service

proto/ (pensado para definir contratos entre servicios)

## ⚙️ Stack técnico

🐹 Golang 1.20+

🧪 testing + testify para tests

📦 mux como router HTTP

🐳 Docker (configurado pero no ejecutado)

💾 PostgreSQL simulado (con posibilidad de real integración)

## 📂 Estructura del proyecto
```bash
wallet-service/
├── cmd/
│   └── main.go                # Entry point del servicio
├── internal/
│   ├── model/                 # Entidades de dominio (Wallet)
│   ├── repository/            # Persistencia mockeada
│   ├── service/               # Lógica de negocio
│   └── transport/
│       └── http/              # Handlers HTTP
├── test/
│   └── service_test.go        # Unit tests para lógica de negocio
├── Dockerfile                 # Imagen dockerizada simulada
├── go.mod / go.sum
└── README.md 
```

## 📬 Endpoints disponibles
```bash
POST	  /wallets	                     Crear una billetera
GET	    /wallets/{userID}	             Obtener saldo de usuario
POST	  /wallets/{userID}/deposit	     Realizar un depósito
POST	  /wallets/{userID}/withdraw	   Realizar un retiro
```

## 🔍 Ejemplos de uso (simulado con Postman)

📥 Crear billetera

```bash
POST /wallets
{
  "user_id": 1
}
```



💰 Depositar saldo

```bash
POST /wallets/1/deposit
{
  "amount": 100.0
}

```

💸 Retirar saldo

```bash
POST /wallets/1/withdraw
{
  "amount": 50.0
}

```



📊 Consultar saldo


```bash
GET /wallets/1

```


🧪 Tests
El proyecto incluye tests de unidad en la capa de service, utilizando testify y enfoque TDD:


```bash
go test ./test/...

```



Casos probados:

✅ Creación de billetera

✅ Depósito y retiro exitosos

✅ Manejo de errores: saldo insuficiente, billetera inexistente

## 📄 Dockerfile
FROM golang:1.20

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o wallet-service ./cmd/main.go

EXPOSE 8080

CMD ["./wallet-service"]


## 🧰 Configuración del entorno
Aunque no se utiliza en esta versión, el sistema puede incorporar variables de entorno:
PORT=8080
DB_URL=localhost


📑 Documentación de la API
Esta API REST está documentada en formato OpenAPI v3:

📄 Ver swagger.yaml

👀 Podés visualizarla en editor.swagger.io pegando el contenido.

Incluye endpoints para crear billeteras, consultar saldo, depositar y retirar.


## 🧠 Posibles mejoras / extensiones

✔️ Implementar una base de datos real (PostgreSQL, MYSQL)

✔️ Crear pruebas de integración y mocks con interfaces

✔️ Agregar un API Gateway y autenticación

✔️ Implementar user-service, transaction-service, notification-service

✔️ Documentar con Swagger / OpenAPI

✔️ Deploy en Docker Compose / Kubernetes


## 👩🏻‍💻💻 Autora
Desarrollado por @HuilenVilches 🤓 (aka SpookyCoincidence 🕸️🕷️)
Golang Backend Developer 

Este proyecto forma parte de un challenge técnico autogestionado, simulado como si fuera parte de un entorno productivo real, pero completamente desarrollado sin ejecución ni dependencias externas.

## 📚 Licencia
MIT License © 2025 – Libre para aprender, compartir y adaptar 🚀
