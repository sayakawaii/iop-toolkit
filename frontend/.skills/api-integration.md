---
description: Integrate backend APIs with proper TypeScript typing, error handling, loading states, and axios best practices for IOP Toolkit.
applyTo:
  - "**/*.tsx"
  - "**/*.ts"
  - "src/config/**"
---

# API Integration Skill

This skill provides patterns for integrating backend APIs with proper typing, error handling, and state management in the IOP Toolkit.

## When to Use This Skill

- Fetching data from backend APIs
- Submitting forms or data to the server
- Implementing real-time updates
- Handling file uploads/downloads
- Managing API authentication and errors

## API Configuration

### 1. Backend Proxy Setup

The project uses Vite's proxy configuration for API requests:

**vite.config.ts:**
```typescript
export default defineConfig({
  server: {
    proxy: {
      "/api": {
        target: "http://localhost:8080",
        changeOrigin: true,
      },
    },
  },
});
```

**What this means:**
- Development: `/api/*` requests go to `http://localhost:8080/api/*`
- Production: Configure backend URL via environment variables

### 2. Centralized API Configuration (Recommended)

Create a centralized API config file:

```typescript
// src/config/api.ts
import axios from "axios";

// Base configuration
export const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || "/api";

// Create axios instance with defaults
export const apiClient = axios.create({
  baseURL: API_BASE_URL,
  timeout: 10000, // 10 seconds
  headers: {
    "Content-Type": "application/json",
  },
});

// Request interceptor (for adding auth tokens, etc.)
apiClient.interceptors.request.use(
  (config) => {
    // Add auth token if needed
    // const token = localStorage.getItem("authToken");
    // if (token) {
    //   config.headers.Authorization = `Bearer ${token}`;
    // }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// Response interceptor (for global error handling)
apiClient.interceptors.response.use(
  (response) => response,
  (error) => {
    // Global error handling
    if (error.response?.status === 401) {
      // Handle unauthorized (e.g., redirect to login)
      console.error("Unauthorized access - redirect to login");
    }
    if (error.response?.status === 500) {
      console.error("Server error:", error.response.data);
    }
    return Promise.reject(error);
  }
);

// API endpoint constants
export const API_ENDPOINTS = {
  // OMCI Analyzer
  OMCI_ONUS: "/omci/onus",
  OMCI_DETAILS: (id: string) => `/omci/details/${id}`,
  
  // Collector
  COLLECTOR_LOGS: "/collector/logs",
  COLLECTOR_START: "/collector/start",
  
  // Library
  LIBRARY_ITEMS: "/library/items",
  
  // Add more endpoints as needed
} as const;
```

### 3. Environment Variables

Create a `.env` file for environment-specific config:

```bash
# .env.development
VITE_API_BASE_URL=http://localhost:8080/api
VITE_ENABLE_DEBUG=true

# .env.production
VITE_API_BASE_URL=https://api.production.com/api
VITE_ENABLE_DEBUG=false
```

**Access in code:**
```typescript
const apiUrl = import.meta.env.VITE_API_BASE_URL;
const isDebug = import.meta.env.VITE_ENABLE_DEBUG === "true";
```

## API Request Patterns

### 1. Basic GET Request

```typescript
import { useState, useEffect } from "react";
import axios from "axios";

interface DataItem {
  id: number;
  name: string;
  status: string;
}

const MyComponent = () => {
  const [data, setData] = useState<DataItem[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        setIsLoading(true);
        setError(null);

        const response = await axios.get<DataItem[]>("/api/items");
        setData(response.data);
      } catch (err) {
        const errorMessage = err instanceof Error 
          ? err.message 
          : "Failed to fetch data";
        setError(errorMessage);
        console.error("Error fetching data:", err);
      } finally {
        setIsLoading(false);
      }
    };

    fetchData();
  }, []);

  if (isLoading) return <Spinner label="Loading..." />;
  if (error) return <div className="text-red-500">{error}</div>;

  return (
    <div>
      {data.map((item) => (
        <div key={item.id}>{item.name}</div>
      ))}
    </div>
  );
};
```

### 2. POST Request (Form Submission)

```typescript
import { useState } from "react";
import { Button, Input } from "@heroui/react";
import axios from "axios";

interface FormData {
  name: string;
  email: string;
}

interface ApiResponse {
  success: boolean;
  message: string;
  id?: number;
}

const FormComponent = () => {
  const [formData, setFormData] = useState<FormData>({
    name: "",
    email: "",
  });
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState<string | null>(null);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    
    try {
      setIsSubmitting(true);
      setError(null);
      setSuccess(null);

      const response = await axios.post<ApiResponse>(
        "/api/submit",
        formData
      );

      setSuccess(response.data.message);
      // Reset form
      setFormData({ name: "", email: "" });
    } catch (err) {
      if (axios.isAxiosError(err)) {
        setError(err.response?.data?.message || "Submission failed");
      } else {
        setError("An unexpected error occurred");
      }
      console.error("Error submitting form:", err);
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <form onSubmit={handleSubmit} className="space-y-4">
      <Input
        label="Name"
        value={formData.name}
        onChange={(e) => setFormData({ ...formData, name: e.target.value })}
        isRequired
      />
      <Input
        label="Email"
        type="email"
        value={formData.email}
        onChange={(e) => setFormData({ ...formData, email: e.target.value })}
        isRequired
      />
      
      {error && <p className="text-red-500">{error}</p>}
      {success && <p className="text-green-500">{success}</p>}
      
      <Button
        type="submit"
        color="primary"
        isLoading={isSubmitting}
        isDisabled={!formData.name || !formData.email}
      >
        Submit
      </Button>
    </form>
  );
};
```

### 3. PUT/PATCH Request (Update)

```typescript
const updateItem = async (id: number, updates: Partial<DataItem>) => {
  try {
    const response = await axios.patch<DataItem>(
      `/api/items/${id}`,
      updates
    );
    return response.data;
  } catch (err) {
    console.error("Error updating item:", err);
    throw err;
  }
};

// Usage
const handleUpdate = async () => {
  try {
    setIsLoading(true);
    const updated = await updateItem(itemId, { status: "active" });
    setItem(updated);
  } catch (err) {
    setError("Failed to update item");
  } finally {
    setIsLoading(false);
  }
};
```

### 4. DELETE Request

```typescript
const deleteItem = async (id: number) => {
  try {
    await axios.delete(`/api/items/${id}`);
  } catch (err) {
    console.error("Error deleting item:", err);
    throw err;
  }
};

// Usage with confirmation
const handleDelete = async (id: number) => {
  if (!confirm("Are you sure you want to delete this item?")) {
    return;
  }

  try {
    setIsLoading(true);
    await deleteItem(id);
    // Remove from local state
    setItems(items.filter((item) => item.id !== id));
  } catch (err) {
    setError("Failed to delete item");
  } finally {
    setIsLoading(false);
  }
};
```

### 5. File Upload

```typescript
const uploadFile = async (file: File) => {
  const formData = new FormData();
  formData.append("file", file);
  formData.append("description", "File description");

  try {
    const response = await axios.post<{ fileId: string; url: string }>(
      "/api/upload",
      formData,
      {
        headers: {
          "Content-Type": "multipart/form-data",
        },
        onUploadProgress: (progressEvent) => {
          const percentCompleted = progressEvent.total
            ? Math.round((progressEvent.loaded * 100) / progressEvent.total)
            : 0;
          console.log(`Upload progress: ${percentCompleted}%`);
          setUploadProgress(percentCompleted);
        },
      }
    );
    return response.data;
  } catch (err) {
    console.error("Error uploading file:", err);
    throw err;
  }
};

// Usage in component
const FileUploadComponent = () => {
  const [uploadProgress, setUploadProgress] = useState(0);
  const [isUploading, setIsUploading] = useState(false);

  const handleFileChange = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (!file) return;

    try {
      setIsUploading(true);
      const result = await uploadFile(file);
      console.log("File uploaded:", result);
    } catch (err) {
      setError("Upload failed");
    } finally {
      setIsUploading(false);
      setUploadProgress(0);
    }
  };

  return (
    <div>
      <input type="file" onChange={handleFileChange} disabled={isUploading} />
      {isUploading && (
        <div className="mt-2">
          <p>Uploading: {uploadProgress}%</p>
          <progress value={uploadProgress} max={100} />
        </div>
      )}
    </div>
  );
};
```

### 6. File Download

```typescript
const downloadFile = async (fileId: string, filename: string) => {
  try {
    const response = await axios.get(`/api/download/${fileId}`, {
      responseType: "blob",
    });

    // Create blob link to download
    const url = window.URL.createObjectURL(new Blob([response.data]));
    const link = document.createElement("a");
    link.href = url;
    link.setAttribute("download", filename);
    document.body.appendChild(link);
    link.click();
    link.remove();
    window.URL.revokeObjectURL(url);
  } catch (err) {
    console.error("Error downloading file:", err);
    throw err;
  }
};

// Usage
<Button onPress={() => downloadFile("file123", "data.csv")}>
  Download CSV
</Button>
```

## Custom Hooks for API Calls

### 1. Generic Data Fetching Hook

```typescript
// src/hooks/useFetch.ts
import { useState, useEffect } from "react";
import axios, { AxiosRequestConfig } from "axios";

interface FetchState<T> {
  data: T | null;
  isLoading: boolean;
  error: string | null;
  refetch: () => void;
}

export const useFetch = <T,>(
  url: string,
  config?: AxiosRequestConfig
): FetchState<T> => {
  const [data, setData] = useState<T | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const fetchData = async () => {
    try {
      setIsLoading(true);
      setError(null);
      const response = await axios.get<T>(url, config);
      setData(response.data);
    } catch (err) {
      setError(err instanceof Error ? err.message : "Failed to fetch");
      console.error("Fetch error:", err);
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    fetchData();
  }, [url]);

  return { data, isLoading, error, refetch: fetchData };
};

// Usage
const MyComponent = () => {
  const { data, isLoading, error, refetch } = useFetch<DataItem[]>("/api/items");

  if (isLoading) return <Spinner />;
  if (error) return <div>Error: {error}</div>;

  return (
    <div>
      {data?.map((item) => <div key={item.id}>{item.name}</div>)}
      <Button onPress={refetch}>Refresh</Button>
    </div>
  );
};
```

### 2. Mutation Hook

```typescript
// src/hooks/useMutation.ts
import { useState } from "react";
import axios, { AxiosRequestConfig } from "axios";

interface MutationState<T> {
  data: T | null;
  isLoading: boolean;
  error: string | null;
  mutate: (payload: any) => Promise<T | undefined>;
}

export const useMutation = <T,>(
  url: string,
  method: "POST" | "PUT" | "PATCH" | "DELETE" = "POST"
): MutationState<T> => {
  const [data, setData] = useState<T | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const mutate = async (payload: any): Promise<T | undefined> => {
    try {
      setIsLoading(true);
      setError(null);

      const response = await axios.request<T>({
        url,
        method,
        data: payload,
      });

      setData(response.data);
      return response.data;
    } catch (err) {
      const errorMsg = err instanceof Error ? err.message : "Mutation failed";
      setError(errorMsg);
      console.error("Mutation error:", err);
      throw err;
    } finally {
      setIsLoading(false);
    }
  };

  return { data, isLoading, error, mutate };
};

// Usage
const FormComponent = () => {
  const { mutate, isLoading, error } = useMutation<ApiResponse>("/api/submit");

  const handleSubmit = async (formData: FormData) => {
    try {
      const result = await mutate(formData);
      console.log("Success:", result);
    } catch (err) {
      // Error already logged in hook
    }
  };

  return (
    <Button onPress={() => handleSubmit(data)} isLoading={isLoading}>
      Submit
    </Button>
  );
};
```

## Error Handling Patterns

### 1. Comprehensive Error Handling

```typescript
import axios, { AxiosError } from "axios";

interface ApiError {
  message: string;
  code?: string;
  details?: unknown;
}

const handleApiError = (err: unknown): string => {
  if (axios.isAxiosError(err)) {
    const axiosError = err as AxiosError<ApiError>;
    
    // Server responded with error
    if (axiosError.response) {
      const status = axiosError.response.status;
      const data = axiosError.response.data;
      
      switch (status) {
        case 400:
          return data?.message || "Invalid request";
        case 401:
          return "Unauthorized - please login";
        case 403:
          return "Access forbidden";
        case 404:
          return "Resource not found";
        case 500:
          return "Server error - please try again later";
        default:
          return data?.message || `Error ${status}`;
      }
    }
    
    // Request made but no response
    if (axiosError.request) {
      return "No response from server - check your connection";
    }
  }
  
  // Something else went wrong
  return err instanceof Error ? err.message : "An unexpected error occurred";
};

// Usage
try {
  const response = await axios.get("/api/data");
  setData(response.data);
} catch (err) {
  const errorMessage = handleApiError(err);
  setError(errorMessage);
}
```

### 2. Retry Logic

```typescript
const fetchWithRetry = async <T,>(
  url: string,
  maxRetries = 3,
  delay = 1000
): Promise<T> => {
  for (let attempt = 1; attempt <= maxRetries; attempt++) {
    try {
      const response = await axios.get<T>(url);
      return response.data;
    } catch (err) {
      if (attempt === maxRetries) {
        throw err;
      }
      
      console.log(`Retry attempt ${attempt} of ${maxRetries}`);
      await new Promise((resolve) => setTimeout(resolve, delay * attempt));
    }
  }
  
  throw new Error("Max retries exceeded");
};

// Usage
try {
  const data = await fetchWithRetry<DataItem[]>("/api/items", 3, 1000);
  setData(data);
} catch (err) {
  setError("Failed after multiple attempts");
}
```

## Loading States

### 1. Table with Loading State

```typescript
import { Table, TableHeader, TableColumn, TableBody, TableRow, TableCell, Spinner } from "@heroui/react";

const DataTable = () => {
  const { data, isLoading } = useFetch<DataItem[]>("/api/items");

  return (
    <Table aria-label="Data table">
      <TableHeader>
        <TableColumn>NAME</TableColumn>
        <TableColumn>STATUS</TableColumn>
      </TableHeader>
      <TableBody
        isLoading={isLoading}
        loadingContent={<Spinner label="Loading..." />}
        emptyContent={!isLoading && "No data available"}
      >
        {(data || []).map((item) => (
          <TableRow key={item.id}>
            <TableCell>{item.name}</TableCell>
            <TableCell>{item.status}</TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  );
};
```

## Type Safety Best Practices

```typescript
// Define API response types
interface ApiListResponse<T> {
  data: T[];
  total: number;
  page: number;
  pageSize: number;
}

interface ApiDetailResponse<T> {
  data: T;
  metadata?: Record<string, unknown>;
}

interface ApiErrorResponse {
  error: string;
  code: string;
  details?: unknown;
}

// Use generics for reusable API functions
const fetchList = async <T,>(endpoint: string): Promise<ApiListResponse<T>> => {
  const response = await axios.get<ApiListResponse<T>>(endpoint);
  return response.data;
};

const fetchDetail = async <T,>(endpoint: string, id: string): Promise<T> => {
  const response = await axios.get<ApiDetailResponse<T>>(`${endpoint}/${id}`);
  return response.data.data;
};

// Usage
const items = await fetchList<DataItem>("/api/items");
const item = await fetchDetail<DataItem>("/api/items", "123");
```

---

**Remember:** Always handle loading states, errors, and empty states gracefully for better user experience.
