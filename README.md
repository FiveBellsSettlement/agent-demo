# agent-demo

Public repository for the Five Bells agent demo.

This repository is intentionally small. Its purpose is to provide a stable
GitHub repository where an agent can open a pull request and the Five Bells demo can decide whether the task was delivered.

## Demo Contract

For the first version of the demo, a task is considered delivered when all of
the following are true:

- a pull request exists against `main`
- the pull request branch is `task/<task-id>`
- the pull request title contains `[task:<task-id>]`
- the required GitHub checks pass

The required checks are:

- `lint`
- `test`

There is intentionally no subjective review step in the success condition for
this demo.

## Repository Layout

```text
.github/
  pull_request_template.md
  workflows/
    pr-checks.yml
demo/
  taskref.go
  taskref_test.go
tasks/
  README.md
  example/
    prompt.md
    task.json
.gitignore
README.md
go.mod
```

## Local Use

Run the same checks locally with:

```bash
gofmt -w .
go test -v ./...
```

## Task Files

Task artifacts live under:

```text
tasks/<task-id>/
```

Each task directory should contain:

- `task.json`
- `prompt.md`

The `tasks/example` directory is only a reference shape for the public demo.
