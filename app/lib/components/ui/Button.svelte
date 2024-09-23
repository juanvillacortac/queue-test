<script lang="ts">
  import type { HTMLButtonAttributes } from "svelte/elements";

  let {
    class: classNames,
    children,
    ...props
  }: HTMLButtonAttributes & {
    value?: string;
  } = $props();
</script>

<button
  {...props}
  class="bg-dark-900 text-white font-bold p4 rounded-8px flex transition-colors duration-200 hover:bg-dark-400 disabled:bg-dark-50 text-center items-center justify-center {classNames}"
  >{@render children?.()}</button
>
