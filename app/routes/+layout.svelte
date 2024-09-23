<script lang="ts">
  import BottomSheet from "$lib/components/ui/BottomSheet.svelte";
  import getBottomsheetContext from "$lib/components/ui/bottomsheetContext.svelte";
  import "$lib/styles";

  let { children } = $props();
  let bottomsheetCtx = getBottomsheetContext();
</script>

<div
  class="flex flex-col min-h-screen w-full relative"
  bind:clientHeight={bottomsheetCtx.containerHeight}
>
  <div
    class="absolute w-full h-full flex-col flex inset-0 z20 pointer-events-none"
  >
    <BottomSheet />
  </div>
  <div
    class="w-full h-full flex-grow flex-col bg-gradient-to-rb from-blue-300 to-purple-900 flex flex-centered"
  >
    {@render children?.()}
  </div>
</div>
