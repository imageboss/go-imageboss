# Changesets

This project uses [Changesets](https://github.com/changesets/changesets) to manage versioning and changelogs.

## Prerequisites

Install dependencies so the `changeset` CLI is available:

```bash
npm install
```

## Adding a changeset

When you make a change that should be released, run:

```bash
npm run changeset
# or: npx changeset add
```

Then:

1. Select the version bump type (patch / minor / major).
2. Write a short summary of the change for the changelog.
3. Commit the new file under `.changeset/`.

If the change doesn’t need a release (e.g. docs-only), run:

```bash
npx changeset add --empty
```

## Releasing a new version

1. Merge all changesets into `main` (or your release branch).
2. Run:

   ```bash
   npm run changeset:version
   # or: CI=1 npx changeset version
   ```

   Setting `CI=1` avoids TTY-related warnings in CI or non-interactive shells. This will:

   - Bump the version in `package.json`.
   - Update `CHANGELOG.md` with the new changesets.
   - Remove the used changeset files.
   - Update the `LibVersion` constant in `builder.go` to match (via `scripts/update-lib-version.js`).

3. For the Go module, tag and push:

   - Create a git tag: `git tag v1.0.1` and push: `git push origin v1.0.1`.

4. Commit the version and changelog updates, then push.

## Checking status

To see whether there are unreleased changesets or unrecorded changes:

```bash
npm run changeset:status
# or: npx changeset status
```

If you see “Some packages have been changed but no changesets were found”, add a changeset with `npm run changeset` or run `npx changeset add --empty` if you don’t intend to release yet.

## CI

You can automate versioning and tagging with a CI job (e.g. [changesets/action](https://github.com/changesets/action)). Use `CI=1` (or run `npm run changeset:version`) when calling `changeset version` in CI.
