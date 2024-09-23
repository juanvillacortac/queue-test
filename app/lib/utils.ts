import { env } from "$env/dynamic/public";

export function getAPIEndpoint(path: string): string {
  return `${env.PUBLIC_API_PREFIX || ""}/${path.replace(/^(\/)/, "")}`;
}

export function getWebsocketAPIEndpoint(path: string): string {
  const base = new URL(
    "/",
    env.PUBLIC_API_PREFIX || window.location.toString(),
  );
  base.protocol = "ws:";
  return `${base.toString()}${path.replace(/^(\/)/, "")}`;
}
