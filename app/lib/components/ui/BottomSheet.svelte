<script lang="ts">
  import { fade, fly } from "svelte/transition";
  import getBottomsheetContext from "./bottomsheetContext.svelte";

  const ctx = getBottomsheetContext();
</script>

{#if ctx.open}
  <div class="h-full w-full flex flex-col pointer-events-auto">
    <div
      class="absolute bg-dark-500 bg-opacity-50 flex w-full h-full inset-0 pointer-events-auto z20 cursor-pointer"
      role="none"
      transition:fade={{ duration: 200 }}
      onclick={() => {
        ctx.open = false;
        ctx.title = undefined;
        ctx.body = undefined;
      }}
    ></div>
    <div
      class="absolute w-full h-full flex px4 flex-col z20 justify-end pointer-events-none overflow-hidden"
    >
      <div
        class="flex flex-col w-full items-center"
        transition:fly={{ y: "50%", duration: 200 }}
      >
        <div
          class="rounded-full flex w-32 min-h-2 bg-gray-100 opacity-50 mb-4"
        ></div>
        <div
          class="rounded-t-xl bg-white flex flex-col w-full p4 shadow-xl pointer-events-auto space-y-4 lg:max-w-prose"
        >
          {@render ctx.title?.()}
          {#if ctx.body}
            <div
              class="w-full flex-col flex"
              style:max-height="{ctx.containerHeight * 0.65}px"
            >
              <div
                class="flex flex-col w-full overflow-auto body-container space-y-2"
              >
                {@render ctx.body?.({
                  maxHeight: ctx.containerHeight * 0.65,
                })}
              </div>
            </div>
          {/if}
        </div>
      </div>
    </div>
  </div>
{/if}

<style>
  .body-container {
    -ms-overflow-style: none; /* Internet Explorer 10+ */
    scrollbar-width: none; /* Firefox */
  }
  .body-container::-webkit-scrollbar {
    display: none; /* Safari and Chrome */
  }
</style>
