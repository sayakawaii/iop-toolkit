// Central API base URL.
//
// Default is empty string => all requests use relative paths (same-origin),
// which works both behind the nginx reverse proxy (production) and with the
// Vite dev-server proxy (development). Override with VITE_API_BASE (e.g.
// "http://10.101.15.238") only when the frontend must target a remote
// backend from a different origin.
export const API_BASE = (import.meta.env.VITE_API_BASE ?? "") as string;

// Build a WebSocket base URL from API_BASE (or the current page origin when
// API_BASE is relative), converting the http(s) scheme to ws(s).
export function wsBaseUrl(): string {
  const base = API_BASE || globalThis.location.origin;
  return base.replace(/^http/, "ws");
}
