import UnoCSS from "unocss/vite";
import extractorSvelte from "@unocss/extractor-svelte";
import { sveltekit } from "@sveltejs/kit/vite";
import { defineConfig } from "vite";

export default defineConfig({
	plugins: [
		UnoCSS({
			extractors: [extractorSvelte()],
		}),
		sveltekit(),
	],
});
