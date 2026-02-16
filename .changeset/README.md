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

1. Merge all changesets into `main` (or your release branch) and push.
2. **With CI (default):** The [release workflow](.github/workflows/release.yml) runs on push to `main`. If there are unreleased changesets, it runs the version bump, commits and pushes, and creates and pushes the tag. You don’t need to run anything else.
3. **Without CI:** Run `npm run release` (it versions, commits, pushes to `main`, and creates/pushes the tag). Or do it step by step: `npm run changeset:version`, then commit and push the changed files, then `npm run tag:push`.

## Checking status

To see whether there are unreleased changesets or unrecorded changes:

```bash
npm run changeset:status
# or: npx changeset status
```

If you see “Some packages have been changed but no changesets were found”, add a changeset with `npm run changeset` or run `npx changeset add --empty` if you don’t intend to release yet.

