<script lang="ts">
  import { goto } from "$app/navigation";
  import { env } from "$env/dynamic/public";
  import Button from "$lib/components/ui/Button.svelte";
  import TextField from "$lib/components/ui/TextField.svelte";
  import { slide } from "svelte/transition";

  let formData = $state({
    email: "",
    password: "",
  });
  let loading = $state(false);
  let error: string = $state("");

  async function login() {
    if (loading) return;
    loading = true;
    let res: Response | undefined;
    try {
      res = await fetch(`${env.PUBLIC_API_PREFIX || ""}/api/auth/login`, {
        method: "POST",
        credentials: "include",
        body: JSON.stringify($state.snapshot(formData)),
      });
    } catch (err) {
      loading = false;
      error = `${err}`;
    }
    if (!res?.ok) {
      loading = false;
      const data = await res?.json();
      error = data?.error || "";
    } else {
      await goto("/", { invalidateAll: true, replaceState: true });
    }
  }
</script>

<form
  class="p4 rounded-lg bg-white shadow-lg max-w-36ch w-full space-y-2 flex flex-col"
  onsubmit={login}
>
  <TextField
    placeholder="ej. admin@example.com"
    title="Correo electrónico"
    type="email"
    required
    bind:value={formData.email}
  />
  <TextField
    title="Contraseña"
    type="password"
    required
    bind:value={formData.password}
  />
  {#if error}
    <p class="text-xs text-red-500" transition:slide={{ duration: 200 }}>
      {error}
    </p>
  {/if}
  <Button disabled={loading} class="w-full">Iniciar sesión</Button>
</form>
