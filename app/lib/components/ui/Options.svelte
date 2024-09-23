<script lang="ts">
  import type { HTMLSelectAttributes } from "svelte/elements";

  let {
    class: classNames,
    value = $bindable(),
    title,
    options = [],
    ...props
  }: HTMLSelectAttributes & {
    value?: any;
    options?: {
      text: string;
      value: any;
    }[];
  } = $props();
</script>

{#snippet inputSnippet()}
  <select
    {...props}
    class="flex p-16px rounded-8px bg-white w-full border-2px leading-none {classNames}"
    bind:value
  >
    <option hidden>Ninguno</option>
    {#each options as item}
      <option value={item.value}>{item.text}</option>
    {/each}
  </select>
{/snippet}

{#if title}
  <label class="flex flex-col space-y-1">
    <span class="text-sm">{title}</span>
    {@render inputSnippet()}
  </label>
{:else}
  {@render inputSnippet()}
{/if}

<style>
  select:focus {
    outline: none;
  }
</style>
