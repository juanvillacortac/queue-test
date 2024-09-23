<script>
  import Clients from "$lib/components/Clients.svelte";
  import History from "$lib/components/History.svelte";
  import Queue from "$lib/components/Queue.svelte";
  import getBottomsheetContext from "$lib/components/ui/bottomsheetContext.svelte";

  getBottomsheetContext();
</script>

<div class="grid w-full h-full grid-cols-1 sm:grid-cols-2 gap-4">
  <Clients />
  <dig class="grid w-full h-full grid-cols-1 gap-4 grid-rows-2">
    <Queue />
    <History />
  </dig>
</div>
