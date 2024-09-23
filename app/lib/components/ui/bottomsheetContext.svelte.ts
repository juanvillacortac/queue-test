import { getContext, setContext, type Snippet } from "svelte";

export class BottomsheetContext {
  open: boolean = $state(false);
  title?: Snippet = $state();
  body?: Snippet<[{ maxHeight: number }]> = $state();
  containerHeight = $state(0);

  close() {
    this.open = false;
    this.title = undefined;
    this.body = undefined;
  }
}

const bottomsheetContextKey = Symbol("bottomsheetContextKey");

export default function getBottomsheetContext(
  key: any = bottomsheetContextKey,
): BottomsheetContext {
  let ctx = getContext<BottomsheetContext | undefined>(key);
  if (!ctx) {
    ctx = setContext(key, new BottomsheetContext());
  }
  return ctx;
}
