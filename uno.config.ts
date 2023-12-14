// uno.config.ts
import { defineConfig } from "unocss";

export default defineConfig({
  // ...UnoCSS options
  shortcuts: {
    "bg-main": "bg-[rgba(207,240,218,0.1)]",
    "title-main": "text-[#3ac569]",
    fcc: "flex justify-center items-center",
    /* NavList */
    "navlist-btn": "[&>li]:(px-2 text-4)",
    /* FoundMe */
    "found-me-btn": "[&>li]:(text-6 fcc p-1 rd-full w-8 h-8 my-2) [&>li>a]:(fcc w-full h-full inline-block)"
  },
});
