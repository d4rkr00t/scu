## TODO

For Release

- [ ] Extract evo config to a separate file to avoid cache misses for package managers
- [ ] Cache improvements
  - [ ] Custom Hash Input for e.g. environment, etc...
  - [ ] Remote cache (MVP)
  - [ ] Cache rotation – e.g. `evo clear-cache` clears stale cache
  - [ ] Cache diagnostics
    - [ ] List of files that were used to create hash
    - [ ] Diff to show what affected cache
- [ ] More commands
  - [ ] `evo clear-cache <package>`
  - [ ] `evo clear-output <package>` – clears all outputs from packages
- [ ] `--force` to force run a command, ignoring cache
- [ ] Fix all issues highlighted by `--race` flag
- [ ] Improve CLI Output
  - [ ] Show how many tasks were run and how many restored from cache/ignored
- [ ] Watch mode
- [ ] Tests
  - [ ] TBD
- [ ] Website
  - [ ] Documentation
- [ ] Refactor
  - [x] Upgrade Go
  - [ ] Split code in packages
  - [ ] Try DoD
  - [ ] Generics -> Go 1.8
  - [ ] REFACTOR: create tasks from workspaces
  - [ ] REFACTOR: Generic create task from a rule
  - [ ] REFACTOR: Refactor cache to interfaces
  - [ ] REFACTOR: Refactor logger
    - [ ] to generics/interfaces
    - [x] Cleaner API
    - [ ] Add debug level to the logger

---------------

After Release

- [ ] Add .env support
- [ ] Support custom package manager (?)
- [ ] Support: task deps: ["file:somefilepath"]
- [ ] Watchman integration and persistent graph
- [ ] More sophisticated logic around package managers
  - [ ] include lock files into cache hash
- [ ] Clean up only once before any of the tasks run (?)
  - [ ] Might have to restore all of the other outputs
  - [ ] Maybe make it manual?
    - [ ] Otherwise need to hash contents of output
- [ ] Improve unknown dependency error message
- [ ] SANDBOXING !!!
- [ ] Rules
  - [ ] check that command exist
  - [ ] check all required fields
    - [ ] provide schema (?)
- [ ] Different info
  - [x] Show what's included in hash for a workspace – `evo show-hash pkg-a`
  - [x] Show all rules for a WS with overrides – `evo show-rules pkg-a`
  - [ ] Why a package is affected – `evo why pkg-a`
- [ ] Make sure target outputs exist
  - [x] After task run
  - [ ] On package invalidate – [if not invalidated restore outputs]
- [ ] More commands
  - [x] `evo show-scope <package>` – show what packages are in scope
  - [ ] `evo add dep@ver` to add a dependency
  - [ ] `evo remove dep` to remove a dependency
  - [ ] `evo import`
    - [ ] makes sure that all dependencies are on expected versions
    - [ ] makes sure that all dependencies are lifted to the top level package.json
- [ ] Put overrides to file system
- [ ] Use .gitignore in addition to include/exclude in evo config
- [ ] FileSystem
  - [ ] Cache file check-sums
  - [ ] Update file check sums only when modified time is after previous update
  - [ ] Watch file changes during task run
    - [ ] Update FileSystem cache only when a file changes
- [ ] Explore not managing dependencies for workspaces
  - [ ] Let yarn/pnpm/npm manage dependencies for workspaces, drop custom linking logic
- [ ] Generators / Templates
  - [ ] Generate a project from pnpm/yarn workspace and npm scripts
- [ ] Use semver comparison for dependencies (?)
- [ ] https://github.com/adrg/xdg
- [ ] https://github.com/karrick/godirwalk
- [ ] https://github.com/sabhiram/go-gitignore (?)

---

## DONE

- [x] evo run to show a list of available commands
- [x] Support empty commands
- [x] CI Publishing
- [x] Add --since flag for CI
- [x] Add --version command
- [x] CI
  - [x] Publishing to NPM
  - [x] Test in PRs
  - [x] Use evo to manage evo
- [x] Use sync.Map for stats
- [x] Tasks Map similar to Workspace Map
  - [x] Use sync.Map
- [x] REFACTOR: Workspaces map to use the DAG library
- [x] Fix that – Didn't produce outputs still caches and doesn't fail next time
- [x] Cache Improvements
  - [x] Tasks cache should only depend on the cache of the outputs of its dependencies
  - [x] Tasks cache should not depend on other tasks in the workspace
- [x] Support yarn
- [x] Support npm
- [x] Validations
  - [x] Validate external dependencies
  - [x] Validate dep cycles
  - [x] Duplicate WS
  - [x] Validate task cycles
- [x] https://github.com/google/chrometracing
- [x] Fix race condition in task runner
- [x] Add random colors for badges
- [x] https://github.com/zenthangplus/goccm
- [x] Limit concurrency setting
- [x] Investigate rebuild cache miss
- [x] REFACTOR: Task runner
- [x] https://github.com/deckarep/golang-set
- [x] https://github.com/pyr-sh/dag
- [x] Pretty print duration
- [x] Detect scope from `cwd`
- [x] Scoped runs – `evo run build pkg-a`
- [x] Add dependency install and linking
  - [x] Cache project state
  - [x] Install packages
  - [x] Check if node_modules folder exists
  - [x] Link packages
  - [x] Link dev dependencies
  - [x] Link binaries
- [x] Clean task outputs before task run
- [x] Simple Test infra
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
- [x] Local dev loop
