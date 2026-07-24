# IOP Toolkit - Copilot Harness Documentation

This directory contains documentation and skills to enhance GitHub Copilot's understanding of the IOP Toolkit project.

## 📚 Documentation Structure

### Main Harness Document
- **[copilot-instructions.md](../.github/copilot-instructions.md)** - Main instruction file for GitHub Copilot
  - Context Engineering: Technology stack, project structure, and design philosophy
  - Architectural Constraints: Code organization, naming conventions, and performance rules
  - Entropy Management: Strategies to prevent code degradation over time

### Skills Directory

Located in `.skills/`, these documents provide specialized guidance for common development tasks:

1. **[component-creation.md](./component-creation.md)**
   - Creating React components with HeroUI + Tailwind CSS
   - Component styling patterns using tailwind-variants
   - TypeScript prop interfaces and type safety
   - Custom hooks for complex logic
   - Accessibility best practices

2. **[page-creation.md](./page-creation.md)**
   - Creating new pages with proper routing
   - Integrating with React Router DOM
   - Adding navigation links to site config
   - URL parameters and query strings
   - Sub-page navigation patterns

3. **[api-integration.md](./api-integration.md)**
   - Backend API integration patterns
   - Axios request/response handling
   - Custom hooks for data fetching (useFetch, useMutation)
   - Error handling and retry logic
   - File upload/download patterns

4. **[typescript-patterns.md](./typescript-patterns.md)**
   - TypeScript strict mode best practices
   - Component prop types and interfaces
   - API response type definitions
   - Generic types and utility types
   - Type guards and discriminated unions

## 🚀 How Copilot Uses This

GitHub Copilot automatically reads these files to understand:
- Project conventions and patterns
- Technology stack configuration
- Code organization principles
- Best practices and anti-patterns

When you ask Copilot to create a component, page, or API integration, it will follow the patterns defined in these documents.

## 📖 Quick Reference

### Technology Stack Summary

| Category | Technology | Version |
|----------|-----------|---------|
| Framework | React | 18.3 |
| Language | TypeScript | 5.6 (strict mode) |
| UI Library | HeroUI | 2.8 |
| Styling | Tailwind CSS | 4.1 |
| Build Tool | Vite | 6.4 |
| Routing | React Router DOM | 6.23 |
| HTTP Client | Axios | 1.12 |
| Animation | Framer Motion | 11.18 |
| Data Viz | Recharts, Mermaid | Latest |

### Project Structure

```
webioptoolkit/
├── .github/
│   └── copilot-instructions.md  # Main harness document
├── .skills/             # Copilot skill documents
├── src/
│   ├── components/      # Shared UI components
│   ├── config/          # App configuration
│   ├── layouts/         # Page layouts
│   ├── pages/           # Route pages (feature-organized)
│   ├── styles/          # Global styles
│   └── types/           # TypeScript type definitions
└── ...
```

### Key Conventions

- **Import Paths**: Use `@/` alias for src imports
- **Component Style**: Functional components with TypeScript
- **Styling**: Tailwind CSS classes + tailwind-variants for variants
- **Type Safety**: No `any` type, strict TypeScript enabled
- **File Naming**: 
  - Components: PascalCase (e.g., `NavBar.tsx`)
  - Pages: lowercase folders, PascalCase files
  - Utilities: camelCase (e.g., `useTheme.ts`)

## 🛠️ Development Workflow

### Creating a New Component

1. Reference: [component-creation.md](.skills/component-creation.md)
2. Determine location: `src/components/` (shared) or `src/pages/<feature>/` (feature-specific)
3. Define TypeScript interface for props
4. Use HeroUI components when available
5. Apply Tailwind CSS for styling
6. Use tailwind-variants for complex variants

### Creating a New Page

1. Reference: [page-creation.md](.skills/page-creation.md)
2. Create page component in `src/pages/<feature>/`
3. Add route in `src/App.tsx`
4. Update navigation in `src/config/site.ts`
5. Wrap in `DefaultLayout` for consistency

### Integrating APIs

1. Reference: [api-integration.md](.skills/api-integration.md)
2. Define TypeScript types for request/response
3. Use axios with proper error handling
4. Implement loading and error states
5. Consider creating custom hooks for reusable logic

### Type Safety

1. Reference: [typescript-patterns.md](.skills/typescript-patterns.md)
2. Define explicit interfaces for all props
3. Type API responses properly
4. Use discriminated unions for variant data
5. Leverage TypeScript utility types

## 📝 Maintaining Harness Documents

### When to Update

Update harness documents when:
- Adding new major dependencies
- Establishing new coding patterns
- Changing architectural decisions
- Discovering common mistakes to avoid

### Document Versioning

Current version: **1.0**  
Last updated: **2026-03-17**

Update the version and date in `.github/copilot-instructions.md` when making significant changes.

## 🎯 Benefits

These harness documents provide:

1. **Consistent Code Generation**: Copilot follows established patterns
2. **Reduced Errors**: Architectural constraints prevent common mistakes
3. **Faster Onboarding**: New developers (and Copilot) understand conventions quickly
4. **Better Maintainability**: Entropy management prevents code degradation
5. **Type Safety**: Strict TypeScript patterns reduce runtime errors

## 🔍 Examples

### Ask Copilot to:

- "Create a new component for displaying OMCI data in a table"
  - Copilot will use HeroUI Table, proper TypeScript types, and Tailwind styling

- "Add a new page for network diagnostics"
  - Copilot will create the page, add routing, and update navigation

- "Implement API call to fetch user settings"
  - Copilot will use axios, proper types, error handling, and loading states

- "Add a custom hook for managing form state"
  - Copilot will follow TypeScript patterns and React hooks best practices

## 📚 Additional Resources

- [HeroUI Documentation](https://heroui.com)
- [Tailwind CSS Documentation](https://tailwindcss.com)
- [React TypeScript Cheatsheet](https://react-typescript-cheatsheet.netlify.app)
- [Vite Documentation](https://vitejs.dev)

---

**Note**: These harness documents are living documentation. Update them as the project evolves to keep Copilot aligned with current best practices.
