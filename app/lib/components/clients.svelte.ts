import type { Client, HistoryEntry } from "$lib/types";
import { getAPIEndpoint } from "$lib/utils";

export class ClientsContext {
  clients: Client[] = $state([]);
  history: HistoryEntry[] | undefined = $state();
  loading = $state(false);
  loadingHistory = $state(false);

  async fetchClients() {
    this.loading = true;
    try {
      const res = await fetch(getAPIEndpoint("/api/clients"), {
        method: "GET",
        credentials: "include",
      });
      if (res?.ok) {
        const data = await res?.json();
        this.clients = data;
      }
    } finally {
      this.loading = false;
    }
  }

  async fetchHistory() {
    const res = await fetch(getAPIEndpoint("/api/history"), {
      method: "GET",
      credentials: "include",
    });
    if (res?.ok) {
      const data = await res?.json();
      this.history = data;
    }
  }
}

export const clientsCtx = new ClientsContext();
