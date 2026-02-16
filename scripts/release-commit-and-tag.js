#!/usr/bin/env node
"use strict";

const { execSync } = require("child_process");
const fs = require("fs");
const path = require("path");

// Check if there are version-bump changes to commit
const status = execSync("git status --porcelain package.json builder.go CHANGELOG.md .changeset/", {
  encoding: "utf8",
});
if (!status.trim()) {
  console.log("No version bump (no changesets consumed). Nothing to commit or tag.");
  process.exit(0);
}

const pkg = JSON.parse(fs.readFileSync(path.join(__dirname, "..", "package.json"), "utf8"));
const version = pkg.version;
if (!version) {
  console.error("Could not read version from package.json");
  process.exit(1);
}

const tag = `v${version}`;
try {
  execSync(`git rev-parse ${tag}`, { stdio: "pipe" });
  console.error(`Tag ${tag} already exists.`);
  process.exit(1);
} catch {
  /* tag does not exist, continue */
}

console.log(`Committing release ${tag}, then tagging and pushing...`);
execSync("git add package.json builder.go CHANGELOG.md .changeset/", { stdio: "inherit" });
execSync(`git commit -m "chore: release ${tag}"`, { stdio: "inherit" });
execSync("git push origin main", { stdio: "inherit" });
execSync(`git tag ${tag}`, { stdio: "inherit" });
execSync(`git push origin ${tag}`, { stdio: "inherit" });
console.log(`Done. Tag ${tag} pushed.`);