---
description: TypeScript type safety patterns, utility types, and best practices for IOP Toolkit strict mode development.
applyTo:
  - "**/*.ts"
  - "**/*.tsx"
  - "src/types/**"
---

# TypeScript Type Safety Skill

This skill provides patterns for maintaining strict type safety, defining robust types, and leveraging TypeScript's advanced features in IOP Toolkit.

## When to Use This Skill

- Defining component prop types
- Creating API response interfaces
- Working with complex data structures
- Type-safe event handlers and callbacks
- Avoiding `any` type and improving type inference

## Project TypeScript Configuration

**Current tsconfig.json settings:**
```json
{
  "compilerOptions": {
    "strict": true,                      // All strict checks enabled
    "noUnusedLocals": true,              // Error on unused local variables
    "noUnusedParameters": true,          // Error on unused parameters
    "noFallthroughCasesInSwitch": true,  // Error on switch fallthrough
    "target": "ES2020",
    "lib": ["ES2020", "DOM", "DOM.Iterable"],
    "jsx": "react-jsx",
    "paths": {
      "@/*": ["./src/*"]                 // Path alias for imports
    }
  }
}
```

This means:
- ❌ **No `any` type** - use proper types or `unknown`
- ✅ **All variables must be typed** explicitly or by inference
- ✅ **Null checks required** (`strictNullChecks`)
- ✅ **Function parameters and returns must be typed**

## Type Definition Patterns

### 1. Component Props

```typescript
// ✅ GOOD: Explicit interface
interface ButtonProps {
  label: string;
  onClick: () => void;
  variant?: "primary" | "secondary";
  isDisabled?: boolean;
  size?: "sm" | "md" | "lg";
  icon?: React.ReactNode;
}

export const Button: FC<ButtonProps> = ({
  label,
  onClick,
  variant = "primary",
  isDisabled = false,
  size = "md",
  icon,
}) => {
  return (
    <button onClick={onClick} disabled={isDisabled}>
      {icon && <span className="mr-2">{icon}</span>}
      {label}
    </button>
  );
};

// ❌ AVOID: Using 'any' or no types
const BadButton = ({ label, onClick }: any) => { ... };
```

### 2. API Response Types

```typescript
// src/types/api.ts

// Base response wrapper
export interface ApiResponse<T> {
  data: T;
  status: number;
  message?: string;
  timestamp: string;
}

// Paginated response
export interface PaginatedResponse<T> {
  items: T[];
  total: number;
  page: number;
  pageSize: number;
  hasMore: boolean;
}

// Error response
export interface ApiError {
  error: string;
  code: string;
  details?: Record<string, string[]>;
  timestamp: string;
}

// Domain-specific types
export interface OmciOnu {
  id: string;
  serialNumber: string;
  status: "online" | "offline" | "unknown";
  firmwareVersion: string;
  lastSeen: string;
  meData?: Record<string, unknown>;
}

export interface LogEntry {
  id: number;
  timestamp: string;
  level: "debug" | "info" | "warn" | "error";
  message: string;
  metadata?: Record<string, unknown>;
}

// Usage in components
const fetchOnus = async (): Promise<ApiResponse<OmciOnu[]>> => {
  const response = await axios.get<ApiResponse<OmciOnu[]>>("/api/omci/onus");
  return response.data;
};
```

### 3. Discriminated Unions

For data that can have different shapes based on a type field:

```typescript
// ✅ GOOD: Discriminated union
type DataState =
  | { status: "idle" }
  | { status: "loading" }
  | { status: "success"; data: DataItem[] }
  | { status: "error"; error: string };

const DataComponent = () => {
  const [state, setState] = useState<DataState>({ status: "idle" });

  // TypeScript knows the shape based on status
  if (state.status === "loading") {
    return <Spinner />;
  }

  if (state.status === "error") {
    // TypeScript knows 'error' property exists
    return <div>Error: {state.error}</div>;
  }

  if (state.status === "success") {
    // TypeScript knows 'data' property exists
    return <div>{state.data.map(...)}</div>;
  }

  return <div>Click to load</div>;
};

// Another example: Different chart types
type ChartConfig =
  | { type: "bar"; values: number[]; labels: string[] }
  | { type: "line"; points: Array<{ x: number; y: number }> }
  | { type: "pie"; segments: Array<{ label: string; value: number; color: string }> };

const renderChart = (config: ChartConfig) => {
  switch (config.type) {
    case "bar":
      // TypeScript knows config has values and labels
      return <BarChart data={config.values} labels={config.labels} />;
    case "line":
      // TypeScript knows config has points
      return <LineChart points={config.points} />;
    case "pie":
      // TypeScript knows config has segments
      return <PieChart segments={config.segments} />;
  }
};
```

### 4. Utility Types

```typescript
// Built-in TypeScript utility types

// Partial - makes all properties optional
interface User {
  id: number;
  name: string;
  email: string;
  role: string;
}

const updateUser = (id: number, updates: Partial<User>) => {
  // updates can have any subset of User properties
};

updateUser(1, { email: "new@email.com" }); // ✅ Valid

// Required - makes all properties required
type RequiredUser = Required<Partial<User>>; // Back to all required

// Pick - select specific properties
type UserCredentials = Pick<User, "email" | "password">;
// { email: string; password: string }

// Omit - exclude specific properties
type UserWithoutId = Omit<User, "id">;
// { name: string; email: string; role: string }

// Record - create object type with specific keys
type StatusMap = Record<string, "active" | "inactive">;
// { [key: string]: "active" | "inactive" }

// Example usage
const statusMap: StatusMap = {
  user1: "active",
  user2: "inactive",
};

// Readonly - make all properties readonly
const config: Readonly<{ apiUrl: string; timeout: number }> = {
  apiUrl: "https://api.example.com",
  timeout: 5000,
};
// config.apiUrl = "..."; // ❌ Error: readonly

// ReturnType - extract return type of function
const fetchData = async () => {
  return { data: [], total: 0 };
};
type FetchDataReturn = ReturnType<typeof fetchData>;
// Promise<{ data: any[]; total: number }>

// Parameters - extract parameter types
const myFunc = (name: string, age: number) => {};
type MyFuncParams = Parameters<typeof myFunc>;
// [string, number]
```

### 5. Generic Types

```typescript
// Generic component props
interface ListProps<T> {
  items: T[];
  renderItem: (item: T) => React.ReactNode;
  keyExtractor: (item: T) => string | number;
  emptyMessage?: string;
}

export const List = <T,>({
  items,
  renderItem,
  keyExtractor,
  emptyMessage = "No items",
}: ListProps<T>) => {
  if (items.length === 0) {
    return <div>{emptyMessage}</div>;
  }

  return (
    <div>
      {items.map((item) => (
        <div key={keyExtractor(item)}>{renderItem(item)}</div>
      ))}
    </div>
  );
};

// Usage with type inference
<List
  items={users}  // TypeScript infers T = User
  renderItem={(user) => <div>{user.name}</div>}
  keyExtractor={(user) => user.id}
/>

// Generic API function
const fetchItems = async <T,>(endpoint: string): Promise<T[]> => {
  const response = await axios.get<T[]>(endpoint);
  return response.data;
};

// Usage
const users = await fetchItems<User>("/api/users");
const products = await fetchItems<Product>("/api/products");
```

### 6. Type Guards

```typescript
// Type guard functions for runtime type checking

// Primitive type guard
const isString = (value: unknown): value is string => {
  return typeof value === "string";
};

// Object type guard
interface ApiSuccess {
  data: unknown;
  status: number;
}

interface ApiError {
  error: string;
  code: string;
}

const isApiError = (response: ApiSuccess | ApiError): response is ApiError => {
  return "error" in response;
};

// Usage
const handleResponse = (response: ApiSuccess | ApiError) => {
  if (isApiError(response)) {
    // TypeScript knows response is ApiError
    console.error(response.error);
  } else {
    // TypeScript knows response is ApiSuccess
    console.log(response.data);
  }
};

// Array type guard
const isStringArray = (value: unknown): value is string[] => {
  return Array.isArray(value) && value.every((item) => typeof item === "string");
};

// Null/undefined check
const isDefined = <T,>(value: T | null | undefined): value is T => {
  return value !== null && value !== undefined;
};

// Usage with array filtering
const values: (string | null)[] = ["a", null, "b", undefined, "c"];
const definedValues = values.filter(isDefined); // string[]
```

### 7. Event Handler Types

```typescript
// Form events
const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
  e.preventDefault();
  // ... form logic
};

const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
  const value = e.target.value;
  // ... handle change
};

const handleSelectChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
  const selectedValue = e.target.value;
};

const handleTextareaChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
  const text = e.target.value;
};

// Mouse events
const handleClick = (e: React.MouseEvent<HTMLButtonElement>) => {
  e.stopPropagation();
  // ... handle click
};

const handleDivClick = (e: React.MouseEvent<HTMLDivElement>) => {
  const { clientX, clientY } = e;
};

// Keyboard events
const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
  if (e.key === "Enter") {
    // ... submit form
  }
};

// Generic event handler
const handleEvent = <T extends HTMLElement>(
  e: React.SyntheticEvent<T>
) => {
  // ... handle any event
};

// Callback types
type OnChange = (value: string) => void;
type OnSubmit = (data: FormData) => Promise<void>;
type OnSelect<T> = (item: T) => void;

interface FormProps {
  onSubmit: OnSubmit;
  onChange?: OnChange;
}
```

### 8. State Types

```typescript
// Simple state
const [count, setCount] = useState<number>(0);
const [name, setName] = useState<string>("");
const [isOpen, setIsOpen] = useState<boolean>(false);

// Object state
interface FormState {
  email: string;
  password: string;
  remember: boolean;
}

const [formState, setFormState] = useState<FormState>({
  email: "",
  password: "",
  remember: false,
});

// Update object state
setFormState((prev) => ({ ...prev, email: "new@email.com" }));

// Array state
const [items, setItems] = useState<DataItem[]>([]);

// Add item
setItems((prev) => [...prev, newItem]);

// Remove item
setItems((prev) => prev.filter((item) => item.id !== removeId));

// Update item
setItems((prev) =>
  prev.map((item) => (item.id === updateId ? { ...item, ...updates } : item))
);

// Nullable state
const [selectedItem, setSelectedItem] = useState<DataItem | null>(null);

// Union state
const [status, setStatus] = useState<"idle" | "loading" | "success" | "error">("idle");

// Complex state with discriminated union
type LoadState<T> =
  | { type: "idle" }
  | { type: "loading" }
  | { type: "success"; data: T }
  | { type: "error"; error: string };

const [dataState, setDataState] = useState<LoadState<DataItem[]>>({
  type: "idle",
});

// Set loading
setDataState({ type: "loading" });

// Set success
setDataState({ type: "success", data: fetchedData });

// Set error
setDataState({ type: "error", error: "Failed to load" });
```

### 9. Ref Types

```typescript
// Element refs
const inputRef = useRef<HTMLInputElement>(null);
const divRef = useRef<HTMLDivElement>(null);

useEffect(() => {
  inputRef.current?.focus(); // ✅ Type-safe with optional chaining
}, []);

// Generic ref
const buttonRef = useRef<HTMLButtonElement>(null);

// Value refs (not DOM elements)
const timerRef = useRef<NodeJS.Timeout | null>(null);
const countRef = useRef<number>(0);

useEffect(() => {
  timerRef.current = setTimeout(() => {
    console.log("Timer fired");
  }, 1000);

  return () => {
    if (timerRef.current) {
      clearTimeout(timerRef.current);
    }
  };
}, []);
```

### 10. Context Types

```typescript
// Define context value type
interface ThemeContextValue {
  theme: "light" | "dark";
  toggleTheme: () => void;
}

// Create context with default value
const ThemeContext = createContext<ThemeContextValue | null>(null);

// Provider component
export const ThemeProvider: FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  const [theme, setTheme] = useState<"light" | "dark">("light");

  const toggleTheme = () => {
    setTheme((prev) => (prev === "light" ? "dark" : "light"));
  };

  const value: ThemeContextValue = { theme, toggleTheme };

  return <ThemeContext.Provider value={value}>{children}</ThemeContext.Provider>;
};

// Custom hook with null check
export const useTheme = (): ThemeContextValue => {
  const context = useContext(ThemeContext);
  
  if (!context) {
    throw new Error("useTheme must be used within ThemeProvider");
  }
  
  return context;
};

// Usage
const MyComponent = () => {
  const { theme, toggleTheme } = useTheme();
  // ... use theme
};
```

## Advanced Patterns

### 1. Conditional Types

```typescript
// Extract array element type
type ArrayElement<T> = T extends (infer U)[] ? U : never;

type NumArray = ArrayElement<number[]>; // number
type StrArray = ArrayElement<string[]>; // string

// Extract promise result type
type Awaited<T> = T extends Promise<infer U> ? U : T;

type PromiseResult = Awaited<Promise<string>>; // string
type NotPromise = Awaited<number>; // number

// Conditional prop types
interface BaseProps {
  variant: "simple" | "complex";
}

type ConditionalProps<T extends BaseProps> = T["variant"] extends "simple"
  ? { data: string }
  : { data: ComplexData; config: ComplexConfig };
```

### 2. Mapped Types

```typescript
// Make all properties optional recursively
type DeepPartial<T> = {
  [P in keyof T]?: T[P] extends object ? DeepPartial<T[P]> : T[P];
};

interface NestedConfig {
  database: {
    host: string;
    port: number;
    credentials: {
      username: string;
      password: string;
    };
  };
}

const partialConfig: DeepPartial<NestedConfig> = {
  database: {
    credentials: {
      username: "admin", // Can omit password
    },
  },
};

// Make properties mutable
type Mutable<T> = {
  -readonly [P in keyof T]: T[P];
};

// Make properties readonly
type DeepReadonly<T> = {
  readonly [P in keyof T]: T[P] extends object ? DeepReadonly<T[P]> : T[P];
};
```

### 3. Template Literal Types

```typescript
// Event names
type EventName = "click" | "focus" | "blur";
type EventHandler = `on${Capitalize<EventName>}`;
// "onClick" | "onFocus" | "onBlur"

// CSS properties
type CSSProperty = "color" | "background-color" | "font-size";
type CSSValue<T extends CSSProperty> = T extends "color"
  ? string
  : T extends "background-color"
  ? string
  : T extends "font-size"
  ? string | number
  : never;

// API routes
type HttpMethod = "GET" | "POST" | "PUT" | "DELETE";
type ApiRoute = `/api/${string}`;
type ApiEndpoint = `${HttpMethod} ${ApiRoute}`;

const endpoint: ApiEndpoint = "GET /api/users"; // ✅ Valid
const invalid: ApiEndpoint = "GET /users"; // ❌ Error
```

## Type Safety Checklist

Before committing code, ensure:

- [ ] No `any` types (search for `: any` in your code)
- [ ] All function parameters and returns are typed
- [ ] All component props have interfaces
- [ ] API responses have defined types
- [ ] Event handlers use proper React event types
- [ ] State variables have explicit types when inference isn't clear
- [ ] No TypeScript errors in the editor
- [ ] `npm run build` passes without type errors

## Common Type Errors and Solutions

### Error: "Object is possibly 'null' or 'undefined'"

```typescript
// ❌ BAD
const value = myRef.current.value; // Error if ref can be null

// ✅ GOOD: Optional chaining
const value = myRef.current?.value;

// ✅ GOOD: Null check
if (myRef.current) {
  const value = myRef.current.value;
}

// ✅ GOOD: Non-null assertion (only if you're sure)
const value = myRef.current!.value;
```

### Error: "Type 'X' is not assignable to type 'Y'"

```typescript
// ❌ BAD
const items: DataItem[] = fetchedData; // Error if types don't match

// ✅ GOOD: Use proper type assertion
const items = fetchedData as DataItem[];

// ✅ BETTER: Type the API response
const response = await axios.get<DataItem[]>("/api/items");
const items = response.data; // Correctly typed
```

### Error: "Property 'X' does not exist on type 'Y'"

```typescript
// ❌ BAD
const value = data.unknownProperty; // Error

// ✅ GOOD: Check if property exists
if ("unknownProperty" in data) {
  const value = data.unknownProperty;
}

// ✅ GOOD: Use optional chaining
const value = data?.unknownProperty;

// ✅ GOOD: Define proper interface
interface DataWithProperty {
  unknownProperty: string;
}
const value = (data as DataWithProperty).unknownProperty;
```

---

**Remember:** Embrace TypeScript's strictness - it catches bugs before runtime and improves code maintainability.
