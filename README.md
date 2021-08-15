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
- [x] Proper dependency tracking for deciding when ws is updated
  - [x] Cache WS state
  - [x] Use ws_files_hash + dependencies_hash as a ws_hash
  - [x] Use ws states to build dependencies cache
- [ ] Workspaces struct
  - [ ] Store all WS
  - [ ] Store all updated ws
  - [ ] Store all affected ws
- [ ] Add reusable commands definition to a config – should be able to run a command defined in a config in multiple rules like `@typescript <params>`
- [ ] Run multiple commands e.g. `cmd: ["tsc", "echo"]`
- [ ] Rebuild examples with a real world use cases
- [ ] Add dependency install and linking
  - [x] Cache project state
  - [x] Install packages
  - [x] Check if node_modules folder exists
  - [x] Link packages
  - [ ] Link dev dependencies
  - [ ] Link binaries
- [ ] Validations
  - [ ] Validate external dependencies
  - [ ] Validate dep cycles
  - [ ] Duplicate WS
- [ ] Rebuild affected by a rule change
- [ ] Per rule inputs config (?)
- [ ] Error handling
  - [ ] Task execution
  - [ ] Task dependencies
- [ ] REFACTOR: Generic create a task from a rule
- [ ] TESTS!!!
- [ ] Watch mode
- [ ] Different info
  - [ ] Show path to a cache output – `scu show-cache 058a068`
  - [ ] Show what's included in hash for a workspace – `scu show-hash pkg-a` or `scy show-hash pkg-a target`
  - [ ] Show dependencies of a workspace – `scu show-deps pkg-a`
  - [ ] Show rule with all overrides – `scu show-rule build pkg-a`
- [ ] Scoped runs – `scu run build pkg-a`
- [ ] REFACTOR: Refactor logger to interfaces
- [ ] REFACTOR: Refactor cache to interfaces
- [ ] Throw an error when not in a SCU project
- [ ] Pretty print duration
- [ ] Watch file changes during task run
  - [ ] Update cache only when a file changes
  - [ ] Derive outputs automatically (?)
- [ ] More commands
  - [ ] `scu add dep@ver` to add a dependency
  - [ ] `scu remove dep` to remove a dependency
  - [ ] `scu clear cache`
  - [ ] `scu clear output` – clears all outputs from packages
