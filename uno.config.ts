// uno.config.ts
import { defineConfig } from "unocss";

export default defineConfig({
  // ...UnoCSS options
  shortcuts: {
    "bg-main": "bg-[rgba(207,240,218,0.1)]",
    "title-main": "text-[#3ac569]",
    fcc: "flex justify-center items-center",
    /* NavList */
    "navlist-btn": "[&>li]:(px-2 text-5)",
  },
});
