<script lang="ts">
  import type { HTMLInputAttributes } from "svelte/elements";

  let {
    class: classNames,
    value = $bindable(),
    type = "text",
    title,
    ...props
  }: HTMLInputAttributes & {
    value?: string;
  } = $props();
</script>

{#snippet inputSnippet()}
  <input
    {...props}
    {type}
    class="flex p-16px rounded-8px bg-white w-full border-2px leading-none {classNames}"
    bind:value
  />
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
  input:focus {
    outline: none;
  }
</style>
