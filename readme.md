# NDN - Notification Delivery Network

---
habilitar tudo por api

---
habilitar permissões e rules

---
POST /notifications
GET  /notifications

POST   /inputs
DELETE /inputs/:id

POST   /outputs
DELETE /outputs/:id


PUT|GET /outputs/http.pull/offset

---
GET    /
GET    /health
GET    /workers
GET    /tasks
DELETE /tasks/:id

# daqui pra baixo são dados de as pessoas vão controlar
GET    /inputs
POST   /inputs
DELETE /inputs/:id

# caso o http esteja ligado
POST /inputs/http.push/notifications

GET    /outputs
POST   /outputs
DELETE /outputs/:id

# caso o http esteja ligado
GET  /outputs?type="http.pull" # filtra todos que são http.pull e podemos verificar o offset pra identificar quem está atrasando tudo
GET  /outputs/http.pull/notifications?entity=product&id=123&version=456
GET  /outputs/http.pull/notifications/checkpoint # exibe a id do ultimo registro processado
POST /outputs/http.pull/notifications/checkpoint # armazena a id do ultimo registro processado

---
notifications # guarda todas as notificações
outputs # guarda todos os outputs
notifications_outputs # guarda o relacionamento das notificações com os outputs

quando o menor valor avançar, fazer o cleanup.

---
criar uma tarefa que avisa quem está atrasado em 1 channel















- recebemos as mensagens de alteracao e notificacao
- enviamos para outro servidor

como fazer o expire?
o server retorna um header contendo o que tem dentro da mensagem

Cache: product:123:2134
Cache: sku:323:1234

Temos as chaves de expiração e temos o endpoint.
Assim que algo tiver que ser expirado, vamos no servidor e fazemos a expiração.

// essa é a chamada para o servidor
{
  "entity": "product",
  "id": "123",
  "version": 456
}

// suportar vários tipos de notificação


as apis chamam enviando os dados
