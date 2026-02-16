#!/usr/bin/env node
"use strict";

const { execSync } = require("child_process");
const fs = require("fs");
const path = require("path");

const pkgPath = path.join(__dirname, "..", "package.json");
const pkg = JSON.parse(fs.readFileSync(pkgPath, "utf8"));
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
  // Tag does not exist, continue
}

console.log(`Creating and pushing tag ${tag}`);
execSync(`git tag ${tag}`, { stdio: "inherit" });
execSync(`git push origin ${tag}`, { stdio: "inherit" });
