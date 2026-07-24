---
description: Create new React components following IOP Toolkit's HeroUI + Tailwind patterns, including type-safe props, variant styling with tailwind-variants, and accessibility best practices.
applyTo:
  - "**/*.tsx"
  - "**/*.ts"
---

# Component Creation Skill

This skill guides you through creating new React components that follow IOP Toolkit's established patterns and best practices.

## When to Use This Skill

- Creating new reusable UI components
- Building feature-specific components for pages
- Implementing complex component variants with multiple states
- Setting up accessible, type-safe components

## Component Creation Checklist

### 1. Determine Component Location

```
src/
├── components/          # ← Shared, reusable components (navbar, icons, theme-switch)
└── pages/
    └── <feature>/       # ← Feature-specific components (e.g., omcianalyzer/LogTable.tsx)
```

**Decision Tree:**
- Used in 2+ pages? → `src/components/`
- Specific to one feature? → `src/pages/<feature>/`
- Is it a primitive/variant? → Add to `src/components/primitives.ts`

### 2. Component Template

```typescript
import { FC } from "react";
import { Button } from "@heroui/react";
import { tv } from "tailwind-variants";

// 1. Define TypeScript interface for props
interface MyComponentProps {
  title: string;
  description?: string;
  variant?: "default" | "highlighted";
  onAction?: () => void;
  isLoading?: boolean;
}

// 2. Create styled variants using tailwind-variants (if needed)
const componentStyles = tv({
  base: "rounded-lg border p-4",
  variants: {
    variant: {
      default: "border-gray-200 bg-white",
      highlighted: "border-blue-500 bg-blue-50",
    },
  },
  defaultVariants: {
    variant: "default",
  },
});

// 3. Component implementation
export const MyComponent: FC<MyComponentProps> = ({
  title,
  description,
  variant = "default",
  onAction,
  isLoading = false,
}) => {
  return (
    <div className={componentStyles({ variant })}>
      <h3 className="text-lg font-semibold text-gray-900">{title}</h3>
      {description && (
        <p className="mt-2 text-sm text-gray-600">{description}</p>
      )}
      {onAction && (
        <Button
          className="mt-4"
          color="primary"
          isLoading={isLoading}
          onPress={onAction}
        >
          Action
        </Button>
      )}
    </div>
  );
};
```

### 3. HeroUI Component Integration

**Always prefer HeroUI components for common UI elements:**

```typescript
import {
  Button,
  Input,
  Table,
  TableHeader,
  TableColumn,
  TableBody,
  TableRow,
  TableCell,
  Card,
  CardHeader,
  CardBody,
  CardFooter,
  Modal,
  ModalContent,
  ModalHeader,
  ModalBody,
  ModalFooter,
  Dropdown,
  DropdownTrigger,
  DropdownMenu,
  DropdownItem,
  Switch,
  Spinner,
  Pagination,
  Link,
  Code,
  Snippet,
} from "@heroui/react";
```

**HeroUI Component Patterns:**

```typescript
// Button variants
<Button color="primary" size="lg" variant="solid">Primary</Button>
<Button color="default" size="md" variant="bordered">Secondary</Button>
<Button color="danger" size="sm" variant="light">Delete</Button>
<Button isLoading>Processing...</Button>
<Button isDisabled>Disabled</Button>

// Input with validation
<Input
  label="Email"
  type="email"
  placeholder="Enter your email"
  isRequired
  errorMessage="Please enter a valid email"
  isInvalid={hasError}
/>

// Table with pagination
<Table aria-label="Data table">
  <TableHeader>
    <TableColumn>NAME</TableColumn>
    <TableColumn>STATUS</TableColumn>
  </TableHeader>
  <TableBody>
    <TableRow key="1">
      <TableCell>Item 1</TableCell>
      <TableCell>Active</TableCell>
    </TableRow>
  </TableBody>
</Table>

// Modal pattern
const { isOpen, onOpen, onClose } = useDisclosure();
<Button onPress={onOpen}>Open Modal</Button>
<Modal isOpen={isOpen} onClose={onClose}>
  <ModalContent>
    <ModalHeader>Modal Title</ModalHeader>
    <ModalBody>Content here</ModalBody>
    <ModalFooter>
      <Button onPress={onClose}>Close</Button>
    </ModalFooter>
  </ModalContent>
</Modal>
```

### 4. Tailwind Styling Guidelines

**Class Organization Pattern:**
```typescript
className="
  flex items-center justify-between     // Layout
  w-full max-w-4xl mx-auto             // Sizing & Spacing
  px-6 py-4 gap-4                      // Padding & Gap
  bg-white dark:bg-gray-800            // Colors
  border border-gray-200               // Borders
  rounded-lg shadow-sm                 // Effects
  hover:shadow-md transition-shadow    // Interactive States
"
```

**Responsive Design:**
```typescript
className="
  text-sm md:text-base lg:text-lg     // Responsive text
  grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3  // Responsive grid
  px-4 md:px-6 lg:px-8                // Responsive spacing
"
```

**Dark Mode Support:**
```typescript
className="
  bg-white dark:bg-gray-800
  text-gray-900 dark:text-gray-100
  border-gray-200 dark:border-gray-700
"
```

### 5. Component Variants with tailwind-variants

For components with multiple style variations:

```typescript
import { tv } from "tailwind-variants";

const button = tv({
  base: "inline-flex items-center justify-center rounded-md font-medium transition-colors",
  variants: {
    intent: {
      primary: "bg-blue-500 text-white hover:bg-blue-600",
      secondary: "bg-gray-500 text-white hover:bg-gray-600",
      danger: "bg-red-500 text-white hover:bg-red-600",
    },
    size: {
      sm: "px-3 py-1.5 text-sm",
      md: "px-4 py-2 text-base",
      lg: "px-6 py-3 text-lg",
    },
    fullWidth: {
      true: "w-full",
    },
  },
  compoundVariants: [
    {
      intent: "primary",
      size: "lg",
      class: "font-bold",
    },
  ],
  defaultVariants: {
    intent: "primary",
    size: "md",
  },
});

// Usage
<button className={button({ intent: "danger", size: "sm" })}>
  Delete
</button>
```

### 6. Handling State & Effects

```typescript
import { useState, useEffect } from "react";

export const DataComponent: FC = () => {
  const [data, setData] = useState<DataType[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        setIsLoading(true);
        const response = await axios.get<DataType[]>("/api/data");
        setData(response.data);
        setError(null);
      } catch (err) {
        setError(err instanceof Error ? err.message : "Failed to load data");
      } finally {
        setIsLoading(false);
      }
    };

    fetchData();
  }, []); // Empty dependency array for mount-only

  if (isLoading) {
    return <Spinner label="Loading..." />;
  }

  if (error) {
    return (
      <div className="text-center text-red-500">
        <p>Error: {error}</p>
      </div>
    );
  }

  return <div>{/* Render data */}</div>;
};
```

### 7. Custom Hooks for Complex Logic

Extract logic when a component exceeds 200 lines:

```typescript
// hooks/useTableData.ts
import { useState, useEffect } from "react";
import axios from "axios";

export const useTableData = <T>(endpoint: string) => {
  const [data, setData] = useState<T[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const fetchData = async () => {
    try {
      setIsLoading(true);
      const response = await axios.get<T[]>(endpoint);
      setData(response.data);
      setError(null);
    } catch (err) {
      setError(err instanceof Error ? err.message : "Failed to load");
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    fetchData();
  }, [endpoint]);

  return { data, isLoading, error, refetch: fetchData };
};

// Usage in component
const MyTable: FC = () => {
  const { data, isLoading, error, refetch } = useTableData<ItemType>("/api/items");
  
  // ... render logic
};
```

### 8. Accessibility Best Practices

```typescript
// ✅ GOOD: Semantic HTML + ARIA labels
<button
  className="..."
  onClick={handleClick}
  aria-label="Close dialog"
  aria-pressed={isActive}
>
  <CloseIcon />
</button>

// ✅ GOOD: Form accessibility
<Input
  label="Username"
  type="text"
  isRequired
  aria-describedby="username-help"
/>
<span id="username-help" className="text-sm text-gray-500">
  Username must be 3-20 characters
</span>

// ✅ GOOD: Focus management
<Modal isOpen={isOpen} onClose={onClose} autoFocus>
  <ModalContent>
    {/* Content */}
  </ModalContent>
</Modal>
```

### 9. Performance Optimization

```typescript
import { memo, useMemo, useCallback } from "react";

// Memoize expensive renders
export const ExpensiveComponent = memo<ExpensiveComponentProps>(
  ({ data, onAction }) => {
    // Memoize computed values
    const processedData = useMemo(() => {
      return data.map((item) => ({
        ...item,
        computed: expensiveCalculation(item),
      }));
    }, [data]);

    // Memoize callbacks to prevent child re-renders
    const handleAction = useCallback(() => {
      onAction();
    }, [onAction]);

    return <div>{/* Render */}</div>;
  }
);

ExpensiveComponent.displayName = "ExpensiveComponent";
```

### 10. Error Boundaries (if needed)

```typescript
// components/ErrorBoundary.tsx
import { Component, ReactNode } from "react";

interface Props {
  children: ReactNode;
  fallback?: ReactNode;
}

interface State {
  hasError: boolean;
  error?: Error;
}

export class ErrorBoundary extends Component<Props, State> {
  constructor(props: Props) {
    super(props);
    this.state = { hasError: false };
  }

  static getDerivedStateFromError(error: Error): State {
    return { hasError: true, error };
  }

  render() {
    if (this.state.hasError) {
      return (
        this.props.fallback || (
          <div className="p-4 text-center text-red-500">
            <h2>Something went wrong</h2>
            <p>{this.state.error?.message}</p>
          </div>
        )
      );
    }

    return this.props.children;
  }
}
```

## Anti-Patterns to Avoid

❌ **Avoid inline styles**
```typescript
// BAD
<div style={{ padding: "10px", color: "red" }}>Content</div>

// GOOD
<div className="p-2.5 text-red-500">Content</div>
```

❌ **Avoid component duplication**
```typescript
// BAD - Creating custom button when HeroUI has one
const MyButton = ({ onClick, children }) => (
  <button className="px-4 py-2 bg-blue-500" onClick={onClick}>
    {children}
  </button>
);

// GOOD - Use HeroUI Button
import { Button } from "@heroui/react";
<Button color="primary" onPress={onClick}>{children}</Button>
```

❌ **Avoid any type**
```typescript
// BAD
const processData = (data: any) => { ... };

// GOOD
interface DataItem {
  id: number;
  name: string;
}
const processData = (data: DataItem[]) => { ... };
```

❌ **Avoid magic numbers/strings**
```typescript
// BAD
if (status === 200) { ... }
setTimeout(() => {}, 3000);

// GOOD
const HTTP_OK = 200;
const DEBOUNCE_DELAY_MS = 3000;
if (status === HTTP_OK) { ... }
setTimeout(() => {}, DEBOUNCE_DELAY_MS);
```

## Component Testing Considerations

```typescript
// Example test structure (using Vitest)
import { describe, it, expect } from "vitest";
import { render, screen } from "@testing-library/react";
import { MyComponent } from "./MyComponent";

describe("MyComponent", () => {
  it("renders with required props", () => {
    render(<MyComponent title="Test" />);
    expect(screen.getByText("Test")).toBeInTheDocument();
  });

  it("calls onAction when button clicked", () => {
    const handleAction = vi.fn();
    render(<MyComponent title="Test" onAction={handleAction} />);
    
    const button = screen.getByRole("button");
    button.click();
    
    expect(handleAction).toHaveBeenCalledOnce();
  });
});
```

## Quick Reference

### Component File Template
```typescript
import { FC } from "react";
import { tv } from "tailwind-variants";

interface ComponentNameProps {
  // Props here
}

const styles = tv({
  // Variants here
});

export const ComponentName: FC<ComponentNameProps> = ({ }) => {
  return (
    <div className={styles()}>
      {/* JSX here */}
    </div>
  );
};
```

### Import Pattern
```typescript
// External
import { FC, useState } from "react";
import { Button, Input } from "@heroui/react";
import { tv } from "tailwind-variants";

// Internal (@/ alias)
import DefaultLayout from "@/layouts/default";
import { title } from "@/components/primitives";
import type { DataItem } from "@/types";

// Relative
import { LocalHelper } from "./helpers";
```

---

**Remember:** When in doubt, check existing components in `src/components/` or `src/pages/` for reference patterns.
