import { defineFormKitConfig } from "@formkit/vue";
import { rootClasses } from "./formkit.theme";
import {
  createAutoAnimatePlugin,
  createFloatingLabelsPlugin,
  createLocalStoragePlugin,
} from "@formkit/addons";

export default defineFormKitConfig({
  config: {
    rootClasses,
  },
  plugins: [
    createAutoAnimatePlugin(),
    createFloatingLabelsPlugin({ useAsDefault: true }),
    createLocalStoragePlugin({}),
  ],
});
