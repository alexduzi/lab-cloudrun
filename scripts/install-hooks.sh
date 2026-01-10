#!/bin/bash

# Install git hooks for the project
# Run this script after cloning the repository: ./scripts/install-hooks.sh

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"
HOOKS_DIR="$PROJECT_ROOT/.git/hooks"

echo "🔧 Installing git hooks..."
echo ""

# Check if we're in a git repository
if [ ! -d "$PROJECT_ROOT/.git" ]; then
    echo "❌ Error: Not a git repository. Please run this from the project root."
    exit 1
fi

# Create pre-push hook
cat > "$HOOKS_DIR/pre-push" << 'EOF'
#!/bin/bash

# Git pre-push hook
# Runs linter and tests before allowing push

echo "🔍 Running pre-push checks..."
echo ""

# Run linter
echo "📋 Running linter..."
if ! make lint; then
    echo ""
    echo "❌ Linter failed! Please fix the issues before pushing."
    exit 1
fi

echo "✅ Linter passed!"
echo ""

# Run tests
echo "🧪 Running tests..."
if ! make test; then
    echo ""
    echo "❌ Tests failed! Please fix the failing tests before pushing."
    exit 1
fi

echo "✅ Tests passed!"
echo ""
echo "✨ All pre-push checks passed! Proceeding with push..."
echo ""

exit 0
EOF

# Make the hook executable
chmod +x "$HOOKS_DIR/pre-push"

echo "✅ Git hooks installed successfully!"
echo ""
echo "Installed hooks:"
echo "  - pre-push: Runs linter and tests before push"
echo ""
echo "To skip hooks temporarily, use: git push --no-verify"
