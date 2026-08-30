# Active Identity Allowlist

This file records the old product and package identities that are still valid in
maintained documentation after the Mnemoteca rename.

Do not use this file as a broad exemption. Each entry applies only to the named
file and section or literal. Any new match must be reviewed before it is kept.

## Reviewed literals in this file

- `Mnemosyne`, `mnemosyne`, `MNEMOSYNE_`, `opencode-mnemosyne`,
  `pi-mnemosyne`, `claudecode-mnemosyne`, and
  `https://github.com/gandazgul/mnemosyne` appear in this allowlist so reviewers
  know which old identities the scan must catch. These literals are not active
  commands, package recommendations, or repository links.

## `cli/README.md`

- Section `Install the CLI`: `Mnemosyne` appears only to explain that the POSIX
  installer can detect a final old command and offer migration.
- Section `Migrating from Mnemosyne`: the heading and link text name the source
  product so users can find the transition guide.
- Section `Connect Mnemoteca to your agent`: `upgrade-from-Mnemosyne` appears
  only to describe integration replacement documentation.
- Section `Configuration`: `Mnemosyne` appears only to state that custom old
  settings must be mapped to new Mnemoteca settings explicitly.
- Section `Benchmarks`: `Mnemosyne` appears only for the dated pre-rename
  LongMemEval results. Current reproduction commands and configuration use
  Mnemoteca.

## `cli/docs/migrate-from-mnemosyne.md`

- Whole file purpose: this is the authoritative source-product migration guide.
  Old product names, old commands, old paths, and old environment prefixes are
  permitted only when they identify the source installation, export command,
  cleanup candidate, or forbidden compatibility behavior.
- Sections `Safety rules`, `Path 2`, and `Path 3`: `mnemosyne export --all --yes
  --output`, `mnemosyne stats`, and `mnemosyne.exe` commands are the required
  source export and source inspection commands.
- Section `What Mnemoteca does not do`: `~/.config/mnemosyne`,
  `~/.local/share/mnemosyne`, `MNEMOSYNE_CONFIG`, `MNEMOSYNE_DB_PATH`, and
  `MNEMOSYNE_*` are old settings that users must replace. They are not accepted
  as new Mnemoteca settings.
- Section `Path 1`: `mnemosyne -> mnemoteca` is the optional POSIX compatibility
  symlink implemented by `cli/install.sh` after verified migration only.
- Section `Path 3`: `mnemosyne.cmd`, `mnemosyne.exe`, and alias text appear only
  to prohibit a Windows compatibility shim.
- Section `Cleanup after successful verification`: old POSIX and Windows
  `mnemosyne` paths are cleanup candidates that require inspection after
  successful verification.
- Section `Integration replacement order`: old integration identity appears only
  as the item that users remove or disable after the new Mnemoteca integration is
  verified.
- Section `Integration replacement order`: repository folder names in relative
  links, including `claudecode-mnemosyne`, `opencode-mnemosyne`, and
  `pi-mnemosyne`, are current local checkout paths. The link labels are
  Mnemoteca identities.

## `cli/CHANGELOG.md`

- Section `Unreleased`: `Mnemosyne` appears only as the source product for the
  rename migration.
- Section `v0.2.3 (2026-06-23)`: old config paths and
  `https://github.com/gandazgul/mnemosyne` commit links are historical release
  evidence and must not be rewritten as Mnemoteca history.

## `cli/ROADMAP.md`

- Section `Completed`: `mnemosyne init` appears only to describe behavior that
  shipped before the rename and to state that the same behavior remains in
  `mnemoteca init`.

## `cli/benchmarks/README.md`

- Sections `Headline Results`, `Comparison Notes`, and related result prose:
  `Mnemosyne` appears only to attribute dated pre-rename benchmark results to the
  product that produced them.
- Generated result files under `cli/benchmarks/results/` are not maintained
  active documentation for this scan. Their old names and absolute paths are
  immutable historical artifacts.

## `opencode-mnemosyne/README.md`

- Section `Upgrade from opencode-mnemosyne`: `opencode-mnemosyne` is the old
  plugin entry that users remove after `opencode-mnemoteca` is verified.
- Section `Upgrade from opencode-mnemosyne`: `mnemosyne` appears only in the
  old plugin identity and the Windows no-shim warning.

## `pi-mnemosyne/README.md`

- Section `Upgrade from pi-mnemosyne`: `pi-mnemosyne` is the old package that
  users remove after `pi-mnemoteca` is verified.
- Section `Upgrade from pi-mnemosyne`: `mnemosyne` appears only in the old
  package name and the Windows no-shim warning.

## `claudecode-mnemosyne/README.md`

- Section `Upgrade from claudecode-mnemosyne`: `claudecode-mnemosyne` is the old
  skill identity that users replace after `mnemoteca` is verified.
- Section `Upgrade from claudecode-mnemosyne`: `~/.claude/skills/mnemosyne` and
  the equivalent PowerShell path are old skill directories that users may remove
  after verifying the new skill.
- Section `Upgrade from claudecode-mnemosyne`: `mnemosyne` appears only in the
  old skill directory paths and the Windows no-shim warning.

## `claudecode-mnemosyne/skills/mnemoteca/SKILL.md`

- No old product, command, environment-prefix, package, or URL identity is
  allowed in this installed user-facing skill document.
