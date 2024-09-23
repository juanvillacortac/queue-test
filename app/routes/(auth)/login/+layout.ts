import { redirect } from "@sveltejs/kit";
import { browser } from "$app/environment";

export async function load({ parent }) {
  if (browser) {
    const data = await parent();
    if (data.user) {
      throw redirect(301, "/");
    }
  }
}
