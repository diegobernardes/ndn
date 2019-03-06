# Any valid value for funciton https://godoc.org/time#LoadLocation.
timezone = "UTC"

api {
  enabled = true

  transport "http" {
    port = 80
  }
}

worker {
  enabled = true

  task "cleanup" {
    # Only execute the task if the server load is equal or lower the value.
    # Optional parameter and has no default value.
    load = 2

    # The time when the task should be executed.
    # Required parameter and has no default value.
    cron = "0 1 * * *"
  }

  task "overdue" {
    # Trigger this function after the given task. If these value is set, 'cron' should be unset.
    # Optional parameter and has no default value.
    after = "cleanup"

    # Only execute the task if the server load is equal or lower the value.
    # Optional parameter and has no default value.
    load = 2

    # The time when the task should be executed. If these value is set, 'after' should be unset.
    # Required parameter and has no default value.
    cron = "0 1 * * *" # executa 1h da manhã e respeita a regra anterior
  }
}

storage {
  type            = "postgres"
  host            = "localhost"
  port            = 5439
  database        = "ndn"
  user            = "postgres"
  password        = "postgres"
  max-connections = 100
  acquire-timeout = "1s"
}
