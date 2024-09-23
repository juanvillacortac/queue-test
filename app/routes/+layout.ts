import { browser } from "$app/environment";
import { getAPIEndpoint } from "$lib/utils";

export const prerender = true;

export async function load({ fetch }) {
  if (browser) {
    const res = await fetch(getAPIEndpoint("/api/auth/whoami"), {
      credentials: "include",
    });
    if (res.ok) {
      const user = await res.json();
      return {
        user: {
          id: (user?.id as string) || "",
          email: (user?.email as string) || "",
        },
      };
    }
  }
  return {};
}
