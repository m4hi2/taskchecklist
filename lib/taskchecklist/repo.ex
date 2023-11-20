defmodule Taskchecklist.Repo do
  use Ecto.Repo,
    otp_app: :taskchecklist,
    adapter: Ecto.Adapters.Postgres
end
