```
/*
 * ┌────────────┐   ┌────────────┐       ┌─────────────────────────┐
 * │            │   │            │       │                         │
 * │ evo config ├──►│ workspaces ├──────►│                         │
 * │            │   │            │       │         labels          │
 * └─────────┬──┘   └────────────┘       │                         │
 *           │                           │    workspace::target    │
 *           │      ┌────────────┐       │        ::target         │
 *           │      │            │       │                         │
 *           └─────►│ targets    ├──────►│                         │
 *                  │            │       │                         │
 *                  └────────────┘       └─────────────────────────┘
 */
```

## TODO

- [ ] FileSystem <- 🦄
  - [ ] Cache file check-sums
  - [ ] Update file check sums only when modified time is after previous update
  - [ ] Watch file changes during task run
    - [ ] Update FileSystem cache only when a file changes
- [ ] Cache improvements
  - [ ] Custom Hash Input for e.g. environment, etc...
  - [ ] Remote cache (MVP)
  - [ ] Cache rotation – e.g. `evo clear-cache` clears stale cache
  - [ ] `evo clear-cache <label>`
  - [ ] `evo clear-output <label>` – clears all outputs from packages
- [ ] Validate missing target dependencies
- [ ] Workspace tags
  - [ ] Add rules to define what tags can depend on what tags, e.g. library and app, app can depend on a library but not the other way around
- [ ] Private targets – Only can be run as a dependency of another target in the same workspace
- [ ] Filter target inputs, apply globs to workspace files, depends on ((FileSystem))
- [ ] Should manage toolchain (?)
- [ ] Support sub workspaces (?)
  - [ ] Pros / Cons
  - [ ] Automatic ignores from top level
- [ ] Fix linking peer dependencies (?)
- [ ] Watch mode
- [ ] Evo status
    - [ ] Show some stats info about the repo
- [ ] Reporter
  - [ ] Spinner disappears
  - [ ] Useless when a lot of tasks
    - [ ] Show only top level and failed. In spinner report progress with – succeeded (+from cache) / failed (+from cache)
  - [ ] On error need to report status – succeeded (+from cache) / failed (+from cache)
  - [ ] Move reporter locking into a separate go routine
- [ ] Website
  - [ ] Documentation
- [ ] Cache debugging
  - [x] Centered around tasks
  - [ ] Shows what task deps have changed
  - [ ] Shows what external workspace deps have changed
  - [x] Shows if workspace files have changed
  - [x] Shows everything that is included in building a cache
  - [x] Shows if command changed
- [ ] Refactor
  - [x] Upgrade Go
  - [x] Split code in packages
  - [x] Generics -> Go 1.8
  - [x] REFACTOR: create tasks from workspaces
  - [x] REFACTOR: Generic create task from a rule
  - [ ] REFACTOR: Refactor cache to interfaces
  - [ ] REFACTOR: Refactor logger
    - [ ] to generics/interfaces
    - [x] Cleaner API
    - [x] Add debug level to the logger
- [ ] Add .env support
- [ ] Watchman integration and persistent graph
- [ ] SANDBOXING !!!
- [ ] Targets
  - [ ] check that command exist
  - [ ] check all required fields
    - [ ] provide schema (?)
- [ ] Different info
  - [x] Show what's included in hash for a target – `evo show-hash pkg-a::target`
  - [x] Show all rules for a WS with overrides – `evo show-rules pkg-a`
  - [ ] Why a target is affected – `evo why pkg-a::build`
- [ ] More commands
  - [x] `evo show-scope <package>` – show what packages are in scope
  - [ ] `evo import`
    - [ ] makes sure that all dependencies are on expected versions
    - [ ] makes sure that all dependencies are lifted to the top level package.json
- [ ] Use .gitignore in addition to include/exclude in evo config
- [ ] Generators / Templates
  - [ ] Generate a project from pnpm/yarn workspace and npm scripts
- [ ] Use semver comparison for dependencies (?)
- [ ] Task Shell Like commands:
    - [ ] https://github.com/denoland/deno_task_shell
    - [ ] https://deno.land/manual/tools/task_runner#built-in-commands
- [ ] https://github.com/karrick/godirwalk
- [ ] https://github.com/sabhiram/go-gitignore (?)
- [ ] Support target deps: ["package::sometarget"]
- [ ] Support file deps: ["file::somefilepath"]
- [ ] Linker: error when duplicate/overriding bins
- [ ] Support project tasks, e.g. run pnpm install for the whole project ((Support file deps))

---

## DONE

- [x] Rewrite show-affected to use :: syntax for targets and workspaces
- [x] Task references workspace directly instead of WsName, WsPath, etc...
- [x] Only care about tasks, workspaces are just meta information and a part of inputs
- [x] Stop swallowing stderr or stdout in case of an error
- [x] Migrate to workspace::target convention
  - [x] evo run ::build
  - [x] evo run workspace::build
- [x] CI should use already published @evobuild/cli package
- [x] Fix GOCCM with empty queue
- [x] Profile long workspaces discovery
- [x] Put overrides to file system
- [x] Re-design CLI output
- [x] Fix error reporting
- [x] Cache diagnostics
  - [x] List of files that were used to create hash
  - [x] Diff to show what affected cache
- [x] Fix all issues highlighted by `--race` flag
- [x] Stop swallowing errors at cmd execution time
- [x] Extract evo config to a separate file to avoid cache misses for package managers
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
