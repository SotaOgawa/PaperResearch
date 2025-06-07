export const BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL || "https://paperresearch-production.up.railway.app/api";

export async function fetchJSON<T>(url: string): Promise<T> {
  const response = await fetch(`${BASE_URL}${url}`);
  if (!response.ok) {
    // サーバーからのエラーをパースして投げる
    const errorBody = await response.json().catch(() => ({}));
    const message = errorBody.error || `Fetch failed with ${response.status}`;
    throw new Error(message);
  }
  return response.json();
}

export async function postJSON<T>(path: string, body: any): Promise<T> {
  const response = await fetch(`${BASE_URL}${path}`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(body),
  });

  if (!response.ok) {
    const errorBody = await response.json().catch(() => ({}));
    const message = errorBody.error || `POST failed with ${response.status}`;
    throw new Error(message);
  }

  return response.json();
}

export async function putJSON<T>(path: string, body: any): Promise<T> {
  const response = await fetch(`${BASE_URL}${path}`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(body),
  });

  if (!response.ok) {
    const errorBody = await response.json().catch(() => ({}));
    const message = errorBody.error || `PUT failed with ${response.status}`;
    throw new Error(message);
  }

  return response.json();
}

export async function deleteJSON<T>(path: string): Promise<T> {
  const response = await fetch(`${BASE_URL}${path}`, {
    method: "DELETE",
  });

  if (!response.ok) {
    const errorBody = await response.json().catch(() => ({}));
    const message = errorBody.error || `DELETE failed with ${response.status}`;
    throw new Error(message);
  }

  return response.json();
}