import type { Writable } from "svelte/store";
import { writable } from "svelte/store";
import type { GenericData } from "../schemas/generic/GenericData";
import { Themes } from "../configs/Themes";

export const genericDataStore: Writable<GenericData> = writable({
    "themeName": "dark",
    "theme": Themes["dark"],
    "showHamburger": false,
    mainPageSectionOneVersion: 2
});
