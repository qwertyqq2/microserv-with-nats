## Http сервис для обработки заказов
Сервис способен  получать и выдавать заказы в json формате. После получения заказы отправляются на брокер Nats, при помощи утилиты nats-streaming.

### Запуск

#### Docker

     docker-compose up -d

#### Локально

Необходимо установить nats-streaming:0.25.5:, postgres:14

     git clone https://github.com/qwertyqq2/microserv-with-nats

     cd microserv-with-nats

     go build main.go

     ./main

В другом терминале

     nats-streaming --clu ster_name test-cluster --cluster nats://0.0.0.0:8222





