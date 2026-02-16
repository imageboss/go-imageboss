# Changesets

This project uses [Changesets](https://github.com/changesets/changesets) to manage versioning and changelogs.

## Adding a changeset

When you make a change that should be released, run:

```bash
npx changeset
```

Then:

1. Select the version bump type (patch / minor / major).
2. Write a short summary of the change for the changelog.
3. Commit the new file under `.changeset/`.

## Releasing a new version

1. Merge all changesets into `main` (or your release branch).
2. Run:

   ```bash
   npx changeset version
   ```

   This will:

   - Bump the version in `package.json` (and any other configured packages).
   - Update `CHANGELOG.md` with the new changesets.
   - Remove the used changeset files.

3. For the Go module, align the version:

   - Update the `LibVersion` constant in `builder.go` to match the new version, then create a git tag: `git tag v1.0.1` and push: `git push origin v1.0.1`.

4. Commit the version and changelog updates, then push.

You can automate step 2–4 with a CI job (e.g. [changesets/action](https://github.com/changesets/action)).
