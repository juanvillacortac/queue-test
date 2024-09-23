<script lang="ts">
  import { clientTypeString } from "$lib/types";
  import { untrack } from "svelte";
  import { clientsCtx } from "./clients.svelte";

  $effect(() => {
    untrack(() => {
      clientsCtx.fetchHistory();
    });
  });
</script>

<div class="w-full h-full flex flex-col space-y-4 flex-grow">
  <div class="flex w-full justify-between items-center">
    <div class="flex items-center space-x-2">
      <span class="i-material-symbols:history text-2xl"></span>
      <h2 class="text-lg font-bold leading-none">Historial de atención</h2>
    </div>
  </div>
  <div
    class="flex rounded-lg flex-grow skeleton w-full border border-light-900 overflow-auto divide-y divide-light-900 px4 flex-col h-0"
    class:border-transparent={!clientsCtx.history}
    class:skeleton={!clientsCtx.history}
  >
    {#if clientsCtx.history}
      {#each clientsCtx.history || [] as entry}
        <div class="flex flex-col space-y-2 py-2 justify-start">
          <p class="text-sm font-bold space-x-1 items-center">
            <span>{entry.client.name}</span>
          </p>
          <p class="text-sm font-bold space-x-1 items-center">
            <span>DPI:</span>
            <span
              class="font-mono font-normal flex-inline p1 rounded bg-light-900 text-xs"
              >{entry.client.dpi}</span
            >
          </p>
          <p class="text-sm font-bold space-x-1 items-center">
            <span>Tipo:</span>
            <span
              class="font-mono font-normal flex-inline p1 rounded bg-light-900 text-xs"
              >{clientTypeString(entry.client.clientType)}</span
            >
          </p>
          <p class="text-sm font-bold space-x-1 items-center">
            <span>Atendido por:</span>
            <span
              class="font-normal flex-inline p1 rounded bg-light-900 text-xs"
              >{entry.attendedBy.email}</span
            >
          </p>
          <p class="text-sm font-bold space-x-1 items-center">
            <span>Tiempo de atención:</span>
            <span
              class="font-normal flex-inline p1 rounded bg-light-900 text-xs"
              >{new Date(entry.attendedAt).toLocaleString()}</span
            >
          </p>
        </div>
      {:else}
        <div class="flex w-full h-full items-center justify-center">
          <p class="text-lg font-bold">Sin clientes en la cola</p>
        </div>
      {/each}
    {/if}
  </div>
</div>
