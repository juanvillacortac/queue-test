import {
  defineConfig,
  presetAttributify,
  presetIcons,
  presetUno,
  transformerDirectives,
  transformerVariantGroup,
} from "unocss";
import presetAnimations from "unocss-preset-animations";
import presetWebFonts from "@unocss/preset-web-fonts";

export default defineConfig({
  shortcuts: {
    "flex-centered": "justify-center items-center",
    "size-screen": "w-full min-h-screen",
  },
  presets: [
    presetIcons(),
    presetUno({ dark: "class" }),
    presetAnimations(),
    presetAttributify({
      prefixedOnly: true,
      prefix: "data-un-",
    }),
    presetWebFonts({
      fonts: {
        sans: "Roboto",
      },
    }),
  ],
  transformers: [transformerVariantGroup(), transformerDirectives()],
});
