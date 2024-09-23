<script lang="ts">
  import { clientTypeString, type Queue } from "$lib/types";
  import { getWebsocketAPIEndpoint } from "$lib/utils";
  import { untrack } from "svelte";
  import { clientsCtx } from "./clients.svelte";

  let loading = $state(true);

  let socket: WebSocket | undefined = $state();
  let data: Queue | undefined = $state();

  function connectWebSocket() {
    socket = new WebSocket(getWebsocketAPIEndpoint("/api/queue/ws"));

    socket.onmessage = (event) => {
      loading = false;
      data = JSON.parse(event.data);
      clientsCtx.fetchHistory();
    };

    socket.onclose = (event) => {
      console.log("Conexión cerrada", event);
      connectWebSocket();
    };

    socket.onerror = (error) => {
      console.error("Error en la conexión", error);
    };
  }

  $effect(() => {
    untrack(() => {
      connectWebSocket();
    });
  });
</script>

<div class="w-full h-full flex flex-col space-y-4 flex-grow">
  <div class="flex w-full justify-between items-center">
    <div class="flex items-center space-x-2">
      <span class="i-material-symbols:person-raised-hand text-2xl"></span>
      <h2 class="text-lg font-bold leading-none">Cola de atención</h2>
    </div>
  </div>
  <div
    class="flex rounded-lg flex-grow skeleton w-full border border-light-900 overflow-auto divide-y divide-light-900 px4 flex-col h-0"
    class:border-transparent={loading}
    class:skeleton={loading}
  >
    {#if !loading}
      {#each data?.items || [] as client, idx}
        <div class="flex flex-col space-y-2 py-2 justify-start">
          <p class="text-sm font-bold space-x-1 items-center">
            <span>{client.name}</span>
          </p>
          <p class="text-sm font-bold space-x-1 items-center">
            <span>DPI:</span>
            <span
              class="font-mono font-normal flex-inline p1 rounded bg-light-900 text-xs"
              >{client.dpi}</span
            >
          </p>
          <p class="text-sm font-bold space-x-1 items-center">
            <span>Tipo:</span>
            <span
              class="font-mono font-normal flex-inline p1 rounded bg-light-900 text-xs"
              >{clientTypeString(client.clientType)}</span
            >
          </p>
          <p class="text-sm font-bold space-x-1 items-center">
            <span>Tiempo de llegada:</span>
            <span
              class="font-normal flex-inline p1 rounded bg-light-900 text-xs"
              >{new Date(client.arrival).toLocaleString()}</span
            >
          </p>
          {#if idx == 0}
            <div
              class="font-mono font-normal p1 rounded bg-green-100 text-xs text-left justify-self-start flex mr-auto"
            >
              En atención
            </div>
          {/if}
        </div>
      {:else}
        <div class="flex w-full h-full items-center justify-center">
          <p class="text-lg font-bold">Sin clientes en la cola</p>
        </div>
      {/each}
    {/if}
  </div>
</div>
