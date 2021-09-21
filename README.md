## TODO

- [ ] Test infra
- [ ] Make sure target outputs exist
- [ ] Scoped runs – `evo run build pkg-a`
- [ ] Put overrides to file system
- [ ] Improve unknown dependency error message
- [ ] REFACTOR: create tasks from workspaces
  - [ ] REFACTOR: Generic create a task from a rule
- [ ] REFACTOR: Refactor cache to interfaces
- [ ] REFACTOR: Task runner
- [ ] REFACTOR: Refactor logger
  - [ ] to interfaces
  - [x] Cleaner API
- [ ] Add dependency install and linking
  - [x] Cache project state
  - [x] Install packages
  - [x] Check if node_modules folder exists
  - [x] Link packages
  - [x] Link dev dependencies
  - [ ] Link binaries
  - [ ] Link peer dependencies (?)
- [ ] Validations
  - [x] Validate external dependencies
  - [x] Validate dep cycles
  - [x] Duplicate WS
  - [ ] Rules
    - [ ] check cycles
    - [ ] check that dependencies exist
    - [ ] check that command exist
    - [ ] check all required fields
      - [ ] provide schema (?)
- [ ] Use .gitignore in addition to include/exclude in evo config
- [ ] Pretty print duration
- [ ] Multi-spinner
- [ ] SANDBOXING !!!
- [ ] Different info
  - [x] Show what's included in hash for a workspace – `evo show-hash pkg-a`
  - [x] Show all rules for a WS with overrides – `evo show-rules pkg-a`
  - [ ] Why a package is affected – `evo why pkg-a`
- [ ] FileSystem
  - [ ] Cache file check-sums
  - [ ] Update file check sums only when modified time is after previous update
  - [ ] Watch file changes during task run
    - [ ] Update FileSystem cache only when a file changes
- [ ] More commands
  - [ ] `evo add dep@ver` to add a dependency
  - [ ] `evo remove dep` to remove a dependency
  - [ ] `evo clear-cache <package>`
  - [ ] `evo clear-output <package>` – clears all outputs from packages
- [ ] Remote cache
- [ ] Watch mode
- [ ] Generators / Templates
  - [ ] Generate a project from pnpm/yarn workspace and npm scripts
- [ ] `--force` to force run a command, ignoring cache (?)
- [ ] Use semver comparison for dependencies (?)

---

## DONE

- [x] Throw an error when not in an EVO project
- [x] Fail when not in `evo` project
- [x] Cache only outputs of the target
- [x] Get rid of `cacheOutputs`
- [x] Pre-hash workspace
- [x] Rule `outputs`
- [x] Skip WS if target doesn't exist
- [x] Running multiple targets `evo run build test lint`
- [x] Non-zero exit code when had failing tasks
- [x] FileSystem module
  - [x] In memory cache of file checksums, update only when update time of a file changed
    - [x] Preserve cache on disk
  - [x] Add / remove files from cache
  - [x] Error handling
- [x] Replace `build` with `run`
- [x] Add configuration
  - [x] Namespace config
- [x] Move more caching logic to a task
  - [x] Cache tasks with a task name prefix
- [x] Add logging and progress
- [x] Build stats
- [x] REFACTOR: Get rid of runner and project in favor of context
  - [x] Remove runner struct
  - [x] Remove project struct
  - [x] Split Run function into stages
  - [x] Create task runner
- [x] Proper dependency tracking for deciding when WS is updated
  - [x] Cache WS state
  - [x] Use WS_files_hash + dependencies_hash as a WS_hash
  - [x] Use WS states to build dependencies cache
- [x] Rebuild affected by a rule change
  - [x] Use a hash of all rules that apply to a WS
  - [x] Move preprocessed rules to WS
- [x] Add reusable command definition to a config – should be able to run a command defined in a config in multiple rules like `@typescript <params>`
- [x] Store STDOUT + STDERR of a command and replay output
- [x] Rename to "evo" from "evoke"
- [x] Throw an error when pnpm errors
- [x] Throw if target doesn't exist
- [x] Stricter overrides
  - [x] Replace glob with a relative path to a group or a certain package
- [x] Simple Error handling
  - [x] Task execution
  - [x] Task dependencies
- [x] Cache tasks afterwards
  - [x] Re-hash once per WS
  - [x] Top sort tasks
- [x] Race condition leads to cache pollution
- [x] Workspaces struct
  - [x] Store all WS
  - [x] Store all updated WS
  - [x] Store all affected WS
- [x] Replace task runner output with a spinner
- [x] Accumulate errors and print them at the end
