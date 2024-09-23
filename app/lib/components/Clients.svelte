<script lang="ts">
  import { slide } from "svelte/transition";
  import getBottomsheetContext from "./ui/bottomsheetContext.svelte";
  import Button from "./ui/Button.svelte";
  import TextField from "./ui/TextField.svelte";
  import { getAPIEndpoint } from "$lib/utils";
  import Options from "./ui/Options.svelte";
  import { clientTypeString, type Client } from "$lib/types";
  import { search } from "$lib/search";
  import { clientsCtx } from "./clients.svelte";

  const bottomsheet = getBottomsheetContext();
  let saving = $state(false);
  let registerError = $state("");
  let queueError = $state("");

  let formData = $state({
    name: "",
    dpi: "",
    clientType: 0,
  });

  let queueData: {
    client?: Client;
    requiredOperations: string;
  } = $state({
    requiredOperations: "10",
  });

  $effect(() => {
    if (!bottomsheet.open) {
      registerError = "";
      queueError = "";
      formData = {
        name: "",
        dpi: "",
        clientType: 0,
      };
      queueData = {
        requiredOperations: "10",
      };
    }
  });

  async function registerClient() {
    saving = true;
    try {
      const res = await fetch(getAPIEndpoint("/api/clients"), {
        method: "POST",
        credentials: "include",
        body: JSON.stringify($state.snapshot(formData)),
      });
      if (!res?.ok) {
        const data = await res?.json();
        registerError = data?.error || "";
      } else {
        bottomsheet.close();
        searchStr = "";
        clientsCtx.fetchClients();
      }
    } catch (err) {
      registerError = `${err}`;
    } finally {
      saving = false;
    }
  }

  async function sendToQueue() {
    saving = true;
    try {
      const res = await fetch(getAPIEndpoint("/api/queue"), {
        method: "POST",
        credentials: "include",
        body: JSON.stringify({
          dpi: queueData.client?.dpi,
          requiredOperations: +queueData.requiredOperations,
        }),
      });
      if (!res?.ok) {
        const data = await res?.json();
        queueError = data?.error || "";
      } else {
        bottomsheet.close();
      }
    } catch (err) {
      queueError = `${err}`;
    } finally {
      saving = false;
    }
  }

  $effect(() => {
    clientsCtx.fetchClients();
  });

  let searchStr = $state("");
  let filteredItems = $derived.by(() => {
    if (!searchStr?.trim()) return clientsCtx.clients || [];
    return search(clientsCtx.clients, searchStr, ["name", "dpi"]);
  });

  $inspect(clientsCtx.clients);
</script>

{#snippet bottomsheetTitle()}
  <div class="flex w-full justify-between items-center">
    <h2 class="text-lg font-bold">Registro de cliente</h2>
  </div>
{/snippet}

{#snippet bottomsheetQueueTitle()}
  <div class="flex w-full justify-between items-center">
    <h2 class="text-lg font-bold">Procesar cliente</h2>
  </div>
{/snippet}

{#snippet bottomsheetBody()}
  <form class="w-full space-y-4 flex flex-col" onsubmit={registerClient}>
    <TextField
      placeholder="ej. Sarah Connor"
      title="Nombre"
      required
      bind:value={formData.name}
    />
    <TextField
      placeholder="ej. 3050824830116"
      title="DPI del cliente"
      required
      bind:value={formData.dpi}
    />
    <Options
      required
      title="Tipo de cliente"
      bind:value={formData.clientType}
      options={[
        {
          value: 0,
          text: "VIP",
        },
        {
          value: 1,
          text: "Regular",
        },
      ]}
    />
    {#if registerError}
      <p class="text-xs text-red-500" transition:slide={{ duration: 200 }}>
        {registerError}
      </p>
    {/if}
    <Button disabled={saving} class="w-full">Registrar</Button>
  </form>
{/snippet}

{#snippet bottomsheetQueueBody()}
  <form class="w-full space-y-4 flex flex-col" onsubmit={sendToQueue}>
    {#if queueData.client}
      {@render clientData(queueData.client)}
    {/if}
    <TextField
      placeholder="ej. 10"
      title="Operaciones requeridas"
      required
      type="number"
      min={10}
      bind:value={queueData.requiredOperations}
    />
    {#if queueError}
      <p class="text-xs text-red-500" transition:slide={{ duration: 200 }}>
        {queueError}
      </p>
    {/if}
    <Button disabled={saving} class="w-full">Enviar a la cola</Button>
  </form>
{/snippet}

{#snippet clientData(client: Client)}
  <div class="flex flex-col space-y-2 w-full text-left">
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
  </div>
{/snippet}

<div class="w-full h-full flex flex-col space-y-4 flex-grow">
  <div class="flex w-full justify-between items-center">
    <div class="flex items-center space-x-2">
      <span class="i-material-symbols:account-circle-outline text-2xl"></span>
      <h2 class="text-lg font-bold leading-none">Clientes</h2>
    </div>
    <Button
      type="button"
      class="text-xs"
      onclick={() => {
        bottomsheet.open = true;
        bottomsheet.title = bottomsheetTitle;
        bottomsheet.body = bottomsheetBody;
      }}>Registrar cliente</Button
    >
  </div>
  <TextField
    title="Buscar cliente"
    placeholder="Escribe su nombre o DPI"
    type="search"
    bind:value={searchStr}
  />
  <div
    class="flex rounded-lg flex-grow skeleton w-full border border-light-900 overflow-auto divide-y divide-light-900 flex-col h-0"
    class:border-transparent={clientsCtx.loading}
    class:skeleton={clientsCtx.loading}
  >
    {#if !clientsCtx.loading}
      {#each filteredItems as client}
        <button
          class="px-4 py-2 bg-transparent hover:bg-light-500"
          onclick={() => {
            queueData.client = client;
            bottomsheet.open = true;
            bottomsheet.title = bottomsheetQueueTitle;
            bottomsheet.body = bottomsheetQueueBody;
          }}
        >
          {@render clientData(client)}
        </button>
      {:else}
        <div class="flex w-full h-full items-center justify-center">
          <p class="text-lg font-bold">Sin resultados</p>
        </div>
      {/each}
    {/if}
  </div>
</div>
