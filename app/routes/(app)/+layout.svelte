<script lang="ts">
  let { children } = $props();
</script>

<div class="p32px flex w-full flex-grow">
  <div class="p4 rounded-xl bg-white shadow-lg w-full flex flex-col flex-grow">
    {@render children()}
  </div>
</div>
