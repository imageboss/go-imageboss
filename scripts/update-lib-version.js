#!/usr/bin/env node
"use strict";

const fs = require("fs");
const path = require("path");

const pkgPath = path.join(__dirname, "..", "package.json");
const builderPath = path.join(__dirname, "..", "builder.go");

const pkg = JSON.parse(fs.readFileSync(pkgPath, "utf8"));
const version = pkg.version;
if (!version) {
  console.error("scripts/update-lib-version.js: no version in package.json");
  process.exit(1);
}

const libVersion = `go-v${version}`;
let builder = fs.readFileSync(builderPath, "utf8");

const regex = /LibVersion\s*=\s*"go-v[^"]*"/;
if (!regex.test(builder)) {
  console.error("scripts/update-lib-version.js: LibVersion constant not found in builder.go");
  process.exit(1);
}

builder = builder.replace(regex, `LibVersion = "${libVersion}"`);
fs.writeFileSync(builderPath, builder);
console.log(`Updated LibVersion to ${libVersion} in builder.go`);
