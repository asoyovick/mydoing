# MyDoing

**A goal-driven task manager built with Go.**

MyDoing is a command-line application designed to connect everyday tasks with meaningful goals. It helps users organize their objectives, break them into actionable tasks, and track progress toward the outcomes they want to achieve.

Rather than treating tasks as isolated items, MyDoing focuses on the relationship between what you do and what you want to accomplish.

## Table of Contents

- [Overview](#overview)
- [Core Concepts](#core-concepts)
- [Features](#features)
- [How It Works](#how-it-works)
- [Project Structure](#project-structure)
- [Technology Stack](#technology-stack)
- [Getting Started](#getting-started)
- [Usage](#usage)
- [Development](#development)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [License](#license)

## Overview

Traditional to-do lists help you remember what needs to be done. MyDoing takes a goal-oriented approach by connecting tasks to the objectives they support.

For example, imagine your goal is to become proficient in Go.

You could break that goal into tasks such as:

- Learn Go fundamentals.
- Practice loops, functions, and data structures.
- Build command-line applications.
- Complete a larger Go project.

MyDoing is designed to help organize these tasks and track how they contribute to the overall goal.

## Core Concepts

### Goals

Goals represent the outcomes you want to achieve. They provide direction and context for your work.

### Tasks

Tasks represent the individual actions required to make progress toward a goal. Linking tasks to goals helps keep daily work aligned with larger objectives.

### Objectives and Key Results (OKRs)

MyDoing includes a dedicated OKR module for organizing objectives and key results. OKRs provide a framework for defining desired outcomes and measuring progress toward them.

### Progress and Evidence

The progress subsystem contains components for progress evaluation, scoring, and evidence. These components form part of the project's approach to tracking progress.

## Features

The project structure includes components for the following capabilities:

- **Goal management:** Goal models and associated services.
- **Task management:** Task models and task-related services.
- **OKR support:** Objectives, key results, and related services.
- **Progress evaluation:** A progress engine and scoring components.
- **Evidence handling:** Components for working with progress evidence.
- **Persistent storage:** SQLite integration.
- **External integrations:** Dedicated areas for filesystem, Git, and GitHub integrations.

The availability and maturity of individual features depend on their current implementation.

## How It Works

MyDoing is organized around the relationship between goals, tasks, and measurable progress.

The intended workflow is:

1. Define a goal or objective.
2. Break it down into actionable tasks.
3. Associate tasks with the goals they support.
4. Track task completion and relevant evidence.
5. Evaluate progress toward the intended outcome.

The central idea is simple: **your tasks should have a purpose, and your progress should reflect the work you do toward your goals.**

## Project Structure

MyDoing uses a modular Go project structure that separates domain logic, storage, configuration, and integrations.

```text
mydoing/
├── cmd/
│   └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── goals/
│   │   ├── goal.go
│   │   └── service.go
│   ├── integrations/
│   │   ├── filesystem/
│   │   ├── git/
│   │   └── github/
│   ├── okr/
│   │   ├── keyresult.go
│   │   ├── objective.go
│   │   └── service.go
│   ├── storage/
│   │   ├── progress/
│   │   │   ├── engine.go
│   │   │   ├── evidence.go
│   │   │   └── scoring.go
│   │   └── sqlite.go
│   └── tasks/
│       ├── services.go
│       └── task.go
├── go.mod
└── README.md
```

### Directory Responsibilities

| Path | Purpose |
|---|---|
| `cmd/` | Application entry point |
| `internal/config/` | Application configuration |
| `internal/goals/` | Goal models and services |
| `internal/tasks/` | Task models and services |
| `internal/okr/` | Objectives, key results, and services |
| `internal/storage/` | Storage implementation |
| `internal/storage/progress/` | Progress engine, scoring, and evidence |
| `internal/integrations/filesystem/` | Filesystem integration components |
| `internal/integrations/git/` | Git integration components |
| `internal/integrations/github/` | GitHub integration components |
| `go.mod` | Go module definition and dependencies |

The `internal/` directory also establishes package boundaries by restricting imports from outside the enclosing Go module's applicable directory tree.

## Technology Stack

- **Go:** Application implementation.
- **SQLite:** Database technology referenced by the storage implementation.
- **Git:** Source control and the basis for the dedicated Git integration.
- **GitHub:** Platform targeted by the dedicated GitHub integration.

The exact libraries, integration behavior, and database configuration should be documented as their implementation is verified.

## Getting Started

### Prerequisites

- Go installed and configured.
- Git installed.
- A working copy of the MyDoing repository.

### Clone the repository

```bash
git clone <repository-url>
cd mydoing
```

Replace `<repository-url>` with the actual repository URL.

### Build the application

From the repository root, run:

```bash
go build ./cmd
```

This checks and builds the application package. The resulting executable location depends on the Go build command and package configuration.

### Run the application

The application entry point is located at `cmd/main.go`. The exact runtime command and any required configuration should be documented after verifying the current startup implementation.

## Usage

MyDoing is intended to be operated through a command-line interface.

The usage documentation should cover the commands supported by the current implementation, including:

- Viewing available commands or help.
- Creating and listing goals.
- Creating tasks and linking them to goals.
- Updating task status.
- Inspecting goal progress.
- Managing objectives and key results.

These are documentation targets, not a claim that every command is already implemented.

Once the command interface is confirmed, this section should include real commands, arguments, example output, and relevant error cases.

## Development

### Verify the project

Run the Go test suite:

```bash
go test ./...
```

Build all packages:

```bash
go build ./...
```

Format Go source files:

```bash
gofmt -w .
```

Run these commands from the repository root. Formatting changes should be reviewed before committing.

### Design principles

Development should preserve clear boundaries between:

- Domain models and business logic.
- Task and goal management.
- Progress evaluation and scoring.
- Persistent storage.
- External integrations.

This separation makes the code easier to test, maintain, and extend as MyDoing grows.

## Roadmap

Potential areas for future development include:

- A complete goal and task management workflow.
- Automatic goal progress recalculation when linked tasks change.
- Clear progress summaries and reporting.
- Evidence-based progress evaluation.
- Deeper Git and GitHub integrations.
- Improved CLI help and command documentation.
- Automated tests for goal, task, and progress logic.

These are potential directions and should be adjusted to reflect the project's existing implementation and priorities.

## Contributing

Contributions, bug reports, and suggestions are welcome.

1. Fork the repository.
2. Create a branch for your changes.
3. Make a focused change.
4. Run the relevant tests.
5. Open a pull request explaining the change.

Keep changes focused, preserve the project's package boundaries, and document any new commands or behavior.

## License

The project license has not been specified here. Add the chosen license and its corresponding file before distributing the project under an open-source license.

---

**MyDoing: Connect your tasks to your goals. Track the work that moves you forward.**
