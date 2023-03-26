import type { Theme } from "../schemas/generic/Theme";

export const Themes: { [themeName: string]: Theme } = {
    "dark": {
        background: "#000000",
        surface: "#2c0b08",
        primary: "#dd5500",
        secondary: "#3f8d8a",
        onBackground: "#f0cc0f",
        onSurface: "#b3a5d7",
        onPrimary: "#eeffee",
        onSecondary: "#ffffff"
    }
};
