import Config

# Only in tests, remove the complexity from the password hashing algorithm
config :argon2_elixir, t_cost: 1, m_cost: 8

# Configure your database
#
# The MIX_TEST_PARTITION environment variable can be used
# to provide built-in test partitioning in CI environment.
# Run `mix help test` for more information.
config :taskchecklist, Taskchecklist.Repo,
  username: System.get_env("PG_USER", "postgres"),
  password: System.get_env("PG_PASS", "postgres"),
  hostname: "localhost",
  database: "taskchecklist_test#{System.get_env("MIX_TEST_PARTITION")}",
  pool: Ecto.Adapters.SQL.Sandbox,
  pool_size: 10

# We don't run a server during test. If one is required,
# you can enable the server option below.
config :taskchecklist, TaskchecklistWeb.Endpoint,
  http: [ip: {127, 0, 0, 1}, port: 4002],
  secret_key_base: "a2j0eRa4svVnz52luVvKVB0g4luK5Qf8oc/uUsKG1GgmaWqm5jvuft2wHShlnNGF",
  server: false

# In test we don't send emails.
config :taskchecklist, Taskchecklist.Mailer, adapter: Swoosh.Adapters.Test

# Disable swoosh api client as it is only required for production adapters.
config :swoosh, :api_client, false

# Print only warnings and errors during test
config :logger, level: :warning

# Initialize plugs at runtime for faster test compilation
config :phoenix, :plug_init_mode, :runtime
