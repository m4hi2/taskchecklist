# Taskchecklist

## Motivation

I have to a lot of tasks where the task itself repeats and has a lot
of small hard to remember very important steps. I want to create this
tool to keep track of these hard to remember steps and things I need
to complete the tasks.

## Design

### ERD

```mermaid
erDiagram
    USERS {
        uuid user_id PK
        text email
        varchar(30) first_name
        varchar(30) last_name
        varchar(255) hashed_password
        timestamp confirmed_at
        timestamp inserted_at
        timestamp updated_at
    }
    USERS ||--o{ BLUEPRINTS : create
    USERS ||--o{ STEPS_COMPLETIONS : has
    USERS ||--o{ INGREDIENTS_COMPLETIONS : has

    BLUEPRINTS {
      uuid blueprint_id PK
      varchar(255) name
      text description
      bigint user_id FK
      varchar(1) visibility
      int completion_count
      uuid remixed_from_blueprint_id FK
      timestamp inserted_at
      timestamp updated_at
    }
    BLUEPRINTS ||--o{ STEPS_COMPLETIONS : has
    BLUEPRINTS ||--o{ INGREDIENTS_COMPLETIONS : has
    BLUEPRINTS ||--o{ STEPS : has
    BLUEPRINTS ||--o{ INGREDIENTS : has

    STEPS {
        uuid step_id PK
        varchar(255) name
        text description
        int step_number
        uuid blueprint_id FK
        timestamp inserted_at
        timestamp updated_at
    }
    STEPS ||--o| STEPS_COMPLETIONS : has
    INGREDIENTS {
        uuid ingredient_id PK
        varchar(255) name
        int count
        uuid blueprint_id FK
        timestamp inserted_at
        timestamp updated_at
    }
    INGREDIENTS ||--o| INGREDIENTS_COMPLETIONS : has

    STEPS_COMPLETIONS {
        uuid completion_id PK
        uuid user_id FK
        uuid blueprint_id FK 
        uuid step_id FK 
        varchar(1) status
        timestamp inserted_at
        timestamp updated_at
    }


    INGREDIENTS_COMPLETIONS {
        uuid completion_id PK
        uuid user_id FK
        uuid blueprint_id FK 
        uuid ingredient_id FK 
        varchar(1) status
        timestamp inserted_at
        timestamp updated_at
    }

```

To start your Phoenix server:

* Run `mix setup` to install and setup dependencies
* Start Phoenix endpoint with `mix phx.server` or inside IEx with `iex -S mix phx.server`

Now you can visit [`localhost:4000`](http://localhost:4000) from your browser
