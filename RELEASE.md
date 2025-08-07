# Release Process

This document describes the automated release process for `randpass`.

## GitHub Actions Workflows

### CI/CD Pipeline (`ci.yml`)

Automatically runs on every push and pull request:

1. **Test Job**: Runs all tests with coverage reporting
2. **Lint Job**: Runs golangci-lint for code quality
3. **Build Job**: Cross-compiles binaries for multiple platforms
4. **Release Job**: Uploads binaries when a release is created
5. **Benchmark Job**: Runs performance benchmarks on main branch

### Manual Release (`release.yml`)

Creates a new release with binaries for all supported platforms:

- **Trigger**: Manual workflow dispatch
- **Inputs**: Version number and pre-release flag
- **Outputs**: Release with binaries and checksums

## Supported Platforms

| Platform | Architecture | Binary Name |
|----------|-------------|-------------|
| Linux | amd64 | `randpass-linux-amd64` |
| Linux | arm64 | `randpass-linux-arm64` |
| macOS | amd64 | `randpass-macos-amd64` |
| macOS | arm64 | `randpass-macos-arm64` |
| Windows | amd64 | `randpass-windows-amd64.exe` |

## Local Release Testing

Build all release binaries locally:

```bash
make build-release
```

Generate checksums:

```bash
make release-checksums
```

## Creating a Release

### Option 1: GitHub UI (Recommended)

1. Go to Actions tab in GitHub
2. Select "Release" workflow
3. Click "Run workflow"
4. Enter version (e.g., `v1.0.0`)
5. Choose if it's a pre-release
6. Click "Run workflow"

### Option 2: Git Tags

```bash
# Create and push a tag
git tag v1.0.0
git push origin v1.0.0

# Then manually create release in GitHub UI
```

## Release Artifacts

Each release includes:

- **Binaries**: Cross-compiled for all supported platforms
- **Archives**: `.tar.gz` for Unix, `.zip` for Windows
- **Checksums**: SHA256 hashes for verification
- **Release Notes**: Automated with download instructions

## Verification

Users can verify downloads using checksums:

```bash
# Download binary and checksums
wget https://github.com/OlegChuev/randpass/releases/download/v1.0.0/randpass-linux-amd64.tar.gz
wget https://github.com/OlegChuev/randpass/releases/download/v1.0.0/checksums.txt

# Verify
sha256sum -c checksums.txt
```

## Continuous Integration

- **Tests**: Run on every commit
- **Benchmarks**: Track performance over time
- **Coverage**: Uploaded to Codecov
- **Linting**: Ensures code quality

## Security

- **Static Analysis**: golangci-lint with security rules
- **Dependency Scanning**: Automatic vulnerability detection
- **Signed Releases**: Checksums for integrity verification
- **Minimal Permissions**: GitHub Actions use least privilege
