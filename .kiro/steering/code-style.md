# Code Style

## Fundamental Rules

### NEVER add comments to code
- Code should be self-explanatory through clear names
- Do not add comments to functions, structs, or variables
- Do not use documentation comments (// or /* */)
- Clean code eliminates the need for additional explanations

### Clean Code Principles
- Variable and function names must be descriptive
- Small functions with single responsibility
- Avoid complex logic in a single line
- Prefer clarity over brevity

### Formatting
- Use `go fmt` for automatic formatting
- Use `goimports` for import organization
- Keep lines under 120 characters when possible
- Consistent spacing between logical blocks

### Data Structures
- Structs must have clear and descriptive names
- JSON tags always in snake_case
- Validations using validator tags
- Optional fields with `omitempty`

### Error Handling
- Always check and handle errors
- Use Go's standard error return pattern
- Structured logging with Zap for errors
- Clear error messages for users