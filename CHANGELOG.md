# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Composite response codes (`BuildResponseCode`, `ParseResponseCode`) with service and case constants.
- Gin writers for success, error, validation (HTTP 422), and cursor/offset pagination envelopes.
- `FormatValidationError` for Laravel-style field maps from `validator.ValidationErrors`.
- Package godoc, Example tests, README, CONTRIBUTING, and llms.txt.
