timezone = "utc" # default

api {
  enabled = true

  transport "http" {
    port = 8080
  }

  transport "https" {
    port = 443
    # passar os certificados ou algo que faça o certificado automático como o certmagic
  }
}

worker {
  enabled = true

  task "cleanup" {
    load = 2           # só ativa se o load estiver com um valor menor que o parametro
    cron = "0 1 * * *" # executa 1h da manhã e respeita a regra anterior
  }

  task "overdue" {
    after = "cleanup"   # vai automaticamente executar depois do cleanup
    load  = 2           # só ativa se o load estiver com um valor menor que o parametro
    cron  = "0 1 * * *" # executa 1h da manhã e respeita a regra anterior
  }
}

storage {
  host     = "localhost"
  port     = 5439
  database = "ndn"
  user     = "postgres"
  port     = "postgres"
}
