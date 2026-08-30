# Migrate from Mnemosyne to Mnemoteca

Mnemoteca replaces Mnemosyne as the active CLI and integration identity.

Use this guide before you remove old data or old integrations. The safe transfer
is JSONL export from the final Mnemosyne CLI, followed by one import into an
empty Mnemoteca destination.

## Safety rules

Do these steps for every migration path:

1. Stop all agent hosts that can use memory.
2. Create a new persistent empty export directory. Do not use a temporary
   directory that cleanup can remove.
3. Export with the final Mnemosyne CLI:
   ```bash
   mnemosyne export --all --yes --output ./mnemosyne-jsonl-export
   ```
   The export includes vectors by default. Do not add `--no-embeddings` for this
   migration.
4. Check that Mnemoteca points to a different, empty destination:
   ```bash
   mnemoteca stats --format json
   ```
   The JSON output includes `database_path`, `collection_count`, and
   `document_count`. `collection_count` and `document_count` must both be `0`.
5. Import once:
   ```bash
   mnemoteca import --dir ./mnemosyne-jsonl-export
   ```
6. Compare collection and document counts exactly.
7. Run a representative search and check that the expected memory appears:
   ```bash
   mnemoteca search --fts-only --no-rerank --limit 5 "your known memory text"
   ```
8. Keep the JSONL export directory after migration. It is your recovery copy.
9. Upgrade integrations.
10. Clean legacy files only after count and search verification succeed.

Do not copy or rename the SQLite database. Mnemoteca does not discover, copy, or
open legacy data automatically.

## What Mnemoteca does not do

The Mnemoteca runtime does not:

- discover legacy Mnemosyne databases;
- copy legacy databases;
- prompt for rename migration;
- fall back to `~/.config/mnemosyne` or `~/.local/share/mnemosyne`;
- read `MNEMOSYNE_*` environment variables;
- recognize old integration package names as active aliases.

Map custom settings explicitly to the new names:

| Old setting | New setting |
| --- | --- |
| `MNEMOSYNE_CONFIG` | `MNEMOTECA_CONFIG` or `MNEMOTECA_CONFIG_PATH` |
| `MNEMOSYNE_DB_PATH` | `MNEMOTECA_DB_PATH` |
| `~/.config/mnemosyne/config.yaml` | `~/.config/mnemoteca/config.yaml` |
| `~/.local/share/mnemosyne/mnemosyne.db` | `<Mnemoteca data dir>/mnemoteca.db` |

Default Mnemoteca data locations are:

- macOS/Linux: `$XDG_DATA_HOME/mnemoteca` or `~/.local/share/mnemoteca`;
- Windows: `%LOCALAPPDATA%\mnemoteca`, then `%APPDATA%\mnemoteca`, then
  `~/AppData/Local/mnemoteca`;
- database: `<data-dir>/mnemoteca.db`.

## Path 1: installer-guided macOS/Linux migration

Use this path on macOS or Linux when a final `mnemosyne` command is still
available on `PATH`.

```bash
curl -fsSL https://raw.githubusercontent.com/gandazgul/mnemoteca/main/install.sh | sh
```

The installer:

1. installs `mnemoteca` into `~/.local/bin` unless `INSTALL_DIR` is set;
2. looks for a qualifying `mnemosyne` command;
3. asks if all agents are stopped and if you want to migrate;
4. creates a retained export directory under `~/.mnemoteca-migration-exports`;
5. runs `mnemosyne export --all --yes --output <export-dir>`;
6. verifies that the export contains one JSONL file per collection and counts
   collection headers plus document lines;
7. reads the old CLI's reported database path;
8. runs `mnemoteca stats --format json` and requires an empty destination;
9. blocks migration if the old and new database paths resolve to the same file;
10. warns that stopped agents are required because the empty check and import are
    not transactionally locked together;
11. runs `mnemoteca import --dir <export-dir>` once;
12. compares imported collection and document counts exactly;
13. asks for a representative search query;
14. runs `mnemoteca search --fts-only --no-rerank --limit 5 <query>`;
15. asks you to confirm that the search result is correct;
16. offers optional legacy cleanup;
17. offers an optional POSIX `mnemosyne -> mnemoteca` symlink.

Prompts require a terminal. If no terminal is available, or if prompt input ends,
the installer still installs Mnemoteca but skips migration, cleanup, and
compatibility link creation. Rerun `install.sh` from a terminal to use the guided
migration.

If legacy paths exist but no qualifying final `mnemosyne` command is found, the
installer does not migrate or clean those paths. Install the final Mnemosyne
release and rerun the installer, or use the manual path.

## Path 2: manual macOS/Linux migration

Use this path when you want to control each command or when the installer cannot
qualify the old command.

1. Stop all agent hosts.
2. Install Mnemoteca:
   ```bash
   curl -fsSL https://raw.githubusercontent.com/gandazgul/mnemoteca/main/install.sh | sh
   ```
3. Create a persistent export directory:
   ```bash
   mkdir -p "$HOME/.mnemoteca-migration-exports/manual-mnemosyne-export"
   ```
4. Export from the final Mnemosyne CLI:
   ```bash
   mnemosyne export --all --yes --output "$HOME/.mnemoteca-migration-exports/manual-mnemosyne-export"
   ```
5. Record the old database path:
   ```bash
   mnemosyne stats
   ```
   Use the single `Database Path:` line. If this line is absent or appears more
   than once, stop and inspect the old installation before you continue.
6. Inspect the new destination:
   ```bash
   mnemoteca stats --format json
   ```
   Require `collection_count: 0` and `document_count: 0`. Compare the Mnemoteca
   `database_path` with the old `Database Path:`. If they are the same file or
   resolve through links to the same file, stop.
7. Import once:
   ```bash
   mnemoteca import --dir "$HOME/.mnemoteca-migration-exports/manual-mnemosyne-export"
   ```
8. Count the export. Each `.jsonl` file is one collection. The first line is the
   collection header. Remaining lines are documents. Empty collections still
   count as collections.
9. Run `mnemoteca stats --format json` again and compare collection and document
   counts exactly.
10. Run a representative search:
    ```bash
    mnemoteca search --fts-only --no-rerank --limit 5 "your known memory text"
    ```
11. Keep the JSONL export.
12. Upgrade integrations before you restart agents.
13. Clean legacy files only after verification succeeds.

## Path 3: manual Windows migration

Windows uses the release ZIP. There is no Windows compatibility shim.

1. Stop all agent hosts.
2. Download the Windows release artifact from the Mnemoteca GitHub release. The
   file name is `mnemoteca_VERSION_windows_amd64.zip`.
3. Extract it in PowerShell:
   ```powershell
   $Version = "v0.3.0"
   $Zip = "$env:TEMP\mnemoteca_$Version`_windows_amd64.zip"
   $InstallDir = "$env:USERPROFILE\bin"
   New-Item -ItemType Directory -Force $InstallDir | Out-Null
   Invoke-WebRequest `
     -Uri "https://github.com/gandazgul/mnemoteca/releases/download/$Version/mnemoteca_$Version`_windows_amd64.zip" `
     -OutFile $Zip
   Expand-Archive -Force $Zip $InstallDir
   ```
4. Add the directory that contains `mnemoteca.exe` to `PATH`, then open a new
   PowerShell session.
5. Verify and set up Mnemoteca:
   ```powershell
   mnemoteca.exe --help
   mnemoteca.exe setup
   ```
6. Create a persistent empty export directory:
   ```powershell
   $ExportDir = "$env:USERPROFILE\mnemoteca-migration\mnemosyne-jsonl-export"
   New-Item -ItemType Directory -Force $ExportDir | Out-Null
   Get-ChildItem $ExportDir -Force
   ```
   The directory must be empty.
7. Export from the final Mnemosyne CLI:
   ```powershell
   mnemosyne.exe export --all --yes --output $ExportDir
   ```
8. Record the old database path:
   ```powershell
   mnemosyne.exe stats
   ```
9. Confirm that Mnemoteca is empty and uses a different database path:
   ```powershell
   mnemoteca.exe stats --format json
   ```
   Require `collection_count` and `document_count` to be `0`. Compare the
   `database_path` with the old `Database Path:`.
10. Import once:
    ```powershell
    mnemoteca.exe import --dir $ExportDir
    ```
11. Count collections and documents in the export:
    ```powershell
    $Files = Get-ChildItem $ExportDir -Filter *.jsonl
    $Collections = $Files.Count
    $Documents = ($Files | ForEach-Object { (Get-Content $_.FullName).Count - 1 } | Measure-Object -Sum).Sum
    "$Collections collections, $Documents documents"
    ```
12. Run `mnemoteca.exe stats --format json` again and compare counts exactly.
13. Run a representative search:
    ```powershell
    mnemoteca.exe search --fts-only --no-rerank --limit 5 "your known memory text"
    ```
14. Keep the JSONL export.
15. Upgrade or remove old integrations before you restart agents.
16. Clean legacy files only after verification succeeds.

Do not create `mnemosyne.cmd`, a PowerShell alias, a copied `mnemosyne.exe`, or a
renamed executable that points to Mnemoteca. On Windows, update integrations so
they call `mnemoteca.exe` directly.

## Partial import and failure recovery

Import appends to the destination. It is not idempotent.

If import fails, is interrupted, or the final counts do not match:

- treat the Mnemoteca destination as partial;
- do not retry `mnemoteca import --dir` into that destination;
- do not clean legacy data;
- do not create or rely on compatibility behavior;
- keep the JSONL export;
- keep the legacy store;
- remove or otherwise resolve the partial Mnemoteca destination before a new
  import attempt.

A failed import does not change the legacy database or the retained export.

## Cleanup after successful verification

Only clean legacy files after these conditions are true:

- the export completed;
- Mnemoteca was empty before import;
- import ran once;
- collection and document counts match exactly;
- a representative search returns the expected memory;
- integrations have been upgraded or are stopped.

Inspect every path before deletion.

Canonical legacy paths on macOS/Linux:

- old binary when it is exactly `~/.local/bin/mnemosyne`;
- `~/.config/mnemosyne/config.yaml`;
- `~/.local/share/mnemosyne/mnemosyne.db`;
- `~/.local/share/mnemosyne/models`;
- `~/.local/share/mnemosyne/lib`.

Canonical legacy paths on Windows:

- old binary path reported by your installation records or `Get-Command
  mnemosyne.exe`;
- `%USERPROFILE%\.config\mnemosyne\config.yaml`, if you used this path;
- `%LOCALAPPDATA%\mnemosyne\mnemosyne.db`;
- `%APPDATA%\mnemosyne\mnemosyne.db`;
- `%USERPROFILE%\AppData\Local\mnemosyne\mnemosyne.db`.

Do not delete:

- a path you did not inspect;
- a custom old path automatically;
- the retained JSONL export;
- a Mnemoteca path;
- a broad parent directory such as `~/.local/share`, `%LOCALAPPDATA%`, or
  `%APPDATA%`.

## Integration replacement order

Use this order for each agent host:

1. Stop the agent.
2. Install the new Mnemoteca integration.
3. Verify that the integration exposes the same `memory_*` capabilities.
4. Remove or disable the old Mnemosyne integration at the same scope.
5. Restart the agent.

Host-specific steps live in the integration READMEs:

- [Claude Code Mnemoteca Skill](../../claudecode-mnemosyne/README.md)
- [OpenCode Mnemoteca Plugin](../../opencode-mnemosyne/README.md)
- [Pi Mnemoteca Extension](../../pi-mnemosyne/README.md)

The `memory_recall`, `memory_recall_global`, `memory_store`,
`memory_store_global`, and `memory_delete` capability names are unchanged.
