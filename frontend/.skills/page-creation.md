---
description: Create new pages in IOP Toolkit with proper routing, layout integration, navigation setup, and TypeScript configuration.
applyTo:
  - "src/pages/**/*.tsx"
  - "src/App.tsx"
  - "src/config/site.ts"
---

# Page Creation Skill

This skill guides you through creating new pages with proper routing, navigation, and layout integration in the IOP Toolkit.

## When to Use This Skill

- Adding a new feature page to the application
- Creating sub-pages for existing features
- Setting up multi-step workflows with separate pages
- Establishing new navigation menu items

## Page Creation Process

### Step 1: Create Page Component

**File Location Pattern:**
```
src/pages/
├── <feature-name>/
│   ├── index.tsx          # Main page (e.g., /featurename)
│   ├── SubPage.tsx        # Sub-pages (e.g., /featurename/subpage)
│   └── components/        # Page-specific components (optional)
│       └── LocalComponent.tsx
```

**Basic Page Template:**

```typescript
// src/pages/myfeature/index.tsx
import { FC } from "react";
import DefaultLayout from "@/layouts/default";
import { title } from "@/components/primitives";

const MyFeaturePage: FC = () => {
  return (
    <DefaultLayout>
      <section className="flex flex-col items-center justify-center gap-4 py-8 md:py-10">
        <div className="inline-block max-w-lg text-center justify-center">
          <h1 className={title()}>My Feature</h1>
          <p className="mt-4 text-lg text-default-600">
            Description of what this page does
          </p>
        </div>

        <div className="mt-8 w-full max-w-4xl">
          {/* Main content here */}
        </div>
      </section>
    </DefaultLayout>
  );
};

export default MyFeaturePage;
```

**Page with Data Fetching:**

```typescript
import { FC, useState, useEffect } from "react";
import { Spinner } from "@heroui/react";
import axios from "axios";
import DefaultLayout from "@/layouts/default";

interface PageData {
  id: number;
  name: string;
  // ... other fields
}

const DataFetchingPage: FC = () => {
  const [data, setData] = useState<PageData[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        setIsLoading(true);
        const response = await axios.get<PageData[]>("/api/myfeature/data");
        setData(response.data);
        setError(null);
      } catch (err) {
        setError(err instanceof Error ? err.message : "Failed to load data");
        console.error("Error fetching data:", err);
      } finally {
        setIsLoading(false);
      }
    };

    fetchData();
  }, []);

  if (isLoading) {
    return (
      <DefaultLayout>
        <div className="flex h-screen items-center justify-center">
          <Spinner label="Loading..." size="lg" />
        </div>
      </DefaultLayout>
    );
  }

  if (error) {
    return (
      <DefaultLayout>
        <div className="flex h-screen items-center justify-center">
          <div className="text-center">
            <p className="text-lg text-red-500">Error: {error}</p>
          </div>
        </div>
      </DefaultLayout>
    );
  }

  return (
    <DefaultLayout>
      <section className="flex flex-col gap-4 py-8 md:py-10">
        {/* Render data */}
      </section>
    </DefaultLayout>
  );
};

export default DataFetchingPage;
```

**Full-Width Page (without container):**

```typescript
import DefaultLayout from "@/layouts/default";

const FullWidthPage: FC = () => {
  return (
    <DefaultLayout fullWidth>
      {/* Content spans full width */}
      <div className="p-6">
        <h1>Full Width Content</h1>
        {/* Large tables, diagrams, etc. */}
      </div>
    </DefaultLayout>
  );
};

export default FullWidthPage;
```

### Step 2: Add Route to App.tsx

**Location:** `src/App.tsx`

```typescript
// 1. Import the new page component
import MyFeaturePage from "@/pages/myfeature/index";
import MyFeatureSubPage from "@/pages/myfeature/SubPage";

function App() {
  return (
    <Routes>
      {/* Existing routes... */}
      
      {/* 2. Add new routes */}
      <Route element={<MyFeaturePage />} path="/myfeature" />
      <Route element={<MyFeatureSubPage />} path="/myfeature/subpage" />
    </Routes>
  );
}
```

**Route Pattern Examples:**

```typescript
// Simple route
<Route element={<AboutPage />} path="/about" />

// Nested routes
<Route element={<FeaturePage />} path="/feature" />
<Route element={<FeatureDetailPage />} path="/feature/detail" />
<Route element={<FeatureStatsPage />} path="/feature/stats" />

// Dynamic routes (with params)
<Route element={<UserProfilePage />} path="/users/:userId" />

// Default/fallback route
<Route element={<IndexPage />} path="/" />
<Route element={<NotFoundPage />} path="*" />
```

### Step 3: Add Navigation Links

**Location:** `src/config/site.ts`

```typescript
export const siteConfig = {
  name: "IOP Toolkit",
  description: "Make Beautiful and Modernized Debug Tools.",
  navItems: [
    {
      label: "Home",
      href: "/index",
    },
    // ... existing items
    {
      label: "My Feature",  // ← Add your new page
      href: "/myfeature",
    },
  ],
  navMenuItems: [
    {
      label: "Home",
      href: "/index",
    },
    // ... existing items
    {
      label: "My Feature",  // ← Add to mobile menu too
      href: "/myfeature",
    },
  ],
};
```

**Tips:**
- Keep labels concise (2-3 words max)
- Use consistent capitalization
- Order items logically (by usage frequency or feature grouping)
- Consider hiding experimental features initially (comment out navigation)

### Step 4: Page with Sub-Navigation

For features with multiple sub-pages, create local navigation:

```typescript
// src/pages/myfeature/index.tsx
import { FC, useState } from "react";
import { useNavigate, useLocation } from "react-router-dom";
import { Button } from "@heroui/react";
import DefaultLayout from "@/layouts/default";

const MyFeaturePage: FC = () => {
  const navigate = useNavigate();
  const location = useLocation();

  const tabs = [
    { name: "Overview", path: "/myfeature" },
    { name: "Details", path: "/myfeature/details" },
    { name: "Analytics", path: "/myfeature/analytics" },
  ];

  return (
    <DefaultLayout>
      <section className="flex flex-col gap-4 py-8">
        {/* Tab navigation */}
        <div className="flex gap-2 border-b border-gray-200 pb-2">
          {tabs.map((tab) => (
            <Button
              key={tab.path}
              color={location.pathname === tab.path ? "primary" : "default"}
              variant={location.pathname === tab.path ? "solid" : "light"}
              onPress={() => navigate(tab.path)}
            >
              {tab.name}
            </Button>
          ))}
        </div>

        {/* Page content */}
        <div className="mt-4">
          {/* Content based on current tab */}
        </div>
      </section>
    </DefaultLayout>
  );
};

export default MyFeaturePage;
```

### Step 5: Using URL Parameters

**Route with parameters:**
```typescript
// src/App.tsx
<Route element={<ItemDetailPage />} path="/items/:itemId" />
```

**Accessing parameters in page:**
```typescript
// src/pages/items/ItemDetail.tsx
import { FC, useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import axios from "axios";
import DefaultLayout from "@/layouts/default";

interface Item {
  id: number;
  name: string;
  description: string;
}

const ItemDetailPage: FC = () => {
  const { itemId } = useParams<{ itemId: string }>();
  const navigate = useNavigate();
  const [item, setItem] = useState<Item | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const fetchItem = async () => {
      try {
        setIsLoading(true);
        const response = await axios.get<Item>(`/api/items/${itemId}`);
        setItem(response.data);
      } catch (err) {
        console.error("Error fetching item:", err);
        // Redirect to list if item not found
        navigate("/items");
      } finally {
        setIsLoading(false);
      }
    };

    if (itemId) {
      fetchItem();
    }
  }, [itemId]);

  if (isLoading) return <DefaultLayout><Spinner /></DefaultLayout>;
  if (!item) return null;

  return (
    <DefaultLayout>
      <section className="py-8">
        <h1 className="text-3xl font-bold">{item.name}</h1>
        <p className="mt-4">{item.description}</p>
      </section>
    </DefaultLayout>
  );
};

export default ItemDetailPage;
```

### Step 6: Using Query Parameters

**Setting query parameters:**
```typescript
import { useNavigate } from "react-router-dom";

const MyComponent = () => {
  const navigate = useNavigate();

  const handleFilter = (filter: string) => {
    navigate(`/mypage?filter=${filter}&sort=asc`);
  };

  return <Button onPress={() => handleFilter("active")}>Filter</Button>;
};
```

**Reading query parameters:**
```typescript
import { useLocation } from "react-router-dom";

const MyPage: FC = () => {
  const location = useLocation();
  const params = new URLSearchParams(location.search);
  
  const filter = params.get("filter") || "all";
  const sort = params.get("sort") || "desc";

  // Use filter and sort in your logic
  useEffect(() => {
    fetchData(filter, sort);
  }, [filter, sort]);

  return (
    <DefaultLayout>
      <div>
        <p>Active filter: {filter}</p>
        <p>Sort order: {sort}</p>
      </div>
    </DefaultLayout>
  );
};
```

### Step 7: Programmatic Navigation

```typescript
import { useNavigate } from "react-router-dom";

const FormPage: FC = () => {
  const navigate = useNavigate();

  const handleSubmit = async (data: FormData) => {
    try {
      await axios.post("/api/submit", data);
      // Navigate to success page
      navigate("/success");
    } catch (err) {
      // Stay on page and show error
      setError("Submission failed");
    }
  };

  const handleCancel = () => {
    // Go back to previous page
    navigate(-1);
    // Or navigate to specific page
    // navigate("/home");
  };

  return (
    <DefaultLayout>
      <form onSubmit={handleSubmit}>
        {/* Form fields */}
        <div className="flex gap-2">
          <Button type="submit" color="primary">Submit</Button>
          <Button onPress={handleCancel}>Cancel</Button>
        </div>
      </form>
    </DefaultLayout>
  );
};
```

### Step 8: Protected/Conditional Routes

If you need authentication or conditional access:

```typescript
// src/components/ProtectedRoute.tsx
import { FC, ReactNode } from "react";
import { Navigate } from "react-router-dom";

interface ProtectedRouteProps {
  children: ReactNode;
  isAllowed: boolean;
  redirectTo?: string;
}

export const ProtectedRoute: FC<ProtectedRouteProps> = ({
  children,
  isAllowed,
  redirectTo = "/",
}) => {
  if (!isAllowed) {
    return <Navigate to={redirectTo} replace />;
  }

  return <>{children}</>;
};

// Usage in App.tsx
import { ProtectedRoute } from "@/components/ProtectedRoute";

<Route
  path="/admin"
  element={
    <ProtectedRoute isAllowed={isAdmin}>
      <AdminPage />
    </ProtectedRoute>
  }
/>
```

## Page Layout Patterns

### Standard Content Page
```typescript
<DefaultLayout>
  <section className="flex flex-col items-center justify-center gap-4 py-8 md:py-10">
    <div className="inline-block max-w-lg text-center">
      <h1 className={title()}>Page Title</h1>
    </div>
    <div className="mt-8 w-full max-w-4xl">
      {/* Content */}
    </div>
  </section>
</DefaultLayout>
```

### Full-Width Data Table Page
```typescript
<DefaultLayout fullWidth>
  <div className="px-6 py-8">
    <h1 className={title()}>Data Overview</h1>
    <Table className="mt-6">
      {/* Table content */}
    </Table>
  </div>
</DefaultLayout>
```

### Two-Column Layout
```typescript
<DefaultLayout>
  <section className="py-8">
    <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
      {/* Sidebar */}
      <aside className="lg:col-span-1">
        <Card>
          <CardBody>
            {/* Filters, navigation, etc. */}
          </CardBody>
        </Card>
      </aside>

      {/* Main content */}
      <main className="lg:col-span-2">
        {/* Primary content */}
      </main>
    </div>
  </section>
</DefaultLayout>
```

### Dashboard Layout
```typescript
<DefaultLayout fullWidth>
  <div className="px-6 py-8">
    <h1 className={title()}>Dashboard</h1>
    
    {/* Stats cards */}
    <div className="mt-6 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
      {stats.map((stat) => (
        <Card key={stat.label}>
          <CardBody>
            <p className="text-sm text-default-600">{stat.label}</p>
            <p className="text-3xl font-bold mt-2">{stat.value}</p>
          </CardBody>
        </Card>
      ))}
    </div>

    {/* Charts/Tables */}
    <div className="mt-8 grid grid-cols-1 lg:grid-cols-2 gap-6">
      <Card>
        <CardHeader>Chart 1</CardHeader>
        <CardBody>{/* Chart */}</CardBody>
      </Card>
      <Card>
        <CardHeader>Chart 2</CardHeader>
        <CardBody>{/* Chart */}</CardBody>
      </Card>
    </div>
  </div>
</DefaultLayout>
```

## SEO & Meta Information

For better browser tab titles and meta tags:

```typescript
import { useEffect } from "react";

const MyFeaturePage: FC = () => {
  useEffect(() => {
    // Set page title
    document.title = "My Feature | IOP Toolkit";
    
    // Optionally set meta description
    const metaDescription = document.querySelector('meta[name="description"]');
    if (metaDescription) {
      metaDescription.setAttribute(
        "content",
        "Description of my feature page"
      );
    }
  }, []);

  return <DefaultLayout>{/* Content */}</DefaultLayout>;
};
```

## Checklist for New Pages

Before considering a page complete, verify:

- [ ] Page component created in `src/pages/<feature>/`
- [ ] Route added to `src/App.tsx`
- [ ] Navigation link added to `src/config/site.ts` (if public)
- [ ] Uses `DefaultLayout` for consistency
- [ ] Proper TypeScript types for all data
- [ ] Loading states handled (Spinner component)
- [ ] Error states handled gracefully
- [ ] Page title set (document.title)
- [ ] Responsive design tested (mobile, tablet, desktop)
- [ ] Accessibility checked (semantic HTML, ARIA labels)
- [ ] Dark mode support (if applicable)

## Common Patterns Reference

```typescript
// Navigate to another page
const navigate = useNavigate();
navigate("/target-page");

// Navigate back
navigate(-1);

// Navigate with state
navigate("/details", { state: { itemId: 123 } });

// Access location state
const location = useLocation();
const { itemId } = location.state || {};

// Link component (for anchor tags)
import { Link } from "react-router-dom";
<Link to="/about">About</Link>

// Or using HeroUI Link
import { Link } from "@heroui/react";
<Link href="/about">About</Link>
```

---

**Remember:** Check existing pages in `src/pages/` for reference implementations specific to your use case.
