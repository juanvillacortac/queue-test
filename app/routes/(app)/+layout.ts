import { redirect } from "@sveltejs/kit";

export const ssr = false;

export async function load({ parent }) {
  const data = await parent();
  if (!data.user) {
    throw redirect(301, "/login");
  }
}
