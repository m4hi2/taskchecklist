# Taskchecklist

## Motivation

I have to a lot of tasks where the task itself repeats and has a lot
of small hard to remember very important steps. I want to create this
tool to keep track of these hard to remember steps and things I need
to complete the tasks.

## Design

```mermaid
erDiagram
    USERS ||--o{ BLUEPRINTS : create
    USERS {
        bigint user_id PK
        text email
        varchar(255) hashed_password
        timestamp confirmed_at
        timestamp inserted_at
        timestamp updated_at
    }
    BLUEPRINTS {
      bigint blueprint_id PK
      varchar(255) name
      text description
      int completion_count
      bigint user_id FK
      bigint remixed_from_blueprint_id FK
      timestamp inserted_at
      timestamp updated_at
    }
    BLUEPRINTS ||--o{ STEPS : has
    STEPS {
        bigint step_id PK
        varchar(255) name
        text description
        int step_number
        bigint blueprint_id FK
        timestamp inserted_at
        timestamp updated_at
    }
    BLUEPRINTS ||--o{ INGREDIENTS : has
    INGREDIENTS {
        bigint ingredient_id PK
        varchar(255) name
        int count
        bigint blueprint_id FK
        timestamp inserted_at
        timestamp updated_at
    }
```

To start your Phoenix server:

* Run `mix setup` to install and setup dependencies
* Start Phoenix endpoint with `mix phx.server` or inside IEx with `iex -S mix phx.server`

Now you can visit [`localhost:4000`](http://localhost:4000) from your browser
